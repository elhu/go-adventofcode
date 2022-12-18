package main

import (
	"adventofcode/utils/coords/coords3d"
	"adventofcode/utils/files"
	"fmt"
	"os"
)

const (
	AIR   = 0
	LAVA  = 1
	STEAM = 2
)

func parseCubes(data []string) [][][]int {
	cubes := make([]coords3d.Coords3d, len(data))
	var maxX, maxY, maxZ int
	for i, d := range data {
		cube := coords3d.Coords3d{}
		fmt.Sscanf(d, "%d,%d,%d", &cube.X, &cube.Y, &cube.Z)
		if cube.X > maxX {
			maxX = cube.X
		}
		if cube.Y > maxY {
			maxY = cube.Y
		}
		if cube.Z > maxZ {
			maxZ = cube.Z
		}
		cubes[i] = cube
	}
	// expand map of the area by 3
	// offset everything by 2 to create room for the steam to go in
	scan := make([][][]int, maxX+1+3)
	for i := 0; i < len(scan); i++ {
		scan[i] = make([][]int, maxY+1+3)
		for j := 0; j < len(scan[i]); j++ {
			scan[i][j] = make([]int, maxZ+1+3)
		}
	}
	for _, c := range cubes {
		scan[c.X+2][c.Y+2][c.Z+2] = LAVA
	}
	return scan
}

var adjacentCoord = []coords3d.Coords3d{
	{X: -1, Y: 0, Z: 0},
	{X: 1, Y: 0, Z: 0},
	{X: 0, Y: -1, Z: 0},
	{X: 0, Y: 1, Z: 0},
	{X: 0, Y: 0, Z: -1},
	{X: 0, Y: 0, Z: 1},
}

func adjacentCubes(cubes [][][]int, x, y, z int) int {
	c := coords3d.Coords3d{X: x, Y: y, Z: z}
	res := 0
	for _, av := range adjacentCoord {
		a := coords3d.Add(c, av)
		if a.X >= 0 && a.X < len(cubes) &&
			a.Y >= 0 && a.Y < len(cubes[x]) &&
			a.Z >= 0 && a.Z < len(cubes[x][y]) {
			if cubes[a.X][a.Y][a.Z] == LAVA {
				res++
			}
		}
	}
	return res
}

func solve(cubes [][][]int) int {
	res := 0
	queue := []coords3d.Coords3d{{X: 0, Y: 0, Z: 0}}
	var head coords3d.Coords3d
	cubes[0][0][0] = STEAM
	for len(queue) > 0 {
		head, queue = queue[0], queue[1:]
		for _, av := range adjacentCoord {
			a := coords3d.Add(head, av)
			if a.X >= 0 && a.X < len(cubes) &&
				a.Y >= 0 && a.Y < len(cubes[0]) &&
				a.Z >= 0 && a.Z < len(cubes[0][0]) {
				if cubes[a.X][a.Y][a.Z] == LAVA {
					res++
				} else if cubes[a.X][a.Y][a.Z] == AIR {
					cubes[a.X][a.Y][a.Z] = STEAM
					queue = append(queue, a)
				}
			}
		}
	}
	return res
}

func main() {
	data := files.ReadLines(os.Args[1])
	cubes := parseCubes(data)
	fmt.Println(solve(cubes))
}
