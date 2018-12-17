package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var inputExp = regexp.MustCompile(`(x|y)=(\d+), (x|y)=(\d+)..(\d+)`)

type inputDesc struct {
	xmin, xmax, ymin, ymax int
}

const Source = '+'
const Sand = '.'
const Clay = '#'
const Water = '~'
const WetSand = '|'
const Boundary = 'X'

const maxInt = int(^uint(0) >> 1)

func findBounds(descriptions []inputDesc) (int, int, int, int) {
	minX := maxInt
	minY := maxInt
	maxX := -1
	maxY := -1

	for _, d := range descriptions {
		if d.xmax > maxX {
			maxX = d.xmax
		}
		if d.ymax > maxY {
			maxY = d.ymax
		}
		if d.xmin < minX {
			minX = d.xmin
		}
		if d.ymin < minY {
			minY = d.ymin
		}
	}
	return minX, maxX, minY, maxY
}

func parse(lines [][]byte) [][]byte {
	descriptions := make([]inputDesc, 0, len(lines))
	for _, l := range lines {
		match := inputExp.FindStringSubmatch(string(l))
		if match[1] == "x" {
			x, _ := strconv.Atoi(match[2])
			ymin, _ := strconv.Atoi(match[4])
			ymax, _ := strconv.Atoi(match[5])
			descriptions = append(descriptions, inputDesc{xmax: x, xmin: x, ymax: ymax, ymin: ymin})
		} else {
			y, _ := strconv.Atoi(match[2])
			xmin, _ := strconv.Atoi(match[4])
			xmax, _ := strconv.Atoi(match[5])
			descriptions = append(descriptions, inputDesc{xmax: xmax, xmin: xmin, ymax: y, ymin: y})
		}
	}
	minX, maxX, minY, maxY := findBounds(descriptions)
	// Leave space at the stop for overflowing
	minY--
	minX -= 2
	maxX += 2
	maxY++
	plan := make([][]byte, (maxY-minY)+1)

	for i := 0; i < (maxY-minY)+1; i++ {
		plan[i] = make([]byte, (maxX-minX)+1)
		for j := 0; j < (maxX-minX)+1; j++ {
			plan[i][j] = Sand
		}
	}
	for _, d := range descriptions {
		for i := d.ymin; i <= d.ymax; i++ {
			for j := d.xmin; j <= d.xmax; j++ {
				plan[i-minY][j-minX] = Clay
			}
		}
	}
	for i := 0; i < len(plan[0]); i++ {
		plan[0][i] = Boundary
		plan[len(plan)-1][i] = Boundary
	}
	for i := 0; i < len(plan); i++ {
		plan[i][0] = Boundary
		plan[i][len(plan[0])-1] = Boundary
	}
	plan[0][500-minX] = Source
	return plan
}

func printPlan(plan [][]byte) {
	for _, l := range plan {
		fmt.Println(string(l))
	}
}

func countWater(plan [][]byte) int {
	res := 0
	for i := 1; i < len(plan); i++ {
		for j := 0; j < len(plan[i]); j++ {
			if plan[i][j] == Water || plan[i][j] == WetSand {
				res++
			}
		}
	}
	return res
}

func countStillWater(plan [][]byte) int {
	res := 0
	for i := 1; i < len(plan); i++ {
		for j := 0; j < len(plan[i]); j++ {
			if plan[i][j] == Water {
				res++
			}
		}
	}
	return res
}

type Location struct {
	x, y int
}

const offset = 100000

func (l *Location) toKey() int {
	return l.y*offset + l.x
}

func fromKey(k int) *Location {
	return &Location{
		y: k / offset,
		x: k % offset,
	}
}

func fillLine(plan [][]byte, e *Location) bool {
	leftIndex := bytes.LastIndexByte(plan[e.y][:e.x], Clay)
	rightIndex := bytes.IndexByte(plan[e.y][e.x:], Clay)
	if leftIndex == -1 || rightIndex == -1 {
		return false
	}
	for i := leftIndex; i < e.x+rightIndex; i++ {
		if !(plan[e.y+1][i] == Water || plan[e.y+1][i] == Clay) {
			return false
		}
	}
	for i := leftIndex + 1; i < e.x+rightIndex; i++ {
		plan[e.y][i] = Water
	}
	return true
}

func solve(plan [][]byte) {
	newCount := countWater(plan)
	prevCount := -1
	sourceLocation := bytes.IndexByte(plan[0], Source)
	edges := make(map[int]struct{}, 1)
	edges[(&Location{x: sourceLocation, y: 0}).toKey()] = struct{}{}
	for prevCount != newCount {
		prevCount = newCount
		for k, _ := range edges {
			e := fromKey(k)
			if plan[e.y+1][e.x] == Sand {
				plan[e.y+1][e.x] = WetSand
				edges[(&Location{x: e.x, y: e.y + 1}).toKey()] = struct{}{}
			} else if plan[e.y+1][e.x] == Clay || plan[e.y+1][e.x] == Water {
				if !fillLine(plan, e) {
					if plan[e.y][e.x-1] == Sand {
						plan[e.y][e.x-1] = WetSand
						edges[(&Location{x: e.x - 1, y: e.y}).toKey()] = struct{}{}
					}
					if plan[e.y][e.x+1] == Sand {
						plan[e.y][e.x+1] = WetSand
						edges[(&Location{x: e.x + 1, y: e.y}).toKey()] = struct{}{}
					}
				} else {
					delete(edges, e.toKey())
				}
			}
		}
		newCount = countWater(plan)
	}
}

func main() {
	input, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input = bytes.TrimSuffix(input, []byte{'\n'})
	lines := bytes.Split(input, []byte{'\n'})
	plan := parse(lines)
	solve(plan)
	fmt.Println(countStillWater(plan))
}
