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
	res, err := strconv.Atoi(str)
	check(err)
	return res
}

const maxIP = 4294967295

// const maxIP = 10

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	blocks := make([]int, maxIP+2)
	for _, line := range input {
		var left, right int
		fmt.Sscanf(line, "%d-%d", &left, &right)
		blocks[left]++
		blocks[right+1]--
	}
	blockCounter := 0
	for i, b := range blocks {
		blockCounter += b
		if blockCounter == 0 {
			fmt.Println(i)
			return
		}
	}
}
