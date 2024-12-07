package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

type Node struct {
	value         int
	n             int
	add, mul, con *Node
}

func concat(a, b int) int {
	return atoi(fmt.Sprintf("%d%d", a, b))
}

func truthable(numbers []int, goal int) bool {
	if len(numbers) == 1 {
		return numbers[0] == goal
	}

	var leaves []*Node
	root := &Node{value: numbers[0], n: numbers[0]}
	leaves = append(leaves, root)
	for i := 1; i < len(numbers); i++ {
		newLeaves := []*Node{}
		for _, leaf := range leaves {
			add := &Node{value: leaf.value + numbers[i], n: numbers[i]}
			mul := &Node{value: leaf.value * numbers[i], n: numbers[i]}
			con := &Node{value: concat(leaf.value, numbers[i]), n: numbers[i]}
			leaf.add = add
			leaf.mul = mul
			leaf.con = con
			newLeaves = append(newLeaves, add, mul, con)
		}
		leaves = newLeaves
	}
	for _, leaf := range leaves {
		if leaf.value == goal {
			return true
		}
	}
	return false
}

func solve(lines []string) int {
	res := 0
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		goal := atoi(parts[0])
		data := strings.Split(parts[1], " ")
		var numbers []int
		for _, number := range data {
			numbers = append(numbers, atoi(number))
		}
		if truthable(numbers, goal) {
			res += goal
		}
	}

	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	fmt.Println(solve(lines))
}
