package main

import "os"
import "io/ioutil"
import "strings"
import "fmt"
import "strconv"

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func solve(instructions []int) int {
  steps, position := 0, 0
  for ; position >= 0 && position < len(instructions); steps++ {
    move := instructions[position]
    if move >= 3 {
      instructions[position]--
    } else {
      instructions[position]++
    }
    position += move
  }
  return steps
}

func main() {
  data, err := ioutil.ReadFile(os.Args[1]);
  check(err)

  raw_lines := strings.Split(string(data), "\n")
  instructions := make([]int, 0, len(raw_lines))
  for _, line := range raw_lines {
    instruction, _  := strconv.Atoi(line)
    instructions = append(instructions, instruction)
  }

  fmt.Println(solve(instructions))
}
