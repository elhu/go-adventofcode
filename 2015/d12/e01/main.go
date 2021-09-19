package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func atoi(str string) int {
	n, err := strconv.Atoi(str)
	check(err)
	return n
}

func solve(input []byte) int {
	res := 0
	for i := 0; i < len(input); i++ {
		if input[i] >= '0' && input[i] <= '9' || input[i] == '-' {
			j := i + 1
			for j < len(input) && input[j] >= '0' && input[j] <= '9' {
				j++
			}
			res += atoi(string(input[i:j]))
			i = j
		}
	}

	return res
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	fmt.Println(solve(data))
}
