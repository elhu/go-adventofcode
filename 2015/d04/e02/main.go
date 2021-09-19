package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"strings"
)

func hashFor(input string, i int) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", input, i))))
}

func solve(input string) int {
	for i := 1; ; i++ {
		if strings.HasPrefix(hashFor(input, i), "000000") {
			return i
		}
	}
}

func main() {
	input := os.Args[1]
	fmt.Println(solve(input))
}
