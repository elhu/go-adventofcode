package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func atoi(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return n
}

func digitCount(n int) int {
	if n == 0 {
		return 1
	}
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

func split(n int) (int, int) {
	nStr := strconv.Itoa(n)
	mid := len(nStr) / 2
	return atoi(nStr[:mid]), atoi(nStr[mid:])
}

func solve(stones []int) int {
	sm := make(map[int]int)
	for _, s := range stones {
		sm[s]++
	}
	for i := 0; i < BLINKS; i++ {
		nsm := make(map[int]int)
		for s, count := range sm {
			if s == 0 {
				nsm[1] += count
			} else if digitCount(s)%2 == 0 {
				a, b := split(s)
				nsm[a] += count
				nsm[b] += count
			} else {
				nsm[s*2024] += count
			}
		}
		sm = nsm
	}
	res := 0
	for _, count := range sm {
		res += count
	}
	return res
}

const BLINKS = 75

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	parts := strings.Fields(data)
	var stones []int
	for _, p := range parts {
		stones = append(stones, atoi(p))
	}
	fmt.Println(solve(stones))
}
