package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Edges map[string][]byte

type Square struct {
	neighbors    map[string]int
	orientations []int
	data         [][]byte
	edges        []Edges
}

func reverse(s []byte) []byte {
	res := make([]byte, len(s))
	for i, c := range s {
		res[len(s)-i-1] = c
	}
	return res
}

func (s *Square) buildEdges() {
	top := s.data[0]
	bottom := s.data[len(s.data)-1]
	left, right := make([]byte, len(s.data)), make([]byte, len(s.data))
	for i, l := range s.data {
		left[i] = l[0]
		right[i] = l[len(l)-1]
	}
	s.edges[0] = Edges{"top": top, "bottom": bottom, "left": left, "right": right}

	// Flip horizontally
	s.edges[1] = Edges{
		"top":    reverse(top),
		"bottom": reverse(bottom),
		"left":   s.edges[0]["right"],
		"right":  s.edges[0]["left"],
	}
	// Flip vertically
	s.edges[2] = Edges{
		"top":    s.edges[1]["bottom"],
		"bottom": s.edges[1]["top"],
		"left":   reverse(s.edges[1]["left"]),
		"right":  reverse(s.edges[1]["right"]),
	}
	// Flip horizontally again
	s.edges[3] = Edges{
		"top":    reverse(s.edges[2]["top"]),
		"bottom": reverse(s.edges[2]["bottom"]),
		"left":   s.edges[2]["right"],
		"right":  s.edges[2]["left"],
	}
	// Rotate 90 degrees
	s.edges[4] = Edges{
		"top":    reverse(left),
		"bottom": reverse(right),
		"left":   bottom,
		"right":  top,
	}
	// Flip horizontally
	s.edges[5] = Edges{
		"top":    reverse(s.edges[4]["top"]),
		"bottom": reverse(s.edges[4]["bottom"]),
		"left":   s.edges[4]["right"],
		"right":  s.edges[4]["left"],
	}
	// Flip vertically
	s.edges[6] = Edges{
		"top":    s.edges[5]["bottom"],
		"bottom": s.edges[5]["top"],
		"left":   reverse(s.edges[5]["left"]),
		"right":  reverse(s.edges[5]["right"]),
	}
	// Flip horizontally again
	s.edges[7] = Edges{
		"top":    reverse(s.edges[6]["top"]),
		"bottom": reverse(s.edges[6]["bottom"]),
		"left":   s.edges[6]["right"],
		"right":  s.edges[6]["left"],
	}
}

func parseSquare(input string) (int, *Square) {
	lines := strings.Split(input, "\n")
	data := make([][]byte, len(lines)-1)
	var id int
	fmt.Sscanf(lines[0], "Tile %d:", &id)
	for i, s := range lines[1:] {
		data[i] = []byte(s)
	}
	sq := &Square{
		data:         data,
		edges:        make([]Edges, 8),
		neighbors:    make(map[string]int),
		orientations: []int{0, 1, 2, 3, 4, 5, 6, 7},
	}
	sq.buildEdges()
	return id, sq
}

func hasMatch(edge []byte, selfID int, oppositeSide string, squares map[int]*Square) bool {
	for id, sq := range squares {
		if selfID == id {
			continue
		}
		for _, e := range sq.edges {
			if bytes.Compare(edge, e[oppositeSide]) == 0 {
				return true
			}
		}
	}
	return false
}

func boolSliceEquals(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

var borderChecks = [][]string{
	[]string{"top", "bottom"},
	[]string{"right", "left"},
	[]string{"bottom", "top"},
	[]string{"left", "right"},
}

func isCorner(id int, squares map[int]*Square) bool {
	// Check each orientation
	checks := [][]string{
		[]string{"top", "bottom"},
		[]string{"right", "left"},
		[]string{"bottom", "top"},
		[]string{"left", "right"},
	}
	cornerChecks := [][]bool{
		[]bool{true, true, false, false},
		[]bool{false, true, true, false},
		[]bool{false, false, true, true},
		[]bool{true, false, false, true},
	}
	for _, edges := range squares[id].edges {
		matches := make([]bool, len(checks))
		for i, c := range checks {
			if hasMatch(edges[c[0]], id, c[1], squares) {
				matches[i] = true
			}
		}
		for _, cc := range cornerChecks {
			if boolSliceEquals(cc, matches) {
				return true
			}
		}
	}
	return false
}

func solve(squares map[int]*Square) uint64 {
	res := uint64(1)
	for id := range squares {
		if isCorner(id, squares) {
			res *= uint64(id)
		}
	}
	return res
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n\n")
	squares := make(map[int]*Square)
	for _, sq := range input {
		id, square := parseSquare(sq)
		squares[id] = square
	}
	fmt.Println(solve(squares))
}
