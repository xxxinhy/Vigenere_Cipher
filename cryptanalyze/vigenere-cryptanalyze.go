package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var letterFrequency = map[string]float64{
	"A": 0.08167,
	"B": 0.01492,
	"C": 0.02782,
	"D": 0.04253,
	"E": 0.12702,
	"F": 0.02228,
	"G": 0.02015,
	"H": 0.06094,
	"I": 0.06966,
	"J": 0.00153,
	"K": 0.00772,
	"L": 0.04025,
	"M": 0.02406,
	"N": 0.06749,
	"O": 0.07507,
	"P": 0.01929,
	"Q": 0.00095,
	"R": 0.05987,
	"S": 0.06327,
	"T": 0.09056,
	"U": 0.02758,
	"V": 0.00978,
	"W": 0.02360,
	"X": 0.00150,
	"Y": 0.01974,
	"Z": 0.00074,
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func clean(dat []byte) []string {
	var c []string
	for _, ch := range dat {
		if ch >= 65 && ch <= 90 {
			c = append(c, string(ch))
		}
	}
	return c
}
func getCipher(c []string, n int, keylength int) []string {
	var getC []string
	for i := n - 1; i < len(c); i += keylength {
		getC = append(getC, c[i])
	}
	return getC
}
func findKey(c []string, keylength int) []string {

	var key []string
	for k := 1; k <= keylength; k++ {
		getC := getCipher(c, k, keylength)

		sum := make(map[string]float64)
		freq := make(map[string]float64)
		var letterF []float64
		var cipherF []float64
		for i := 0; i < len(getC); i++ {
			sum[getC[i]] = sum[getC[i]] + 1
			freq[getC[i]] = sum[getC[i]] / float64(len(getC))
		}
		var str []string
		for i := 'A'; i <= 'Z'; i++ {
			str = append(str, string(i))
			cipherF = append(cipherF, freq[string(i)])
			letterF = append(letterF, letterFrequency[string(i)])

		}

		cipherLength := len(cipherF)

		shift := 0
		maxProduct := 0.0
		for i := 0; i < cipherLength; i++ {
			product := 0.0
			for j := 0; j < cipherLength; j++ {
				product = product + cipherF[(j+i+cipherLength)%cipherLength]*letterF[j]
			}
			if maxProduct <= product {
				maxProduct = product
				shift = i
			}
		}
		key = append(key, string(rune(shift+65)))

	}
	return key
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Input example: vigenere-cryptanalyze <ciphertext filename> <key length>")
	}
	file := os.Args[1]
	keyL := os.Args[2]
	keylength, err := strconv.Atoi(keyL)
	if err != nil {
		panic(err)
	}

	dat, err := os.ReadFile(file)
	check(err)
	c := clean(dat)

	key := findKey(c, keylength)

	keystr := strings.Join(key, "")
	fmt.Println(keystr)

	//var out []string
	// for i := 0; i < len(dat); i++ {
	// 	out = append(out, string(rune((int(dat[i])-int(keystr[(len(out)%len(keystr))])+26)%26+65)))
	// }
	// fmt.Println(strings.Join(out, ""))

}
