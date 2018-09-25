package main

import (
  "os"
  "bufio"
  "io"
  "fmt"
  "strings"
)

// Arbitrary width to compute hashmap
// Assumes boundaries to a theoritically infinite grid, but speeds things up
const MaxLen = 10000000000

const (
  up = iota
  right = iota
  down = iota
  left = iota
)

type Carrier struct {
  Pos Coord
  Dir int
}

const (
  clean = iota
  weakened = iota
  infected = iota
  flagged = iota
)

type Coord struct {
  X, Y int
}

type Set map[int]int

func (c *Carrier) Move() {
  switch c.Dir {
  case up:
    c.Pos.Y--
  case right:
    c.Pos.X++
  case down:
    c.Pos.Y++
  case left:
    c.Pos.X--
  }
}

func (c *Carrier) TurnLeft() {
  c.Dir--
  if c.Dir == -1 {
    c.Dir = left
  }
}

func (c *Carrier) TurnRight() {
  c.Dir = (c.Dir + 1) % 4
}

func (c Coord) ToKey() int {
  return c.Y * MaxLen + c.X
}

func (s Set) State(k Coord) int {
  state, ok := s[k.ToKey()]
  if ok {
    return state
  }
  return clean
}

func (s Set) Add(k Coord, state int) {
  s[k.ToKey()] = state
}

func (s Set) AdvanceState(k Coord) int {
  state, ok := s[k.ToKey()]
  if ok {
    newState := (state + 1) % 4
    s[k.ToKey()] = newState
    return newState
  } else {
    s[k.ToKey()] = weakened
    return weakened
  }
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func readLines(fh *bufio.Reader, c chan string) {
  for {
    line, err := fh.ReadString('\n')
    c <- strings.TrimSuffix(line, "\n")
    if err == io.EOF {
      break
    }
  }
  close(c)
}

func markViruses(c chan string) (Set, Coord) {
  s := make(Set, 10000)

  i := 0
  var width int
  for line := range c {
    width = len(line)
    for j, c := range line {
      if c == '#' {
        s.Add(Coord{j, i}, infected)
      }
    }
    i++
  }
  return s, Coord{X: width / 2, Y: i / 2}
}

func solve(v Set, c *Carrier) int {
  marked := 0
  for i := 0; i < 10000000; i++ {
    switch v.State(c.Pos) {
    case clean:
      c.TurnLeft()
    case infected:
      c.TurnRight()
    case flagged:
      c.TurnRight()
      c.TurnRight()
    }
    newState := v.AdvanceState(c.Pos)
    if newState == infected {
      marked++
    }
    c.Move()
  }
  return marked
}

func main() {
  fh, err := os.Open(os.Args[1])
  check(err)
  defer fh.Close()

  reader := bufio.NewReader(fh)
  c := make(chan string, 100)

  go readLines(reader, c)

  viruses, startPos := markViruses(c)
  fmt.Println(solve(viruses, &Carrier{Pos: startPos, Dir: up}))
}
