package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

type Instruction struct {
  inst string
  x byte
  y string
}

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

func initRegisters() map[byte]int {
  r := make(map[byte]int)
  for i := byte('a'); i <= byte('h'); i++ {
    r[i] = 0
  }

  return r
}

func convertVal(y string, r map[byte]int) int {
  if y[0] >= 'a' && y[0] <= 'h' {
    return r[byte(y[0])]
  }
  val, _ := strconv.Atoi(y)
  return val
}

func parseInst(data string) Instruction {
  inst := Instruction{}
  parts := strings.Split(data, " ")
  inst.inst = parts[0]
  inst.x = byte(parts[1][0])
  if len(parts) == 3 {
    inst.y = parts[2]
  }
  return inst
}

func solve(instructions []Instruction, registers map[byte]int) int {
  mulCount := 0

  for pos := 0; pos < len(instructions) && pos >= 0; pos++ {
    inst := instructions[pos]
    switch inst.inst {
    case "set":
      registers[inst.x] = convertVal(inst.y, registers)
    case "sub":
      registers[inst.x] -= convertVal(inst.y, registers)
    case "mul":
      mulCount++
      registers[inst.x] *= convertVal(inst.y, registers)
    case "jnz":
      if convertVal(string(inst.x), registers) != 0 {
        pos += convertVal(inst.y, registers) - 1
      }
    }
  }
  return mulCount
}

func main() {
  fh, err := os.Open(os.Args[1])
  check(err)
  defer fh.Close()

  reader := bufio.NewReader(fh)

  c := make(chan string, 100)

  go readLines(reader, c)
  instructions := make([]Instruction, 0)

  for i := range c {
    instructions = append(instructions, parseInst(i))
  }
  fmt.Println(solve(instructions, initRegisters()))
}
