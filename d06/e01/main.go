package main

import "os"
import "io/ioutil"
import "fmt"
import "strings"
import "strconv"

type IntSliceSet struct {
  set map[string]struct{}
}

func newIntSliceSet() *IntSliceSet {
  set := IntSliceSet{}
  set.set = make(map[string]struct{})
  return &set
}

func (set *IntSliceSet) Add(i []int) bool {
  key := strings.Trim(strings.Replace(fmt.Sprint(i), " ", ":", -1), "[]")
  _, found := set.set[key]
  set.set[key] = struct{}{}
  return !found
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
  seen.Add(blocks)

  for ;; count++ {
    currIdx := maxIndex(blocks)
    toRealloc := blocks[currIdx]
    blocks[currIdx] = 0
    currIdx++
    for i := 0; i < toRealloc; i++ {
      blocks[currIdx % len(blocks)]++
      currIdx++
    }
    if !seen.Add(blocks) {
      break
    }
  }
  return count
}

func main() {
  data, err := ioutil.ReadFile(os.Args[1]);
  check(err);

  blocks := getBlocks(string(data))

  fmt.Println(solve(blocks));
}
