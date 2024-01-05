package main

import (
	"adventofcode/utils/files"
	"adventofcode/utils/graphs"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func graphSize(graph *graphs.Graph[string, string], start string) int {
	visited := make(map[string]bool)
	queue := []string{start}
	var curr string
	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		if !visited[curr] {
			visited[curr] = true
			for n := range graph.Edges[curr] {
				if !visited[n] {
					queue = append(queue, n)
				}
			}
		}
	}
	return len(visited)
}

func nodeRandomizer(graph *graphs.Graph[string, string]) func(string) string {
	var names []string
	for name := range graph.Vertices {
		names = append(names, name)
	}
	return func(exclude string) string {
		var n string
		for n = graph.Vertices[names[rand.Intn(len(names))]]; n == exclude; {
			n = graph.Vertices[names[rand.Intn(len(names))]]
		}
		return n
	}
}

// Arbitrary number, but it's enough to find the three bridges
const PATHS_TO_CHECK = 42

func key(left, right string) [2]string {
	if left < right {
		return [2]string{left, right}
	}
	return [2]string{right, left}
}

func topN(freqs map[[2]string]int, n int) [][2]string {
	edges := make([][2]string, 0, len(freqs))
	for k := range freqs {
		edges = append(edges, k)
	}
	sort.Slice(edges, func(i, j int) bool { return freqs[edges[i]] > freqs[edges[j]] })
	return edges[0:n]
}

func solve(graph *graphs.Graph[string, string]) int {
	getRandomNode := nodeRandomizer(graph)
	for {
		freqs := make(map[[2]string]int)
		for i := 0; i < PATHS_TO_CHECK; i++ {
			left := getRandomNode("")
			right := getRandomNode(left)
			path, error := graph.ShortestPath(left, right)
			check(error)
			for i := 1; i < len(path); i++ {
				freqs[key(path[i-1], path[i])]++
			}
		}
		candidates := topN(freqs, 3)
		for _, edge := range candidates {
			graph.RemoveEdge(edge[0], edge[1])
			graph.RemoveEdge(edge[1], edge[0])
		}
		if graphSize(graph, candidates[0][0]) != len(graph.Vertices) {
			// We found the three bridges
			return graphSize(graph, candidates[0][0]) * graphSize(graph, candidates[0][1])
		}
		// Add the edges back in, try again
		for _, edge := range candidates {
			graph.AddEdge(edge[0], edge[1])
			graph.AddEdge(edge[1], edge[0])
		}
	}
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")

	graph := graphs.NewGraph[string, string]()
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		name := parts[0]
		others := strings.Split(parts[1], " ")
		graph.AddVertex(name, name)
		for _, other := range others {
			graph.AddVertex(other, other)
			graph.AddEdge(name, other)
			graph.AddEdge(other, name)
		}
	}
	fmt.Println(solve(graph))
}
