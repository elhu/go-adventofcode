package main

import (
  "os"
  "io/ioutil"
  "fmt"
  "strconv"
  "strings"
  "sync/atomic"
  "time"
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

func (s *Set) Len() int {
  return len(s.set)
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

func AddComponent(currPin int, parts []*Component, used *Set, resChan chan [2]int, counter *int64) {
  atomic.AddInt64(counter, 1)
  defer func() {
    // Hide exceptions about sending on closed channel
    recover()
  }()
  candidates := FindCandidates(currPin, parts, used)
  for _, c := range candidates {
    cUsed := used.Copy()
    newPin := c.Left
    if currPin == c.Left {
      newPin = c.Right
    }
    cUsed.Add(c)
    go AddComponent(newPin, parts, cUsed, resChan, counter)
  }
  atomic.AddInt64(counter, -1)
  resChan <- [2]int{used.Len(), used.Sum}
  if *counter == 0 {
    time.Sleep(time.Second)
    if *counter == 0 {
      close(resChan)
    }
  }
}

func Solve(parts []*Component) int {
  used := NewSet()
  maxLen, max := 0, 0
  resChan := make(chan [2]int)
  var counter int64
  go AddComponent(0, parts, used, resChan, &counter)

  for res := range resChan {
    if res[0] > maxLen {
      maxLen = res[0]
      max = 0
    }
    if res[1] > max {
      max = res[1]
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
