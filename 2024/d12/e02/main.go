package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"adventofcode/utils/sets"
	"fmt"
	"os"
	"strings"
)

var N = coords2d.Coords2d{0, -1}
var E = coords2d.Coords2d{1, 0}
var S = coords2d.Coords2d{0, 1}
var W = coords2d.Coords2d{-1, 0}

var ADJ = [4]coords2d.Coords2d{N, S, W, E}
var EMPTY_SET = sets.New[coords2d.Coords2d]()

func computeSides(pl sets.Set[State]) int {
	var res int
	for _, s := range pl.Members() {
		pos := s.borderPos
		vec := s.vec
		if vec == N || vec == S {
			if !pl.HasMember(State{borderPos: coords2d.Add(pos, E), vec: vec}) {
				res++
			}
			if !pl.HasMember(State{borderPos: coords2d.Add(pos, W), vec: vec}) {
				res++
			}
		} else {
			if !pl.HasMember(State{borderPos: coords2d.Add(pos, S), vec: vec}) {
				res++
			}
			if !pl.HasMember(State{borderPos: coords2d.Add(pos, N), vec: vec}) {
				res++
			}
		}
	}
	return res / 2
}

type State struct {
	borderPos coords2d.Coords2d
	vec       coords2d.Coords2d
}

func processRegion(pt map[coords2d.Coords2d]byte, pm map[coords2d.Coords2d]bool, start coords2d.Coords2d) int {
	var area int
	queue := []coords2d.Coords2d{start}
	var curr coords2d.Coords2d
	rt := pt[start]
	pl := *sets.New[State]()
	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		if pm[curr] {
			continue
		}
		pm[curr] = true
		area++
		for _, v := range ADJ {
			if pt[coords2d.Add(curr, v)] == rt {
				queue = append(queue, coords2d.Add(curr, v))
			} else {
				s := State{borderPos: coords2d.Add(curr, v), vec: v}
				pl.Add(s)
			}
		}
	}
	sides := computeSides(pl)
	return area * sides
}

func solve(data []string) int {
	pm := make(map[coords2d.Coords2d]bool)
	pt := make(map[coords2d.Coords2d]byte)
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			pm[coords2d.Coords2d{X: j, Y: i}] = false
			pt[coords2d.Coords2d{X: j, Y: i}] = data[i][j]
		}
	}
	res := 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if !pm[coords2d.Coords2d{X: j, Y: i}] {
				res += processRegion(pt, pm, coords2d.Coords2d{X: j, Y: i})
			}
		}
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	raw := strings.Split(data, "\n")
	fmt.Println(solve(raw))
}
