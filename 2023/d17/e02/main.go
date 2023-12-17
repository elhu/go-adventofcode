package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"container/heap"
	"fmt"
	"os"
	"strings"
)

var north = coords2d.Coords2d{X: 0, Y: -1} // North
var east = coords2d.Coords2d{X: 1, Y: 0}   // East
var south = coords2d.Coords2d{X: 0, Y: 1}  // South
var west = coords2d.Coords2d{X: -1, Y: 0}  // West

var directions = [4]coords2d.Coords2d{north, east, south, west}

type State struct {
	pos         coords2d.Coords2d
	dirIdx      int
	moveCounter int
	priority    int
	index       int
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*State)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

const maxMoves = 10
const minMoves = 4

func inBounds(grid []string, pos coords2d.Coords2d) bool {
	return pos.X >= 0 && pos.X < len(grid[0]) && pos.Y >= 0 && pos.Y < len(grid)
}

func toVisit(grid []string, seen map[[4]int]int, s *State) bool {
	k := [4]int{s.pos.X, s.pos.Y, s.dirIdx, s.moveCounter}
	if val, found := seen[k]; found && val <= s.priority {
		return false
	}
	seen[k] = s.priority
	return true
}

func solve(grid []string) int {
	var pq PriorityQueue
	heap.Init(&pq)
	heap.Push(&pq, &State{pos: coords2d.Coords2d{X: 0, Y: 0}, dirIdx: 2, moveCounter: 0, priority: 0}) // start facing south
	heap.Push(&pq, &State{pos: coords2d.Coords2d{X: 0, Y: 0}, dirIdx: 1, moveCounter: 0, priority: 0}) // start facing east
	target := coords2d.Coords2d{X: len(grid[0]) - 1, Y: len(grid) - 1}
	seen := make(map[[4]int]int)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*State)
		if item.pos == target && item.moveCounter >= minMoves {
			return item.priority
		}

		seen[[4]int{item.pos.X, item.pos.Y, item.dirIdx, item.moveCounter}] = item.priority
		// 0..10 -> move straight
		// 4..10 -> turn
		if item.moveCounter < maxMoves {
			if newPos := coords2d.Add(item.pos, directions[item.dirIdx]); inBounds(grid, newPos) {
				newState := &State{pos: newPos, dirIdx: item.dirIdx, moveCounter: item.moveCounter + 1, priority: item.priority + int(grid[newPos.Y][newPos.X]-'0')}
				if toVisit(grid, seen, newState) {
					heap.Push(&pq, newState)
				}
			}
		}
		if item.moveCounter >= minMoves {
			for _, i := range []int{-1, 1} {
				newDir := (item.dirIdx + i) % len(directions)
				if newDir < 0 {
					newDir = len(directions) - 1
				}
				if newPos := coords2d.Add(item.pos, directions[newDir]); inBounds(grid, newPos) {
					newState := &State{pos: newPos, dirIdx: newDir, moveCounter: 1, priority: item.priority + int(grid[newPos.Y][newPos.X]-'0')}
					if toVisit(grid, seen, newState) {
						heap.Push(&pq, newState)
					}
				}
			}
		}
	}
	panic("WTF")
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	grid := strings.Split(data, "\n")
	fmt.Println(solve(grid))
}
