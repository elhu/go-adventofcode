package main

import (
	"adventofcode/utils/files"
	"adventofcode/utils/sets/string_set"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func mostFreq(input []string, pos int) byte {
	bits := make(map[byte]int)
	for _, s := range input {
		bits[s[pos]]++
	}
	if bits['0'] > bits['1'] {
		return '0'
	}
	return '1'
}

func reduce(input []string, cmpFn func(byte, byte) bool) int64 {
	set := string_set.New()
	for _, l := range input {
		set.Add(l)
	}
	for i := range input[0] {
		mf := mostFreq(set.Members(), i)
		for _, l := range input {
			if cmpFn(l[i], mf) {
				set.Remove(l)
				if set.Len() == 1 {
					n, err := strconv.ParseInt(string(set.Members()[0]), 2, 64)
					check(err)
					return n
				}
			}
		}
	}
	panic("WTF")
}

func solve(input []string) int64 {
	oxy := reduce(input, func(a, b byte) bool { return a == b })
	carb := reduce(input, func(a, b byte) bool { return a != b })
	return oxy * carb
}

func main() {
	lines := files.ReadLines(os.Args[1])
	fmt.Println(solve(lines))
}
