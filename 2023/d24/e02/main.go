package main

import (
	"adventofcode/utils/coords/coords3d"
	"adventofcode/utils/files"
	"fmt"
	"math"
	"os"
	"strings"
)

type Hail struct {
	pos coords3d.Coords3d
	vec coords3d.Coords3d
}

func parseHail(line string) Hail {
	var h Hail
	fmt.Sscanf(line, "%d, %d, %d @ %d, %d, %d", &h.pos.X, &h.pos.Y, &h.pos.Z, &h.vec.X, &h.vec.Y, &h.vec.Z)
	return h
}

func toMatrix(hailstones []Hail, fn func(h Hail) (int, int, int, int, int)) [][5]float64 {
	var res [][5]int
	for _, h := range hailstones {
		a, b, c, d, e := fn(h)
		res = append(res, [5]int{a, b, c, d, e})
	}
	// At this point, matrix is in the form:
	// [-dy, dx, y, -x, y*dx - x*dy]
	last := res[len(res)-1]
	var rr [][5]float64
	for _, row := range res[0:4] {
		rr = append(rr, [5]float64{
			float64(row[0] - last[0]), float64(row[1] - last[1]),
			float64(row[2] - last[2]), float64(row[3] - last[3]),
			float64(row[4] - last[4]),
		})
	}
	// After this transformation, matrix is in the form:
	// (' denotes the last row of the matrix, arbitrarily chosen)
	// [-dy - -dy', dx - dx', y - y', -x - -x', (y*dx - x*dy) - (y'*dx' - x'*dy')]
	return rr
}

func gaussianElimination(mat [][5]float64) [][5]float64 {
	for i := 0; i < len(mat); i++ {
		factor := mat[i][i]
		for j := 0; j < len(mat[i]); j++ {
			mat[i][j] /= factor
		}
		for j := i + 1; j < len(mat); j++ {
			factor = mat[j][i]
			for k := 0; k < len(mat[j]); k++ {
				mat[j][k] -= factor * mat[i][k]
			}
		}
	}
	for i := len(mat) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			factor := mat[j][i]
			for k := 0; k < len(mat[j]); k++ {
				mat[j][k] -= factor * mat[i][k]
			}
		}
	}
	return mat
}

func solve3(hailStones []Hail) int {
	getXY := func(h Hail) (int, int, int, int, int) {
		return -h.vec.Y, h.vec.X, h.pos.Y, -h.pos.X, h.pos.Y*h.vec.X - h.pos.X*h.vec.Y
	}
	getZY := func(h Hail) (int, int, int, int, int) {
		return -h.vec.Y, h.vec.Z, h.pos.Y, -h.pos.Z, h.pos.Y*h.vec.Z - h.pos.Z*h.vec.Y
	}
	xyMat := toMatrix(hailStones, getXY)
	xyMat = gaussianElimination(xyMat)
	zyMat := toMatrix(hailStones, getZY)
	zyMat = gaussianElimination(zyMat)
	x := int(math.Round(xyMat[0][len(xyMat[0])-1]))
	y := int(math.Round(xyMat[1][len(xyMat[0])-1]))
	z := int(math.Round(zyMat[0][len(zyMat[0])-1]))
	return x + y + z
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	var hailStones []Hail
	for _, line := range lines {
		hailStones = append(hailStones, parseHail(line))
	}
	// fmt.Println(solve(hailStones))
	// fmt.Println(solve2(hailStones))
	fmt.Println(solve3(hailStones))
}
