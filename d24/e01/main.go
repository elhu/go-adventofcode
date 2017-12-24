package main

import (
  "os"
  "io/ioutil"
  "fmt"
  "strconv"
  "strings"
)

type Component struct {
  Left, Right, Id int
}

type Set struct {
  set map[int]struct{}
  Sum int
}

func NewSet() *Set {
  s := make(map[int]struct{})
  return &Set{set: s}
}

func (s *Set) Add(c *Component) {
  s.set[c.Id] = struct{}{}
  s.Sum += c.Left + c.Right
}

func (s *Set) Present(c *Component) bool {
  _, ok := s.set[c.Id]
  return ok
}

func (s *Set) Copy() *Set {
  newSet := make(map[int]struct{}, len(s.set))
  for k, v := range s.set {
    newSet[k] = v
  }
  return &Set{set: newSet, Sum: s.Sum}
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func ParseParts(data string) []*Component {
  data = strings.Trim(data, "\n")
  lines := strings.Split(data, "\n")
  parts := make([]*Component, len(lines))
  for i, l := range lines {
    pins := strings.Split(l, "/")
    left, _ := strconv.Atoi(pins[0])
    right, _ := strconv.Atoi(pins[1])
    parts[i] = &Component{Left: left, Right: right, Id: i}
  }

  return parts
}

func FindCandidates(pins int, parts []*Component, used *Set) []*Component {
  candidates := make([]*Component, 0)
  for _, c := range parts {
    if !used.Present(c) && (c.Left == pins || c.Right == pins) {
      candidates = append(candidates, c)
    }
  }

  return candidates
}

func AddComponent(currPin int, parts []*Component, used *Set, resChan chan int) {
  candidates := FindCandidates(currPin, parts, used)
  for _, c := range candidates {
    cUsed := used.Copy()
    newPin := c.Left
    if currPin == c.Left {
      newPin = c.Right
    }
    cUsed.Add(c)
    go AddComponent(newPin, parts, cUsed, resChan)
  }
  fmt.Printf("%d\n", used.Sum)
  resChan <- used.Sum
}

func Solve(parts []*Component) int {
  used := NewSet()
  max := 0
  resChan := make(chan int)
  go AddComponent(0, parts, used, resChan)

  for sum := range resChan {
    if sum > max {
      max = sum
    }
  }
  return max
}

func main() {
  data, err := ioutil.ReadFile(os.Args[1])
  check(err)

  parts := ParseParts(string(data))
  fmt.Println(Solve(parts))
}
