package eccpow

import (
	"math"
	"math/big"
	"log"
	"github.com/ethereum/go-ethereum/core/types"
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
	
	//100번 전 부모의 타임스탬프.
	GrandParentTime = new(big.Int).SetUint64(0)
	
	//initLevel int = 10
	minLevel  int = 10

	//count  int = -1
	init_c int = 2
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

func MakeLDPCDifficultyCalculator_Seoul() func(time uint64, parent *types.Header) *big.Int {
	return func(time uint64, parent *types.Header) *big.Int {
		bigTime := new(big.Int).SetUint64(time)
		bigParentTime := new(big.Int).SetUint64(parent.Time)
		x := new(big.Int)
		diff := new(big.Int)
		index := int(parent.Number.Uint64()) % 100

		if index == 0 { // 난이도 변경 시점인 경우.
			
			level := SearchLevel(parent.Difficulty)

			//처음인 경우.
			if GrandParentTime.Uint64() == 0 {
				GrandParentTime = bigParentTime
				diff = ProbToDifficulty(Table[level].miningProb)
				return diff
			} 

			x.Sub(bigTime, GrandParentTime)
			avgTime := int(x.Uint64()) / 100

			if avgTime < stTime {
				level += 1
			} else {
				if level > minLevel {
					level -= 1
				}
			}
			diff = ProbToDifficulty(Table[level].miningProb)
			//할아버지 시간 갱신.
			GrandParentTime = bigParentTime

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
