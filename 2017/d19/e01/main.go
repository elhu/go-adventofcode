package main

import (
  "os"
  "io/ioutil"
  "fmt"
  "strings"
)

const (
  up = iota
  right = iota
  down = iota
  left = iota
)

type Pos struct {
  x, y int
  dir int
}

func (p *Pos) NextPos() {
  switch p.dir {
  case up:
    p.y--
  case right:
    p.x++
  case down:
    p.y++
  case left:
    p.x--
  }
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func buildDiagram(data []byte) ([]string, *Pos) {
  diagram := strings.Split(string(data), "\n")
  pos := Pos{x: strings.Index(diagram[0], "|"), y: 0, dir: down}

  return diagram, &pos
}

func findNewDir(prev byte, pos *Pos, grid []string) int {
  posToCheck := make(map[int][2]int)
  posToCheck[up] = [2]int{pos.y - 1, pos.x}
  posToCheck[right] = [2]int{pos.y, pos.x + 1}
  posToCheck[down] = [2]int{pos.y + 1, pos.x}
  posToCheck[left] = [2]int{pos.y, pos.x - 1}

  for k, v := range posToCheck {
    if v[0] >= 0 && v[0] < len(grid) && v[1] >= 0 && v[1] < len(grid[0]) {
      if grid[v[0]][v[1]] != prev && grid[v[0]][v[1]] != ' ' {
        return k
      }
    }
  }
  panic("WTF")
}

func solve(grid []string, pos *Pos) {
  for {
    prevChar := grid[pos.y][pos.x]
    pos.NextPos()
    if pos.y >= len(grid) || pos.x >= len(grid[0]) || grid[pos.y][pos.x] == ' ' {
      break
    }
    if grid[pos.y][pos.x] == '+' {
      pos.dir = findNewDir(prevChar, pos, grid)
    }
    if grid[pos.y][pos.x] >= 'A' && grid[pos.y][pos.x] <= 'Z' {
      fmt.Printf(string(grid[pos.y][pos.x]))
    }
  }
  fmt.Printf("\n")
}

func main() {
  data, err := ioutil.ReadFile(os.Args[1])
  check(err)

  grid, startPos := buildDiagram(data)
  solve(grid, startPos)
}
