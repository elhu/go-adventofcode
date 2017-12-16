package main

import (
  "bytes"
  "fmt"
  "io/ioutil"
  "os"
  "strconv"
  "strings"
)

const programsCount = 16

const dances = 1000000000

const (
  spin = iota
  exchange = iota
  partner = iota
)

type Move struct{
  move int
  argA byte
  argB byte
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func parseMoves(data string) []Move {
  moves := make([]Move, 0, 500)
  for _, d := range strings.Split(data, ",") {
    move := Move{}
    switch d[0] {
    case 's':
      move.move = spin
      a, _ := strconv.Atoi(d[1:])
      move.argA = byte(a)
    case 'x':
      move.move = exchange
      args := strings.Split(d[1:], "/")
      a, _ := strconv.Atoi(args[0])
      b, _ := strconv.Atoi(args[1])
      move.argA = byte(a)
      move.argB = byte(b)
    case 'p':
      move.move = partner
      move.argA = d[1]
      move.argB = d[3]
    }
    moves = append(moves, move)
  }

  return moves
}

func initPrograms() []byte {
  p := make([]byte, programsCount)
  for i := 0; i < programsCount; i++ {
    p[i] = byte(int('a') + i)
  }
  return p
}

func findPos(p []byte, a byte) int {
  for i := 0; i < len(p); i++ {
    if p[i] == a {
      return i
    }
  }
  panic("WTF")
}

func doSpin(p []byte, c int) {
  for i := 0; i < c; i++ {
    buff := p[len(p) - 1]
    for j := len(p) - 1; j > 0; j-- {
      p[j] = p[j - 1]
    }
    p[0] = buff
  }
}

func doExchange(p []byte, a, b int) {
  p[a], p[b] = p[b], p[a]
}

func doPartner(p []byte, a, b byte) {
  doExchange(p, findPos(p, a), findPos(p, b))
}

func doDance(p []byte, moves []Move) {
  for _, m := range moves {
    switch m.move {
    case spin:
      doSpin(p, int(m.argA))
    case exchange:
      doExchange(p, int(m.argA), int(m.argB))
    case partner:
      doPartner(p, m.argA, m.argB)
    }
  }
}

func findPeriod(p []byte, m []Move) int {
  cmp := initPrograms()
  for i := 0; i < dances; i++ {
    doDance(p, m)
    if bytes.Equal(p, cmp) {
      return i + 1
    }
  }
  panic("Period not found")
}

func solve(p []byte, m []Move) {
  period := findPeriod(p, m)
  for i := 0; i < dances % period; i++ {
    doDance(p, m)
  }
}

func main() {
  data, err := ioutil.ReadFile(os.Args[1]);
  check(err);

  programs := initPrograms()
  moves := parseMoves(string(data))
  solve(programs, moves)
  fmt.Println(string(programs))
}
