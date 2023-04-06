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
	Sensitivity         = big.NewInt(8)  // 8 ) // 1 ) // 8)
)

// MakeLDPCDifficultyCalculator calculate difficulty using difficulty table
func MakeLDPCDifficultyCalculator() func(time uint64, parent *types.Header) *big.Int {
	return func(time uint64, parent *types.Header) *big.Int {
		bigTime := new(big.Int).SetUint64(time)
		bigParentTime := new(big.Int).SetUint64(parent.Time)

		x := new(big.Int)
		x.Sub(bigTime, bigParentTime)

		diff := new(big.Int)

		// 이전블록이 제니시스 블록인 경우 난이도를 100으로 조정.
		if (parent.Difficulty).Cmp(big.NewInt(500000)) > 0 {
			diff = big.NewInt(10)

		} else {

			if x.Cmp(BlockGenerationTime) < 0 { // 36초 보다 빠르게 생성
				diff.Add(parent.Difficulty, big.NewInt(1)) // 난이도 1 증가
			} else {
				diff.Sub(parent.Difficulty, big.NewInt(1)) // 아니면, 난이도 1 감소
			}

		}

		// fix 최소 난이도 추가해야함

		return diff
	}
}

// SearchLevel return next level by using currentDifficulty of header
// Type of Ethereum difficulty is *bit.Int so arg is *big.int
func SearchLevel(difficulty *big.Int) int {
	// foo := MakeLDPCDifficultyCalculator()
	// Next level := SearchNextLevel(foo(currentBlock's time stamp, parentBlock))

	// var currentProb = DifficultyToProb(difficulty)
	// var level int

	// 이전블록이 제니시스 블록인 경우 난이도를 100으로 조정.
	if difficulty.Cmp(big.NewInt(500000)) > 0 {
		difficulty = big.NewInt(10)
	}

	var level int = int(difficulty.Uint64())
	// if difficulty.Cmp(big.NewInt(2)) < 0 {
	// 	level = level + 1
	// } else {bl
	// 	if level != 0 {
	// 		level = level - 1
	// 	}
	// }

	// distance := 1.0
	// for i := range Table {
	// 	if math.Abs(currentProb-Table[i].miningProb) <= distance {
	// 		level = Table[i].level
	// 		distance = math.Abs(currentProb - Table[i].miningProb)
	// 	} else {
	// 		break
	// 	}
	// }

	return level
}
