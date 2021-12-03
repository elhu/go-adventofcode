package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
)

func minMax(input []string, pos int) (byte, byte) {
	bits := make(map[byte]int)
	for _, s := range input {
		bits[s[pos]]++
	}
	if bits['0'] > bits['1'] {
		return '1', '0'
	}
	return '0', '1'
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func solve(input []string) int64 {
	gamma := make([]byte, len(input[0]))
	epsilon := make([]byte, len(input[0]))
	for i := range input[0] {
		min, max := minMax(input, i)
		gamma[i] = max
		epsilon[i] = min
	}
	g, err := strconv.ParseInt(string(gamma), 2, 64)
	check(err)
	e, err := strconv.ParseInt(string(epsilon), 2, 64)
	check(err)
	return g * e
}

func main() {
	lines := files.ReadLines(os.Args[1])
	fmt.Println(solve(lines))
}
