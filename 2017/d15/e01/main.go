package main

import (
  "bytes"
  "fmt"
  "os"
  "strconv"
)

const iterations = 40000000

func compareBits(a, b int64) bool {
  bitsA := []byte(strconv.FormatInt(a, 2))
  bitsB := []byte(strconv.FormatInt(b, 2))

  if len(bitsA) > 16 {
    bitsA = bitsA[len(bitsA) - 16:]
  }
  if len(bitsB) > 16 {
    bitsB = bitsB[len(bitsB) - 16:]
  }
  return bytes.Equal(bitsA, bitsB)
}

func solve(a, b int64) int {
  res := 0
  for i := 0; i < iterations; i++ {
    a = a * 16807 % 2147483647
    b = b * 48271 % 2147483647
    if compareBits(a, b) == true {
      res++
    }
  }
  return res
}

func main() {
  seedA, _ := strconv.Atoi(os.Args[1])
  seedB, _ := strconv.Atoi(os.Args[2])
  fmt.Println(solve(int64(seedA), int64(seedB)))
}
