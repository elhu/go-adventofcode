package main

import (
	"bytes"
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

func sum(numbers []int) int {
	s := 0
	for _, i := range numbers {
		s += i
	}
	return s
}

type node struct {
	childCount, metadataCount int
	children                  []node
	metadata                  []int
}

func processNode(n node, data []int, pos int) (node, int) {
	for i := 0; i < n.childCount; i++ {
		child := node{}
		child.childCount = data[pos]
		child.metadataCount = data[pos+1]
		pos += 2
		children, newPos := processNode(child, data, pos)
		n.children = append(n.children, children)
		pos = newPos
	}
	n.metadata = append(n.metadata, data[pos:pos+n.metadataCount]...)
	pos += n.metadataCount
	return n, pos
}

func solve(n node) int {
	if n.childCount == 0 {
		return sum(n.metadata)
	}
	res := 0
	for _, mIdx := range n.metadata {
		if mIdx > 0 && mIdx <= len(n.children) {
			res += solve(n.children[mIdx-1])
		}
	}
	return res
}

func main() {
	input, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input = bytes.TrimSuffix(input, []byte{'\n'})
	numbers := bytes.Split(input, []byte{' '})
	rawTree := make([]int, 0, len(numbers))
	for _, n := range numbers {
		i, _ := strconv.Atoi(string(n))
		rawTree = append(rawTree, i)
	}
	node := node{childCount: rawTree[0], metadataCount: rawTree[1]}
	node, _ = processNode(node, rawTree, 2)
	fmt.Println(solve(node))
}
