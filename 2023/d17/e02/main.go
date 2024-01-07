package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"adventofcode/utils/pqueue"
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
}

const maxMoves = 10
const minMoves = 4

func inBounds(grid []string, pos coords2d.Coords2d) bool {
	return pos.X >= 0 && pos.X < len(grid[0]) && pos.Y >= 0 && pos.Y < len(grid)
}

func toVisit(grid []string, seen map[[4]int]int, s *pqueue.Item[*State]) bool {
	k := [4]int{s.Value.pos.X, s.Value.pos.Y, s.Value.dirIdx, s.Value.moveCounter}
	if val, found := seen[k]; found && val <= s.Priority {
		return false
	}
	seen[k] = s.Priority
	return true
}

func solve(grid []string) int {
	var pq pqueue.PriorityQueue[*State]
	heap.Init(&pq)
	heap.Push(&pq, &pqueue.Item[*State]{Value: &State{pos: coords2d.Coords2d{X: 0, Y: 0}, dirIdx: 2, moveCounter: 0}, Priority: 0}) // start facing south
	heap.Push(&pq, &pqueue.Item[*State]{Value: &State{pos: coords2d.Coords2d{X: 0, Y: 0}, dirIdx: 1, moveCounter: 0}, Priority: 0}) // start facing east
	target := coords2d.Coords2d{X: len(grid[0]) - 1, Y: len(grid) - 1}
	seen := make(map[[4]int]int)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*pqueue.Item[*State])
		if item.Value.pos == target && item.Value.moveCounter >= minMoves {
			return item.Priority
		}

		seen[[4]int{item.Value.pos.X, item.Value.pos.Y, item.Value.dirIdx, item.Value.moveCounter}] = item.Priority
		// 0..10 -> move straight
		// 4..10 -> turn
		if item.Value.moveCounter < maxMoves {
			if newPos := coords2d.Add(item.Value.pos, directions[item.Value.dirIdx]); inBounds(grid, newPos) {
				newState := &State{pos: newPos, dirIdx: item.Value.dirIdx, moveCounter: item.Value.moveCounter + 1}
				newItem := &pqueue.Item[*State]{Value: newState, Priority: item.Priority + int(grid[newPos.Y][newPos.X]-'0')}
				if toVisit(grid, seen, newItem) {
					heap.Push(&pq, newItem)
				}
			}
		}
		if item.Value.moveCounter >= minMoves {
			for _, i := range []int{-1, 1} {
				newDir := (item.Value.dirIdx + i) % len(directions)
				if newDir < 0 {
					newDir = len(directions) - 1
				}
				if newPos := coords2d.Add(item.Value.pos, directions[newDir]); inBounds(grid, newPos) {
					newState := &State{pos: newPos, dirIdx: newDir, moveCounter: 1}
					newItem := &pqueue.Item[*State]{Value: newState, Priority: item.Priority + int(grid[newPos.Y][newPos.X]-'0')}
					if toVisit(grid, seen, newItem) {
						heap.Push(&pq, newItem)
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
