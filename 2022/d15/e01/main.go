package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"adventofcode/utils/sets/intset"
	"fmt"
	"os"
)

func amplitude(readings [][2]coords2d.Coords2d) (int, int) {
	minX, maxX := readings[0][0].X, readings[0][0].X
	for _, r := range readings {
		if r[0].X < minX {
			minX = r[0].X
		}
		if r[1].X < minX {
			minX = r[1].X
		}
		if r[0].X > maxX {
			maxX = r[0].X
		}
		if r[1].X > maxX {
			maxX = r[1].X
		}
	}
	return minX * 2, maxX * 2
}

func solve(readings [][2]coords2d.Coords2d, targetRow int) int {
	res := intset.New()
	minX, maxX := amplitude(readings)
	for _, reading := range readings {
		dist := coords2d.Distance(reading[0], reading[1])
		for i := minX; i <= maxX; i++ {
			if coords2d.Distance(reading[0], coords2d.Coords2d{X: i, Y: targetRow}) <= dist {
				res.Add(i)
			}
		}
	}
	for _, reading := range readings {
		if reading[1].Y == targetRow {
			res.Remove(reading[1].X)
		}
	}
	return res.Len()
}

func main() {
	data := files.ReadLines(os.Args[1])
	bs := make([][2]coords2d.Coords2d, len(data))
	var sx, sy, bx, by int
	for i, line := range data {
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		bs[i] = [2]coords2d.Coords2d{{X: sx, Y: sy}, {X: bx, Y: by}}
	}
	fmt.Println(solve(bs, 2000000))
}
