package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

var ADJ = [4]coords2d.Coords2d{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func processRegion(pt map[coords2d.Coords2d]byte, pm map[coords2d.Coords2d]bool, start coords2d.Coords2d) int {
	var area, perimeter int
	queue := []coords2d.Coords2d{start}
	var curr coords2d.Coords2d
	rt := pt[start]
	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		if pm[curr] {
			continue
		}
		pm[curr] = true
		area++
		for _, v := range ADJ {
			if pt[coords2d.Add(curr, v)] == rt {
				queue = append(queue, coords2d.Add(curr, v))
			} else {
				perimeter++
			}
		}
	}
	return area * perimeter
}

func solve(data []string) int {
	pm := make(map[coords2d.Coords2d]bool)
	pt := make(map[coords2d.Coords2d]byte)
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			pm[coords2d.Coords2d{X: j, Y: i}] = false
			pt[coords2d.Coords2d{X: j, Y: i}] = data[i][j]
		}
	}
	res := 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if !pm[coords2d.Coords2d{X: j, Y: i}] {
				res += processRegion(pt, pm, coords2d.Coords2d{X: j, Y: i})
			}
		}
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	raw := strings.Split(data, "\n")
	fmt.Println(solve(raw))
}
