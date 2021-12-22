package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
)

type Cuboid struct {
	x, y, z [2]int
	charge  int
}

func newCuboid(minX, maxX, minY, maxY, minZ, maxZ, charge int) Cuboid {
	return Cuboid{
		[2]int{minX, maxX},
		[2]int{minY, maxY},
		[2]int{minZ, maxZ},
		charge,
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func intersection(a, b Cuboid) (bool, Cuboid) {
	xmin := max(a.x[0], b.x[0])
	xmax := min(a.x[1], b.x[1])
	ymin := max(a.y[0], b.y[0])
	ymax := min(a.y[1], b.y[1])
	zmin := max(a.z[0], b.z[0])
	zmax := min(a.z[1], b.z[1])
	if xmin <= xmax && ymin <= ymax && zmin <= zmax {
		return true, newCuboid(xmin, xmax, ymin, ymax, zmin, zmax, -a.charge)
	}
	return false, Cuboid{}
}

func size(c Cuboid) int {
	return (c.x[1] - c.x[0] + 1) * (c.y[1] - c.y[0] + 1) * (c.z[1] - c.z[0] + 1)
}

func solve(input []string) int {
	cuboids := make([]Cuboid, 0)
	for _, l := range input {
		var inst string
		var minX, maxX, minY, maxY, minZ, maxZ int
		fmt.Sscanf(l, "%s x=%d..%d,y=%d..%d,z=%d..%d", &inst, &minX, &maxX, &minY, &maxY, &minZ, &maxZ)
		nc := newCuboid(minX, maxX, minY, maxY, minZ, maxZ, 1)

		for _, c := range cuboids {
			if collides, ic := intersection(c, nc); collides {
				cuboids = append(cuboids, ic)
			}
		}
		if inst == "on" {
			cuboids = append(cuboids, nc)
		}
	}
	res := 0
	for _, c := range cuboids {
		res += size(c) * c.charge
	}
	return res
}

func main() {
	data := files.ReadLines(os.Args[1])
	fmt.Println(solve(data))
}
