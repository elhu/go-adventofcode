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
	opcodes[1] = 12
	opcodes[2] = 2
	solve(opcodes)
	fmt.Println(opcodes[0])
}
