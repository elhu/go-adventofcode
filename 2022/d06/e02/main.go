package main

import (
	"adventofcode/utils/files"
	"adventofcode/utils/sets/byteset"
	"bytes"
	"fmt"
	"os"
)

func allDifferent(data []byte) bool {
	return byteset.NewFromSlice(data).Len() == len(data)
}

func solve(data []byte) int {
	for i := 0; i < len(data)-13; i++ {
		if allDifferent(data[i : i+14]) {
			return i + 14
		}
	}
	panic("wtf")
}

func main() {
	data := files.ReadFile(os.Args[1])
	data = bytes.TrimRight(data, "\n")
	fmt.Println(solve(data))
}
