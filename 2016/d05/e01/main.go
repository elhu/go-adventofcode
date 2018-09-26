package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"os"
)

func hashFor(input string, i int) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", input, i))))
}

func solve(input string) []byte {
	res := make([]byte, 8)

	for found, i := 0, 0; found < 8; i++ {
		h := hashFor(input, i)
		if bytes.Compare([]byte(h[0:5]), []byte{'0', '0', '0', '0', '0'}) == 0 {
			res[found] = h[5]
			found++
		}
	}
	return res
}

func main() {
	input := os.Args[1]
	fmt.Println(string(solve(input)))
}
