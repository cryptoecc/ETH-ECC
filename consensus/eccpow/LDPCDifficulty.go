package eccpow

import (
	"fmt"
	"math"
	"math/big"

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
	BlockGenerationTime = big.NewInt(10) // 36) // 10 ) // 36)
	Sensitivity         = big.NewInt(8)

	stTime      int = 12
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

	initLevel int = 30
	minLevel  int = 5

	count  int = -1
	init_c int = 2
)

// MakeLDPCDifficultyCalculator calculate difficulty using difficulty table
func MakeLDPCDifficultyCalculator() func(time uint64, parent *types.Header) *big.Int {
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
}

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
