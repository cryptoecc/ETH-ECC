package eccpow

import (
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

	initDiff = big.NewInt(1)
	minDiff  = big.NewInt(5)

	count int = 0
)

// MakeLDPCDifficultyCalculator calculate difficulty using difficulty table
func MakeLDPCDifficultyCalculator() func(time uint64, parent *types.Header) *big.Int {
	return func(time uint64, parent *types.Header) *big.Int {
		bigTime := new(big.Int).SetUint64(time)
		bigParentTime := new(big.Int).SetUint64(parent.Time)

		x := new(big.Int)
		x.Sub(bigTime, bigParentTime)

		diff := new(big.Int)

		if count > 1000 {
			count = count - 1000
		}
		count = count + 1

		if (parent.Difficulty).Cmp(big.NewInt(500000)) > 0 {
			diff = big.NewInt(120)
			return diff
		}

		bIndex := (parent.Number).Uint64() % 100
		avgTimeList[bIndex] = int(x.Uint64())

		totalTime := 0
		for i := 0; i < 100; i++ {
			totalTime += avgTimeList[i]
		}
		avgTime := totalTime / 100

		if count%100 == 42 {
			if avgTime < stTime {
				diff.Add(parent.Difficulty, big.NewInt(1))
			} else {
				if (parent.Difficulty).Cmp(minDiff) > 0 {
					diff.Sub(parent.Difficulty, big.NewInt(1))
				}
			}
		} else {
			diff = parent.Difficulty
		}

		return diff
	}
}

// SearchLevel return next level by using currentDifficulty of header
// Type of Ethereum difficulty is *bit.Int so arg is *big.int
func SearchLevel(difficulty *big.Int) int {
	if difficulty.Cmp(big.NewInt(500000)) > 0 {
		difficulty = big.NewInt(120)
	}

	var level int = int(difficulty.Uint64())

	return level
}
