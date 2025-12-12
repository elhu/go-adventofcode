package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Shape struct {
	lines    []string
	areaUsed int
}

func parseShape(data string) Shape {
	lines := strings.Split(data, "\n")
	areaUsed := 0
	for _, line := range lines[1:] {
		areaUsed += strings.Count(line, "#")
	}
	return Shape{
		lines:    lines[1:],
		areaUsed: areaUsed,
	}
}

type Area struct {
	width  int
	height int
	reqs   []int
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func parseArea(data string) Area {
	parts := strings.Split(data, " ")
	xy := strings.Split(parts[0], "x")
	area := Area{
		width:  atoi(xy[0]),
		height: atoi(xy[1][:len(xy[1])-1]),
	}
	for _, req := range parts[1:] {
		area.reqs = append(area.reqs, atoi(req))
	}
	return area
}

func computeFillRate(area Area, shapes []Shape) float64 {
	shapeArea := 0.0
	for i, req := range area.reqs {
		shapeArea += float64(req) * float64(shapes[i].areaUsed)
	}
	return shapeArea / float64(area.width*area.height)
}

func solve(shapes []Shape, areas []Area) int {
	fits := 0
	for _, area := range areas {
		fillRate := computeFillRate(area, shapes)
		if fillRate < 1.0 { // Fill rates are either > 1.0 or < 0.75, taking a guess that < 1.0 means it fits
			fits++
		}
	}
	return fits
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	sections := strings.Split(data, "\n\n")
	shapes := make([]Shape, len(sections)-1)
	for i := range sections[:len(sections)-1] {
		shapes[i] = parseShape(sections[i])
	}
	var areas []Area
	for _, areaData := range strings.Split(sections[len(sections)-1], "\n") {
		areas = append(areas, parseArea(areaData))
	}
	fmt.Println(solve(shapes, areas))
}
