package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func solve(data []byte) int {
	floor := 0
	for _, c := range data {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}
	}
	return floor
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	fmt.Println(solve(data))
}
