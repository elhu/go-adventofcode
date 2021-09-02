package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Coord struct {
	x, y int
}

type Vertex struct {
	distance int
	target   *Node
}

type Node struct {
	name   byte
	coords Coord
	edges  map[byte]Vertex
}

func copyMap(input [][]byte) [][]byte {
	res := make([][]byte, len(input))
	for i := range input {
		res[i] = make([]byte, len(input[i]))
		copy(res[i], input[i])
	}
	return res
}

func neighbors(loc Coord, ducts [][]byte) []Coord {
	var res []Coord
	for _, offset := range []Coord{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		target := Coord{loc.x + offset.x, loc.y + offset.y}
		if ducts[target.y][target.x] != '#' {
			res = append(res, target)
		}
	}
	return res
}

func buildVertices(from *Node, nodes map[byte]*Node, input [][]byte) {
	ducts := copyMap(input)
	queue := []Coord{from.coords}
	ducts[from.coords.y][from.coords.x] = '#'
	var nextRound []Coord
	distance := 0

	for len(queue) > 0 {
		loc := queue[0]
		queue = queue[1:]

		for _, n := range neighbors(loc, ducts) {
			if bytes.ContainsAny([]byte{input[n.y][n.x]}, "1234567890") {
				if _, f := from.edges[input[n.y][n.x]]; !f {
					from.edges[input[n.y][n.x]] = Vertex{distance + 1, nodes[input[n.y][n.x]]}
				}
			}
			nextRound = append(nextRound, n)
			ducts[n.y][n.x] = '#'
		}
		if len(queue) == 0 {
			distance++
			queue = append(queue, nextRound...)
			nextRound = nil
		}
	}
}

func printNodes(nodes map[byte]*Node) {
	for _, n := range nodes {
		fmt.Printf("Node %c:", n.name)
		for n, e := range n.edges {
			fmt.Printf("  %c (%d)", n, e.distance)
		}
		fmt.Println("")
	}
}

func solve(nodes map[byte]*Node) int {
	visited := make(map[byte]struct{})
	curr := nodes['0']
	visited['0'] = struct{}{}
	res := 0
	for len(visited) != len(nodes) {
		minDist := 999999999
		var closest byte
		for c, e := range curr.edges {
			if _, f := visited[c]; !f && e.distance < minDist {
				closest = c
				minDist = e.distance
			}
		}
		curr = nodes[closest]
		visited[closest] = struct{}{}
		res += minDist
	}
	return res
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := bytes.Split(bytes.TrimRight(data, "\n"), []byte("\n"))
	nodes := make(map[byte]*Node)

	for i := range input {
		for j := range input[i] {
			if bytes.ContainsAny([]byte{input[i][j]}, "1234567890") {
				nodes[input[i][j]] = &Node{input[i][j], Coord{j, i}, make(map[byte]Vertex)}
			}
		}
	}
	for _, n := range nodes {
		buildVertices(n, nodes, input)
	}
	printNodes(nodes)
	fmt.Println(solve(nodes))
}
