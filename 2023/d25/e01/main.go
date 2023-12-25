package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Node struct {
	name        string
	connectedTo []*Node
}

func getOrAddNode(nodes map[string]*Node, name string) *Node {
	if node, ok := nodes[name]; ok {
		return node
	}
	node := Node{name: name}
	nodes[name] = &node
	return &node
}

func visualize(nodes map[string]*Node) {
	fmt.Println("digraph {")
	for _, node := range nodes {
		var nn []string
		for _, n := range node.connectedTo {
			nn = append(nn, n.name)
		}
		fmt.Printf("\t%s -> {%s}\n", node.name, strings.Join(nn, " "))
	}
	fmt.Println("}")
}

func graphSize(nodes map[string]*Node, start string) int {
	visited := make(map[string]bool)
	queue := []string{start}
	var curr string
	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		if !visited[curr] {
			visited[curr] = true
			for _, n := range nodes[curr].connectedTo {
				if !visited[n.name] {
					queue = append(queue, n.name)
				}
			}
		}
	}
	return len(visited)
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	nodes := make(map[string]*Node)
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		name := parts[0]
		others := strings.Split(parts[1], " ")
		left := getOrAddNode(nodes, name)
		for _, other := range others {
			right := getOrAddNode(nodes, other)
			left.connectedTo = append(left.connectedTo, right)
			right.connectedTo = append(right.connectedTo, left)
		}
	}
	fmt.Println("Copy the following into a graph.dot file and run `neato -Tpng -o graph.png graph.dot`")
	visualize(nodes)
	fmt.Println("Next, find the three bridges, and replace the values on the next line")
	for left, right := range map[string]string{"dqf": "cbx", "xft": "pzv", "sds": "hbr"} {
		nodes[left].connectedTo = slices.DeleteFunc(nodes[left].connectedTo, func(e *Node) bool { return e.name == right })
		nodes[right].connectedTo = slices.DeleteFunc(nodes[right].connectedTo, func(e *Node) bool { return e.name == left })
	}
	fmt.Println("Finally, change the values on the next line to both sides of one of the bridges")
	fmt.Println(graphSize(nodes, "dqf") * graphSize(nodes, "cbx"))
}
