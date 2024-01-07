package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	set "adventofcode/utils/sets"
	"fmt"
	"os"
	"strings"
)

type Beam struct {
	dirIdx int
	pos    coords2d.Coords2d
	active bool
}

var north = coords2d.Coords2d{X: 0, Y: -1} // North
var east = coords2d.Coords2d{X: 1, Y: 0}   // East
var south = coords2d.Coords2d{X: 0, Y: 1}  // South
var west = coords2d.Coords2d{X: -1, Y: 0}  // West

var directions = [4]coords2d.Coords2d{north, east, south, west}

func activeBeams(beams []*Beam) int {
	res := 0
	for _, beam := range beams {
		if beam.active {
			res++
		}
	}
	return res
}

func pad(lines []string, val string) []string {
	padded := make([]string, len(lines)+2)
	padded[0] = strings.Repeat(val, len(lines[0])+2)
	for i, line := range lines {
		padded[i+1] = val + line + val
	}
	padded[len(lines)+1] = strings.Repeat(val, len(lines[0])+2)
	return padded
}

func move(grid []string, beam *Beam) *Beam {
	dir := directions[beam.dirIdx]
	beam.pos = coords2d.Add(beam.pos, dir)
	var newBeam *Beam
	switch grid[beam.pos.Y][beam.pos.X] {
	case 'O':
		beam.active = false
	case '/': // North -> East, East -> North, South -> West, West -> South
		if dir == north || dir == south {
			beam.dirIdx = (beam.dirIdx + 1) % len(directions)
		} else {
			beam.dirIdx = (beam.dirIdx - 1) % len(directions)
			if beam.dirIdx == -1 {
				beam.dirIdx = len(directions) - 1
			}
		}
	case '\\': // North -> East, East -> South, South -> West, West -> North
		if dir == north || dir == south {
			beam.dirIdx = (beam.dirIdx - 1) % len(directions)
			if beam.dirIdx == -1 {
				beam.dirIdx = len(directions) - 1
			}
		} else {
			beam.dirIdx = (beam.dirIdx + 1) % len(directions)
		}
	case '-':
		if dir.Y != 0 {
			beam.dirIdx = 3
			newBeam = &Beam{active: true, dirIdx: 1, pos: beam.pos}
		}
	case '|':
		if dir.X != 0 {
			beam.dirIdx = 0
			newBeam = &Beam{active: true, dirIdx: 2, pos: beam.pos}

		}
	}
	return newBeam
}

func toKey(beam *Beam) string {
	return fmt.Sprintf("%d:%d:%d", beam.pos.X, beam.pos.Y, beam.dirIdx)
}

func solve(grid []string) int {
	beams := []*Beam{{active: true, dirIdx: 1, pos: coords2d.Coords2d{X: 0, Y: 1}}}
	seen := set.New[string]()
	for activeBeams(beams) > 0 {
		for _, beam := range beams {
			if beam.pos.X == 2 && beam.pos.Y == 10 {
				fmt.Println(beam)
			}
			key := toKey(beam)
			if beam.active == false {
				continue
			}
			if seen.HasMember(key) {
				beam.active = false
			}

			newBeam := move(grid, beam)
			if newBeam != nil {
				beams = append(beams, newBeam)
			}
			seen.Add(key)
		}
	}
	return countEnergized(seen) - 1
}

func countEnergized(seen *set.Set[string]) int {
	visited := set.New[string]()
	for _, key := range seen.Members() {
		var x, y int
		fmt.Sscanf(key, "%d:%d:%d", &x, &y)
		visited.Add(fmt.Sprintf("%d:%d", x, y))
	}
	return visited.Len()
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	grid := strings.Split(data, "\n")
	grid = pad(grid, "O")
	fmt.Println(solve(grid))
}
