package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"adventofcode/utils/graphs"
	set "adventofcode/utils/sets"
	"fmt"
	"os"
	"sort"
	"strings"
)

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

func bfs(start, end coords2d.Coords2d, graph *graphs.Graph[coords2d.Coords2d, struct{}]) int {
	type state struct {
		curr coords2d.Coords2d
		path map[coords2d.Coords2d]int
	}
	var queue []state
	queue = append(queue, state{curr: start, path: map[coords2d.Coords2d]int{start: 0}})
	var head state
	var validPaths []state
	for len(queue) > 0 {
		head, queue = queue[0], queue[1:]
		if head.curr == end {
			validPaths = append(validPaths, head)
			continue
		}
		for n, e := range graph.Edges[head.curr] {
			if _, found := head.path[n]; !found {
				newPath := make(map[coords2d.Coords2d]int)
				for k, v := range head.path {
					newPath[k] = v
				}
				newPath[n] = e.Weight
				queue = append(queue, state{curr: n, path: newPath})
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

func visitedToKey(visited map[coords2d.Coords2d]struct{}) string {
	var keys []string
	for k := range visited {
		keys = append(keys, fmt.Sprintf("%d,%d", k.X, k.Y))
	}
	sort.StringSlice(keys).Sort()
	return strings.Join(keys, ";")
}

func findNeighbors(grid []string, junctions *graphs.Graph[coords2d.Coords2d, struct{}], pos coords2d.Coords2d) map[coords2d.Coords2d]int {
	neighbors := make(map[coords2d.Coords2d]int)
	toVisit := []coords2d.Coords2d{pos}
	var next []coords2d.Coords2d
	dist := 0
	visited := set.New[coords2d.Coords2d]()
	var head coords2d.Coords2d
	for len(toVisit) > 0 {
		head, toVisit = toVisit[0], toVisit[1:]
		visited.Add(head)
		if _, err := junctions.GetVertex(head); err == nil && pos != head {
			neighbors[head] = dist
		} else {
			for _, dir := range []coords2d.Coords2d{north, east, south, west} {
				n := coords2d.Add(head, dir)
				if !visited.HasMember(n) {
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

func buildJunctionGraph(grid []string) *graphs.Graph[coords2d.Coords2d, struct{}] {
	junctions := graphs.NewWeightedGraph[coords2d.Coords2d, struct{}]()
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
					junctions.AddVertex(pos, struct{}{})
				}
			}
		}
	}
	startPos, endPos := coords2d.Coords2d{X: 2, Y: 1}, coords2d.Coords2d{X: len(grid[0]) - 3, Y: len(grid) - 2}
	junctions.AddVertex(startPos, struct{}{})
	junctions.AddVertex(endPos, struct{}{})
	for k := range junctions.Vertices {
		for n, d := range findNeighbors(grid, junctions, k) {
			junctions.AddEdge(k, n, d)
		}
	}
	return junctions
}

func solve(grid []string) int {
	graph := buildJunctionGraph(grid)
	startPos, endPos := coords2d.Coords2d{X: 2, Y: 1}, coords2d.Coords2d{X: len(grid[0]) - 3, Y: len(grid) - 2}
	return bfs(startPos, endPos, graph)
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	lines = pad(lines, "#")
	fmt.Println(solve(lines))
}
