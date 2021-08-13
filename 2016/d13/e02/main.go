package main

import (
	"fmt"
	"math/bits"
)

const targetX = 31
const targetY = 39

type Coord struct {
	x, y int
}

func coordHash(c Coord) int {
	return c.x*1000000 + c.y
}

func isOpen(c Coord) bool {
	if c.x < 0 || c.y < 0 {
		return false
	}
	val := uint(c.x*c.x + 3*c.x + 2*c.x*c.y + c.y + c.y*c.y + magicNumber)
	return bits.OnesCount(val)%2 == 0
}

func neighbors(c Coord) []Coord {
	return []Coord{
		{c.x - 1, c.y},
		{c.x + 1, c.y},
		{c.x, c.y - 1},
		{c.x, c.y + 1},
	}
}

const magicNumber = 1358

func main() {
	start := Coord{1, 1}
	seen := map[int]struct{}{coordHash(start): {}}
	nextQueue := []Coord{start}
	var stepCount = -1

	// Stop at 49 because we flag the tiles for round n+1 as seen in round n
	for stepCount < 49 {
		queue := make([]Coord, len(nextQueue))
		copy(queue, nextQueue)
		nextQueue = nil
		stepCount++

		for _, c := range queue {
			for _, n := range neighbors(c) {
				if _, found := seen[coordHash(n)]; !found && isOpen(n) {
					seen[coordHash(n)] = struct{}{}
					nextQueue = append(nextQueue, n)
				}
			}
		}
	}
	fmt.Println(len(seen))
}
