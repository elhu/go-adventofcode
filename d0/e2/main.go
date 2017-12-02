package main

import "os"
import "io/ioutil"
import "fmt"

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  data, err := ioutil.ReadFile(os.Args[1]);
  check(err);
  sum := 0;
  data = data[:len(data) - 1];

  for i, num := range data {
    next_pos := (i + len(data) / 2) % len(data);
    next := data[next_pos];
    if num == next {
      sum += int(num) - 48;
    }
  }
  fmt.Println(sum);
}
