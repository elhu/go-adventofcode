package main

import (
	"adventofcode/utils/coords/coords3d"
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

type Hail struct {
	pos       coords3d.Coords3d
	vec       coords3d.Coords3d
	slope     float64
	intercept float64
}

func parseHail(line string) Hail {
	var h Hail
	fmt.Sscanf(line, "%d, %d, %d @ %d, %d, %d", &h.pos.X, &h.pos.Y, &h.pos.Z, &h.vec.X, &h.vec.Y, &h.vec.Z)
	nextPos := coords3d.Add(h.pos, h.vec)
	h.slope = float64(nextPos.Y-h.pos.Y) / float64(nextPos.X-h.pos.X)
	h.intercept = float64(h.pos.Y) - h.slope*float64(h.pos.X)
	return h
}

func intersectInRange(a, b Hail) bool {
	left := a.slope - b.slope
	right := b.intercept - a.intercept
	x := right / left
	y := a.slope*x + a.intercept
	if x < float64(MIN) || x > float64(MAX) || y < float64(MIN) || y > float64(MAX) {
		return false
	}
	if ((x >= float64(a.pos.X) && a.vec.X >= 0) || (x <= float64(a.pos.X) && a.vec.X <= 0)) &&
		((x >= float64(b.pos.X) && b.vec.X >= 0) || (x <= float64(b.pos.X) && b.vec.X <= 0)) {
		return true
	}
	return false
}

func solve(hailstones []Hail) int {
	res := 0
	for i, a := range hailstones {
		for _, b := range hailstones[i+1:] {
			if intersectInRange(a, b) {
				res++
			}
		}
	}
	return res
}

// const (
// 	MIN = 7
// 	MAX = 27
// )

// 21868 too high

const (
	MIN = 200000000000000
	MAX = 400000000000000
)

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	var hailStones []Hail
	for _, line := range lines {
		hailStones = append(hailStones, parseHail(line))
	}
	fmt.Println(solve(hailStones))
}
