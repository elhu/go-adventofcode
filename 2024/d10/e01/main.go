package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"adventofcode/utils/sets"
	"fmt"
	"os"
	"strings"
)

func findZeros(topo [][]int) []coords2d.Coords2d {
	var z []coords2d.Coords2d
	for y, l := range topo {
		for x, h := range l {
			if h == 0 {
				z = append(z, coords2d.Coords2d{X: x, Y: y})
			}
		}
	}
	return z
}

func inBounds(c coords2d.Coords2d, topo [][]int) bool {
	return c.X >= 0 && c.X < len(topo[0]) && c.Y >= 0 && c.Y < len(topo)
}

func bfs(topo [][]int, z coords2d.Coords2d) int {
	visited := sets.New[coords2d.Coords2d]()
	toVisit := []coords2d.Coords2d{z}
	nines := sets.New[coords2d.Coords2d]()
	var curr coords2d.Coords2d
	for len(toVisit) > 0 {
		curr, toVisit = toVisit[0], toVisit[1:]
		if topo[curr.Y][curr.X] == 9 {
			nines.Add(curr)
		}
		visited.Add(curr)
		for _, v := range []coords2d.Coords2d{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			newPos := coords2d.Add(curr, v)
			if !visited.HasMember(newPos) && inBounds(newPos, topo) && topo[newPos.Y][newPos.X] == topo[curr.Y][curr.X]+1 {
				toVisit = append(toVisit, newPos)
			}
		}
	}
	return nines.Len()
}

func solve(topo [][]int) int {
	res := 0
	for _, z := range findZeros(topo) {
		res += bfs(topo, z)
	}
	return res
}
func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	topo := make([][]int, len(lines))
	for i, line := range lines {
		topo[i] = make([]int, len(line))
		for j, c := range line {
			topo[i][j] = int(c - '0')
		}
	}
	fmt.Println(solve(topo))
}
