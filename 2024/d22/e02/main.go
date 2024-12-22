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

func genSequences() [][4]int {
	var res [][4]int
	for i := -9; i <= 9; i++ {
		for j := -9; j <= 9; j++ {
			for k := -9; k <= 9; k++ {
				for l := -9; l <= 9; l++ {
					res = append(res, [4]int{i, j, k, l})
				}
			}
		}
	}
	return res
}

func genDeltas(seeds []int, seqMap map[[4]int]map[int]int) {
	for _, seed := range seeds {
		initialVal := seed
		var tmp []int
		for i := 0; i < 2000; i++ {
			next := nexttVal(seed)
			tmp = append(tmp, next%10-seed%10)
			seed = next
			if len(tmp) >= 4 {
				k := [4]int{tmp[len(tmp)-4], tmp[len(tmp)-3], tmp[len(tmp)-2], tmp[len(tmp)-1]}
				if _, found := seqMap[k][initialVal]; !found {
					seqMap[k][initialVal] = next % 10
				}
			}
		}
	}
}

func solve(seeds []int) int {
	sequences := genSequences()
	seqMap := make(map[[4]int]map[int]int)
	for _, sequence := range sequences {
		seqMap[sequence] = make(map[int]int)
	}
	genDeltas(seeds, seqMap)
	max := 0
	for _, data := range seqMap {
		s := 0
		for _, val := range data {
			s += val
		}
		if s > max {
			max = s
		}
	}
	return max
}

// 1892017334 too high
func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	var seeds []int
	for _, line := range lines {
		seeds = append(seeds, atoi(line))
	}
	fmt.Println(solve(seeds))
}
