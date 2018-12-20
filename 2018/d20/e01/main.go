package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	optionStart = '('
	optionEnd   = ')'
	optionSep   = '|'
	expEnd      = '$'
)

const (
	west  = 'W'
	east  = 'E'
	north = 'N'
	south = 'S'
)

var vectors = map[byte]Coordinates{
	west:  Coordinates{x: -1, y: 0},
	east:  Coordinates{x: 1, y: 0},
	north: Coordinates{x: 0, y: -1},
	south: Coordinates{x: 0, y: 1},
}

const CoordsOffset = 1000000

type Coordinates struct {
	x, y int
}

func (c *Coordinates) toKey() int {
	return c.y*CoordsOffset + c.x
}

func (c *Coordinates) add(v Coordinates) {
	c.x += v.x
	c.y += v.y
}

type Stackable struct {
	coords   Coordinates
	distance int
}

func buildDistances(rooms map[int]int, input []byte) {
	coords := Coordinates{}
	distance := 0
	stack := make([]Stackable, 0)
	var last Stackable
	for _, c := range input {
		switch c {
		case optionStart:
			stack = append(stack, Stackable{coords, distance})
		case optionSep:
			last = stack[len(stack)-1]
			coords = last.coords
			distance = last.distance
		case optionEnd:
			last, stack = stack[len(stack)-1], stack[:len(stack)-1]
			coords = last.coords
			distance = last.distance
		default:
			distance++
			coords.add(vectors[c])
			prevDist, exists := rooms[coords.toKey()]
			if !exists || prevDist > distance {
				rooms[coords.toKey()] = distance
			}
		}
	}
}

func solve(rooms map[int]int) int {
	maxDistance := 0

	for _, distance := range rooms {
		if distance > maxDistance {
			maxDistance = distance
		}
	}
	return maxDistance
}

func main() {
	input, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input = bytes.TrimSuffix(input, []byte{'\n'})
	rooms := make(map[int]int)
	buildDistances(rooms, input[1:])
	fmt.Println(solve(rooms))
}
