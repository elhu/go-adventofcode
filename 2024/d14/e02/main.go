package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"fmt"
	"image"
	"image/draw"
	"image/png"
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

func printBots(bots []Bot) {
	grid := make([][]byte, HEIGHT)
	for i := range grid {
		grid[i] = []byte(strings.Repeat(".", WIDTH))
	}
	for _, b := range bots {
		grid[b.pos.Y][b.pos.X] = '#'
	}
	for i := range grid {
		fmt.Println(string(grid[i]))
	}
	fmt.Println("\n\n")
}

func solve(bots []Bot) int {
	for t := 0; t < 10000; t++ {
		img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
		draw.Draw(img, img.Bounds(), &image.Uniform{image.Black}, image.ZP, draw.Src)
		for i := range bots {
			bots[i].pos = coords2d.Add(bots[i].pos, bots[i].vec)
			bots[i].pos.X = (bots[i].pos.X + WIDTH) % WIDTH
			bots[i].pos.Y = (bots[i].pos.Y + HEIGHT) % HEIGHT
			img.Set(bots[i].pos.X, bots[i].pos.Y, image.White)
		}
		f, _ := os.Create("out/" + fmt.Sprintf("%06d", t) + ".png")
		defer f.Close()
		png.Encode(f, img)
	}
	// Now use your eyes
	return 0
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
