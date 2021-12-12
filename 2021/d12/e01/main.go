package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Name  string
	Edges []*Node
}

func parse(data []string) map[string]*Node {
	nodes := make(map[string]*Node)
	for _, l := range data {
		parts := strings.Split(l, "-")
		from, to := parts[0], parts[1]
		if _, exists := nodes[from]; !exists {
			nodes[from] = &Node{Name: from}
		}
		if _, exists := nodes[to]; !exists {
			nodes[to] = &Node{Name: to}
		}
		nodes[from].Edges = append(nodes[from].Edges, nodes[to])
		nodes[to].Edges = append(nodes[to].Edges, nodes[from])
	}
	return nodes
}

func unvisitedSmallCave(path []*Node, n *Node) bool {
	if strings.ToLower(n.Name) != n.Name {
		return true
	}
	for _, p := range path {
		if p == n {
			return false
		}
	}
	return true
}

func pathNames(path []*Node) []string {
	res := make([]string, len(path))
	for i, n := range path {
		res[i] = n.Name
	}
	return res
}

func solve(nodes map[string]*Node) int {
	start, end := nodes["start"], nodes["end"]
	queue := [][]*Node{{start}}
	var validPaths [][]*Node
	var path []*Node
	for len(queue) > 0 {
		path, queue = queue[0], queue[1:]
		node := path[len(path)-1]
		if node != end {
			for _, e := range node.Edges {
				if unvisitedSmallCave(path, e) {
					newPath := make([]*Node, len(path))
					copy(newPath, path)
					newPath = append(newPath, e)
					queue = append(queue, newPath)
				}
			}
		} else {
			validPaths = append(validPaths, path)
		}
	}
	return len(validPaths)
}

func main() {
	data := files.ReadLines(os.Args[1])
	nodes := parse(data)
	fmt.Println(solve(nodes))
}
