package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	pos       coords2d.Coords2d
	neighbors map[*Node]int
}

func pad(lines []string, char string) []string {
	padded := make([]string, len(lines)+2)
	padded[0] = strings.Repeat(char, len(lines[0])+2)
	for i, line := range lines {
		padded[i+1] = char + line + char
	}
	padded[len(lines)+1] = strings.Repeat(char, len(lines[0])+2)
	return padded
}

var (
	north = coords2d.Coords2d{X: 0, Y: -1}
	east  = coords2d.Coords2d{X: 1, Y: 0}
	south = coords2d.Coords2d{X: 0, Y: 1}
	west  = coords2d.Coords2d{X: -1, Y: 0}
)

func findNeighbors(grid []string, junctions map[coords2d.Coords2d]*Node, pos coords2d.Coords2d) map[*Node]int {
	neighbors := make(map[*Node]int)
	toVisit := []coords2d.Coords2d{pos}
	var next []coords2d.Coords2d
	dist := 0
	visited := make(map[coords2d.Coords2d]struct{})
	var head coords2d.Coords2d
	for len(toVisit) > 0 {
		head, toVisit = toVisit[0], toVisit[1:]
		visited[head] = struct{}{}
		if node, found := junctions[head]; pos != head && found {
			neighbors[node] = dist
		} else {
			for _, dir := range []coords2d.Coords2d{north, east, south, west} {
				n := coords2d.Add(head, dir)
				if _, found := visited[n]; !found {
					if grid[n.Y][n.X] != '#' {
						next = append(next, n)
					}
				}
			}
		}
		if len(toVisit) == 0 {
			dist++
			toVisit = make([]coords2d.Coords2d, len(next))
			copy(toVisit, next)
			next = nil
		}
	}
	return neighbors
}

func buildJunctionGraph(grid []string) (*Node, *Node, map[coords2d.Coords2d]*Node) {
	junctions := make(map[coords2d.Coords2d]*Node)
	for y, line := range grid {
		for x, c := range line {
			if c != '#' {
				options := 0
				for _, dir := range []coords2d.Coords2d{north, east, south, west} {
					n := coords2d.Add(coords2d.Coords2d{X: x, Y: y}, dir)
					if grid[n.Y][n.X] != '#' {
						options++
					}
				}
				if options > 2 {
					pos := coords2d.Coords2d{X: x, Y: y}
					junctions[pos] = &Node{pos: pos, neighbors: make(map[*Node]int)}
				}
			}
		}
	}
	startPos, endPos := coords2d.Coords2d{X: 2, Y: 1}, coords2d.Coords2d{X: len(grid[0]) - 3, Y: len(grid) - 2}
	junctions[startPos] = &Node{pos: startPos, neighbors: make(map[*Node]int)}
	junctions[endPos] = &Node{pos: endPos, neighbors: make(map[*Node]int)}
	for k, v := range junctions {
		v.neighbors = findNeighbors(grid, junctions, k)
		var ns []string
		for n, d := range v.neighbors {
			ns = append(ns, fmt.Sprintf("%v(%d)", n.pos, d))
		}
	}
	return junctions[startPos], junctions[endPos], junctions
}

type State struct {
	currNode *Node
	path     map[*Node]int
}

func bfs(start, end *Node) int {
	var queue []State
	queue = append(queue, State{currNode: start, path: map[*Node]int{start: 0}})
	var head State
	var validPaths []State
	for len(queue) > 0 {
		head, queue = queue[0], queue[1:]
		if head.currNode == end {
			validPaths = append(validPaths, head)
			continue
		}
		for n, d := range head.currNode.neighbors {
			if _, found := head.path[n]; !found {
				newPath := make(map[*Node]int)
				for k, v := range head.path {
					newPath[k] = v
				}
				newPath[n] = d
				queue = append(queue, State{currNode: n, path: newPath})
			}
		}
	}
	max := 0
	for _, vp := range validPaths {
		len := 0
		for _, d := range vp.path {
			len += d
		}
		if len > max {
			max = len
		}
	}
	return max
}

func solve(grid []string) int {
	startNode, endNode, nodes := buildJunctionGraph(grid)
	nn := make(map[coords2d.Coords2d]struct{})
	for n := range nodes {
		nn[n] = struct{}{}
	}
	return bfs(startNode, endNode)
}

var target = coords2d.Coords2d{X: 0, Y: 0}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	lines = pad(lines, "#")
	fmt.Println(solve(lines))
}
