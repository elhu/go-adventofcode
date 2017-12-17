package main

import (
  "fmt"
  "os"
  "strconv"
)

const cycles = 50000000

func solve(steps int) int {
  pos := 0
  val := -1
  for i := 1; i <= cycles; i++ {
    pos = (pos + steps + 1) % i
    if pos == 0 {
      val = i
    }
  }
  return val
}

func main() {
  input, _ := strconv.Atoi(os.Args[1])
  fmt.Println(solve(input))
}
