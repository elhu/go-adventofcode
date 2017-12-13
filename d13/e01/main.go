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

func solve(c chan string) int {
  cost := 0

  for str := range c {
    d, r := parse(str)
    if posAt(d, r) == 0 {
      cost += d * r
    }
  }
  return cost
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
