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

func compute(n int64, factor int64, cond int64) []int64 {
  res := make([]int64, 0, iterations)

  count := 0
  for {
    n = n * factor % 2147483647
    if n % cond == 0 {
      count++
      res = append(res, n)
    }
    if count == iterations - 1 {
      return res
    }
  }
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
  resA := compute(a, 16807, 4)
  resB := compute(b, 48271, 8)

  res := 0
  for i := 0; i < min(len(resA), len(resB)); i++ {
    if compareBits(resA[i], resB[i]) {
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
