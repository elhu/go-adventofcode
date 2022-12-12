package main

import (
	"adventofcode/utils/files"
	"adventofcode/utils/sets/stringset"
	"fmt"
	"os"
)

type Node struct {
	id         string
	height     byte
	neighbours []*Node
}

func parseGraph(data []string) (*Node, *Node) {
	var start, end *Node
	nodes := make(map[string]*Node)
	for i, line := range data {
		for j, cell := range line {
			newNode := &Node{id: fmt.Sprintf("%d:%d", i, j)}
			if cell == 'S' {
				cell = 'a'
				start = newNode
			} else if cell == 'E' {
				cell = 'z'
				end = newNode
			}
			newNode.height = byte(cell)
			nodes[newNode.id] = newNode
		}
	}
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			self := nodes[fmt.Sprintf("%d:%d", i, j)]
			candidates := []string{
				fmt.Sprintf("%d:%d", i-1, j),
				fmt.Sprintf("%d:%d", i+1, j),
				fmt.Sprintf("%d:%d", i, j-1),
				fmt.Sprintf("%d:%d", i, j+1),
			}
			for _, cid := range candidates {
				candidate, exists := nodes[cid]
				if exists && candidate.height <= self.height+1 {
					self.neighbours = append(self.neighbours, candidate)
				}
			}
		}
	}
	return start, end
}

func bfs(start, end *Node) int {
	rounds := 0
	queue, newQueue := []*Node{start}, make([]*Node, 0)
	visited := stringset.New()
	var head *Node
	for len(queue) > 0 {
		fmt.Println(len(queue), len(newQueue))
		head, queue = queue[0], queue[1:]
		visited.Add(head.id)
		if head == end {
			return rounds
		}
		for _, n := range head.neighbours {
			if !visited.HasMember(n.id) {
				fmt.Printf("Adding candidate step from %c to %c\n", head.height, n.height)
				newQueue = append(newQueue, n)
			}
		}
		if len(queue) == 0 {
			queue = newQueue
			newQueue = newQueue[:0]
			rounds++
		}
	}
	panic("Didn't find end node")
}

func main() {
	data := files.ReadLines(os.Args[1])
	start, end := parseGraph(data)
	fmt.Println(start)
	fmt.Println(end)
	fmt.Println(bfs(start, end))
}
