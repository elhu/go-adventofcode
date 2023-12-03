package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func pad(lines []string) []string {
	padded := make([]string, len(lines)+2)
	padded[0] = strings.Repeat(".", len(lines[0])+2)
	for i, line := range lines {
		padded[i+1] = "." + line + "."
	}
	padded[len(lines)+1] = strings.Repeat(".", len(lines[0])+2)
	return padded
}

func adjacentGear(lines []string, i, start, end int) []coords2d.Coords2d {
	gears := make([]coords2d.Coords2d, 0)
	for j := start - 1; j <= end; j++ {
		for x := i - 1; x <= i+1; x++ {
			if j >= start && j < end && x == i {
				continue
			}
			if lines[x][j] == '*' {
				gears = append(gears, coords2d.Coords2d{X: j, Y: x})
			}
		}
	}
	return gears
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	check(err)
	return n
}

func toKey(c coords2d.Coords2d) int {
	return c.Y*1000000 + c.X
}

func solve(lines []string) int {
	re := regexp.MustCompile(`\d+`)
	gears := make(map[int][]int)
	for i := 1; i < len(lines)-1; i++ {
		for _, match := range re.FindAllStringIndex(lines[i], -1) {
			value := atoi(lines[i][match[0]:match[1]])
			for _, gear := range adjacentGear(lines, i, match[0], match[1]) {
				gears[toKey(gear)] = append(gears[toKey(gear)], value)
			}
		}
	}
	res := 0
	for _, values := range gears {
		if len(values) == 2 {
			res += values[0] * values[1]
		}
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := pad(strings.Split(data, "\n"))
	fmt.Println(solve(lines))
}
