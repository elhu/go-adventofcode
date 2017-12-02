package main

import "os"
import "bufio"
import "io"
import "fmt"
import "strings"
import "strconv"
import "sort"

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

func lineIntegers(fh *bufio.Reader, c chan []int) {
  lines_channel := make(chan string)
  go readLines(fh, lines_channel)

  for line := range lines_channel {
    words := strings.Split(line, "\t")
    integers := make([]int, len(words))

    for i, word := range words {
      integer, err := strconv.Atoi(word)
      integers[i] = integer
      check(err)
    }
    c <- integers
  }
  close(c)
}

func main() {
  fh, err := os.Open(os.Args[1])
  check(err)
  reader := bufio.NewReader(fh)
  c := make(chan []int, 100)

  go lineIntegers(reader, c)

  sum := 0

  for integers := range c {
    sort.Sort(sort.Reverse(sort.IntSlice(integers)))
    found := false
    for i, dividend := range integers[:len(integers) - 1] {
      for _, divisor := range integers[i + 1:] {
        if dividend % divisor == 0 {
          sum += dividend / divisor
          found = true
          break
        }
      }
      if found {
        break
      }
    }
  }
  fmt.Println(sum)
}
