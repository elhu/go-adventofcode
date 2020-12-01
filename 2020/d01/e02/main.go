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

func solve(numbers []int) int {
	for i := 0; i < len(numbers); i++ {
		for j := 1; j < len(numbers); j++ {
			for k := 2; k < len(numbers); k++ {
				if numbers[i]+numbers[j]+numbers[k] == 2020 {
					return numbers[i] * numbers[j] * numbers[k]
				}
			}
		}
	}
	return -1
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	var numbers []int
	for _, l := range lines {
		i, err := strconv.Atoi(l)
		check(err)
		numbers = append(numbers, i)
	}
	fmt.Println(solve(numbers))
}
