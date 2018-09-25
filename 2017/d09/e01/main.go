package main

import (
  "fmt"
  "io/ioutil"
  "os"
)

const (
  garbageOpen   = '<'
  garbageClose  = '>'

  garbageIgnore = '!'

  groupOpen     = '{'
  groupClose    = '}'

  separator     = ','
)

const (
  inGroup    = iota
  inGarbage  = iota
  ignoreNext = iota
  unknown    = iota
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func processGroup(c byte, depth int) (int, int) {
  switch c {
  case groupClose:
    return depth - 1, inGroup
  case groupOpen:
    return depth + 1, inGroup
  case garbageOpen:
    return depth, inGarbage
  case separator:
    return depth, inGroup
  default:
    return 0, unknown
  }
}

func processGarbage(c byte, depth int) int {
  switch c {
  case garbageIgnore:
    return ignoreNext
  case garbageOpen:
    return inGarbage
  case garbageClose:
    return inGroup
  default:
    return inGarbage
  }
}

func solve(data []byte) int {
  depth := 1
  totalScore := 1
  state := inGroup
  for _, c := range data[1:] {
    newDepth := depth
    if state == inGroup {
      newDepth, state = processGroup(c, depth)
    } else if state == inGarbage {
      state = processGarbage(c, depth)
    } else if state == unknown {
      panic("WAT")
    } else if state == ignoreNext {
      state = inGarbage
    }
    if newDepth != depth {
      if newDepth > depth {
        totalScore += newDepth
      }
      depth = newDepth
    }
  }
  return totalScore
}

func main() {
  data, e := ioutil.ReadFile(os.Args[1]);
  check(e)
  if data[len(data) - 1] == '\n' {
    data = data[:len(data) - 1]
  }
  fmt.Println(solve(data))
}
