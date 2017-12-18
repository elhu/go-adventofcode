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
  for i := byte('a'); i <= byte('z'); i++ {
    r[i] = 0
  }

  return r
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

func convertY(y string, r map[byte]int) int {
  if y[0] >= 'a' && y[0] < 'z' {
    return r[byte(y[0])]
  }
  val, _ := strconv.Atoi(y)
  return val
}

func solve(instructions []Instruction, registers map[byte]int) int {
  lastSound := -1
  for pos := 0; pos < len(instructions) && pos >= 0; pos++ {
    inst := instructions[pos]
    switch inst.inst {
    case "snd":
      lastSound = registers[inst.x]
    case "set":
      registers[inst.x] = convertY(inst.y, registers)
    case "add":
      registers[inst.x] += convertY(inst.y, registers)
    case "mul":
      registers[inst.x] *= convertY(inst.y, registers)
    case "mod":
      registers[inst.x] %= convertY(inst.y, registers)
    case "rcv":
      if registers[inst.x] != 0 {
        return lastSound
      }
    case "jgz":
      if registers[inst.x] > 0 {
        pos += convertY(inst.y, registers) - 1
      }
    }
  }
  return -1
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
