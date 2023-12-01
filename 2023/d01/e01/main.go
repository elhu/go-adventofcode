package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func solve(strings []string) int {
	res := 0
	for _, s := range strings {
		digits := make([]byte, 0)
		for _, c := range s {
			if c >= '0' && c <= '9' {
				digits = append(digits, byte(c))
			}
		}
		value, err := strconv.Atoi(string([]byte{digits[0], digits[len(digits)-1]}))
		check(err)
		res += value
	}

	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	strings := strings.Split(data, "\n")
	fmt.Println(solve(strings))
}
