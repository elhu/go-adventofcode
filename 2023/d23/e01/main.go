package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"adventofcode/utils/pqueue"
	set "adventofcode/utils/sets"
	"container/heap"
	"fmt"
	"os"
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

type State struct {
	pos     coords2d.Coords2d
	visited *set.Set[coords2d.Coords2d]
}

var (
	north = coords2d.Coords2d{X: 0, Y: -1}
	east  = coords2d.Coords2d{X: 1, Y: 0}
	south = coords2d.Coords2d{X: 0, Y: 1}
	west  = coords2d.Coords2d{X: -1, Y: 0}
)

var vectors = map[byte][]coords2d.Coords2d{
	'.': {north, east, south, west},
	'^': {north},
	'>': {east},
	'v': {south},
	'<': {west},
}

func astar(grid []string, startPos coords2d.Coords2d) int {
	var pq pqueue.PriorityQueue[*State]
	heap.Init(&pq)
	visited := set.New[coords2d.Coords2d]()
	visited.Add(startPos)
	heap.Push(&pq, &pqueue.Item[*State]{Value: &State{pos: startPos, visited: visited}, Priority: 0})
	target := coords2d.Coords2d{X: len(grid[0]) - 3, Y: len(grid) - 2}
	var candidatePaths []*set.Set[coords2d.Coords2d]
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*pqueue.Item[*State])
		curr := item.Value
		curr.visited.Add(curr.pos)
		if curr.pos == target {
			candidatePaths = append(candidatePaths, curr.visited)
			continue
		}
		for _, dir := range vectors[grid[curr.pos.Y][curr.pos.X]] {
			next := coords2d.Add(curr.pos, dir)
			if !curr.visited.HasMember(next) && grid[next.Y][next.X] != '#' {
				visited := set.NewFromSlice(curr.visited.Members())
				visited.Add(next)
				heap.Push(&pq, &pqueue.Item[*State]{Value: &State{pos: next, visited: visited}, Priority: item.Priority + 1})
			}
		}
	}
	max := 0
	for _, path := range candidatePaths {
		if path.Len() > max {
			max = path.Len()
		}
	}
	return max
}

func solve(grid []string) int {
	startPos := coords2d.Coords2d{X: 2, Y: 1}
	return astar(grid, startPos) - 1
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	lines = pad(lines, "#")
	fmt.Println(solve(lines))
}
