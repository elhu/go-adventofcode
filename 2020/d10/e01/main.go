package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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

func solve(adapters []int) int {
	diffs := make(map[int]int)
	sort.Slice(adapters, func(i, j int) bool { return adapters[i] < adapters[j] })
	for i := 1; i < len(adapters); i++ {
		diffs[adapters[i]-adapters[i-1]]++
	}
	// For device adapter
	diffs[3]++
	fmt.Println(diffs)
	return diffs[3] * diffs[1]
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	var adapters []int
	adapters = append(adapters, 0)
	for _, l := range lines {
		adapters = append(adapters, atoi(l))
	}
	fmt.Println(solve(adapters))
}
