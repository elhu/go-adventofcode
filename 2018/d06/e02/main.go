package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type coord struct {
	x, y int
}

func parseCoords(path string) map[int]coord {
	input, err := ioutil.ReadFile(path)
	check(err)
	input = bytes.TrimSuffix(input, []byte{'\n'})
	lines := bytes.Split(input, []byte{'\n'})
	coords := make(map[int]coord)
	for i, l := range lines {
		parts := bytes.Split(l, []byte{',', ' '})
		x, _ := strconv.Atoi(string(parts[0]))
		y, _ := strconv.Atoi(string(parts[1]))
		coords[i] = coord{x, y}
	}
	return coords
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func distance(a, b coord) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

const maxInt = int(^uint(0) >> 1)

func sumDistances(a coord, points map[int]coord) int {
	total := 0
	for _, b := range points {
		total += distance(a, b)
	}
	return total
}

func getMaxCoords(coords map[int]coord) coord {
	maxX, maxY := 0, 0
	for _, c := range coords {
		if c.x > maxX {
			maxX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}
	}
	return coord{maxX, maxY}
}

func getMinCoords(coords map[int]coord) coord {
	minX, minY := maxInt, maxInt
	for _, c := range coords {
		if c.x < minX {
			minX = c.x
		}
		if c.y < minY {
			minY = c.y
		}
	}
	return coord{minX, minY}
}

func isBorder(current, max, min coord) bool {
	return current.x == min.x || current.x == max.x || current.y == min.y || current.y == max.y
}

const threshold = 10000

func main() {
	coords := parseCoords(os.Args[1])
	maxCoords := getMaxCoords(coords)
	minCoords := getMinCoords(coords)
	totalSize := 0

	for i := minCoords.y; i <= maxCoords.y; i++ {
		for j := minCoords.x; j <= maxCoords.x; j++ {
			current := coord{j, i}
			totalDistance := sumDistances(current, coords)
			if totalDistance < threshold {
				totalSize++
			}
		}
	}
	fmt.Println(totalSize)
}
