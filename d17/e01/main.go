package main

import (
  "fmt"
  "os"
  "strconv"
)

const cycles = 2017

func insertAt(b *[]int, item, pos int) {
  *b = append(*b, 0)
  copy((*b)[pos + 2:], (*b)[pos + 1:])
  (*b)[pos + 1] = item
}

func findIndex(b []int, item int) int {
  for i, e := range b {
    if e == item {
      return i
    }
  }
  panic("WTF")
}

func solve(steps int) int {
  buffer := make([]int, 1, cycles)
  buffer[0] = 0
  pos := 0
  for i := 1; i <= cycles; i++ {
    pos = (pos + steps + 1) % len(buffer)
    insertAt(&buffer, i, pos)
  }
  idx := findIndex(buffer, 2017) + 1 % len(buffer)
  return buffer[idx]
}

func main() {
  input, _ := strconv.Atoi(os.Args[1])
  fmt.Println(solve(input))
}
