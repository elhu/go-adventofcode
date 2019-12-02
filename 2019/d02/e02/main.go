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

func solve(opcodes []int) {
	pos := 0
	for opcodes[pos] != 99 {
		a, b, newPos := opcodes[pos+1], opcodes[pos+2], opcodes[pos+3]
		if opcodes[pos] == 1 {
			opcodes[newPos] = opcodes[a] + opcodes[b]
		} else if opcodes[pos] == 2 {
			opcodes[newPos] = opcodes[a] * opcodes[b]
		} else {
			check(fmt.Errorf("found unexpected opcode %d at position %d", opcodes[pos], pos))
		}
		pos += 4
	}
	// return opcodes
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	opcodesStr := strings.Split(strings.TrimRight(string(data), "\n"), ",")
	opcodes := make([]int, 0, len(opcodesStr))
	for _, s := range opcodesStr {
		i, err := strconv.Atoi(s)
		check(err)
		opcodes = append(opcodes, i)
	}
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			dupCodes := append(make([]int, 0, len(opcodes)), opcodes...)
			// fmt.Println(dupCodes)
			dupCodes[1] = i
			dupCodes[2] = j
			solve(dupCodes)
			if dupCodes[0] == 19690720 {
				fmt.Println(100*i + j)
			}
		}
	}
}
