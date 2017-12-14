package main

import (
  "fmt"
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

func makeElements() []int {
  elements := make([]int, listLength, listLength)
  for i := 0; i < listLength; i++ {
    elements[i] = i
  }
  return elements
}

func Hexdigest(data []byte) []byte {
  res := ""
  for _, x := range Digest(data) {
    res = fmt.Sprintf("%s%.2x", res, x)
  }
  return []byte(res)
}

func Digest(data []byte) []int {
  elements := makeElements()
  lengths := getLengths(data)
  pos, skip := 0, 0
  for i := 0; i < 64; i++ {
    pos, skip = hash(elements, lengths, pos, skip)
  }
  return denseHash(elements)
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
