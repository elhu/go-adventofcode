package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

const MAP_SIZE = 1000000

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
	for i := -MAP_SIZE; i < MAP_SIZE; i++ {
		area[(maxY+3)*MAP_SIZE+i] = '#'
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
				continue
			}
			if _, exists := area[(sandY+1)*MAP_SIZE+sandX-1]; !exists {
				sandY++
				sandX--
				continue
			}
			if _, exists := area[(sandY+1)*MAP_SIZE+sandX+1]; !exists {
				sandY++
				sandX++
				continue
			}
			area[sandY*MAP_SIZE+sandX] = 'o'
			atRest++
			if sandX == 500 && sandY == 0 {
				return atRest
			}
			break
		}
	}
}

func main() {
	data := files.ReadLines(os.Args[1])
	area, maxY := buildMap(data)
	fmt.Println(solve(area, maxY))
}
