package main

import (
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

const (
	north = iota
	east  = iota
	south = iota
	west  = iota
)

type position struct {
	direction int
	posX      int
	posY      int
}

func (pos *position) turnLeft() {
	pos.direction--
	if pos.direction < north {
		pos.direction = west
	}
}

func (pos *position) turnRight() {
	pos.direction++
	if pos.direction > west {
		pos.direction = north
	}
}

func (pos *position) move(steps int) {
	switch pos.direction {
	case north:
		pos.posY += steps
	case east:
		pos.posX += steps
	case south:
		pos.posY -= steps
	case west:
		pos.posX -= steps
	}
}

func intAbs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func solve(data string) int {
	instructions := strings.Split(data, ", ")
	pos := position{north, 0, 0}
	for _, inst := range instructions {
		if inst[0] == 'L' {
			pos.turnLeft()
		} else {
			pos.turnRight()
		}
		steps, err := strconv.Atoi(inst[1:])
		check(err)
		pos.move(steps)
	}
	return intAbs(pos.posX) + intAbs(pos.posY)
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	fmt.Println(solve(string(data)))
}
