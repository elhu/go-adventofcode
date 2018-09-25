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

func processGarbage(c byte, depth int) (int, int) {
  switch c {
  case garbageIgnore:
    return 0, ignoreNext
  case garbageOpen:
    return 1, inGarbage
  case garbageClose:
    return 0, inGroup
  default:
    return 1, inGarbage
  }
}

func solve(data []byte) int {
  depth := 1
  state := inGroup
  totalCanceled := 0
  for _, c := range data[1:] {
    newDepth := depth
    canceled := 0
    if state == inGroup {
      newDepth, state = processGroup(c, depth)
    } else if state == inGarbage {
      canceled, state = processGarbage(c, depth)
      totalCanceled += canceled
    } else if state == unknown {
      panic("WAT")
    } else if state == ignoreNext {
      state = inGarbage
    }
    depth = newDepth
  }
  return totalCanceled
}

func main() {
  data, e := ioutil.ReadFile(os.Args[1]);
  check(e)
  if data[len(data) - 1] == '\n' {
    data = data[:len(data) - 1]
  }
  fmt.Println(solve(data))
}
