package eccpow

import (
	"math"
	"math/big"
	"log"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"
	"github.com/ethereum/go-ethereum/consensus"
)

/*
	https://ethereum.stackexchange.com/questions/5913/how-does-the-ethereum-homestead-difficulty-adjustment-algorithm-work?noredirect=1&lq=1
	https://github.com/ethereum/EIPs/issues/100

	Ethereum difficulty adjustment
	 algorithm:
	diff = (parent_diff +
	         (parent_diff / 2048 * max((2 if len(parent.uncles) else 1) - ((timestamp - parent.timestamp) // 9), -99))
	        ) + 2^(periodCount - 2)

	LDPC difficulty adjustment
	algorithm:
	diff = (parent_diff +
			(parent_diff / 256 * max((2 if len(parent.uncles) else 1) - ((timestamp - parent.timestamp) // BlockGenerationTime), -99)))

	Why 8?
	This number is sensitivity of blockgeneration time
	If this number is high, difficulty is not changed much when block generation time is different from goal of block generation time
	But if this number is low, difficulty is changed much when block generatin time is different  from goal of block generation time

*/

// "github.com/ethereum/go-ethereum/consensus/ethash/consensus.go"
// Some weird constants to avoid constant memory allocs for them.
var (
	MinimumDifficulty   = ProbToDifficulty(Table[0].miningProb)
	BlockGenerationTime = big.NewInt(36) // 36) // 10 ) // 36)
	Sensitivity         = big.NewInt(8)

	// BlockGenerationTime for Seoul
	stTime      int = 10

	avgTimeList     = [100]int{stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime,
		stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime,
		stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime,
		stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime,
		stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime,
		stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime,
		stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime,
		stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime,
		stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime,
		stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime, stTime}

	
	//initLevel int = 10
	minLevel  int = 10
	diff_interval = 100

	//count  int = -1
	//init_c int = 2
)

const (
	// frontierDurationLimit is for Frontier:
	// The decision boundary on the blocktime duration used to determine
	// whether difficulty should go up or down.
	frontierDurationLimit = 10
	// minimumDifficulty The minimum that the difficulty may ever be.
	minimumDifficulty = 131072
	// expDiffPeriod is the exponential difficulty period
	expDiffPeriodUint = 100000
	// difficultyBoundDivisorBitShift is the bound divisor of the difficulty (2048),
	// This constant is the right-shifts to use for the division.
	difficultyBoundDivisor = 11
)

// MakeLDPCDifficultyCalculator calculate difficulty using difficulty table
func MakeLDPCDifficultyCalculator() func(time uint64, parent *types.Header) *big.Int {
	return func(time uint64, parent *types.Header) *big.Int {
		bigTime := new(big.Int).SetUint64(time)
		bigParentTime := new(big.Int).SetUint64(parent.Time)

		// holds intermediate values to make the algo easier to read & audit
		x := new(big.Int)
		y := new(big.Int)

		// (2 if len(parent_uncles) else 1) - (block_timestamp - parent_timestamp) // BlockGenerationTime
		x.Sub(bigTime, bigParentTime)
		//fmt.Printf("block_timestamp - parent_timestamp : %v\n", x)

		x.Div(x, BlockGenerationTime)
		//fmt.Printf("(block_timestamp - parent_timestamp) / BlockGenerationTime : %v\n", x)

		if parent.UncleHash == types.EmptyUncleHash {
			//fmt.Printf("No uncle\n")
			x.Sub(big1, x)
		} else {
			//fmt.Printf("Uncle block exists")
			x.Sub(big2, x)
		}
		//fmt.Printf("(2 if len(parent_uncles) else 1) - (block_timestamp - parent_timestamp) / BlockGenerationTime : %v\n", x)

		// max((2 if len(parent_uncles) else 1) - (block_timestamp - parent_timestamp) // 9, -99)
		if x.Cmp(bigMinus99) < 0 {
			x.Set(bigMinus99)
		}
		//fmt.Printf("max(1 - (block_timestamp - parent_timestamp) / BlockGenerationTime, -99) : %v\n", x)

		// parent_diff + (parent_diff / Sensitivity * max((2 if len(parent.uncles) else 1) - ((timestamp - parent.timestamp) // BlockGenerationTime), -99))
		y.Div(parent.Difficulty, Sensitivity)
		//fmt.Printf("parent.Difficulty / 8 : %v\n", y)

		x.Mul(y, x)
		//fmt.Printf("parent.Difficulty / 8 * max(1 - (block_timestamp - parent_timestamp) / BlockGenerationTime, -99) : %v\n", x)

		x.Add(parent.Difficulty, x)
		//fmt.Printf("parent.Difficulty - parent.Difficulty / 8 * max(1 - (block_timestamp - parent_timestamp) / BlockGenerationTime, -99) : %v\n", x)

		// minimum difficulty can ever be (before exponential factor)
		if x.Cmp(MinimumDifficulty) < 0 {
			x.Set(MinimumDifficulty)
		}

		//fmt.Printf("x : %v, Minimum difficulty : %v\n", x, MinimumDifficulty)

		return x
	}
}

