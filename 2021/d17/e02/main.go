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
	hits := 0
	for y := ymin * 10; y < 100; y++ {
		for x := 1; x < xmax*10; x++ {
			pos := coords2d.Coords2d{X: 0, Y: 0}
			velocity := coords2d.Coords2d{X: x, Y: y}
			for i := 0; i < 1000; i++ {
				pos = coords2d.Add(pos, velocity)
				if posInTarget(pos, xmin, xmax, ymin, ymax) {
					hits++
					break
				}
				if velocity.X > 0 {
					velocity.X--
				} else if velocity.X < 0 {
					velocity.X++
				}
				velocity.Y--
			}
		}
	}
	return hits
}

func main() {
	data := files.ReadFile(os.Args[1])
	var xmin, xmax, ymin, ymax int
	fmt.Sscanf(string(data), "target area: x=%d..%d, y=%d..%d", &xmin, &xmax, &ymin, &ymax)
	fmt.Printf("target area: x=%d..%d, y=%d..%d\n", xmin, xmax, ymin, ymax)
	fmt.Println(solve(xmin, xmax, ymin, ymax))
}
