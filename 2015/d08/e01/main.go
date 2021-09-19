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

func memLen(str string) int {
	res := 0
	for i := 1; i < len(str)-1; i++ {
		if str[i] == '\\' {
			if str[i+1] == '\\' {
				i++
			} else if str[i+1] == 'x' {
				i += 3
			} else if str[i+1] == '"' {
				i++
			} else {
				fmt.Println(str[i:])
				panic("wtf")
			}
		}
		res++
	}
	return res
}

func solve(input []string) int {
	var codeLength, memLength int
	for _, l := range input {
		codeLength += len(l)
		memLength += memLen(l)
	}
	return codeLength - memLength
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	fmt.Println(solve(input))
}
