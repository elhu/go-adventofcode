package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func readLines(fh *bufio.Reader, c chan string) {
  for {
    line, err := fh.ReadString('\n')
    c <- strings.TrimSuffix(line, "\n")
    if err == io.EOF {
      break
    }
  }
  close(c)
}

func parse(str string) (int, int) {
  values := strings.Split(str, ": ")
  d, _ := strconv.Atoi(values[0])
  r, _ := strconv.Atoi(values[1])
  return d, r
}

func posAt(time, r int) int {
  length := r - 1
  pos := time % (2 * length)
  if pos > length {
    pos = length - (pos - length)
  }
  return pos
}

func caught(layers [][2]int, offset int) bool {
  for _, layer := range layers {
    d, r := layer[0], layer[1]
    if posAt(d + offset, r) == 0 {
      return true
    }
  }
  return false
}

func getLayers(c chan string) ([][2]int) {
  layers := make([][2]int, 0)
  for str := range c {
    d, r := parse(str)
    layers = append(layers, [2]int{d, r})
  }
  return layers
}

func solve(c chan string) int {
  layers := getLayers(c)
  i := 0
  for ;;i++ {
    if caught(layers, i) == false {
      break
    }
  }
  return i
}

func main() {
  fh, err := os.Open(os.Args[1])
  check(err)
  defer fh.Close()

  reader := bufio.NewReader(fh)

  c := make(chan string, 100)

  go readLines(reader, c)
  fmt.Println(solve(c))
}
