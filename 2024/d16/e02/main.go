package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/coords/coords3d"
	"adventofcode/utils/files"
	"adventofcode/utils/graphs"
	"adventofcode/utils/sets"
	"fmt"
	"os"
	"strings"
)

var EAST, SOUTH, WEST, NORTH = coords3d.Coords3d{X: 1, Y: 0, Z: 0}, coords3d.Coords3d{X: 0, Y: 1, Z: 0}, coords3d.Coords3d{X: -1, Y: 0, Z: 0}, coords3d.Coords3d{X: 0, Y: -1, Z: 0}
var dirs = []coords3d.Coords3d{EAST, SOUTH, WEST, NORTH}

func findEdge(lines []string, curr coords3d.Coords3d) (coords3d.Coords3d, bool) {
	next := coords3d.Add(curr, dirs[curr.Z])
	for next = coords3d.Add(curr, dirs[curr.Z]); ; next = coords3d.Add(next, dirs[curr.Z]) {
		if lines[next.Y][next.X] == '#' {
			return coords3d.Coords3d{}, false
		} else if lines[next.Y][next.X] == 'E' || lines[next.Y][next.X] == 'S' {
			return next, true
		}

		fc := 0
		for i, d := range dirs {
			if i == curr.Z || i == (curr.Z+2)%4 {
				continue
			}
			n := coords3d.Add(next, d)
			if lines[n.Y][n.X] == '.' || lines[n.Y][n.X] == 'E' || lines[n.Y][n.X] == 'S' {
				fc++
			}
		}
		if fc >= 1 {
			return next, true
		}
	}
}

const TURN_COST = 1000

func buildGraph(lines []string) *graphs.Graph[coords3d.Coords3d, coords3d.Coords3d] {
	graph := graphs.NewWeightedGraph[coords3d.Coords3d, coords3d.Coords3d]()
	startCoords := coords3d.Coords3d{X: 1, Y: len(lines) - 2, Z: 0}
	queue := []coords3d.Coords3d{startCoords}
	visited := sets.New[coords3d.Coords3d]()
	var curr coords3d.Coords3d
	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		if visited.HasMember(curr) {
			continue
		}
		visited.Add(curr)
		graph.AddVertex(curr, curr)
		e, f := findEdge(lines, curr)
		if f {
			if _, err := graph.GetVertex(e); err != nil {
				graph.AddVertex(e, e)
			}
			graph.AddEdge(curr, e, coords3d.Distance(curr, e))
			queue = append(queue, e)
		}
		// Rotate and add the nodes
		for i := 0; i < len(dirs); i++ {
			a, b := coords3d.Coords3d{X: curr.X, Y: curr.Y, Z: (i - 1 + len(dirs)) % len(dirs)}, coords3d.Coords3d{X: curr.X, Y: curr.Y, Z: i}
			if _, err := graph.GetVertex(a); err != nil {
				queue = append(queue, a)
			}
			if _, err := graph.GetVertex(b); err != nil {
				queue = append(queue, b)
			}
			graph.AddEdge(a, b, TURN_COST)
			graph.AddEdge(b, a, TURN_COST)
		}
	}
	return graph
}

func coords3dto2d(c coords3d.Coords3d) coords2d.Coords2d {
	return coords2d.Coords2d{X: c.X, Y: c.Y}
}

func solve(graph *graphs.Graph[coords3d.Coords3d, coords3d.Coords3d], from, to coords3d.Coords3d) int {
	tl := 99999999999999999
	var tt coords3d.Coords3d
	for d := range dirs {
		e := coords3d.Coords3d{X: to.X, Y: to.Y, Z: d}
		dist, err := graph.ShortestDistance(from, e)
		if err == nil && dist < tl {
			tl = dist
			tt = e
		}
	}
	paths, _ := graph.AllShortestPaths(from, tt)
	res := sets.New[coords2d.Coords2d]()
	for _, path := range paths {
		for i := len(path) - 1; i > 0; i-- {
			for c := path[i]; coords3dto2d(c) != coords3dto2d(path[i-1]); c = coords3d.Add(c, dirs[c.Z]) {
				res.Add(coords3dto2d(c))
			}
		}
		res.Add(coords3dto2d(path[0]))
	}
	return res.Len()
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	g := buildGraph(lines)
	s := coords3d.Coords3d{X: 1, Y: len(lines) - 2, Z: 0}
	e := coords3d.Coords3d{X: len(lines[0]) - 2, Y: 1, Z: 0}
	fmt.Println(solve(g, s, e))
}
