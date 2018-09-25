package main

import "os"
import "bufio"
import "io"
import "fmt"
import "strings"
import "strconv"

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

func main() {
  fh, err := os.Open(os.Args[1])
  check(err)
  defer fh.Close()

  reader := bufio.NewReader(fh)
  c := make(chan string, 100)

  go readLines(reader, c)

  sum := 0

  for line := range c {
    min, max := -1, -1
    for i, raw_cell := range strings.Split(line, "\t") {
      cell, err := strconv.Atoi(raw_cell)
      check(err)
      if i == 0 {
        min = cell
        max = cell
      } else {
        if cell < min {
          min = cell
        }
        if cell > max {
          max = cell
        }
      }
    }
    sum += max - min
  }
  fmt.Println(sum)
}
