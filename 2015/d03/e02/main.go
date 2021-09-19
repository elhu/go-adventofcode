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
	currRobot := Coord{0, 0}
	currSanta := Coord{0, 0}
	seen[toKey(currRobot)] = struct{}{}
	for i, b := range data {
		if i%2 == 0 {
			currSanta.apply(vectors[b])
			seen[toKey(currSanta)] = struct{}{}
		} else {
			currRobot.apply(vectors[b])
			seen[toKey(currRobot)] = struct{}{}
		}
	}

	return len(seen)
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	fmt.Println(solve(data))
}
