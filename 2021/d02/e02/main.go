package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"fmt"
	"os"
)

type AimCoords struct {
	coords2d.Coords2d
	Aim int
}

func solve(lines []string) int {
	c := coords2d.Coords2d{X: 0, Y: 0}
	pos := AimCoords{Coords2d: c, Aim: 0}
	for _, l := range lines {
		var action string
		var val int
		fmt.Sscanf(l, "%s %d", &action, &val)
		switch action {
		case "down":
			pos.Aim += val
		case "up":
			pos.Aim -= val
		case "forward":
			pos.X += val
			pos.Y += val * pos.Aim
		}
	}
	return pos.X * pos.Y
}

func main() {
	lines := files.ReadLines(os.Args[1])
	fmt.Println(solve(lines))
}
