package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	name  string
	left  *Node
	right *Node
}

func solve(nodes map[string]*Node, directions string) int {
	currNode := nodes["AAA"]
	var i int
	for i = 0; ; i++ {
		if currNode.name == "ZZZ" {
			return i
		}

		if directions[i%len(directions)] == 'R' {
			currNode = currNode.right
		} else {
			currNode = currNode.left
		}
	}
}

func fetchOrCreateNode(nodes map[string]*Node, name string) *Node {
	if node, found := nodes[name]; found {
		return node
	}
	nodes[name] = &Node{name: name}
	return nodes[name]
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	parts := strings.Split(data, "\n\n")
	directions := parts[0]
	nodes := make(map[string]*Node)
	for _, line := range strings.Split(parts[1], "\n") {
		var name, lft, rgt string
		fmt.Sscanf(line, "%s = (%s %s)", &name, &lft, &rgt)
		lft = strings.TrimRight(lft, ",)")
		rgt = strings.TrimRight(rgt, ",)")
		node := fetchOrCreateNode(nodes, name)
		node.left = fetchOrCreateNode(nodes, lft)
		node.right = fetchOrCreateNode(nodes, rgt)
	}
	fmt.Println(solve(nodes, directions))
}
