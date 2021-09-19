package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Coord struct {
	x, y int
}

type LightField3d struct {
	data [][]int
}

func new(x, y int) LightField3d {
	data := make([][]int, y)
	for i := 0; i < y; i++ {
		data[i] = make([]int, x)
	}
	return LightField3d{data}
}

func (b *LightField3d) set(tl, br Coord, val int) {
	for i := tl.y; i <= br.y; i++ {
		for j := tl.x; j <= br.x; j++ {
			b.data[i][j] += val
			if b.data[i][j] < 0 {
				b.data[i][j] = 0
			}
		}
	}
}

func (b *LightField3d) countBrightness() int {
	res := 0
	for _, l := range b.data {
		for _, c := range l {
			res += c
		}
	}
	return res
}

func solve(input []string) int {
	b := new(1000, 1000)
	for _, s := range input {
		var tl, br Coord
		if strings.HasPrefix(s, "turn on") {
			fmt.Sscanf(s, "turn on %d,%d through %d,%d", &tl.x, &tl.y, &br.x, &br.y)
			b.set(tl, br, 1)
		} else if strings.HasPrefix(s, "turn off") {
			fmt.Sscanf(s, "turn off %d,%d through %d,%d", &tl.x, &tl.y, &br.x, &br.y)
			b.set(tl, br, -1)
		} else {
			fmt.Sscanf(s, "toggle %d,%d through %d,%d", &tl.x, &tl.y, &br.x, &br.y)
			b.set(tl, br, 2)
		}
	}
	return b.countBrightness()
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	fmt.Println(solve(input))
}
