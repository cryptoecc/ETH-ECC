package main

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
	"time"
	"fmt"
)

func main() { // MonteCarlo simulation for probability that LDPC decoding success
	for i := 120; i <= 256; i += 4 {
		test(i)
	}
}

func test(numN int){
	parameters := Parameters { 
		n:  numN,
		wc: 3,
		wr: 4,
	}
	parameters.m = int(parameters.n * parameters.wc / parameters.wr)
	numSim := 100000000 // TODO: set as proper value, > 1e+9
	//numSim := 1000000 
	//numSim := 100 // TODO: set as proper value, > 1e+9
	decodeSuccess := 0
	hammingWeightHistogram := make([]int, numN+1)//initialized by zeros
	messageSeed := rand.NewSource(time.Now().UnixNano()) // TODO: check type cast
    messageRng := rand.New(messageSeed)
    startMilliSec := time.Now().UnixNano() / 1000000
	for cnt_sim := 0; cnt_sim < numSim ; cnt_sim++ {
		parameters.seed = int(time.Now().UnixNano()) // int64
		H := generateH(parameters)
		colInRow, rowInCol := generateQ(parameters, H)
		hashVector := make([]int, parameters.n) // randomly
		decoderFlag := true
		for i := 0; i < parameters.n; i++ {
			hashVector[i] = messageRng.Intn(2)
		}
        
		_, outputWord, _ := OptimizedDecoding(parameters, hashVector, H, rowInCol, colInRow)
		
		if cnt_sim % 1000000 == 0 {
			nowMilliSec := time.Now().UnixNano() / 1000000
			elapse := nowMilliSec - startMilliSec
			elapseSec := float64(elapse) / float64(1000)
			fmt.Printf("%d th simulation... (elapse: %.3f sec)\n", cnt_sim, elapseSec)
		}
		for i := 0; i < parameters.m; i++ {
			sum := 0
			for j := 0; j < parameters.wr; j++ {
				sum = sum + outputWord[colInRow[j][i]]
			}
			if sum %2 == 1 { // check if outputWord is codeword
				decoderFlag = false
				break;
			}
		}
		if decoderFlag {
			decodeSuccess++ // mark as success case
			numOfOnes := 0
			for _, val := range outputWord {
				numOfOnes += val
			}
			hammingWeightHistogram[numOfOnes] += 1 // store in histogram
		}
	}
	// TODO: save numSim, decodeSuccess, hammingWeightHistogram
	/*
	fmt.Println("number of simulations: ", numN)
	fmt.Println("number of simulations: ", numSim)
	fmt.Println("number of successive decoding: ", decodeSuccess)*/
	fmt.Println(numN, numSim, decodeSuccess)
	//fmt.Println("Hamming weight distribution of decoded codeword: ")
	fmt.Println(hammingWeightHistogram)
}

//OptimizedDecoding return hashVector, outputWord, LRrtl
func OptimizedDecoding(parameters Parameters, hashVector []int, H, rowInCol, colInRow [][]int) ([]int, []int, [][]float64) {
	outputWord := make([]int, parameters.n)
	LRqtl := make([][]float64, parameters.n)
	LRrtl := make([][]float64, parameters.n)
	LRft := make([]float64, parameters.n)

	for i := 0; i < parameters.n; i++ {
		LRqtl[i] = make([]float64, parameters.m)
		LRrtl[i] = make([]float64, parameters.m)
		LRft[i] = math.Log((1-crossErr)/crossErr) * float64((hashVector[i]*2 - 1))
	}
	LRpt := make([]float64, parameters.n)

	for ind := 1; ind <= maxIter; ind++ {
		for t := 0; t < parameters.n; t++ {
			temp3 := 0.0

			for mp := 0; mp < parameters.wc; mp++ {
				temp3 = infinityTest(temp3 + LRrtl[t][rowInCol[mp][t]])
			}
			for m := 0; m < parameters.wc; m++ {
				temp4 := temp3
				temp4 = infinityTest(temp4 - LRrtl[t][rowInCol[m][t]])
				LRqtl[t][rowInCol[m][t]] = infinityTest(LRft[t] + temp4)
			}
		}

		for k := 0; k < parameters.m; k++ {
			for l := 0; l < parameters.wr; l++ {
				temp3 := 0.0
				sign := 1.0
				tempSign := 0.0
				for m := 0; m < parameters.wr; m++ {
					if m != l {
						temp3 = temp3 + funcF(math.Abs(LRqtl[colInRow[m][k]][k]))
						if LRqtl[colInRow[m][k]][k] > 0.0 {
							tempSign = 1.0
						} else {
							tempSign = -1.0
						}
						sign = sign * tempSign
					}
				}
				magnitude := funcF(temp3)
				LRrtl[colInRow[l][k]][k] = infinityTest(sign * magnitude)
			}
		}

		for t := 0; t < parameters.n; t++ {
			LRpt[t] = infinityTest(LRft[t])
			for k := 0; k < parameters.wc; k++ {
				LRpt[t] += LRrtl[t][rowInCol[k][t]]
				LRpt[t] = infinityTest(LRpt[t])
			}

			if LRpt[t] >= 0 {
				outputWord[t] = 1
			} else {
				outputWord[t] = 0
			}
		}
	}

	return hashVector, outputWord, LRrtl
}

