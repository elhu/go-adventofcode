package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Coord struct {
	x, y int
}

func toKey(c Coord) int {
	return c.y*100000000 + c.x
}

func (a *Coord) apply(v Coord) {
	a.x += v.x
	a.y += v.y
}

var vectors = map[byte]Coord{
	'^': {0, -1},
	'v': {0, 1},
	'>': {1, 0},
	'<': {-1, 0},
}

func solve(data []byte) int {
	seen := make(map[int]struct{})
	curr := Coord{0, 0}
	seen[toKey(curr)] = struct{}{}
	for _, b := range data {
		curr.apply(vectors[b])
		seen[toKey(curr)] = struct{}{}
	}

	return len(seen)
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	fmt.Println(solve(data))
}