func MakeLDPCDifficultyCalculator_Seoul() func(chain consensus.ChainHeaderReader, time uint64, parent *types.Header) *big.Int {
	return func(chain consensus.ChainHeaderReader, time uint64, parent *types.Header) *big.Int {

		if parent.Number.Uint64() < uint64(diff_interval){
			return parent.Difficulty
		}

		bigTime := new(big.Int).SetUint64(time)
		bigParentTime := new(big.Int).SetUint64(parent.Time)
		x := new(big.Int)
		diff := new(big.Int)
		index := int(parent.Number.Uint64()) % diff_interval
		
		if index == 0 { // 난이도 변경 시점인 경우.
			level := SearchLevel(parent.Difficulty)
			grandParent := chain.GetHeaderByNumber(parent.Number.Uint64() - uint64(diff_interval))
			grandParentTime := new(big.Int).SetUint64(grandParent.Time)
			
			x.Sub(bigTime, grandParentTime)
			avgTime := int(x.Uint64()) / 100

			if avgTime < stTime {
				level += 1
			} else {
				if level > minLevel {
					level -= 1
				}
			}
			diff = ProbToDifficulty(Table[level].miningProb)

			//테스트 출력.
			log.Println("Level: ", level)
			log.Println("Average Time: ", avgTime, "s")
			log.Println("Time List: ", avgTimeList)
		} else { // 난이도 변경 시점이 아닌 경우.
			diff = parent.Difficulty
		}
		
		//분포 측정을 위한 코드.
		blocktime := new(big.Int)
		blocktime.Sub(bigTime, bigParentTime)
		avgTimeList[index] = int(blocktime.Uint64())

		return diff
	}
}

/*
func MakeLDPCDifficultyCalculator_Seoul2() func(time uint64, parent *types.Header) *big.Int {
	return func(time uint64, parent *types.Header) *big.Int {
		bigTime := new(big.Int).SetUint64(time)
		bigParentTime := new(big.Int).SetUint64(parent.Time)

		x := new(big.Int)
		x.Sub(bigTime, bigParentTime)

		level := SearchLevel(parent.Difficulty)
		diff := new(big.Int)

		count = count + 1

		if count < 1 {
			level = initLevel
		} else if count > 1000 {
			count = count - 1000
		}

		if init_c < 5 {
			level = initLevel
			x = big.NewInt(12)
			init_c += 1
		}

		bIndex := count % 100
		avgTimeList[bIndex] = int(x.Uint64())

		totalTime := 0
		for i := 0; i < 100; i++ {
			totalTime += avgTimeList[i]
		}
		avgTime := totalTime / 100

		if count%100 == 90 {
			if avgTime < stTime {
				level += 1
			} else {
				if level > minLevel {
					level -= 1
				}
			}
		}

		diff = ProbToDifficulty(Table[level].miningProb)

		fmt.Println("Index: ", bIndex)
		fmt.Println("Level: ", level)
		fmt.Println("Average Time: ", avgTime, "s")
		fmt.Println("Time List: ", avgTimeList)

		return diff
	}
}*/

