package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isNice(str string) bool {
	pairs := make(map[string][]int)
	repeatLetters := false
	for i := 1; i < len(str); i++ {
		pair := str[i-1 : i+1]
		pairs[pair] = append(pairs[pair], i-1)
		if i > 2 && str[i] == str[i-2] {
			repeatLetters = true
		}
	}
	repeatPairs := false
	for _, positions := range pairs {
		if len(positions) > 2 {
			repeatPairs = true
		}
		if len(positions) == 2 {
			repeatPairs = positions[1] > positions[0]+1
		}
	}
	return repeatLetters && repeatPairs
}

func solve(input []string) int {
	nice := 0
	for _, l := range input {
		if isNice(l) {
			nice++
		}
	}
	return nice
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	fmt.Println(solve(input))
}
