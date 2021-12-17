package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"fmt"
	"os"
)

func posInTarget(pos coords2d.Coords2d, xmin, xmax, ymin, ymax int) bool {
	return pos.X >= xmin && pos.X <= xmax && pos.Y >= ymin && pos.Y <= ymax
}

func solve(xmin, xmax, ymin, ymax int) int {
	maxY := 0
	for y := 1; y < 100; y++ {
		for x := 1; x < xmax; x++ {
			localMaxY := 0
			targetHit := false
			pos := coords2d.Coords2d{X: 0, Y: 0}
			velocity := coords2d.Coords2d{X: x, Y: y}
			for pos.X < xmax && pos.Y > ymax {
				pos = coords2d.Add(pos, velocity)
				if pos.Y >= localMaxY {
					localMaxY = pos.Y
				}
				if posInTarget(pos, xmin, xmax, ymin, ymax) {
					targetHit = true
					break
				}
				if velocity.X > 0 {
					velocity.X--
				} else if velocity.X < 0 {
					velocity.X++
				}
				velocity.Y--
			}
			if targetHit && localMaxY > maxY {
				maxY = localMaxY
			}
		}
	}
	return maxY
}

func main() {
	data := files.ReadFile(os.Args[1])
	var xmin, xmax, ymin, ymax int
	fmt.Sscanf(string(data), "target area: x=%d..%d, y=%d..%d", &xmin, &xmax, &ymin, &ymax)
	fmt.Printf("target area: x=%d..%d, y=%d..%d\n", xmin, xmax, ymin, ymax)
	fmt.Println(solve(xmin, xmax, ymin, ymax))
}
