package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	check(err)
	return n
}

func allZero(sequence []int) bool {
	for _, i := range sequence {
		if i != 0 {
			return false
		}
	}
	return true
}

func resolve(sequence []int) []int {
	sequences := [][]int{sequence}
	for !allZero(sequences[len(sequences)-1]) {
		var nextSeq []int
		prevSeq := sequences[len(sequences)-1]
		for i := 1; i < len(prevSeq); i++ {
			nextSeq = append(nextSeq, prevSeq[i]-prevSeq[i-1])
		}
		sequences = append(sequences, nextSeq)
	}
	sequences[len(sequences)-1] = append(sequences[len(sequences)-1], 0)
	for i := len(sequences) - 2; i >= 0; i-- {
		sequences[i] = append(sequences[i], sequences[i][len(sequences[i])-1]+sequences[i+1][len(sequences[i+1])-1])
	}
	return sequences[0]
}

func solve(sequences [][]int) int {
	var res int
	for _, seq := range sequences {
		res += resolve(seq)[len(seq)]
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	var sequences [][]int
	for _, l := range lines {
		var seq []int
		for _, p := range strings.Fields(l) {
			seq = append(seq, atoi(p))
		}
		slices.Reverse(seq)
		sequences = append(sequences, seq)
	}
	fmt.Println(solve(sequences))
}
