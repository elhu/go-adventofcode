package main

import "fmt"

func main() {
  b, c, d, e, f, h := 0, 0, 0, 0, 0, 0

  c = 123700
  b = 106700
  for ; b <= c; b += 17 {
    f = 1
    for d = 2; d != b; d++ {
      for e = 2; e != b; e++ {
        if d * e > b {
          break
        }
        if d * e == b {
          f = 0
        }
      }
    }
    if f == 0 {
      h += 1
    }
  }

  fmt.Println(h)
}
