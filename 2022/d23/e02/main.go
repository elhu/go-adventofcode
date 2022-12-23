package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"fmt"
	"os"
)

func toKey(c coords2d.Coords2d) string {
	return fmt.Sprintf("%d:%d", c.X, c.Y)
}

func fromKey(k string) coords2d.Coords2d {
	var x, y int
	fmt.Sscanf(k, "%d:%d", &x, &y)
	return coords2d.Coords2d{
		X: x, Y: y,
	}
}

var quadrants = [4][3]coords2d.Coords2d{
	{{X: -1, Y: -1}, {X: 0, Y: -1}, {X: 1, Y: -1}}, // NW, N, NE
	{{X: -1, Y: 1}, {X: 0, Y: 1}, {X: 1, Y: 1}},    // SW, S, SE
	{{X: -1, Y: -1}, {X: -1, Y: 0}, {X: -1, Y: 1}}, // NW, W, SW
	{{X: 1, Y: -1}, {X: 1, Y: 0}, {X: 1, Y: 1}},    // NE, E, SE
}

func shouldMove(elves map[string]int, pos coords2d.Coords2d) bool {
	for _, q := range quadrants {
		for _, c := range q {
			if _, exists := elves[toKey(coords2d.Add(pos, c))]; exists {
				return true
			}
		}
	}
	return false
}

func copyMap(m map[string]int) map[string]int {
	d := make(map[string]int)
	for k, v := range m {
		d[k] = v
	}
	return d
}

func equals(m, o map[string]int) bool {
	for k := range m {
		if _, exists := o[k]; !exists {
			return false
		}
	}
	return true
}

func solve(elves map[string]int) int {
	qID := 0
	prevElves := make(map[string]int)
	for !equals(elves, prevElves) {
		prevElves = copyMap(elves)
		newPos := make(map[string][]string)
		for k := range elves {
			pos := fromKey(k)
			if !shouldMove(elves, pos) {
				continue
			}
			for q := 0; q < len(quadrants); q++ {
				q1 := quadrants[(qID+q)%len(quadrants)][0]
				q2 := quadrants[(qID+q)%len(quadrants)][1]
				q3 := quadrants[(qID+q)%len(quadrants)][2]
				_, check1 := elves[toKey(coords2d.Add(pos, q1))]
				_, check2 := elves[toKey(coords2d.Add(pos, q2))]
				_, check3 := elves[toKey(coords2d.Add(pos, q3))]
				if !check1 && !check2 && !check3 {
					newPos[toKey(coords2d.Add(pos, q2))] = append(newPos[toKey(coords2d.Add(pos, q2))], k)
					break
				}
			}
		}
		for k, v := range newPos {
			if len(v) == 1 {
				delete(elves, v[0])
				elves[k] = 1
			}
		}
		qID++
	}
	return qID
}

func main() {
	data := files.ReadLines(os.Args[1])
	elves := make(map[string]int)
	for y, l := range data {
		for x, c := range l {
			if c == '#' {
				elves[toKey(coords2d.Coords2d{X: x, Y: y})] = 1
			}
		}
	}
	fmt.Println(solve(elves))
}
