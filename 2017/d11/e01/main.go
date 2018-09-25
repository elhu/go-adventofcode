package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "strings"
)

const delim = ","

type Coord struct {
  x int
  y int
}

func (c *Coord) Add(other *Coord) {
  c.x += other.x
  c.y += other.y
}

func intAbs(a int) int {
  if a < 0 {
    return -a
  } else {
    return a
  }
}

func (c *Coord) Distance() int {
  if intAbs(c.x) > intAbs(c.y) {
    return intAbs(c.x)
  } else {
    return intAbs(c.y)
  }
}

func (c *Coord) Print() {
  fmt.Printf("x: %d, y: %d\n", c.x, c.y)
}

func NewCoord(x, y int) *Coord {
  coord := Coord{x: x, y: y}
  return &coord
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func countMoves(moves []string) int {
  coords := make(map[string]*Coord)
  coords["sw"] = NewCoord(-1, 1)
  coords["nw"] = NewCoord(-1, 0)
  coords["n"] = NewCoord(0, -1)
  coords["ne"] = NewCoord(1, -1)
  coords["se"] = NewCoord(1, 0)
  coords["s"] = NewCoord(0, 1)

  currPos := NewCoord(0, 0)

  for _, move := range moves {
    currPos.Add(coords[move])
  }
  return currPos.Distance()
}

func solve(moves []string) int {
  return countMoves(moves)
}

func main() {
  data, e := ioutil.ReadFile(os.Args[1]);
  check(e)
  cleanData := strings.Trim(string(data), " \t\n")
  fmt.Println(solve(strings.Split(cleanData, delim)))
}
