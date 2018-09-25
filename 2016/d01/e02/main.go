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

func (pos *position) move(steps int, visited map[string]struct{}) bool {
	switch pos.direction {
	case north:
		for i := 0; i < steps; i++ {
			pos.posY++
			key := pos.toKey()
			_, seen := visited[key]
			if seen {
				return true
			}
			visited[key] = struct{}{}
		}
	case east:
		for i := 0; i < steps; i++ {
			pos.posX++
			key := pos.toKey()
			_, seen := visited[key]
			if seen {
				return true
			}
			visited[key] = struct{}{}
		}
	case south:
		for i := 0; i < steps; i++ {
			pos.posY--
			key := pos.toKey()
			_, seen := visited[key]
			if seen {
				return true
			}
			visited[key] = struct{}{}
		}
	case west:
		for i := 0; i < steps; i++ {
			pos.posX--
			key := pos.toKey()
			_, seen := visited[key]
			if seen {
				return true
			}
			visited[key] = struct{}{}
		}
	}
	return false
}

func (pos *position) toKey() string {
	return fmt.Sprintf("%d:%d", pos.posX, pos.posY)
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
	visited := make(map[string]struct{})
	visited[pos.toKey()] = struct{}{}
	for _, inst := range instructions {
		if inst[0] == 'L' {
			pos.turnLeft()
		} else {
			pos.turnRight()
		}
		steps, err := strconv.Atoi(inst[1:])
		check(err)
		stopped := pos.move(steps, visited)
		if stopped {
			return intAbs(pos.posX) + intAbs(pos.posY)
		}
	}
	panic("Couldn't find duplicate position")
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	fmt.Println(solve(string(data)))
}
