package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	check(err)
	return i
}

func solve(startNumbers []int) int {
	lastSeenAt := make(map[int]int)
	for i, n := range startNumbers {
		lastSeenAt[n] = i + 1
	}
	prevNumber := startNumbers[len(startNumbers)-1]
	for i := len(startNumbers) + 1; i <= 2020; i++ {
		var newNumber int
		if lastSeen, found := lastSeenAt[prevNumber]; found {
			delta := (i - 1) - lastSeen
			newNumber = delta
		} else {
			newNumber = 0
		}
		lastSeenAt[prevNumber] = i - 1
		prevNumber = newNumber
	}
	return prevNumber
}

func main() {
	data := os.Args[1]
	input := strings.Split(strings.TrimRight(string(data), "\n"), ",")
	var startNumbers []int
	for _, n := range input {
		startNumbers = append(startNumbers, atoi(n))
	}
	fmt.Println(solve(startNumbers))
}
