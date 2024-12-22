package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const PRUNE = 16777216
const F1 = 64
const DIV = 32
const F2 = 2048

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func nexttVal(seed int) int {
	seed = (seed ^ seed*F1) % PRUNE
	seed = (seed ^ seed/DIV) % PRUNE
	seed = (seed ^ seed*F2) % PRUNE
	return seed
}

func solve(seeds []int) int {
	res := 0
	for _, seed := range seeds {
		for i := 0; i < 2000; i++ {
			seed = nexttVal(seed)
		}
		res += seed
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	var seeds []int
	for _, line := range lines {
		seeds = append(seeds, atoi(line))
	}
	fmt.Println(solve(seeds))
}