// SearchLevel return next level by using currentDifficulty of header
// Type of Ethereum difficulty is *bit.Int so arg is *big.int
func SearchLevel(difficulty *big.Int) int {
	// foo := MakeLDPCDifficultyCalculator()
	// Next level := SearchNextLevel(foo(currentBlock's time stamp, parentBlock))

	var currentProb = DifficultyToProb(difficulty)
	var level int

	distance := 1.0
	for i := range Table {
		if math.Abs(currentProb-Table[i].miningProb) <= distance {
			level = Table[i].level
			distance = math.Abs(currentProb - Table[i].miningProb)
		} else {
			break
		}
	}

	return level
}

// CalcDifficultyFrontierU256 is the difficulty adjustment algorithm. It returns the
// difficulty that a new block should have when created at time given the parent
// block's time and difficulty. The calculation uses the Frontier rules.
func MakeLDPCDifficultyCalculator2() func(time uint64, parent *types.Header) *big.Int {
	return func(time uint64, parent *types.Header) *big.Int {
		/*
			Algorithm
			block_diff = pdiff + pdiff / 2048 * (1 if time - ptime < 13 else -1) + int(2^((num // 100000) - 2))

			Where:
			- pdiff  = parent.difficulty
			- ptime = parent.time
			- time = block.timestamp
			- num = block.number
		*/

		pDiff, _ := uint256.FromBig(parent.Difficulty) // pDiff: pdiff
		adjust := pDiff.Clone()
		adjust.Rsh(adjust, difficultyBoundDivisor) // adjust: pDiff / 2048

		if time-parent.Time < frontierDurationLimit {
			pDiff.Add(pDiff, adjust)
		} else {
			pDiff.Sub(pDiff, adjust)
		}
		if pDiff.LtUint64(minimumDifficulty) {
			pDiff.SetUint64(minimumDifficulty)
		}
		// 'pdiff' now contains:
		// pdiff + pdiff / 2048 * (1 if time - ptime < 13 else -1)

		if periodCount := (parent.Number.Uint64() + 1) / expDiffPeriodUint; periodCount > 1 {
			// diff = diff + 2^(periodCount - 2)
			expDiff := adjust.SetOne()
			expDiff.Lsh(expDiff, uint(periodCount-2)) // expdiff: 2 ^ (periodCount -2)
			pDiff.Add(pDiff, expDiff)
		}
		return pDiff.ToBig()
	}
}

