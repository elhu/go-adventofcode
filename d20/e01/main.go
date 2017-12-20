package main

import (
  "os"
  "io/ioutil"
  "fmt"
  "regexp"
  "sort"
  "strconv"
  "strings"
)

var exp = regexp.MustCompile(`p=<(-?\d+),(-?\d+),(-?\d+)>, v=<(-?\d+),(-?\d+),(-?\d+)>, a=<(-?\d+),(-?\d+),(-?\d+)>`)

type Coords struct {
  X, Y, Z int64
}

type Particle struct {
  Id int
  P, V, A *Coords
}

func intAbs(i int64) int64 {
  if i < 0 {
    return -i
  }
  return i
}

func (p *Particle) Distance() int64 {
  return intAbs(p.P.X) + intAbs(p.P.Y) + intAbs(p.P.Z)
}

func (p *Particle) Acceleration() int64 {
  return intAbs(p.A.X) + intAbs(p.A.Y) + intAbs(p.A.Z)
}

func (p *Particle) Velocity() int64 {
  return intAbs(p.V.X) + intAbs(p.V.Y) + intAbs(p.V.Z)
}

func MatchToCoord(match []string, idx int) *Coords {
  x, _ := strconv.ParseInt(match[idx + 1], 10, 64)
  y, _ := strconv.ParseInt(match[idx + 2], 10, 64)
  z, _ := strconv.ParseInt(match[idx + 3], 10, 64)

  return &Coords{X: x, Y: y, Z: z}
}

func ParseParticles(data string) []*Particle {
  lines := strings.Split(strings.TrimSuffix(data, "\n"), "\n")
  particles := make([]*Particle, len(lines))
  for i, line := range lines {
    match := exp.FindStringSubmatch(line)
    p := MatchToCoord(match, 0)
    v := MatchToCoord(match, 3)
    a := MatchToCoord(match, 6)
    particles[i] = &Particle{P: p, V: v, A: a, Id: i}
  }
  return particles
}

func solve(particles []*Particle) int {
  sort.SliceStable(particles, func(i, j int) bool {
    if particles[i].Acceleration() != particles[j].Acceleration() {
      return particles[i].Acceleration() < particles[j].Acceleration()
    } else if particles[i].Velocity() != particles[j].Velocity() {
      return particles[i].Velocity() < particles[j].Velocity()
    } else {
      return particles[i].Distance() < particles[j].Distance()
    }
  })
  return particles[0].Id
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  data, err := ioutil.ReadFile(os.Args[1])
  check(err)

  particles := ParseParticles(string(data))
  fmt.Println(solve(particles))
}
