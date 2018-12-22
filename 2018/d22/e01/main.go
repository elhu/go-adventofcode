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
	plan[y-1][x-1] = computeErosionLevel(0, depth)
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

func solve(plan [][]byte) int {
	res := 0
	for i := 0; i < len(plan); i++ {
		for j := 0; j < len(plan[i]); j++ {
			switch plan[i][j] {
			case rocky:
				res += 0
			case wet:
				res += 1
			case narrow:
				res += 2
			}
		}
	}
	return res
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
	x++
	y++
	erosionPlan := prepPlan(x, y, depth)
	plan := convertToType(erosionPlan)
	// for _, l := range plan {
	// 	fmt.Println(string(l))
	// }
	fmt.Println(solve(plan))
}
