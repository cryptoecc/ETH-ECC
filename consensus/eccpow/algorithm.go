package eccpow

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"math/rand"

	//"reflect"

	"sync"
	"time"

	"github.com/Onther-Tech/go-ethereum/common"
	"github.com/Onther-Tech/go-ethereum/consensus"
	"github.com/Onther-Tech/go-ethereum/core/types"
	"github.com/Onther-Tech/go-ethereum/crypto"
	"github.com/Onther-Tech/go-ethereum/metrics"
	"github.com/Onther-Tech/go-ethereum/rpc"
)

type ECC struct {
	config Config

	// Mining related fields
	rand     *rand.Rand    // Properly seeded random source for nonces
	threads  int           // Number of threads to mine on if mining
	update   chan struct{} // Notification channel to update mining parameters
	hashrate metrics.Meter // Meter tracking the average hashrate

	// Remote sealer related fields
	workCh       chan *sealTask   // Notification channel to push new work and relative result channel to remote sealer
	fetchWorkCh  chan *sealWork   // Channel used for remote sealer to fetch mining work
	submitWorkCh chan *mineResult // Channel used for remote sealer to submit their mining result
	fetchRateCh  chan chan uint64 // Channel used to gather submitted hash rate for local or remote sealer.
	submitRateCh chan *hashrate   // Channel used for remote sealer to submit their mining hashrate

	shared    *ECC          // Shared PoW verifier to avoid cache regeneration
	fakeFail  uint64        // Block number which fails PoW check even in fake mode
	fakeDelay time.Duration // Time delay to sleep for before returning from verify

	lock      sync.Mutex      // Ensures thread safety for the in-memory caches and mining fields
	closeOnce sync.Once       // Ensures exit channel will not be closed twice.
	exitCh    chan chan error // Notification channel to exiting backend threads

}

type Mode uint

const (
	ModeNormal Mode = iota
	//ModeShared
	ModeTest
	ModeFake
	ModeFullFake
)

// Config are the configuration parameters of the ethash.
type Config struct {
	PowMode Mode
}

// sealTask wraps a seal block with relative result channel for remote sealer thread.
type sealTask struct {
	block   *types.Block
	results chan<- *types.Block
}

// mineResult wraps the pow solution parameters for the specified block.
type mineResult struct {
	nonce     types.BlockNonce
	mixDigest common.Hash
	hash      common.Hash

	errc chan error
}

// hashrate wraps the hash rate submitted by the remote sealer.
type hashrate struct {
	id   common.Hash
	ping time.Time
	rate uint64

	done chan struct{}
}

// sealWork wraps a seal work package for remote sealer.
type sealWork struct {
	errc chan error
	res  chan [4]string
}

// hasher is a repetitive hasher allowing the same hash data structures to be
// reused between hash runs instead of requiring new ones to be created.
//var hasher func(dest []byte, data []byte)

var (
	two256 = new(big.Int).Exp(big.NewInt(2), big.NewInt(256), big.NewInt(0))

	sharedECC = New(Config{ModeNormal}, nil, false)
)

type verifyParameters struct {
	n          uint64
	m          uint64
	wc         uint64
	wr         uint64
	seed       uint64
	outputWord []uint64
}

//const cross_err = 0.01

//type (
//	intMatrix   [][]int
//	floatMatrix [][]float64
//)

