package main

import (
  "fmt"
  "io/ioutil"
  "os"
)

const listLength = 256

var lengthSuffix = []int{17, 31, 73, 47, 23}

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

func hash(elements []int, lengths []int, pos int, skip int) (int, int) {
  for _, length := range lengths {
    reverse(elements, pos, length)
    pos += skip + length
    skip++
  }
  return pos, skip
}

func xor(arr []int) int {
  acc := arr[0] ^ arr[1]
  for i := 2; i < len(arr); i++ {
    acc = acc ^ arr[i]
  }
  return acc
}

func denseHash(elements []int) []int {
  res := make([]int, 0, 16)
  for i := 0; i < len(elements); i += 16 {
    res = append(res, xor(elements[i:i+16]))
  }
  return res
}

func solve(elements []int, lengths []int) {
  pos, skip := 0, 0
  for i := 0; i < 64; i++ {
    pos, skip = hash(elements, lengths, pos, skip)
  }
  for _, x := range denseHash(elements) {
    fmt.Printf("%.2x", x)
  }
  fmt.Printf("\n")
}

func getLengths(data []byte) []int {
  lengths := make([]int, 0, len(data) + len(lengthSuffix))

  for _, item := range data {
    lengths = append(lengths, int(item))
  }
  for _, item := range lengthSuffix {
    lengths = append(lengths, item)
  }

  return lengths
}

func main() {
  data, e := ioutil.ReadFile(os.Args[1]);
  check(e)
  if data[len(data) - 1] == '\n' {
    data = data[:len(data) - 1]
  }
  lengths := getLengths(data)
  elements := make([]int, listLength, listLength)
  for i := 0; i < listLength; i++ {
    elements[i] = i
  }
  solve(elements, lengths)
}
