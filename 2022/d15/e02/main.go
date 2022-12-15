package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func solve(readings [][2]coords2d.Coords2d, maxCoord int) int {
	// res := intset.New()
	for _, reading := range readings {
		dist := coords2d.Distance(reading[0], reading[1])
		for i := 0; i <= maxCoord; i++ {
			dy := abs(reading[0].Y - i)
			dxAbs := dist - dy
			dxLow := reading[0].X - dxAbs
			dxHi := reading[0].X - -dxAbs
			if i == 11 {
				fmt.Println(dxLow, dxHi)
			}
			// res.Add(i*TUNING_OFFSET + j)
		}
	}
	// for i := 0; i <= maxCoord; i++ {
	// 	for j := 0; j <= maxCoord; j++ {
	// 		if !res.HasMember(i*TUNING_OFFSET + j) {
	// 			return i*TUNING_OFFSET + j
	// 		}
	// 	}
	// }
	panic("Couldn't find beacon")
}

const TUNING_OFFSET = 4000000

func main() {
	data := files.ReadLines(os.Args[1])
	bs := make([][2]coords2d.Coords2d, len(data))
	var sx, sy, bx, by int
	for i, line := range data {
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		bs[i] = [2]coords2d.Coords2d{{X: sx, Y: sy}, {X: bx, Y: by}}
	}
	fmt.Println(solve(bs, 20))
}