// CalcDifficultyHomesteadU256 is the difficulty adjustment algorithm. It returns
// the difficulty that a new block should have when created at time given the
// parent block's time and difficulty. The calculation uses the Homestead rules.
func MakeLDPCDifficultyCalculator3() func(time uint64, parent *types.Header)*big.Int {
	
	return func(time uint64, parent *types.Header) *big.Int {
		/*
			https://github.com/ethereum/EIPs/blob/master/EIPS/eip-2.md
			Algorithm:
			block_diff = pdiff + pdiff / 2048 * max(1 - (time - ptime) / 10, -99) + 2 ^ int((num / 100000) - 2))

			Our modification, to use unsigned ints:
			block_diff = pdiff - pdiff / 2048 * max((time - ptime) / 10 - 1, 99) + 2 ^ int((num / 100000) - 2))

			Where:
			- pdiff  = parent.difficulty
			- ptime = parent.time
			- time = block.timestamp
			- num = block.number
		*/

		pDiff, _ := uint256.FromBig(parent.Difficulty) // pDiff: pdiff
		adjust := pDiff.Clone()
		adjust.Rsh(adjust, difficultyBoundDivisor) // adjust: pDiff / 2048

		x := (time - parent.Time) / 10 // (time - ptime) / 10)
		var neg = true
		if x == 0 {
			x = 1
			neg = false
		} else if x >= 100 {
			x = 99
		} else {
			x = x - 1
		}
		z := new(uint256.Int).SetUint64(x)
		adjust.Mul(adjust, z) // adjust: (pdiff / 2048) * max((time - ptime) / 10 - 1, 99)
		if neg {
			pDiff.Sub(pDiff, adjust) // pdiff - pdiff / 2048 * max((time - ptime) / 10 - 1, 99)
		} else {
			pDiff.Add(pDiff, adjust) // pdiff + pdiff / 2048 * max((time - ptime) / 10 - 1, 99)
		}
		if pDiff.LtUint64(minimumDifficulty) {
			pDiff.SetUint64(minimumDifficulty)
		}
		// for the exponential factor, a.k.a "the bomb"
		// diff = diff + 2^(periodCount - 2)
		if periodCount := (1 + parent.Number.Uint64()) / expDiffPeriodUint; periodCount > 1 {
			expFactor := adjust.Lsh(adjust.SetOne(), uint(periodCount-2))
			pDiff.Add(pDiff, expFactor)
		}
		return pDiff.ToBig()
	}
}

// MakeDifficultyCalculatorU256 creates a difficultyCalculator with the given bomb-delay.
// the difficulty is calculated with Byzantium rules, which differs from Homestead in
// how uncles affect the calculation
func MakeLDPCDifficultyCalculator4() func(time uint64, parent *types.Header)*big.Int {
	// Note, the calculations below looks at the parent number, which is 1 below
	// the block number. Thus we remove one from the delay given
	return func(time uint64, parent *types.Header) *big.Int {
		/*
			https://github.com/ethereum/EIPs/issues/100
			pDiff = parent.difficulty
			BLOCK_DIFF_FACTOR = 9
			a = pDiff + (pDiff // BLOCK_DIFF_FACTOR) * adj_factor
			b = min(parent.difficulty, MIN_DIFF)
			child_diff = max(a,b )
		*/
		x := (time - parent.Time) / 9 // (block_timestamp - parent_timestamp) // 9
		c := uint64(1)                // if parent.unclehash == emptyUncleHashHash
		if parent.UncleHash != types.EmptyUncleHash {
			c = 2
		}
		xNeg := x >= c
		if xNeg {
			// x is now _negative_ adjustment factor
			x = x - c // - ( (t-p)/p -( 2 or 1) )
		} else {
			x = c - x // (2 or 1) - (t-p)/9
		}
		if x > 99 {
			x = 99 // max(x, 99)
		}
		// parent_diff + (parent_diff / 2048 * max((2 if len(parent.uncles) else 1) - ((timestamp - parent.timestamp) // 9), -99))
		y := new(uint256.Int)
		y.SetFromBig(parent.Difficulty)    // y: p_diff
		pDiff := y.Clone()                 // pdiff: p_diff
		z := new(uint256.Int).SetUint64(x) //z : +-adj_factor (either pos or negative)
		y.Rsh(y, difficultyBoundDivisor)   // y: p__diff / 2048
		z.Mul(y, z)                        // z: (p_diff / 2048 ) * (+- adj_factor)

		if xNeg {
			y.Sub(pDiff, z) // y: parent_diff + parent_diff/2048 * adjustment_factor
		} else {
			y.Add(pDiff, z) // y: parent_diff + parent_diff/2048 * adjustment_factor
		}
		// minimum difficulty can ever be (before exponential factor)
		if y.LtUint64(minimumDifficulty) {
			y.SetUint64(minimumDifficulty)
		}
		// calculate a fake block number for the ice-age delay
		// Specification: https://eips.ethereum.org/EIPS/eip-1234
		return y.ToBig()
	}
}

