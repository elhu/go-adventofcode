package main

import (
  "fmt"
  "os"
)

var hexLookup = map[byte][4]int{
  '0': [4]int{0, 0, 0, 0},
  '1': [4]int{0, 0, 0, -1},
  '2': [4]int{0, 0, -1, 0},
  '3': [4]int{0, 0, -1, -1},
  '4': [4]int{0, -1, 0, 0},
  '5': [4]int{0, -1, 0, -1},
  '6': [4]int{0, -1, -1, 0},
  '7': [4]int{0, -1, -1, -1},
  '8': [4]int{-1, 0, 0, 0},
  '9': [4]int{-1, 0, 0, -1},
  'a': [4]int{-1, 0, -1, 0},
  'b': [4]int{-1, 0, -1, -1},
  'c': [4]int{-1, -1, 0, 0},
  'd': [4]int{-1, -1, 0, -1},
  'e': [4]int{-1, -1, -1, 0},
  'f': [4]int{-1, -1, -1, -1},
}

func getMap(input string) [128][128]int {
  res := [128][128]int{}
  for i := 0; i < 128; i++ {
    res[i] = [128]int{}

    toHash := fmt.Sprintf("%s-%d", input, i)
    for j, b := range Hexdigest([]byte(toHash)) {
      value, _ := hexLookup[b]
      for k := 0; k < 4; k++ {
        res[i][4 * j + k] = value[k]
      }
    }
  }
  return res
}

// func printMap(m [128][128]byte) {
//   for i := 0; i < 128; i++ {
//     fmt.Println(string(m[i][:]))
//   }
// }


func chartZone(m *[128][128]int, i, j, label int) {
  m[i][j] = label
  if i > 0 && m[i - 1][j] == -1 {
    chartZone(m, i - 1, j, label)
  }
  if i < 127 && m[i + 1][j] == -1 {
    chartZone(m, i + 1, j, label)
  }
  if j > 0 && m[i][j - 1] == -1 {
    chartZone(m, i, j - 1, label)
  }
  if j < 127 && m[i][j + 1] == -1 {
    chartZone(m, i, j + 1, label)
  }
}

func countZones(m [128][128]int) int {
  curr := 0
  for i := 0; i < 128; i++ {
    for j := 0; j < 128; j++ {
      if m[i][j] == -1 {
        curr++
        chartZone(&m, i, j, curr)
      }
    }
  }
  return curr
}

func solve(input string) int {
  m := getMap(input)
  return countZones(m)
}

func main() {
  input := os.Args[1]
  fmt.Println(solve(input))
}
