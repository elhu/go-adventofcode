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
  Alive bool
}

func intAbs(i int64) int64 {
  if i < 0 {
    return -i
  }
  return i
}

func (c *Coords) Equal(d *Coords) bool {
  return c.X == d.X && c.Y == d.Y && c.Z == d.Z
}

func (p *Particle) Distance() int64 {
  return intAbs(p.P.X) + intAbs(p.P.Y) + intAbs(p.P.Z)
}

func (p *Particle) Update() {
  p.V.X += p.A.X
  p.V.Y += p.A.Y
  p.V.Z += p.A.Z

  p.P.X += p.V.X
  p.P.Y += p.V.Y
  p.P.Z += p.V.Z
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
    particles[i] = &Particle{P: p, V: v, A: a, Id: i, Alive: true}
  }
  return particles
}

func solve(particles []*Particle) int {
  prevLen := len(particles)
  for i := 0;; i++ {
    // Update particles
    for _, p := range particles {
      p.Update()
    }
    // Check for collisions
    sort.Slice(particles, func(i, j int) bool { return particles[i].Distance() < particles[j].Distance() })
    // Flag collisions
    for i := 1; i < len(particles); i++ {
      if particles[i].P.Equal(particles[i - 1].P) {
        particles[i].Alive = false
        particles[i - 1].Alive = false
      }
    }
    // Delete flagged collisions
    for i := 0; i < len(particles); i++ {
      if !particles[i].Alive {
        particles = append(particles[:i], particles[i + 1:]...)
      }
    }
    if i > 0 && i % 10000 == 0 {
      if prevLen == len(particles) {
        break
      }
      prevLen = len(particles)
    }
  }
  return len(particles)
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
