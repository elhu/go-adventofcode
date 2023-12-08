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

func findStartNodes(nodes map[string]*Node) []*Node {
	sn := make([]*Node, 0)
	for k, n := range nodes {
		if k[len(k)-1] == 'A' {
			sn = append(sn, n)
		}
	}

	return sn
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func solve(nodes map[string]*Node, directions string) int {
	currNodes := findStartNodes(nodes)
	periods := make([]int, 0, len(currNodes))

	var i int
	for i = 0; ; i++ {
		for j, n := range currNodes {
			if directions[i%len(directions)] == 'R' {
				currNodes[j] = n.right
			} else {
				currNodes[j] = n.left
			}
			// Works because all start find a stop before the first one loops again
			if currNodes[j].name[2] == 'Z' {
				periods = append(periods, i+1)
			}
		}
		if len(periods) == len(currNodes) {
			break
		}
	}
	return LCM(periods[0], periods[1], periods...)
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
