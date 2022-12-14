package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

const MAP_SIZE = 100000000

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func buildMap(data []string) (map[int]byte, int) {
	area := make(map[int]byte)
	maxY := 0
	var startX, startY, endX, endY int
	for _, path := range data {
		parts := strings.Split(path, " -> ")
		for i := 0; i < len(parts)-1; i++ {
			fmt.Sscanf(parts[i], "%d,%d", &startX, &startY)
			fmt.Sscanf(parts[i+1], "%d,%d", &endX, &endY)
			if startY > maxY {
				maxY = startY
			}
			if endY > maxY {
				maxY = endY
			}
			if startX == endX {
				s, e := minMax(startY, endY)
				for ; s <= e; s++ {
					area[MAP_SIZE*s+startX] = '#'
				}
			} else {
				s, e := minMax(startX, endX)
				for ; s <= e; s++ {
					area[MAP_SIZE*startY+s] = '#'
				}
			}
		}
	}
	return area, maxY
}

func solve(area map[int]byte, maxY int) int {
	atRest := 0
	for {
		sandX, sandY := 500, 0
		for {
			if _, exists := area[(sandY+1)*MAP_SIZE+sandX]; !exists {
				sandY++
				if sandY > maxY+1 {
					return atRest
				}
				continue
			}
			if _, exists := area[(sandY+1)*MAP_SIZE+sandX-1]; !exists {
				sandY++
				sandX--
				if sandY > maxY+1 {
					return atRest
				}
				continue
			}
			if _, exists := area[(sandY+1)*MAP_SIZE+sandX+1]; !exists {
				sandY++
				sandX++
				if sandY > maxY+1 {
					return atRest
				}
				continue
			}
			area[sandY*MAP_SIZE+sandX] = 'o'
			atRest++
			break
		}
	}
}

func main() {
	data := files.ReadLines(os.Args[1])
	area, maxY := buildMap(data)
	fmt.Println(solve(area, maxY))
}
