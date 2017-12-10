package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "strconv"
  "strings"
)

const listLength = 256

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func idxConv(idx int) int {
  return idx % listLength
}

func reverse(slice []int, start int, length int) {
  for i := 0; i < length / 2; i++ {
    slice[idxConv(i + start)], slice[idxConv(start + length - i - 1)] = slice[idxConv(start + length - i - 1)], slice[idxConv(i + start)]
  }
}

func solve(elements []int, lengths []int) int {
  pos := 0
  skip := 0

  for _, length := range lengths {
    reverse(elements, pos, length)
    pos += skip + length
    skip++
  }
  return elements[0] * elements[1]
}

func getLengths(data string) []int {
  stringLengths := strings.Split(data, ",")
  lengths := make([]int, 0, len(stringLengths))
  for _, length := range stringLengths {
    l, _ := strconv.Atoi(strings.Trim(length, " \t"))
    lengths = append(lengths, l)
  }
  return lengths
}

func main() {
  data, e := ioutil.ReadFile(os.Args[1]);
  check(e)
  if data[len(data) - 1] == '\n' {
    data = data[:len(data) - 1]
  }
  lengths := getLengths(string(data))
  elements := make([]int, listLength, listLength)
  for i := 0; i < listLength; i++ {
    elements[i] = i
  }
  fmt.Println(solve(elements, lengths))
}