//Parameters for matrix and seed
const (
	BigInfinity = 1000000.0
	Inf         = 64.0
	MaxNonce    = 1<<32 - 1

	// These parameters are only used for the decoding function.
	maxIter  = 20   // The maximum number of iteration in the decoder
	crossErr = 0.01 // A transisient error probability. This is also fixed as a small value
)

type Parameters struct {
	n    int
	m    int
	wc   int
	wr   int
	seed int
}

//generateRandomNonce generate 64bit random nonce with similar way of ethereum block nonce
func generateRandomNonce() uint64 {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	source := rand.New(rand.NewSource(seed.Int64()))

	return uint64(source.Int63())
}

func funcF(x float64) float64 {
	if x >= BigInfinity {
		return 1.0 / BigInfinity
	} else if x <= (1.0 / BigInfinity) {
		return BigInfinity
	} else {
		return math.Log((math.Exp(x) + 1) / (math.Exp(x) - 1))
	}
}

func infinityTest(x float64) float64 {
	if x >= Inf {
		return Inf
	} else if x <= -Inf {
		return -Inf
	} else {
		return x
	}
}

//generateSeed generate seed using previous hash vector
func generateSeed(phv [32]byte) int {
	sum := 0
	for i := 0; i < len(phv); i++ {
		sum += int(phv[i])
	}
	return sum
}

//generateH generate H matrix using parameters
//generateH Cannot be sure rand is same with original implementation of C++
func generateH(parameters Parameters) [][]int {
	var H [][]int
	var hSeed int64
	var colOrder []int

	hSeed = int64(parameters.seed)
	k := parameters.m / parameters.wc

	H = make([][]int, parameters.m)
	for i := range H {
		H[i] = make([]int, parameters.n)
	}

	for i := 0; i < k; i++ {
		for j := i * parameters.wr; j < (i+1)*parameters.wr; j++ {
			H[i][j] = 1
		}
	}

	for i := 1; i < parameters.wc; i++ {
		colOrder = nil
		for j := 0; j < parameters.n; j++ {
			colOrder = append(colOrder, j)
		}

		src := rand.NewSource(hSeed)
		rnd := rand.New(src)
		rnd.Seed(hSeed)
		rnd.Shuffle(len(colOrder), func(i, j int) {
			colOrder[i],colOrder[j] = colOrder[j], colOrder[i]
		})
		hSeed--

		for j := 0; j < parameters.n; j++ {
			index := (colOrder[j]/parameters.wr + k*i)
			H[index][j] = 1
		}
	}

	return H
}

//generateQ generate colInRow and rowInCol matrix using H matrix
func generateQ(parameters Parameters, H [][]int) ([][]int, [][]int) {
	colInRow := make([][]int, parameters.wr)
	for i := 0; i < parameters.wr; i++ {
		colInRow[i] = make([]int, parameters.m)
	}

	rowInCol := make([][]int, parameters.wc)
	for i := 0; i < parameters.wc; i++ {
		rowInCol[i] = make([]int, parameters.n)
	}

	rowIndex := 0
	colIndex := 0

	for i := 0; i < parameters.m; i++ {
		for j := 0; j < parameters.n; j++ {
			if H[i][j] == 1 {
				colInRow[colIndex%parameters.wr][i] = j
				colIndex++

				rowInCol[rowIndex/parameters.n][j] = i
				rowIndex++
			}
		}
	}

	return colInRow, rowInCol
}

//generateHv generate hashvector
//It needs to compare with origin C++ implementation Especially when sha256 function is used
func generateHv(parameters Parameters, encryptedHeaderWithNonce []byte) []int {
	hashVector := make([]int, parameters.n)

	/*
		if parameters.n <= 256 {
			tmpHashVector = sha256.Sum256(headerWithNonce)
		} else {
			/*
				This section is for a case in which the size of a hash vector is larger than 256.
				This section will be implemented soon.
		}
			transform the constructed hexadecimal array into an binary array
			ex) FE01 => 11111110000 0001
	*/

	for i := 0; i < parameters.n/8; i++ {
		decimal := int(encryptedHeaderWithNonce[i])
		for j := 7; j >= 0; j-- {
			hashVector[j+8*(i)] = decimal % 2
			decimal /= 2
		}
	}

	//outputWord := hashVector[:parameters.n]
	return hashVector
}

