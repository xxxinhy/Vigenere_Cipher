package main

import (
	"fmt"
	"math"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func clean(file string) []string {
	dat, err := os.ReadFile(file)
	check(err)
	var c []string
	for _, ch := range dat {
		if ch >= 65 && ch <= 90 {
			c = append(c, string(ch))
		}
	}
	return c
}
func group(num int, c []string) float64 {

	N := float64(len(c)) / float64(num)
	var avg float64
	for i := 1; i <= num; i++ {
		freq := make(map[string]int)
		for j := i - 1; j < len(c); j = j + num {
			freq[c[j]] = freq[c[j]] + 1
		}
		ic := 0.0
		for _, cnt := range freq {
			ic = ic + float64(cnt)*(float64(cnt)-1.0)/(N*(N-1.0))
		}
		avg = avg + ic
		clear(freq)
	}
	return avg / float64(num)
}
func keylength(c []string) int {
	min := 1.0
	keylength := 1
	for i := 2; i <= 20; i++ {
		res := group(i, c)
		if math.Abs(res-0.068) < min {
			if keylength > 2 && i%keylength == 0 && math.Abs(math.Abs(res-0.068)-min) < 0.001 {
				continue
			}
			keylength = i
			min = math.Abs(group(i, c) - 0.068)
		}

		//fmt.Printf("Length:%d:%f\n", i, group(i, c))
	}
	return keylength

}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Input example: vigenere-keylength <ciphertext filename>")
	}
	file := os.Args[1]
	c := clean(file)
	fmt.Println(keylength(c))

}
