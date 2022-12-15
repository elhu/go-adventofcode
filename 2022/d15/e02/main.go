package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"fmt"
	"os"
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func inRange(j int, r [2]int) bool {
	return j >= r[0] && j <= r[1]
}

func remove(slice [][2]int, i int) [][2]int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func mergeRanges(ranges [][2]int) [][2]int {
	reduced := true
	for reduced {
		reduced = false
		for i := 0; i < len(ranges); i++ {
			for j := 0; j < len(ranges); j++ {
				if i == j {
					continue
				}
				if inRange(ranges[i][0], ranges[j]) && !inRange(ranges[i][1], ranges[j]) {
					ranges[j][1] = ranges[i][1]
					ranges = remove(ranges, i)
					reduced = true
					break
				}
				if inRange(ranges[i][1], ranges[j]) && !inRange(ranges[i][0], ranges[j]) {
					ranges[j][0] = ranges[i][0]
					ranges = remove(ranges, i)
					reduced = true
					break
				}
				if inRange(ranges[i][0], ranges[j]) && inRange(ranges[i][1], ranges[j]) {
					ranges = remove(ranges, i)
					reduced = true
					break
				}
			}
		}
		if reduced {
			continue
		}
	}
	return ranges
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solve(readings [][2]coords2d.Coords2d, maxCoord int) int {
	for i := 0; i <= maxCoord; i++ {
		ranges := make([][2]int, 0)
		for _, reading := range readings {
			dist := coords2d.Distance(reading[0], reading[1])
			dy := abs(reading[0].Y - i)
			dxAbs := dist - dy
			dxLow := reading[0].X - dxAbs
			dxHi := reading[0].X - -dxAbs
			if dxLow < dxHi {
				ranges = append(ranges, [2]int{dxLow, dxHi})
			}
		}
		mergedRanges := mergeRanges(ranges)
		if len(mergedRanges) == 2 {
			return (min(mergedRanges[1][1], mergedRanges[0][1])+1)*TUNING_OFFSET + i
		}
	}
	panic("WTF")
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
	fmt.Println(solve(bs, 4000000))
}
