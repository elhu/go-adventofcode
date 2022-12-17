package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Shape []coords2d.Coords2d

var minus = Shape{{0, 0}, {1, 0}, {2, 0}, {3, 0}}
var plus = Shape{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}}
var angle = Shape{{2, 2}, {2, 1}, {2, 0}, {0, 0}, {1, 0}}
var bar = Shape{{0, 0}, {0, 1}, {0, 2}, {0, 3}}
var square = Shape{{0, 0}, {0, 1}, {1, 0}, {1, 1}}

var shapes = []Shape{minus, plus, angle, bar, square}

var windDirection = map[byte]int{'<': -1, '>': 1}

func collides(rock Shape, offset coords2d.Coords2d, grid [][7]bool) bool {
	for _, p := range rock {
		pPosition := coords2d.Add(p, offset)
		if pPosition.X < 0 || pPosition.X >= 7 {
			return true
		}
		if pPosition.Y <= 0 || grid[pPosition.Y][pPosition.X] {
			return true
		}
	}
	return false
}

type State struct {
	top        [][7]bool
	windIndex  int
	shapeIndex int
}

func stateToKey(state State) string {
	var parts []string
	for _, t := range state.top {
		parts = append(parts, fmt.Sprintf("%v", t))
	}
	parts = append(parts, fmt.Sprintf("%d:%d", state.windIndex, state.shapeIndex))
	return strings.Join(parts, ":")
}

// Only check the state for the top N rows
// 10 is arbitrary after manual checks, works on my input
const TOP_LIMIT = 10

func solve(wind []byte, stopAfter int) int {
	max := 0
	grid := make([][7]bool, 0)
	currentWind := 0
	seenStates := make(map[string][2]int)
	var heightOffset int
	for rID := 0; rID < stopAfter; rID++ {
		rock := shapes[rID%len(shapes)]
		for i := 0; len(grid) < max+4+4; i++ {
			grid = append(grid, [7]bool{})
		}
		if len(grid) > TOP_LIMIT {
			state := stateToKey(State{top: grid[len(grid)-TOP_LIMIT:], windIndex: currentWind % len(wind), shapeIndex: rID % len(shapes)})
			if old, exists := seenStates[state]; exists {
				period := rID - old[1]
				growth := max - old[0]
				if heightOffset == 0 {
					heightOffset = growth * ((stopAfter - rID) / period)
				}
				stopAfter = (stopAfter-rID)%period + rID
			} else {
				seenStates[state] = [2]int{max, rID}
			}
		}
		offset := coords2d.Coords2d{X: 2, Y: max + 4}
		for i := 0; ; i++ {
			if i%2 == 0 {
				direction := windDirection[wind[currentWind%len(wind)]]
				newOffset := coords2d.Add(offset, coords2d.Coords2d{X: direction, Y: 0})
				if !collides(rock, newOffset, grid) {
					offset = newOffset
				}
				currentWind++
			} else {
				newOffset := coords2d.Add(offset, coords2d.Coords2d{X: 0, Y: -1})
				if collides(rock, newOffset, grid) {
					for _, p := range rock {
						pPosition := coords2d.Add(p, offset)
						grid[pPosition.Y][pPosition.X] = true
						if pPosition.Y > max {
							max = pPosition.Y
						}
					}
					break
				}
				offset = newOffset
			}
		}
	}
	return max + heightOffset
}

func main() {
	data := bytes.TrimRight(files.ReadFile(os.Args[1]), "\n")
	fmt.Println(solve(data, 1000000000000))
}
