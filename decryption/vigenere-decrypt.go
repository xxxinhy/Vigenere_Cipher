package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Input example: vigenere-decrypt <decipherment key> <ciphertext filename>")
	}
	key := os.Args[1]
	file := os.Args[2]
	var out []string
	dat, err := os.ReadFile(file)
	check(err)

	for i := 0; i < len(dat); i++ {
		if !(dat[i] >= 65 && dat[i] <= 90) {
			if dat[i] >= 97 && dat[i] <= 122 {
				dat[i] = dat[i] - 32
			} else {
				continue
			}
		}

		out = append(out, string(rune((int(dat[i])-int(key[(len(out)%len(key))])+26)%26+65)))

	}
	fmt.Print(strings.Join(out, ""))

}
