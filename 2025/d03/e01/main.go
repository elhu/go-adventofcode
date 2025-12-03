package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func processBattery(b string) string {
	maxA := b[0]
	maxPos := 0
	for i := 1; i < len(b); i++ {
		if b[i] > maxA {
			maxA = b[i]
			maxPos = i
		}
	}
	if maxPos == len(b)-1 {
		return reverse(processBattery(reverse(b)))
	}
	maxB := b[maxPos+1]
	for j := maxPos + 1; j < len(b); j++ {
		if b[j] > maxB {
			maxB = b[j]
		}
	}
	return fmt.Sprintf("%c%c", maxA, maxB)
}

func solve(batteries []string) int {
	res := 0
	for _, b := range batteries {
		n, err := strconv.Atoi(processBattery(b))
		if err != nil {
			panic(err)
		}
		res += n
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	batteries := strings.Split(data, "\n")

	fmt.Println(solve(batteries))
}
