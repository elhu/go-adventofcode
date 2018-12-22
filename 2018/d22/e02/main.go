package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"

	"github.com/RyanCarrier/dijkstra"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var numberToType = map[int]byte{
	0: rocky,
	1: wet,
	2: narrow,
}

const (
	rocky  = '.'
	narrow = '|'
	wet    = '='
	mouth  = 'M'
	target = 'T'
)

const xOffset = 48271
const yOffset = 16807
const offset = 20183

var targetExp = regexp.MustCompile(`target: (\d+),(\d+)`)

func computeErosionLevel(geoIndex, depth int) int {
	return (geoIndex + depth) % offset
}

func prepPlan(x, y, depth int) [][]int {
	targetX, targetY := x, y
	x += 100
	y += 100
	plan := make([][]int, y)
	for i := 0; i < y; i++ {
		plan[i] = make([]int, x)
	}
	for i := 0; i < y; i++ {
		geoIndex := i * xOffset
		plan[i][0] = computeErosionLevel(geoIndex, depth)
	}
	for i := 0; i < x; i++ {
		geoIndex := i * yOffset
		plan[0][i] = computeErosionLevel(geoIndex, depth)
	}
	plan[0][0] = computeErosionLevel(0, depth)
	populate(plan, depth)
	plan[targetY][targetX] = computeErosionLevel(0, depth)
	return plan
}

func populate(plan [][]int, depth int) {
	for i := 1; i < len(plan); i++ {
		for j := 1; j < len(plan[i]); j++ {
			geoIndex := plan[i][j-1] * plan[i-1][j]
			plan[i][j] = computeErosionLevel(geoIndex, depth)
		}
	}
}

func convertToType(plan [][]int) [][]byte {
	res := make([][]byte, len(plan))
	for i := 0; i < len(plan); i++ {
		res[i] = make([]byte, len(plan[i]))
		for j := 0; j < len(plan[i]); j++ {
			res[i][j] = numberToType[plan[i][j]%3]
		}
	}
	return res
}

const (
	none         = iota
	climbingGear = iota
	torch        = iota
)

type DistTool struct {
	distance, tool int
}

func cost(kind byte, tool int) (int, int) {
	switch kind {
	case rocky:
		if tool == climbingGear || tool == torch {
			return 1, tool
		} else {
			return 8, torch
		}
	case wet:
		if tool == climbingGear || tool == none {
			return 1, tool
		} else {
			return 8, climbingGear
		}
	case narrow:
		if tool == torch || tool == none {
			return 1, tool
		} else {
			return 8, none
		}
	}
	panic("WTF")
}

var toolMapping = map[byte][]int{
	rocky:  []int{climbingGear, torch},
	wet:    []int{climbingGear, none},
	narrow: []int{torch, none},
}

var costMapping = map[byte]map[byte][]int{
	rocky:  {rocky: []int{torch, climbingGear}, wet: []int{climbingGear}, narrow: []int{torch}},
	wet:    {rocky: []int{climbingGear}, wet: []int{climbingGear, none}, narrow: []int{none}},
	narrow: {rocky: []int{torch}, wet: []int{none}, narrow: []int{torch, none}},
}

func toKey(x, y, tool int) string {
	return fmt.Sprintf("%d:%d:%d", x, y, tool)
}

func solve(plan [][]byte, x, y int) int64 {
	graph := dijkstra.NewGraph()
	for i := 0; i < len(plan); i++ {
		for j := 0; j < len(plan[i]); j++ {
			added := make([]string, 0)
			for _, t := range toolMapping[plan[i][j]] {
				added = append(added, toKey(j, i, t))
				// Add vertices for each possible tool for the cell
				graph.AddMappedVertex(toKey(j, i, t))
			}
			for a := 0; a < len(added); a++ {
				for b := 0; b < len(added); b++ {
					if added[a] != added[b] {
						// Add tool change between same cell with different tools
						graph.AddMappedArc(added[a], added[b], 7)
					}
				}
			}
		}
	}
	for i := 0; i < len(plan); i++ {
		for j := 0; j < len(plan[i]); j++ {
			neighbours := [][2]int{
				{i - 1, j},
				{i + 1, j},
				{i, j - 1},
				{i, j + 1},
			}
			for _, n := range neighbours {
				ny := n[0]
				nx := n[1]
				if ny < 0 || nx < 0 || ny >= len(plan) || nx >= len(plan[0]) {
					continue
				}
				for _, t := range costMapping[plan[i][j]][plan[ny][nx]] {
					// For each tool in common between src and dst, add an Edge
					src := toKey(j, i, t)
					dst := toKey(nx, ny, t)
					graph.AddMappedArc(src, dst, 1)
				}
			}
		}
	}
	if src, e := graph.GetMapping(toKey(0, 0, torch)); e == nil {
		if dst, f := graph.GetMapping(toKey(x, y, torch)); f == nil {
			best, err := graph.Shortest(src, dst)
			check(err)
			return best.Distance
		}
	}
	panic("WTF")
}

func main() {
	input, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input = bytes.TrimSuffix(input, []byte{'\n'})
	lines := bytes.Split(input, []byte{'\n'})
	depth, _ := strconv.Atoi(string(lines[0][7:]))
	match := targetExp.FindStringSubmatch(string(lines[1]))
	x, _ := strconv.Atoi(match[1])
	y, _ := strconv.Atoi(match[2])
	erosionPlan := prepPlan(x, y, depth)
	plan := convertToType(erosionPlan)
	solve(plan, x, y)
	fmt.Println(solve(plan, x, y))
}
