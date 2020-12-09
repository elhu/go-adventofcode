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

func buildSums(values []int) []int {
	sums := make([]int, len(values))
	sums[0] = values[0]
	for i, v := range values[1:] {
		sums[i+1] = sums[i] + v
	}
	return sums
}

func getRangeMinMax(values []int) (int, int) {
	max := values[0]
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func solve(values []int) int {
	sums := buildSums(values)
	discarded := 0
	for i := 0; i < len(values); i++ {
		for j, v := range sums[i:] {
			if v-discarded == targetNumber {
				min, max := getRangeMinMax(values[i : i+j+1])
				return min + max
			}
		}
		discarded += values[i]
	}
	return -42
}

const targetNumber = 88311122

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	var values []int
	for _, l := range lines {
		values = append(values, atoi(l))
	}
	fmt.Println(solve(values))
}
