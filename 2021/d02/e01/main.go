package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"fmt"
	"os"
)

func solve(lines []string) int {
	pos := coords2d.Coords2d{0, 0}
	for _, l := range lines {
		var action string
		var val int
		fmt.Sscanf(l, "%s %d", &action, &val)
		switch action {
		case "down":
			pos.Y += val
		case "up":
			pos.Y -= val
		case "forward":
			pos.X += val
		}
	}
	return pos.X * pos.Y
}

func main() {
	lines := files.ReadLines(os.Args[1])
	fmt.Println(solve(lines))
}
