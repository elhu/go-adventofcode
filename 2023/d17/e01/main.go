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

const maxMoves = 3

func inBounds(grid []string, pos coords2d.Coords2d) bool {
	return pos.X >= 0 && pos.X < len(grid[0]) && pos.Y >= 0 && pos.Y < len(grid)
}

func solve(grid []string) int {
	var pq pqueue.PriorityQueue[*State]
	heap.Init(&pq)
	heap.Push(&pq, &pqueue.Item[*State]{Value: &State{pos: coords2d.Coords2d{X: 0, Y: 0}, dirIdx: 2, moveCounter: 0}, Priority: 0})
	heap.Push(&pq, &pqueue.Item[*State]{Value: &State{pos: coords2d.Coords2d{X: 0, Y: 0}, dirIdx: 1, moveCounter: 0}, Priority: 0})
	seen := make(map[[4]int]int)
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*pqueue.Item[*State])
		if item.Value.pos.X == len(grid[0])-1 && item.Value.pos.Y == len(grid)-1 {
			return item.Priority
		}

		if val, found := seen[[4]int{item.Value.pos.X, item.Value.pos.Y, item.Value.dirIdx, item.Value.moveCounter}]; found && val <= item.Priority {
			continue
		}
		seen[[4]int{item.Value.pos.X, item.Value.pos.Y, item.Value.dirIdx, item.Value.moveCounter}] = item.Value.moveCounter
		if item.Value.moveCounter < maxMoves {
			if newPos := coords2d.Add(item.Value.pos, directions[item.Value.dirIdx]); inBounds(grid, newPos) {
				heap.Push(&pq, &pqueue.Item[*State]{Value: &State{pos: newPos, dirIdx: item.Value.dirIdx, moveCounter: item.Value.moveCounter + 1}, Priority: item.Priority + int(grid[newPos.Y][newPos.X]-'0')})
			}
		}
		for _, i := range []int{-1, 1} {
			newDir := (item.Value.dirIdx + i) % len(directions)
			if newDir < 0 {
				newDir = len(directions) - 1
			}
			if newPos := coords2d.Add(item.Value.pos, directions[newDir]); inBounds(grid, newPos) {
				heap.Push(&pq, &pqueue.Item[*State]{Value: &State{pos: newPos, dirIdx: newDir, moveCounter: 1}, Priority: item.Priority + int(grid[newPos.Y][newPos.X]-'0')})
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
