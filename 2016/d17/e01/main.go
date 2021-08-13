package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"strings"
)

func computeHash(passcode string, path []byte) string {
	toHash := []byte(fmt.Sprintf("%s%s", passcode, path))
	digest := fmt.Sprintf("%x", md5.Sum(toHash))
	return digest
}

type Coord struct {
	x, y int
}

type Candidate struct {
	coord Coord
	path  []byte
}

func coordEqual(a, b Coord) bool {
	return a.x == b.x && a.y == b.y
}

func outOfBounds(c Coord) bool {
	return c.x < 0 || c.x > 3 || c.y < 0 || c.y > 3
}

var directions = []byte("UDLR")

func neighbors(current Candidate, passcode string) []Candidate {
	var res []Candidate
	hash := computeHash(passcode, current.path)
	for i, m := range [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
		coord := Coord{current.coord.x + m[0], current.coord.y + m[1]}
		if !outOfBounds(coord) && strings.ContainsAny(hash[i:i+1], "bcdef") {
			path := make([]byte, len(current.path)+1)
			copy(path, current.path)
			path[len(current.path)] = directions[i]
			res = append(res, Candidate{coord, path})
		}
	}
	return res
}

func solve(passcode string) []byte {
	target := Coord{3, 3}
	queue := []Candidate{{Coord{0, 0}, []byte{}}}

	var current Candidate
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		if coordEqual(current.coord, target) {
			return current.path
		}
		queue = append(queue, neighbors(current, passcode)...)
	}
	panic("No path found")
}

func main() {
	passcode := os.Args[1]
	fmt.Println(string(solve(passcode)))
}
