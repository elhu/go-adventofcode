package main

import (
  "os"
  "io/ioutil"
  "strings"
  "fmt"
  "regexp"
  "strconv"
)

var beginExp = regexp.MustCompile(`Begin in state (\w)\.`)
var stepsExp = regexp.MustCompile(`Perform a diagnostic checksum after (\d+) steps\.`)

var stateExp = regexp.MustCompile(`In state (\w):`)
// var condExp = regexp.MustCompile(`  If the current value is (\d):`)
var writeExp = regexp.MustCompile(`    - Write the value (\d)\.`)
var moveExp = regexp.MustCompile(`    - Move one slot to the (\w+)\.`)
var contExp = regexp.MustCompile(`    - Continue with state (\w)\.`)

type Machine struct {
  State byte
  Pos int
  Registers map[int]int
  Steps int
  States map[byte]State
}

type State struct {
  Label byte
  Cases [2]Case
}

type Case struct {
  Value int
  Offset int
  NextState byte
}

func NewMachine(s byte) (* Machine) {
  m := Machine{}
  m.State = s
  m.Registers = make(map[int]int)
  m.States = make(map[byte]State)

  return &m
}

func (m *Machine) Checksum() int {
  c := 0
  for _, v := range m.Registers {
    if v == 1 {
      c++
    }
  }
  return c
}

func (m *Machine) Step() {
  m.Steps++
  val := m.Registers[m.Pos]
  c := m.States[m.State].Cases[val]
  m.Registers[m.Pos] = c.Value
  m.Pos += c.Offset
  m.State = c.NextState
}


func check(e error) {
  if e != nil {
    panic(e)
  }
}

func ParseHeader(data string) (byte, int) {
  lines := strings.Split(data, "\n")
  match := beginExp.FindStringSubmatch(lines[0])
  initialStep := byte(match[1][0])

  match = stepsExp.FindStringSubmatch(lines[1])
  rounds, _ := strconv.Atoi(match[1])
  return initialStep, rounds
}

func ParseCase(lines []string) Case {
  ret := Case{}

  match := writeExp.FindStringSubmatch(lines[1])
  ret.Value, _ = strconv.Atoi(match[1])

  match = moveExp.FindStringSubmatch(lines[2])
  switch match[1] {
  case "left":
    ret.Offset = -1
  case "right":
    ret.Offset = 1
  default:
    panic("Ooops: " + match[1])
  }

  match = contExp.FindStringSubmatch(lines[3])
  ret.NextState = byte(match[1][0])

  return ret
}

func ParseState(data string) State {
  ret := State{}
  lines := strings.Split(data, "\n")
  ret.Cases[0] = ParseCase(lines[1:5])
  ret.Cases[1] = ParseCase(lines[5:])

  match := stateExp.FindStringSubmatch(lines[0])
  ret.Label = byte(match[1][0])

  return ret
}

func solve(m *Machine, rounds int) int {
  for i := 0; i < rounds; i++ {
    m.Step()
  }
  return m.Checksum()
}

func main() {
  data, err := ioutil.ReadFile(os.Args[1])
  check(err)

  d := strings.Trim(string(data), "\n")
  parts := strings.Split(d, "\n\n")
  initialStep, rounds := ParseHeader(parts[0])
  m := NewMachine(initialStep)
  for _, p := range parts[1:] {
    step := ParseState(p)
    m.States[step.Label] = step
  }
  fmt.Println(solve(m, rounds))
}
