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

  for i, num := range data {
    next := data[0];
    if i < len(data) - 1 {
      next = data[i + 1];
    }
    if num == next {
      sum += int(num) - 48;
    }
  }
  fmt.Println(sum);
}
