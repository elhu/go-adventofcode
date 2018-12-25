package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Point struct {
	x, y, z, t int
}

func parsePoints(lines [][]byte) []*Point {
	res := make([]*Point, 0, len(lines))
	for _, l := range lines {
		parts := strings.Split(string(l), ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		t, _ := strconv.Atoi(parts[3])
		res = append(res, &Point{x, y, z, t})
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func distance(a, b *Point) int {
	return abs(a.x-b.x) + abs(a.y-b.y) + abs(a.z-b.z) + abs(a.t-b.t)
}

func isNear(p *Point, values []*Point) bool {
	for _, v := range values {
		if distance(p, v) <= 3 {
			return true
		}
	}
	return false
}

func solve(points []*Point) int {
	constellations := make(map[int][]*Point)
	constellationID := 0
	for _, p := range points {
		var candidates []int
		for id, values := range constellations {
			if isNear(p, values) {
				candidates = append(candidates, id)
			}
		}
		if len(candidates) == 0 {
			constellations[constellationID] = append(constellations[constellationID], p)
			constellationID++
		} else {
			sourceID := candidates[0]
			constellations[sourceID] = append(constellations[sourceID], p)
			for _, id := range candidates[1:] {
				constellations[sourceID] = append(constellations[sourceID], constellations[id]...)
				delete(constellations, id)
			}
		}
	}
	return len(constellations)
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	data = bytes.TrimSuffix(data, []byte{'\n'})
	lines := bytes.Split(data, []byte{'\n'})
	points := parsePoints(lines)
	fmt.Println(solve(points))
}
