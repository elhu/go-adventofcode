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

type Program struct {
  pid int
  registers map[byte]int
  pos int
  rpos int
  rq *[]int
  sq *[]int
  state int

  sndCount int
}

const (
  running = iota
  receiving = iota
  finished = iota
)

func NewProgram(pid int, rq, sq *[]int) *Program {
  p := Program{pid: pid}

  p.registers = initRegisters(p.pid)
  p.pos = 0
  p.rpos = 0
  p.state = running

  p.sndCount = 0

  p.rq = rq
  p.sq = sq


  return &p
}

func (p *Program) Process(instructions []Instruction) {
  // fmt.Printf("[%d] Processing %d\n", p.pid, p.pos)
  inst := instructions[p.pos]
  switch inst.inst {
  case "snd":
    *p.sq = append(*p.sq, p.registers[inst.x])
    p.sndCount++
  case "set":
    p.registers[inst.x] = convertVal(inst.y, p.registers)
  case "add":
    p.registers[inst.x] += convertVal(inst.y, p.registers)
  case "mul":
    p.registers[inst.x] *= convertVal(inst.y, p.registers)
  case "mod":
    p.registers[inst.x] %= convertVal(inst.y, p.registers)
  case "rcv":
    if len(*p.rq) > p.rpos {
      p.registers[inst.x] = (*p.rq)[p.rpos]
      p.rpos++
    } else {
      p.state = receiving
      return
    }
  case "jgz":
    if convertVal(string(inst.x), p.registers) > 0 {
      p.pos += convertVal(inst.y, p.registers) - 1
    }
  }
  p.pos++
  if (p.pos >= 0 && p.pos < len(instructions)) {
    p.state = running
  } else {
    p.state = finished
  }
}

func solve(zero, one *Program, instructions []Instruction) int {
  for ;zero.state == running || one.state == running;  {
    if zero.state != finished {
      zero.Process(instructions)
    }
    if one.state != finished {
      one.Process(instructions)
    }
    // fmt.Printf("%d %d\n", zero.state, one.state)
  }
  return one.sndCount
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

func initRegisters(pid int) map[byte]int {
  r := make(map[byte]int)
  for i := byte('a'); i <= byte('z'); i++ {
    r[i] = 0
  }
  r['p'] = pid
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

func convertVal(y string, r map[byte]int) int {
  if y[0] >= 'a' && y[0] < 'z' {
    return r[byte(y[0])]
  }
  val, _ := strconv.Atoi(y)
  return val
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
  zmq := make([]int, 0, 0)
  omq := make([]int, 0, 0)
  zero := NewProgram(0, &zmq, &omq)
  one := NewProgram(1, &omq, &zmq)
  fmt.Println(solve(zero, one, instructions))
}
