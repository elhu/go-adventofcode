package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func solve(input string) int {
	var res []byte
	for r := 0; r < 50; r++ {
		curr := input[0]
		res = nil
		currCount := 1
		for i := 1; i < len(input); i++ {
			if input[i] == curr {
				currCount++
			} else {
				res = append(res, []byte(fmt.Sprintf("%d%c", currCount, curr))...)
				curr = input[i]
				currCount = 1
			}
		}
		res = append(res, []byte(fmt.Sprintf("%d%c", currCount, curr))...)
		input = string(res)
	}
	return len(res)
}

func main() {
	fmt.Println(solve(os.Args[1]))
}
