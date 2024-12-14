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

const WIDTH = 101
const HEIGHT = 103

func variance(vals []int) int {
	avg := 0
	for _, v := range vals {
		avg += v
	}
	avg /= len(vals)
	variance := 0
	for _, v := range vals {
		variance += (v - avg) * (v - avg)
	}
	return variance
}

func solve(bots []Bot) int {
	minVariance := 1000000000
	bestTime := -1
	for t := 0; t < 10000; t++ {
		xs := make([]int, len(bots))
		ys := make([]int, len(bots))
		for i := range bots {
			bots[i].pos = coords2d.Add(bots[i].pos, bots[i].vec)
			bots[i].pos.X = (bots[i].pos.X + WIDTH) % WIDTH
			bots[i].pos.Y = (bots[i].pos.Y + HEIGHT) % HEIGHT
			xs[i] = bots[i].pos.X
			ys[i] = bots[i].pos.Y
		}
		// Since the tree is a tight cluster, this is when the variance is the lowest
		v := variance(xs) + variance(ys)
		if v < minVariance {
			minVariance = v
			bestTime = t + 1
		}
	}
	return bestTime
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
