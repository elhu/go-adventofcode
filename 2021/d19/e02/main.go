package main

import (
	"adventofcode/utils/coords/coords3d"
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

func vector(a, b coords3d.Coords3d) string {
	return fmt.Sprintf("%d:%d:%d", b.X-a.X, b.Y-a.Y, b.Z-a.Z)
}

func buildVectors(coords []coords3d.Coords3d) map[string][]coords3d.Coords3d {
	res := make(map[string][]coords3d.Coords3d)
	for i := range coords {
		for _, c := range coords[i+1:] {
			res[vector(coords[i], c)] = []coords3d.Coords3d{coords[i], c}
			res[vector(c, coords[i])] = []coords3d.Coords3d{c, coords[i]}
		}
	}
	return res
}

func matchingVectors(a, b map[string][]coords3d.Coords3d) map[string]struct{} {
	res := make(map[string]struct{})
	for ka := range a {
		if _, found := b[ka]; found {
			res[ka] = struct{}{}
		}
	}
	return res
}

func buildVariants(coords []coords3d.Coords3d) [][]coords3d.Coords3d {
	variants := make([][]coords3d.Coords3d, 24)
	// {+x,+y,+z},           {+y,+z,+x},            {+z,+x,+y},            {+z,+y,-x},            {+y,+x,-z},            {+x,+z,-y},
	// {+x,-y,-z},           {+y,-z,-x},            {+z,-x,-y},            {+z,-y,+x},            {+y,-x,+z},            {+x,-z,+y},
	// {-x,+y,-z},           {-y,+z,-x},            {-z,+x,-y},            {-z,+y,+x},            {-y,+x,+z},            {-x,+z,+y},
	// {-x,-y,+z},           {-y,-z,+x},            {-z,-x,+y},            {-z,-y,-x},            {-y,-x,-z},            {-x,-z,-y}
	orientations := [][]int{
		{+1, +1, +1, 0, 1, 2}, {+1, +1, +1, 1, 2, 0}, {+1, +1, +1, 2, 0, 1}, {+1, +1, -1, 2, 1, 0}, {+1, +1, -1, 1, 0, 2}, {+1, +1, -1, 0, 2, 1},
		{+1, -1, -1, 0, 1, 2}, {+1, -1, -1, 1, 2, 0}, {+1, -1, -1, 2, 0, 1}, {+1, -1, +1, 2, 1, 0}, {+1, -1, +1, 1, 0, 2}, {+1, -1, +1, 0, 2, 1},
		{-1, +1, -1, 0, 1, 2}, {-1, +1, -1, 1, 2, 0}, {-1, +1, -1, 2, 0, 1}, {-1, +1, +1, 2, 1, 0}, {-1, +1, +1, 1, 0, 2}, {-1, +1, +1, 0, 2, 1},
		{-1, -1, +1, 0, 1, 2}, {-1, -1, +1, 1, 2, 0}, {-1, -1, +1, 2, 0, 1}, {-1, -1, -1, 2, 1, 0}, {-1, -1, -1, 1, 0, 2}, {-1, -1, -1, 0, 2, 1},
	}
	for i, o := range orientations {
		variants[i] = make([]coords3d.Coords3d, len(coords))
		for n, c := range coords {
			cp := []int{c.X, c.Y, c.Z}
			variants[i][n] = coords3d.Coords3d{
				X: o[0] * cp[o[3]],
				Y: o[1] * cp[o[4]],
				Z: o[2] * cp[o[5]],
			}
		}
	}
	return variants
}

// dimension 0 is scanner ID
// dimension 1 is each orientation
// dimension 2 is points for given dimension
func parse(data []string) [][][]coords3d.Coords3d {
	scanners := make([][][]coords3d.Coords3d, 0)
	currentScanner := make([]coords3d.Coords3d, 0)
	currentId := 0
	for _, l := range data[1:] {
		if l == "" {
			continue
		}
		if strings.HasPrefix(l, "---") {
			scanners = append(scanners, buildVariants(currentScanner))
			currentScanner = make([]coords3d.Coords3d, 0)
			currentId++
		} else {
			c := coords3d.Coords3d{}
			fmt.Sscanf(l, "%d,%d,%d", &c.X, &c.Y, &c.Z)
			currentScanner = append(currentScanner, c)
		}
	}
	scanners = append(scanners, buildVariants(currentScanner))
	return scanners
}

func relativePosition(refVecs, cVecs map[string][]coords3d.Coords3d) coords3d.Coords3d {
	for kref, vref := range refVecs {
		if vcoord, found := cVecs[kref]; found {
			return coords3d.Coords3d{X: vref[0].X - vcoord[0].X, Y: vref[0].Y - vcoord[0].Y, Z: vref[0].Z - vcoord[0].Z}
		}
	}
	panic("WTF")
}

func maxDistance(positions map[int]coords3d.Coords3d) int {
	max := 0
	for _, p := range positions {
		for _, o := range positions {
			if d := coords3d.Distance(p, o); d > max {
				max = d
			}
		}
	}
	return max
}

func solve(scanners [][][]coords3d.Coords3d) int {
	orientations := map[int]int{0: 0}
	positions := map[int]coords3d.Coords3d{0: {X: 0, Y: 0, Z: 0}}
	vectors := make([][]map[string][]coords3d.Coords3d, len(scanners))
	for i := 0; i < len(scanners); i++ {
		vectors[i] = make([]map[string][]coords3d.Coords3d, len(scanners[i]))
		for j := 0; j < len(scanners[i]); j++ {
			vectors[i][j] = buildVectors(scanners[i][j])
		}
	}

	for len(orientations) < len(scanners) {
		for i := range orientations {
			for j := 0; j < len(scanners); j++ {
				sb := scanners[j]
				oriA, foundA := orientations[i]
				_, foundB := orientations[j]
				if i == j || !foundA || foundB {
					continue
				}
				for oriId := range sb {
					mv := matchingVectors(vectors[i][oriA], vectors[j][oriId])
					if len(mv) >= 24 {
						rp := relativePosition(vectors[i][oriA], vectors[j][oriId])
						positions[j] = coords3d.Add(rp, positions[i])
						orientations[j] = oriId
						break
					}
				}
			}
		}
	}
	return maxDistance(positions)
}

func main() {
	data := files.ReadLines(os.Args[1])
	scanners := parse(data)
	fmt.Println(solve(scanners))
}
