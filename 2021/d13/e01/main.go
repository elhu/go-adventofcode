package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"adventofcode/utils/sets/intset"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func coordToKey(c coords2d.Coords2d) int {
	return c.Y*100000 + c.X
}

func keyToCoord(k int) coords2d.Coords2d {
	y := k / 100000
	return coords2d.Coords2d{
		Y: y,
		X: k - y*100000,
	}
}

func atoi(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return n
}

func parsePoints(input []string) *intset.IntSet {
	res := intset.New()
	for _, l := range input {
		parts := strings.Split(l, ",")
		res.Add(
			coordToKey(coords2d.Coords2d{X: atoi(parts[0]), Y: atoi(parts[1])}),
		)
	}
	return res
}

func foldLeft(pos int, points *intset.IntSet) {
	points.Each(func(p int) {
		coords := keyToCoord(p)
		if coords.X > pos {
			points.Remove(p)
			coords.X = pos - (coords.X - pos)
			points.Add(
				coordToKey(coords),
			)
		}
	})
}

func foldTop(pos int, points *intset.IntSet) {
	points.Each(func(p int) {
		coords := keyToCoord(p)
		if coords.Y > pos {
			points.Remove(p)
			coords.Y = pos - (coords.Y - pos)
			points.Add(
				coordToKey(coords),
			)
		}
	})
}

func solve(points *intset.IntSet, insts []string) int {
	var axis string
	var pos int
	for _, inst := range insts {
		fmt.Sscanf(strings.ReplaceAll(inst, "=", " "), "fold along %s %d", &axis, &pos)
		if axis == "x" {
			foldLeft(pos, points)
		} else {
			foldTop(pos, points)
		}
		break
	}
	return points.Len()
}

func main() {
	data := files.ReadFile(os.Args[1])
	parts := strings.Split(string(data), "\n\n")
	points := parsePoints(strings.Split(parts[0], "\n"))
	fmt.Println(solve(points, strings.Split(parts[1], "\n")))
}
