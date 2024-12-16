package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

type Bot struct {
	pos coords2d.Coords2d
	vec coords2d.Coords2d
}

func parseBot(data string) Bot {
	var b Bot
	fmt.Sscanf(data, "p=%d,%d v=%d,%d", &b.pos.X, &b.pos.Y, &b.vec.X, &b.vec.Y)
	return b
}

// const WIDTH = 11
// const HEIGHT = 7
const WIDTH = 101
const HEIGHT = 103
const SECONDS = 100

func solve(bots []Bot) int {
	for s := 0; s < SECONDS; s++ {
		for i := range bots {
			bots[i].pos = coords2d.Add(bots[i].pos, bots[i].vec)
			bots[i].pos.X = (bots[i].pos.X + WIDTH) % WIDTH
			bots[i].pos.Y = (bots[i].pos.Y + HEIGHT) % HEIGHT
		}
	}
	quadCount := [4]int{0, 0, 0, 0}
	for _, b := range bots {
		if b.pos.X < WIDTH/2 && b.pos.Y < HEIGHT/2 {
			quadCount[0]++
		} else if b.pos.X > WIDTH/2 && b.pos.Y < HEIGHT/2 {
			quadCount[1]++
		} else if b.pos.X < WIDTH/2 && b.pos.Y > HEIGHT/2 {
			quadCount[2]++
		} else if b.pos.X > WIDTH/2 && b.pos.Y > HEIGHT/2 {
			quadCount[3]++
		}
	}
	return quadCount[0] * quadCount[1] * quadCount[2] * quadCount[3]
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	var bots []Bot
	for _, l := range lines {
		b := parseBot(l)
		bots = append(bots, b)
	}
	fmt.Println(solve(bots))
}