//RunOptimizedConcurrencyLDPC use goroutine for mining block
func RunOptimizedConcurrencyLDPC(header *types.Header, hash []byte) ([]int, []int, uint64, []byte) {
	//Need to set difficulty before running LDPC
	// Number of goroutines : 500, Number of attempts : 50000 Not bad

	var LDPCNonce uint64
	var hashVector []int
	var outputWord []int
	var digest []byte

	var wg sync.WaitGroup
	var outerLoopSignal = make(chan struct{})
	var innerLoopSignal = make(chan struct{})
	var goRoutineSignal = make(chan struct{})

	parameters, _ := setParameters(header)
	H := generateH(parameters)
	colInRow, rowInCol := generateQ(parameters, H)

outerLoop:
	for {
		select {
		// If outerLoopSignal channel is closed, then break outerLoop
		case <-outerLoopSignal:
			break outerLoop

		default:
			// Defined default to unblock select statement
		}

	innerLoop:
		//for i := 0; i < runtime.NumCPU(); i++ {
		for i := 0; i < 1; i++ {
			select {
			// If innerLoop signal is closed, then break innerLoop and close outerLoopSignal
			case <-innerLoopSignal:
				close(outerLoopSignal)
				break innerLoop

			default:
				// Defined default to unblock select statement
			}

			wg.Add(1)
			go func(goRoutineSignal chan struct{}) {
				defer wg.Done()
				//goRoutineNonce := generateRandomNonce()
				//fmt.Printf("Initial goroutine Nonce : %v\n", goRoutineNonce)

				var goRoutineHashVector []int
				var goRoutineOutputWord []int

				select {
				case <-goRoutineSignal:
					break

				default:
				attemptLoop:
					for attempt := 0; attempt < 5000; attempt++ {
						goRoutineNonce := generateRandomNonce()
						seed := make([]byte, 40)
						copy(seed, hash)
						binary.LittleEndian.PutUint64(seed[32:], goRoutineNonce)
						seed = crypto.Keccak512(seed)

						goRoutineHashVector = generateHv(parameters, seed)
						goRoutineHashVector, goRoutineOutputWord, _ = OptimizedDecoding(parameters, goRoutineHashVector, H, rowInCol, colInRow)
						flag := MakeDecision(header, colInRow, goRoutineOutputWord)

						select {
						case <-goRoutineSignal:
							// fmt.Println("goRoutineSignal channel is already closed")
							break attemptLoop
						default:
							if flag {
								close(goRoutineSignal)
								close(innerLoopSignal)
								fmt.Printf("Codeword is founded with nonce = %d\n", goRoutineNonce)
								fmt.Printf("Codeword : %d\n", goRoutineOutputWord)
								hashVector = goRoutineHashVector
								outputWord = goRoutineOutputWord
								LDPCNonce = goRoutineNonce
								digest = seed
								break attemptLoop
							}
						}
						//goRoutineNonce++
					}
				}
			}(goRoutineSignal)
		}
		// Need to wait to prevent memory leak
		wg.Wait()
	}

	return hashVector, outputWord, LDPCNonce, digest
}

//MakeDecision check outputWord is valid or not using colInRow
func MakeDecision(header *types.Header, colInRow [][]int, outputWord []int) bool {
	parameters, difficultyLevel := setParameters(header)
	for i := 0; i < parameters.m; i++ {
		sum := 0
		for j := 0; j < parameters.wr; j++ {
			//	fmt.Printf("i : %d, j : %d, m : %d, wr : %d \n", i, j, m, wr)
			sum = sum + outputWord[colInRow[j][i]]
		}
		if sum%2 == 1 {
			return false
		}
	}

	var numOfOnes int
	for _, val := range outputWord {
		numOfOnes += val
	}

	if numOfOnes >= Table[difficultyLevel].decisionFrom &&
		numOfOnes <= Table[difficultyLevel].decisionTo &&
		numOfOnes%Table[difficultyLevel].decisionStep == 0 {
		return true
	}

	return false
}

//func isRegular(nSize, wCol, wRow int) bool {
//	res := float64(nSize*wCol) / float64(wRow)
//	m := math.Round(res)
//
//	if int(m)*wRow == nSize*wCol {
//		return true
//	}
//
//	return false
//}

//func SetDifficulty(nSize, wCol, wRow int) bool {
//	if isRegular(nSize, wCol, wRow) {
//		n = nSize
//		wc = wCol
//		wr = wRow
//		m = int(n * wc / wr)
//		return true
//	}
//	return false
//}

//func newIntMatrix(rows, cols int) intMatrix {
//	m := intMatrix(make([][]int, rows))
//	for i := range m {
//		m[i] = make([]int, cols)
//	}
//	return m
//}
//
//func newFloatMatrix(rows, cols int) floatMatrix {
//	m := floatMatrix(make([][]float64, rows))
//	for i := range m {
//		m[i] = make([]float64, cols)
//	}
//	return m
//}

// New creates a full sized ethash PoW scheme and starts a background thread for
// remote mining, also optionally notifying a batch of remote services of new work
// packages.
func New(config Config, notify []string, noverify bool) *ECC {
	ecc := &ECC{
		config:       config,
		update:       make(chan struct{}),
		hashrate:     metrics.NewMeterForced(),
		workCh:       make(chan *sealTask),
		fetchWorkCh:  make(chan *sealWork),
		submitWorkCh: make(chan *mineResult),
		fetchRateCh:  make(chan chan uint64),
		submitRateCh: make(chan *hashrate),
		exitCh:       make(chan chan error),
	}
	go ecc.remote(notify, noverify)
	return ecc
}

