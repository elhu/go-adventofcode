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

func encodedLen(str string) int {
	res := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '"' {
			res++
		} else if str[i] == '\\' {
			res++
		}
		res++
	}
	return res + 2 // +2 for encasing double quotes
}

func solve(input []string) int {
	var codeLength, encodedLength int
	for _, l := range input {
		codeLength += len(l)
		encodedLength += encodedLen(l)
	}
	return encodedLength - codeLength
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	fmt.Println(solve(input))
}
