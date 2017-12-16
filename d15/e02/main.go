package main

import (
  "fmt"
  "os"
  "strconv"
)

const iterations = 5000000

func min(a, b int) int {
  if a > b {
    return b
  }
  return a
}

func compute(n int64, factor int64, cond int64, c chan int64) {
  for i := 0; i < iterations; {
    n = n * factor % 2147483647
    if n % cond == 0 {
      i++
      c <- n
    }
  }
  close(c)
}

func compareBits(a, b int64) bool {
  stringA := strconv.FormatInt(a, 2)
  stringB := strconv.FormatInt(b, 2)

  if len(stringA) > 16 {
    stringA = stringA[len(stringA) - 16:]
  }
  if len(stringB) > 16 {
    stringB = stringB[len(stringB) - 16:]
  }

  return stringA == stringB
}

func solve(a, b int64) int {
  cA := make(chan int64, 1000)
  cB := make(chan int64, 1000)
  go compute(a, 16807, 4, cA)
  go compute(b, 48271, 8, cB)

  res := 0
  for i := 0; i < iterations; i++ {
    resA := <- cA
    resB := <- cB
    if compareBits(resA, resB) {
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
