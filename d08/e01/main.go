package main

import "os"
import "bufio"
import "strings"
import "io"
import "fmt"
import "regexp"
import "strconv"

var instructionExp = regexp.MustCompile(`(?P<register>\w+) (?P<operation>(inc|dec)) (?P<value>-?\d+) if (?P<condLeft>\w+) (?P<condOp>[<>!=]=?) (?P<condRight>-?\d+)`)

type Instruction struct {
  register string
  operation string
  value int
  condLeft string
  condOp string
  condRight int
}

func (i Instruction) shouldApply(registers map[string]int) bool {
  switch i.condOp {
  case "<=":
    return registers[i.condLeft] <= i.condRight
  case "<":
    return registers[i.condLeft] < i.condRight
  case "==":
    return registers[i.condLeft] == i.condRight
  case "!=":
    return registers[i.condLeft] != i.condRight
  case ">=":
    return registers[i.condLeft] >= i.condRight
  case ">":
    return registers[i.condLeft] > i.condRight
  }
  return false
}

func (i Instruction) Apply(registers map[string]int) {
  if i.shouldApply(registers) {
    switch i.operation {
    case "dec":
      registers[i.register] -= i.value
    case "inc":
      registers[i.register] += i.value
    }
  }
}

func (i Instruction) Print() {
  fmt.Printf("%s %s %d if %s %s %d\n", i.register, i.operation, i.value, i.condLeft, i.condOp, i.condRight)
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

func parseLine(line string) Instruction {
  match := instructionExp.FindStringSubmatch(line)
  inst := Instruction{}

  inst.register = match[1]
  inst.operation = match[2]
  inst.value, _ = strconv.Atoi(match[4])
  inst.condLeft = match[5]
  inst.condOp = match[6]
  inst.condRight, _ = strconv.Atoi(match[7])

  return inst
}

func solve(c chan string) int {
  registers := make(map[string]int)
  for line := range c {
    inst := parseLine(line)
    inst.Apply(registers)
  }

  var max int
  for _, v := range(registers) {
    if v > max {
      max = v
    }
  }

  return max
}

func main() {
  fh, err := os.Open(os.Args[1])
  check(err)
  defer fh.Close()

  reader := bufio.NewReader(fh)

  c := make(chan string, 100)

  go readLines(reader, c)
  fmt.Println(solve(c))
}
