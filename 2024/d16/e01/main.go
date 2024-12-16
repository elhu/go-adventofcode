package main

import (
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
	// fmt.Println("Trying to find edge from", curr)
	for next = coords3d.Add(curr, dirs[curr.Z]); ; next = coords3d.Add(next, dirs[curr.Z]) {
		// fmt.Printf(" Visiting %d (%c)\n", next, lines[next.Y][next.X])
		if lines[next.Y][next.X] == '#' {
			return coords3d.Coords3d{}, false
		} else if lines[next.Y][next.X] == 'E' || lines[next.Y][next.X] == 'S' {
			// fmt.Printf(" > Edge found %v (%c)!\n", next, lines[next.Y][next.X])
			if lines[next.Y][next.X] == 'E' {
			}
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
			// fmt.Printf(" > Edge found %v (%c)!\n", next, lines[next.Y][next.X])
			return next, true
		}
	}
}

const TURN_COST = 1000

func buildGraph(lines []string) *graphs.Graph[coords3d.Coords3d, coords3d.Coords3d] {
	graph := graphs.NewWeightedGraph[coords3d.Coords3d, coords3d.Coords3d]()
	startCoords := coords3d.Coords3d{X: 1, Y: len(lines) - 2, Z: 0}
	// fmt.Printf("Starting from %v (%c)\n", startCoords, lines[startCoords.Y][startCoords.X])
	queue := []coords3d.Coords3d{startCoords}
	visited := sets.New[coords3d.Coords3d]()
	var curr coords3d.Coords3d
	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		if visited.HasMember(curr) {
			continue
		}
		visited.Add(curr)
		// fmt.Printf("Processing %v (%c)\n", curr, lines[curr.Y][curr.X])
		// if _, err := graph.GetVertex(curr); err == nil {
		// 	continue
		// }
		graph.AddVertex(curr, curr)
		// Find the next edge
		e, f := findEdge(lines, curr)
		if f {
			if _, err := graph.GetVertex(e); err != nil {
				graph.AddVertex(e, e)
			}
			// fmt.Printf("Edge found: %v (%c) [%d]\n", e, lines[e.Y][e.X], coords3d.Distance(curr, e))
			// fmt.Printf("Adding directional edge: %v (%c) -> %v (%c) [%d]\n", curr, lines[curr.Y][curr.X], e, lines[e.Y][e.X], coords3d.Distance(curr, e))
			graph.AddEdge(curr, e, coords3d.Distance(curr, e))
			queue = append(queue, e)
		}
		// Rotate and add the nodes
		for i := 0; i < len(dirs); i++ {
			a, b := coords3d.Coords3d{X: curr.X, Y: curr.Y, Z: (i - 1 + len(dirs)) % len(dirs)}, coords3d.Coords3d{X: curr.X, Y: curr.Y, Z: i}
			// fmt.Printf("Rotating to %v (%c) and %v (%c)\n", a, lines[a.Y][a.X], b, lines[b.Y][b.X])
			if _, err := graph.GetVertex(a); err != nil {
				// graph.AddVertex(a, a)
				// fmt.Println("Adding vertex", a)
				queue = append(queue, a)
			}
			if _, err := graph.GetVertex(b); err != nil {
				// graph.AddVertex(b, b)
				// fmt.Println("Adding vertex", b)
				queue = append(queue, b)
			}
			// fmt.Printf("Adding rotational edge: %v (%c) -> %v (%c): %d\n", a, lines[a.Y][a.Y], b, lines[b.Y][b.Y], TURN_COST)
			graph.AddEdge(a, b, TURN_COST)
			// fmt.Printf("Adding rotational edge: %v (%c) -> %v (%c): %d\n", b, lines[b.Y][b.Y], a, lines[a.Y][a.Y], TURN_COST)
			graph.AddEdge(b, a, TURN_COST)
		}
	}
	// fmt.Println(graph.Vertices)
	// for _, v := range graph.Edges {
	// 	for _, e := range v {
	// 		fmt.Printf("%v (%c) -> %v (%c): %d\n", e.From, lines[e.From.Y][e.From.Y], e.To, lines[e.To.Y][e.To.Y], e.Weight)
	// 	}
	// }
	// fmt.Println(graph.Edges)
	return graph
}

func dbgPath(graph *graphs.Graph[coords3d.Coords3d, coords3d.Coords3d], start, end coords3d.Coords3d, lines []string) {
	path, _ := graph.ShortestPath(start, end)
	grid := make([][]byte, len(lines))
	var chars = []byte{'>', 'v', '<', '^'}
	fmt.Println(path)
	for i, line := range lines {
		grid[i] = []byte(line)
	}
	for i := len(path) - 1; i >= 0; i-- {
		e := path[i]
		grid[e.Y][e.X] = chars[e.Z]
	}
	for _, line := range grid {
		fmt.Println(string(line))
	}
}

func solve(graph *graphs.Graph[coords3d.Coords3d, coords3d.Coords3d], start, end coords3d.Coords3d, lines []string) int {
	_, err := graph.GetVertex(start)
	if err != nil {
		panic(err)
	}
	res := 99999999999999999
	for i := range dirs {
		e := coords3d.Coords3d{X: end.X, Y: end.Y, Z: i}
		if _, err := graph.GetVertex(e); err != nil {
			fmt.Println("Can't find end vertex", e)
		}
		// fmt.Printf("Looking for shortest path from %v (%c) to %v (%c)\n", start, 'S', e, 'E')
		dist, err := graph.ShortestDistance(start, e)
		if err == nil {
			if dist < res {
				res = dist
			}
		}
		// fmt.Printf("Distance: %d\n", dist)
		// dbgPath(graph, start, e, lines)
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	g := buildGraph(lines)
	s := coords3d.Coords3d{X: 1, Y: len(lines) - 2, Z: 0}
	e := coords3d.Coords3d{X: len(lines[0]) - 2, Y: 1, Z: 0}
	fmt.Println(solve(g, s, e, lines))
}
