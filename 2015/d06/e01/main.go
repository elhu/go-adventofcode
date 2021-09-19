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

type BitField3d struct {
	data [][]bool
}

func new(x, y int) BitField3d {
	data := make([][]bool, y)
	for i := 0; i < y; i++ {
		data[i] = make([]bool, x)
	}
	return BitField3d{data}
}

func (b *BitField3d) toggle(tl, br Coord) {
	for i := tl.y; i <= br.y; i++ {
		for j := tl.x; j <= br.x; j++ {
			b.data[i][j] = !b.data[i][j]
		}
	}
}

func (b *BitField3d) set(tl, br Coord, val bool) {
	for i := tl.y; i <= br.y; i++ {
		for j := tl.x; j <= br.x; j++ {
			b.data[i][j] = val
		}
	}
}

func (b *BitField3d) countOnes() int {
	res := 0
	for _, l := range b.data {
		for _, c := range l {
			if c {
				res++
			}
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
			b.set(tl, br, true)
		} else if strings.HasPrefix(s, "turn off") {
			fmt.Sscanf(s, "turn off %d,%d through %d,%d", &tl.x, &tl.y, &br.x, &br.y)
			b.set(tl, br, false)
		} else {
			fmt.Sscanf(s, "toggle %d,%d through %d,%d", &tl.x, &tl.y, &br.x, &br.y)
			b.toggle(tl, br)
		}
	}
	return b.countOnes()
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	fmt.Println(solve(input))
}
