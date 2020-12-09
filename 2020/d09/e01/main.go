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

func isValid(val int, preamble []int) bool {
	for i := 0; i < len(preamble); i++ {
		for j := i + 1; j < len(preamble); j++ {
			if preamble[i] != preamble[j] && preamble[i]+preamble[j] == val {
				return true
			}
		}
	}
	return false
}

func solve(values []int) int {
	i := 0
	var preamble []int
	for ; i < preambleSize; i++ {
		preamble = append(preamble, values[i])
	}
	for _, val := range values[preambleSize:] {
		if !isValid(val, preamble) {
			return val
		}
		preamble = append(preamble[1:], val)
	}
	return -42
}

const preambleSize = 25

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
