package main

import (
  "os"
  "bufio"
  "io"
  "fmt"
  "strings"
)

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

type Coord struct {
  X, Y int
}

type Set map[string]struct{}

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

func (c Coord) ToKey() string {
  return fmt.Sprintf("%d:%d", c.X, c.Y)
}

func (s Set) Present(k Coord) bool {
  _, ok := s[k.ToKey()]
  return ok
}

func (s Set) Add(k Coord) {
  s[k.ToKey()] = struct{}{}
}

func (s Set) Remove(k Coord) {
  delete(s, k.ToKey())
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
  s := make(Set)

  i := 0
  var width int
  for line := range c {
    width = len(line)
    for j, c := range line {
      if c == '#' {
        s.Add(Coord{j, i})
      }
    }
    i++
  }
  return s, Coord{X: width / 2, Y: i / 2}
}

func solve(v Set, c *Carrier) int {
  marked := 0
  for i := 0; i < 10000; i++ {
    if v.Present(c.Pos) {
      c.TurnRight()
      v.Remove(c.Pos)
    } else {
      c.TurnLeft()
      v.Add(c.Pos)
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

