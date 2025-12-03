package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processBattery(b []byte) int {
	fmt.Println(string(b))
	toRemove := len(b) - 12
	var trimmed []byte
	for _, c := range b {
		for toRemove > 0 && len(trimmed) > 0 && trimmed[len(trimmed)-1] < c {
			trimmed = trimmed[:len(trimmed)-1]
			toRemove--
		}
		trimmed = append(trimmed, c)
	}
	returnVal, err := strconv.Atoi(string(trimmed[:12]))
	if err != nil {
		panic(err)
	}
	return returnVal
}

func solve(batteries []string) int {
	res := 0
	for _, b := range batteries {
		res += processBattery([]byte(b))
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	batteries := strings.Split(data, "\n")

	fmt.Println(solve(batteries))
}