func NewTester(notify []string, noverify bool) *ECC {
	ecc := &ECC{
		config:       Config{PowMode: ModeTest},
		update:       make(chan struct{}),
		hashrate:     metrics.NewMeterForced(),
		workCh:       make(chan *sealTask),
		fetchWorkCh:  make(chan *sealWork),
		submitWorkCh: make(chan *mineResult),
		fetchRateCh:  make(chan chan uint64),
		submitRateCh: make(chan *hashrate),
		exitCh:       make(chan chan error),
	}
	go ecc.remote(notify, noverify)
	return ecc
}

// NewFaker creates a ethash consensus engine with a fake PoW scheme that accepts
// all blocks' seal as valid, though they still have to conform to the Ethereum
// consensus rules.
func NewFaker() *ECC {
	return &ECC{
		config: Config{
			PowMode: ModeFake,
		},
	}
}

// NewFakeFailer creates a ethash consensus engine with a fake PoW scheme that
// accepts all blocks as valid apart from the single one specified, though they
// still have to conform to the Ethereum consensus rules.
func NewFakeFailer(fail uint64) *ECC {
	return &ECC{
		config: Config{
			PowMode: ModeFake,
		},
		fakeFail: fail,
	}
}

// NewFakeDelayer creates a ethash consensus engine with a fake PoW scheme that
// accepts all blocks as valid, but delays verifications by some time, though
// they still have to conform to the Ethereum consensus rules.
func NewFakeDelayer(delay time.Duration) *ECC {
	return &ECC{
		config: Config{
			PowMode: ModeFake,
		},
		fakeDelay: delay,
	}
}

// NewFullFaker creates an ethash consensus engine with a full fake scheme that
// accepts all blocks as valid, without checking any consensus rules whatsoever.
func NewFullFaker() *ECC {
	return &ECC{
		config: Config{
			PowMode: ModeFullFake,
		},
	}
}

// NewShared creates a full sized ethash PoW shared between all requesters running
// in the same process.
//func NewShared() *ECC {
//	return &ECC{shared: sharedECC}
//}

// Close closes the exit channel to notify all backend threads exiting.
func (ecc *ECC) Close() error {
	var err error
	ecc.closeOnce.Do(func() {
		// Short circuit if the exit channel is not allocated.
		if ecc.exitCh == nil {
			return
		}
		errc := make(chan error)
		ecc.exitCh <- errc
		err = <-errc
		close(ecc.exitCh)
	})
	return err
}

// Threads returns the number of mining threads currently enabled. This doesn't
// necessarily mean that mining is running!
func (ecc *ECC) Threads() int {
	ecc.lock.Lock()
	defer ecc.lock.Unlock()

	return ecc.threads
}

// SetThreads updates the number of mining threads currently enabled. Calling
// this method does not start mining, only sets the thread count. If zero is
// specified, the miner will use all cores of the machine. Setting a thread
// count below zero is allowed and will cause the miner to idle, without any
// work being done.
func (ecc *ECC) SetThreads(threads int) {
	ecc.lock.Lock()
	defer ecc.lock.Unlock()

	// If we're running a shared PoW, set the thread count on that instead
	if ecc.shared != nil {
		ecc.shared.SetThreads(threads)
		return
	}
	// Update the threads and ping any running seal to pull in any changes
	ecc.threads = threads
	select {
	case ecc.update <- struct{}{}:
	default:
	}
}

// Hashrate implements PoW, returning the measured rate of the search invocations
// per second over the last minute.
// Note the returned hashrate includes local hashrate, but also includes the total
// hashrate of all remote miner.
func (ecc *ECC) Hashrate() float64 {
	// Short circuit if we are run the ecc in normal/test mode.

	var res = make(chan uint64, 1)

	select {
	case ecc.fetchRateCh <- res:
	case <-ecc.exitCh:
		// Return local hashrate only if ecc is stopped.
		return ecc.hashrate.Rate1()
	}

	// Gather total submitted hash rate of remote sealers.
	return ecc.hashrate.Rate1() + float64(<-res)
}

// APIs implements consensus.Engine, returning the user facing RPC APIs.
func (ecc *ECC) APIs(chain consensus.ChainReader) []rpc.API {
	// In order to ensure backward compatibility, we exposes ecc RPC APIs
	// to both eth and ecc namespaces.
	return []rpc.API{
		{
			Namespace: "eth",
			Version:   "1.0",
			Service:   &API{ecc},
			Public:    true,
		},
		{
			Namespace: "ecc",
			Version:   "1.0",
			Service:   &API{ecc},
			Public:    true,
		},
	}
}

//// SeedHash is the seed to use for generating a verification cache and the mining
//// dataset.
func SeedHash(block *types.Block) []byte {
	return block.ParentHash().Bytes()
}
