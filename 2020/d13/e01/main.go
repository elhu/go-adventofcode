package main

import (
	"fmt"
	"io/ioutil"
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

func divmod(a, b int) (int, int) {
	return a / b, a % b
}

func solve(start int, buses []int) int {
	minTime := start + buses[0]
	minBus := buses[0]
	minDelta := buses[0]
	for _, b := range buses {
		q, r := divmod(start, b)
		if r == 0 {
			return 0
		}
		if earliest := b * (q + 1); earliest < minTime {
			minTime = earliest
			minBus = b
			minDelta = earliest - start
		}
	}
	return minDelta * minBus
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	start := atoi(lines[0])
	buses := make([]int, 0)
	for _, t := range strings.Split(lines[1], ",") {
		if t != "x" {
			buses = append(buses, atoi(t))
		}
	}
	fmt.Println(solve(start, buses))
}
