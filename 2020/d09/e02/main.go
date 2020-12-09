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
	left, right := 0, 1
	currSum := values[0] + values[1]
	for left < len(values) && right < len(values) {
		if currSum == targetNumber {
			min, max := getRangeMinMax(values[left : right+1])
			return min + max
		} else if currSum < targetNumber {
			right++
			currSum += values[right]
		} else {
			currSum -= values[left]
			left++
		}
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
