package main

import "os"
import "io/ioutil"
import "fmt"
import "strings"
import "strconv"

type IntSliceSet struct {
  set map[string]int
}

func newIntSliceSet() *IntSliceSet {
  set := IntSliceSet{}
  set.set = make(map[string]int)
  return &set
}

func (set *IntSliceSet) Add(i []int, pos int) (bool, int) {
  key := strings.Trim(strings.Replace(fmt.Sprint(i), " ", ":", -1), "[]")
  previousPos, found := set.set[key]
  set.set[key] = pos
  return !found, previousPos
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func getBlocks(data string) []int {
  blocksAsStrings := strings.Split(strings.TrimSuffix(data, "\n"), "\t")
  res := make([]int, 0, len(blocksAsStrings))

  for _, block := range blocksAsStrings {
    value, err := strconv.Atoi(block)
    check(err)
    res = append(res, value)
  }

  return res
}

func maxIndex(blocks []int) int {
  max := blocks[0]
  maxIdx := 0
  for idx, value := range blocks {
    if value > max {
      max = value
      maxIdx = idx
    }
  }
  return maxIdx
}

func solve(blocks []int) int {
  seen := newIntSliceSet()
  count := 1
  seen.Add(blocks, 0)

  for ;; count++ {
    currIdx := maxIndex(blocks)
    toRealloc := blocks[currIdx]
    blocks[currIdx] = 0
    currIdx++
    for i := 0; i < toRealloc; i++ {
      blocks[currIdx % len(blocks)]++
      currIdx++
    }
    isNewState, pos := seen.Add(blocks, count)
    if !isNewState {
      return count - pos
    }
  }
}

func main() {
  data, err := ioutil.ReadFile(os.Args[1]);
  check(err);

  blocks := getBlocks(string(data))

  fmt.Println(solve(blocks));
}
