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

func inAnyRange(j int, ranges [][2]int) bool {
	for _, r := range ranges {
		if inRange(j, r) {
			return true
		}
	}
	return false
}

func inRange(j int, r [2]int) bool {
	return j >= r[0] && j <= r[1]
}

func rangeMerge(newRange [2]int, ranges [][2]int) [][2]int {
	merged := false
	for i, r := range ranges {
		if inRange(newRange[0], r) && !inRange(newRange[1], r) {
			ranges[i][1] = newRange[1]
			merged = true
		}
		if inRange(newRange[1], r) && !inRange(newRange[0], r) {
			ranges[i][0] = newRange[0]
			merged = true
		}
	}
	if !merged {
		ranges = append(ranges, newRange)
	}
	return ranges
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

func solve(readings [][2]coords2d.Coords2d, maxCoord int) int {
	allRanges := make([][][2]int, maxCoord+1)
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
		allRanges[i] = mergedRanges
	}
	for i := 0; i <= maxCoord; i++ {
		if len(allRanges[i]) > 1 {
			for j := 0; j <= maxCoord; j++ {
				if !inAnyRange(j, allRanges[i]) {
					return (j*TUNING_OFFSET + i)
				}
			}
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
