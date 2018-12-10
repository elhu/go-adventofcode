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

func sum(numbers []int) int {
	s := 0
	for _, i := range numbers {
		s += i
	}
	return s
}

type point struct {
	posX, posY int
	vecX, vecY int
}

var exp = regexp.MustCompile(`position=<\s*(-?\d+),\s*(-?\d+)> velocity=<\s*(-?\d+),\s*(-?\d+)>`)

func parsePoints(data [][]byte) []*point {
	res := make([]*point, 0, len(data))
	for _, line := range data {
		match := exp.FindStringSubmatch(string(line))
		posX, _ := strconv.Atoi(match[1])
		posY, _ := strconv.Atoi(match[2])
		vecX, _ := strconv.Atoi(match[3])
		vecY, _ := strconv.Atoi(match[4])
		res = append(res, &point{posX, posY, vecX, vecY})
	}
	return res
}

func (p *point) move() {
	p.posX += p.vecX
	p.posY += p.vecY
}

func (p *point) rollback() {
	p.posX -= p.vecX
	p.posY -= p.vecY
}

func moveAll(points []*point) {
	for _, p := range points {
		p.move()
	}
}

func rollbackAll(points []*point) {
	for _, p := range points {
		p.rollback()
	}
}

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

func getMinMaxCoords(points []*point) (int, int, int, int) {
	minX, minY := maxInt, maxInt
	maxX, maxY := minInt, minInt
	for _, c := range points {
		if c.posX < minX {
			minX = c.posX
		}
		if c.posY < minY {
			minY = c.posY
		}
		if c.posX > maxX {
			maxX = c.posX
		}
		if c.posY > maxY {
			maxY = c.posY
		}

	}
	return minX, minY, maxX, maxY
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func computeArea(points []*point) int {
	minX, minY, maxX, maxY := getMinMaxCoords(points)
	return abs((maxX - minX) * (maxY - minY))
}

func solve(points []*point) {
	prevArea := computeArea(points)
	time := 0
	for newArea := computeArea(points); newArea <= prevArea; moveAll(points) {
		prevArea = newArea
		time++
		newArea = computeArea(points)
	}
	rollbackAll(points)
	rollbackAll(points)
	fmt.Println(time - 2)
}

func main() {
	input, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input = bytes.TrimSuffix(input, []byte{'\n'})
	data := bytes.Split(input, []byte{'\n'})
	points := parsePoints(data)
	solve(points)
}
