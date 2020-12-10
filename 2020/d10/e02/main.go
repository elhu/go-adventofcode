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
	sort.Slice(adapters, func(i, j int) bool { return adapters[i] < adapters[j] })
	adapters = append(adapters, []int{adapters[len(adapters)-1] + 3, adapters[len(adapters)-1] + 6}...)
	branches := make([]int, len(adapters))
	branches[len(branches)-1] = 1
	branches[len(branches)-2] = 1
	branches[len(branches)-3] = 1
	for i := len(adapters) - 4; i >= 0; i-- {
		for _, dist := range []int{1, 2, 3} {
			if adapters[i+dist]-adapters[i] <= 3 {
				branches[i] += branches[i+dist]
			}
		}
	}
	return branches[0]
}
func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	var adapters []int
	adapters = append(adapters, 0)
	maxJolt := 0
	for _, l := range lines {
		i := atoi(l)
		if i > maxJolt {
			maxJolt = i
		}
		adapters = append(adapters, i)
	}
	adapters = append(adapters, maxJolt+3)
	fmt.Println(solve(adapters))
}
