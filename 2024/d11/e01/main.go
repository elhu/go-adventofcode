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
	var newStones []int
	for i := 0; i < BLINKS; i++ {
		for _, s := range stones {
			if s == 0 {
				newStones = append(newStones, 1)
			} else if digitCount(s)%2 == 0 {
				a, b := split(s)
				newStones = append(newStones, a, b)
			} else {
				newStones = append(newStones, s*2024)
			}
		}
		stones = newStones
		newStones = []int{}
	}
	return len(stones)
}

const BLINKS = 25

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	parts := strings.Fields(data)
	var stones []int
	for _, p := range parts {
		stones = append(stones, atoi(p))
	}
	fmt.Println(stones)
	fmt.Println(solve(stones))
}
