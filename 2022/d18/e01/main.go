package main

import (
	"adventofcode/utils/coords/coords3d"
	"adventofcode/utils/files"
	"fmt"
	"os"
)

func parseCubes(data []string) [][][]bool {
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
	scan := make([][][]bool, maxX+1)
	for i := 0; i < len(scan); i++ {
		scan[i] = make([][]bool, maxY+1)
		for j := 0; j < len(scan[i]); j++ {
			scan[i][j] = make([]bool, maxZ+1)
		}
	}
	for _, c := range cubes {
		scan[c.X][c.Y][c.Z] = true
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

func adjacentCubes(cubes [][][]bool, x, y, z int) int {
	c := coords3d.Coords3d{X: x, Y: y, Z: z}
	res := 0
	for _, av := range adjacentCoord {
		a := coords3d.Add(c, av)
		if a.X >= 0 && a.X < len(cubes) &&
			a.Y >= 0 && a.Y < len(cubes[x]) &&
			a.Z >= 0 && a.Z < len(cubes[x][y]) {
			if cubes[a.X][a.Y][a.Z] {
				res++
			}
		}
	}
	return res
}

func solve(cubes [][][]bool) int {
	res := 0
	for x := 0; x < len(cubes); x++ {
		for y := 0; y < len(cubes[x]); y++ {
			for z := 0; z < len(cubes[x][y]); z++ {
				if cubes[x][y][z] {
					res += 6 - adjacentCubes(cubes, x, y, z)
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
