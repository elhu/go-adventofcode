package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Coordinates in the grid
type Coordinates struct {
	x, y int
}

// Cart struct
type Cart struct {
	id               int
	coord            Coordinates
	direction        int
	nextIntersection int
	alive            bool
}
type sortCarts []*Cart

const (
	goLeft     int = -1
	goStraight int = 0
	goRight    int = 1
)

const (
	left  int = iota
	up    int = iota
	right int = iota
	down  int = iota
)

var vectors = map[int]Coordinates{
	down:  Coordinates{0, 1},
	up:    Coordinates{0, -1},
	left:  Coordinates{-1, 0},
	right: Coordinates{1, 0},
}

func (c *Cart) checkCollision(carts []*Cart) bool {
	for _, o := range carts {
		if c.coord == o.coord && c.id != o.id && c.alive && o.alive {
			c.alive = false
			o.alive = false
			return true
		}
	}
	return false
}

func printState(tracks [][]byte, carts []*Cart) {
	cpy := make([][]byte, len(tracks))
	for i, l := range tracks {
		cpy[i] = make([]byte, len(l))
		copy(cpy[i], l)
	}
	for _, c := range carts {
		if c.alive == true {
			if c.direction == down {
				cpy[c.coord.y][c.coord.x] = 'v'
			} else if c.direction == up {
				cpy[c.coord.y][c.coord.x] = '^'
			} else if c.direction == left {
				cpy[c.coord.y][c.coord.x] = '<'
			} else if c.direction == right {
				cpy[c.coord.y][c.coord.x] = '>'
			}
		}
	}
	for i, l := range cpy {
		fmt.Printf("%3d: %s\n", i, l)
	}
	fmt.Println("--")
}

func (c *Cart) move(tracks [][]byte) {
	vec := vectors[c.direction]
	c.coord.x += vec.x
	c.coord.y += vec.y
	if tracks[c.coord.y][c.coord.x] == '+' {
		c.direction = (c.direction + c.nextIntersection) % 4
		if c.direction == -1 {
			c.direction = 3
		}
		if c.nextIntersection == goRight {
			c.nextIntersection = goLeft
		} else {
			c.nextIntersection++
		}
	} else if tracks[c.coord.y][c.coord.x] == '\\' {
		if c.direction == left || c.direction == right {
			c.direction++
		} else {
			c.direction--
		}
	} else if tracks[c.coord.y][c.coord.x] == '/' {
		if c.direction == left || c.direction == right {
			if c.direction == left {
				c.direction = down
			} else {
				c.direction--
			}
		} else {
			c.direction = (c.direction + 1) % 4
		}
	}
}

func (s sortCarts) Less(i, j int) bool {
	if s[i].coord.y == s[j].coord.y {
		return s[i].coord.x < s[j].coord.x
	}
	return s[i].coord.y < s[j].coord.y
}

func (s sortCarts) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortCarts) Len() int {
	return len(s)
}

func cartDirection(c byte) int {
	if c == '^' {
		return up
	}
	if c == '>' {
		return right
	}
	if c == '<' {
		return left
	}
	if c == 'v' {
		return down
	}
	return -1
}

func findCarts(tracks [][]byte) []*Cart {
	carts := make([]*Cart, 0, 1)
	cartID := 0
	for i := 0; i < len(tracks); i++ {
		for j := 0; j < len(tracks[i]); j++ {
			if dir := cartDirection(tracks[i][j]); dir != -1 {
				c := Cart{cartID, Coordinates{j, i}, dir, goLeft, true}
				cartID++
				carts = append(carts, &c)
				if c.direction == left || c.direction == right {
					tracks[i][j] = '-'
				} else {
					tracks[i][j] = '|'
				}
			}
		}
	}
	return carts
}

func printCarts(carts []*Cart) {
	res := make([]Coordinates, 0, len(carts))
	for _, c := range carts {
		res = append(res, c.coord)
	}
	fmt.Println(res)
}

func lastAliveCoords(carts []*Cart) Coordinates {
	for _, c := range carts {
		if c.alive {
			return c.coord
		}
	}
	return Coordinates{}
}

func tick(tracks [][]byte, carts []*Cart) int {
	processed := 0
	sort.Sort(sortCarts(carts))
	for _, c := range carts {
		if c.alive {
			processed++
			c.move(tracks)
			if c.checkCollision(carts) {
				processed -= 2
			}
		}
	}
	return processed
}

func solve(tracks [][]byte, carts []*Cart) Coordinates {
	for {
		if tick(tracks, carts) == 1 {
			return lastAliveCoords(carts)
		}
	}
}

func main() {
	input, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input = bytes.TrimSuffix(input, []byte{'\n'})
	tracks := bytes.Split(input, []byte{'\n'})
	carts := findCarts(tracks)
	crashCoords := solve(tracks, carts)
	fmt.Printf("%d,%d\n", crashCoords.x, crashCoords.y)
}
