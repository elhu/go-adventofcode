package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"adventofcode/utils/sets/stringset"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	WALL     = iota
	BLIZZARD = iota
	EMPTY    = iota
)

var directions = map[byte]coords2d.Coords2d{
	'>': {1, 0},
	'<': {-1, 0},
	'^': {0, -1},
	'v': {0, 1},
}

var printDirection = map[byte]coords2d.Coords2d{
	'>': {1, 0},
	'<': {-1, 0},
	'^': {0, -1},
	'v': {0, 1},
}

type Cell struct {
	kind      int
	blizzards []coords2d.Coords2d
}

func toString(state State) string {
	parts := make([]string, 0)
	valley := state.valley
	for i, l := range valley {
		for j, c := range l {
			if state.pos.Y == i && state.pos.X == j {
				parts = append(parts, "E")
				continue
			}
			switch c.kind {
			case WALL:
				parts = append(parts, "#")
			case BLIZZARD:
				if len(c.blizzards) > 1 {
					parts = append(parts, strconv.Itoa(len(c.blizzards)))
				} else {
					switch c.blizzards[0] {
					case coords2d.Coords2d{1, 0}:
						parts = append(parts, ">")
					case coords2d.Coords2d{-1, 0}:
						parts = append(parts, "<")
					case coords2d.Coords2d{0, -1}:
						parts = append(parts, "^")
					case coords2d.Coords2d{0, 1}:
						parts = append(parts, "v")
					}
				}
			case EMPTY:
				parts = append(parts, ".")
			}
		}
		parts = append(parts, "\n")
	}
	return strings.Join(parts, "")
}

func emptyValley(valley [][]Cell) [][]Cell {
	empty := make([][]Cell, len(valley))
	for i, l := range valley {
		empty[i] = make([]Cell, len(l))
		for j, c := range l {
			if c.kind == WALL {
				empty[i][j] = Cell{kind: WALL}
			} else {
				empty[i][j] = Cell{kind: EMPTY}
			}
		}
	}
	return empty
}

func moveBlizzards(valley [][]Cell) [][]Cell {
	newState := emptyValley(valley)
	for i, l := range valley {
		for j, c := range l {
			pos := coords2d.Coords2d{X: j, Y: i}
			for _, b := range c.blizzards {
				dest := coords2d.Add(pos, b)
				if newState[dest.Y][dest.X].kind == WALL {
					switch b {
					case coords2d.Coords2d{X: 1, Y: 0}:
						dest = coords2d.Coords2d{X: 1, Y: dest.Y}
					case coords2d.Coords2d{X: -1, Y: 0}:
						dest = coords2d.Coords2d{X: len(l) - 2, Y: dest.Y}
					case coords2d.Coords2d{X: 0, Y: 1}:
						dest = coords2d.Coords2d{X: dest.X, Y: 1}
					case coords2d.Coords2d{X: 0, Y: -1}:
						dest = coords2d.Coords2d{X: dest.X, Y: len(valley) - 2}
					}
				}
				newState[dest.Y][dest.X].kind = BLIZZARD
				newState[dest.Y][dest.X].blizzards = append(newState[dest.Y][dest.X].blizzards, b)
			}
		}
	}
	return newState
}

type State struct {
	pos    coords2d.Coords2d
	valley [][]Cell
	moves  int
}

func candidates(pos coords2d.Coords2d, valley [][]Cell) []coords2d.Coords2d {
	res := make([]coords2d.Coords2d, 0)
	if valley[pos.Y][pos.X].kind == EMPTY {
		res = append(res, pos)
	}
	for _, v := range directions {
		dest := coords2d.Add(pos, v)
		if dest.Y >= 0 && dest.X >= 0 && dest.Y < len(valley) && dest.Y < len(valley[0]) &&
			valley[dest.Y][dest.X].kind == EMPTY {
			res = append(res, dest)
		}
	}
	return res
}

func findPath(valley [][]Cell, start, end coords2d.Coords2d) State {
	queue := []State{{pos: start, valley: valley, moves: 0}}
	seen := stringset.New()
	var head State
	for len(queue) > 0 {
		head, queue = queue[0], queue[1:]
		k := toString(head)
		if seen.HasMember(k) {
			continue
		}
		seen.Add(k)
		if head.pos == end {
			return head
		}

		newValley := moveBlizzards(head.valley)
		for _, c := range candidates(head.pos, newValley) {
			queue = append(queue, State{pos: c, valley: newValley, moves: head.moves + 1})
		}
	}
	panic("WTF")
}

func solve(valley [][]Cell) int {
	res := 0
	start := coords2d.Coords2d{X: 1, Y: 0}
	end := coords2d.Coords2d{X: len(valley[0]) - 2, Y: len(valley) - 1}
	finalState := findPath(valley, start, end)
	fmt.Printf("Go there in %d\n", finalState.moves)
	res += finalState.moves
	finalState = findPath(finalState.valley, end, start)
	fmt.Printf("Go back in %d\n", finalState.moves)
	res += finalState.moves
	finalState = findPath(finalState.valley, start, end)
	fmt.Printf("Finally got out in %d\n", finalState.moves)
	res += finalState.moves
	return res
}

func main() {
	data := files.ReadLines(os.Args[1])
	valley := make([][]Cell, len(data))
	for i, l := range data {
		valley[i] = make([]Cell, len(l))
		for j, c := range l {
			switch c {
			case '#':
				valley[i][j] = Cell{kind: WALL}
			case '.':
				valley[i][j] = Cell{kind: EMPTY}
			default:
				valley[i][j] = Cell{kind: BLIZZARD, blizzards: []coords2d.Coords2d{directions[byte(c)]}}
			}
		}
	}
	fmt.Println(solve(valley))
	// s := coords2d.Coords2d{X: 1, Y: 0}
	// fmt.Println(toString(State{pos: s, valley: valley}))
	// fmt.Println("--")
	// valley = moveBlizzards(valley)
	// fmt.Println(toString(State{pos: s, valley: valley}))
	// fmt.Println("--")
	// valley = moveBlizzards(valley)
	// fmt.Println(toString(State{pos: s, valley: valley}))
	// fmt.Println("--")
	// valley = moveBlizzards(valley)
	// fmt.Println(toString(State{pos: s, valley: valley}))
	// fmt.Println("--")
}
