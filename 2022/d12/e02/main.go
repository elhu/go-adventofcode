package main

import (
	"adventofcode/utils/files"
	"adventofcode/utils/sets/stringset"
	"fmt"
	"os"
)

type Node struct {
	id         string
	height     int
	neighbours []*Node
}

func parseGraph(data []string) ([]*Node, *Node) {
	nodes := make(map[string]*Node)
	var end *Node
	starts := make([]*Node, 0)
	for i, line := range data {
		for j, char := range line {
			node := &Node{id: fmt.Sprintf("%d:%d", i, j)}
			height := int(char - 'a')
			if char == 'S' {
				height = 0
			} else if char == 'E' {
				height = 25
				end = node
			}
			node.height = height
			if node.height == 0 {
				starts = append(starts, node)
			}
			nodes[node.id] = node
		}
	}
	for i := range data {
		for j := range data[i] {
			node := nodes[fmt.Sprintf("%d:%d", i, j)]
			for _, nID := range []string{
				fmt.Sprintf("%d:%d", i+1, j),
				fmt.Sprintf("%d:%d", i-1, j),
				fmt.Sprintf("%d:%d", i, j-1),
				fmt.Sprintf("%d:%d", i, j+1),
			} {
				if neighbour, exists := nodes[nID]; exists && neighbour.height <= node.height+1 {
					node.neighbours = append(node.neighbours, neighbour)
				}
			}
		}
	}
	return starts, end
}

func bfs(start, end *Node) (int, bool) {
	queue := [][]*Node{{start}}
	visited := stringset.New()
	visited.Add(start.id)
	var headPath []*Node
	for len(queue) > 0 {
		headPath, queue = queue[0], queue[1:]
		head := headPath[len(headPath)-1]
		for _, n := range head.neighbours {
			if n == end {
				return len(headPath), true
			}
			if !visited.HasMember(n.id) {
				visited.Add(n.id)
				newPath := make([]*Node, len(headPath))
				copy(newPath, headPath)
				queue = append(queue, append(newPath, n))
			}
		}
	}
	return 0, false
}

func main() {
	data := files.ReadLines(os.Args[1])
	starts, end := parseGraph(data)
	min := 999999999999
	for _, start := range starts {
		distance, success := bfs(start, end)
		if success && distance < min {
			min = distance
		}
	}
	fmt.Println(min)
}
