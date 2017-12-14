package main

import (
  "fmt"
  "math/bits"
  "os"
)

func hexDecimalLookup() map[byte]uint {
  return map[byte]uint{
    '0': 0,
    '1': 1,
    '2': 2,
    '3': 3,
    '4': 4,
    '5': 5,
    '6': 6,
    '7': 7,
    '8': 8,
    '9': 9,
    'a': 10,
    'b': 11,
    'c': 12,
    'd': 13,
    'e': 14,
    'f': 15,
  }
}

func solve(input string) int {
  count := 0
  lookup := hexDecimalLookup()

  for i := 0; i < 128; i++ {
    toHash := fmt.Sprintf("%s-%d", input, i)
    used := 0
    for _, b := range Hexdigest([]byte(toHash)) {
      value, _ := lookup[b]
      used += bits.OnesCount(value)
    }
    count += used
  }
  return count
}

func main() {
  input := os.Args[1]
  fmt.Println(solve(input))
}
