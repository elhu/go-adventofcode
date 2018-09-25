package main

import "os"
import "bufio"
import "strings"
import "io"
import "fmt"

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

func uniq(items []string) []string {
  itemMap := make(map[string]bool)
  for _, item := range(items) {
    itemMap[item] = true
  }
  uniqueItems := make([]string, 0, len(itemMap))
  for key := range itemMap {
    uniqueItems = append(uniqueItems, key)
  }
  return uniqueItems
}

func valid(words []string) bool {
  return len(words) == len(uniq(words))
}

func main() {
  fh, err := os.Open(os.Args[1])
  check(err)
  defer fh.Close()

  reader := bufio.NewReader(fh)

  c := make(chan string, 100)

  go readLines(reader, c)

  validCount := 0

  for line := range c {
    words := strings.Split(line, " ")
    if valid(words) {
      validCount++
    }
  }
  fmt.Println(validCount)
}
