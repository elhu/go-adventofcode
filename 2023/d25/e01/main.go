package main

import (
	"adventofcode/utils/files"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"sort"
	"strings"
)

type Node struct {
	name        string
	id          int
	connectedTo []*Node
}

func getOrAddNode(nodes map[string]*Node, name string) *Node {
	if node, ok := nodes[name]; ok {
		return node
	}
	node := Node{name: name, id: len(nodes)}
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

func nodeRandomizer(nodes map[string]*Node) func(*Node) *Node {
	var names []string
	for name := range nodes {
		names = append(names, name)
	}
	return func(exclude *Node) *Node {
		var n *Node
		for n = nodes[names[rand.Intn(len(names))]]; n == exclude; {
			n = nodes[names[rand.Intn(len(names))]]
		}
		return n
	}
}

// Arbitrary number, but it's enough to find the three bridges
const PATHS_TO_CHECK = 42

func key(left, right *Node) [2]*Node {
	if left.name < right.name {
		return [2]*Node{left, right}
	}
	return [2]*Node{right, left}
}

type State struct {
	currNode *Node
	path     []*Node
}

func bfs(start, end *Node) []*Node {
	toVisit := []State{{currNode: start, path: []*Node{}}}
	visited := make(map[*Node]struct{})
	head := toVisit[0]
	for len(toVisit) > 0 {
		toVisit, head = toVisit[1:], toVisit[0]
		visited[head.currNode] = struct{}{}
		if head.currNode == end {
			return head.path
		}
		for _, n := range head.currNode.connectedTo {
			if _, found := visited[n]; !found {
				toVisit = append(toVisit, State{currNode: n, path: append(head.path, n)})
			}
		}
	}
	return []*Node{}
}

func topN(freqs map[[2]*Node]int, n int) [][2]*Node {
	edges := make([][2]*Node, 0, len(freqs))
	for k := range freqs {
		edges = append(edges, k)
	}
	sort.Slice(edges, func(i, j int) bool { return freqs[edges[i]] > freqs[edges[j]] })
	return edges[0:n]
}

func solve(nodes map[string]*Node) int {
	getRandomNode := nodeRandomizer(nodes)
	for {
		freqs := make(map[[2]*Node]int)
		for i := 0; i < PATHS_TO_CHECK; i++ {
			left := getRandomNode(nil)
			right := getRandomNode(left)

			path := bfs(left, right)
			for i := 1; i < len(path); i++ {
				freqs[key(path[i-1], path[i])]++
			}
		}
		candidates := topN(freqs, 3)
		for _, edge := range candidates {
			edge[0].connectedTo = slices.DeleteFunc(edge[0].connectedTo, func(e *Node) bool { return e == edge[1] })
			edge[1].connectedTo = slices.DeleteFunc(edge[1].connectedTo, func(e *Node) bool { return e == edge[0] })
		}
		if graphSize(nodes, candidates[0][0].name) != len(nodes) {
			// We found the three bridges
			return graphSize(nodes, candidates[0][0].name) * graphSize(nodes, candidates[0][1].name)
		}
		// Add the edges back in, try again
		for _, edge := range candidates {
			edge[0].connectedTo = append(edge[0].connectedTo, edge[1])
			edge[1].connectedTo = append(edge[1].connectedTo, edge[0])
		}
	}
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
	fmt.Println(solve(nodes))
}
