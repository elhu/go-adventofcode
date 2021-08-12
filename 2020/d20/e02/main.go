package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
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
	variants     [][][]byte
	processed    bool
}

func reverse(s []byte) []byte {
	res := make([]byte, len(s))
	for i, c := range s {
		res[len(s)-i-1] = c
	}
	return res
}

// Works because square
func rotateRight(data [][]byte) [][]byte {
	res := make([][]byte, len(data))
	for i := range data {
		res[i] = make([]byte, len(data[i]))
		copy(res[i], data[i])
	}
	for layer := 0; layer < len(data)/2; layer++ {
		from, to := layer, len(data)-layer-1
		for i := from; i < to; i++ {
			offset := i - from
			top := res[from][i]
			res[from][i] = res[to-offset][from]
			res[to-offset][from] = res[to][to-offset]
			res[to][to-offset] = res[i][to]
			res[i][to] = top
		}
	}
	return res
}

func horizontalFlip(data [][]byte) [][]byte {
	res := make([][]byte, len(data))
	for i := range data {
		res[i] = reverse(data[i])
	}
	return res
}

func verticalFlip(data [][]byte) [][]byte {
	res := make([][]byte, len(data))
	for i := range data {
		res[i] = data[len(data)-i-1]
	}
	return res
}

func (s *Square) buildVariants() {
	s.variants[0] = make([][]byte, len(s.data))
	for i := range s.data {
		s.variants[0][i] = make([]byte, len(s.data[i]))
		copy(s.variants[0][i], s.data[i])
	}
	s.variants[1] = horizontalFlip(s.variants[0])
	s.variants[2] = verticalFlip(s.variants[1])
	s.variants[3] = horizontalFlip(s.variants[2])
	s.variants[4] = rotateRight(s.variants[0])
	s.variants[5] = horizontalFlip(s.variants[4])
	s.variants[6] = verticalFlip(s.variants[5])
	s.variants[7] = horizontalFlip(s.variants[6])
}

func (s *Square) buildEdges() {
	for i, variant := range s.variants {
		top := variant[0]
		bottom := variant[len(variant)-1]
		left, right := make([]byte, len(variant)), make([]byte, len(variant))
		for i, l := range variant {
			left[i] = l[0]
			right[i] = l[len(l)-1]
		}
		s.edges[i] = Edges{"top": top, "bottom": bottom, "left": left, "right": right}
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
		variants:     make([][][]byte, 8),
		neighbors:    make(map[string]int),
		orientations: []int{0, 1, 2, 3, 4, 5, 6, 7},
	}
	sq.buildVariants()
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

var borderChecks = [][]string{
	[]string{"top", "bottom"},
	[]string{"right", "left"},
	[]string{"bottom", "top"},
	[]string{"left", "right"},
}

func setTopLeftCorner(selfID int, squares map[int]*Square) []int {
	currSquare := squares[selfID]
	for _, orientation := range currSquare.orientations {
		bordersMatched := 0
		for _, check := range borderChecks[1:3] {
			for id, sq := range squares {
				if selfID == id {
					continue
				}
				for _, o := range sq.orientations {
					if bytes.Compare(currSquare.edges[orientation][check[0]], sq.edges[o][check[1]]) == 0 {
						bordersMatched++
					}
				}
			}
		}
		// Find the right orientation (corner must have two matched borders)
		if bordersMatched == 2 {
			for _, check := range borderChecks[1:3] {
				for id, sq := range squares {
					if selfID == id {
						continue
					}
					for _, o := range sq.orientations {
						if bytes.Compare(currSquare.edges[orientation][check[0]], sq.edges[o][check[1]]) == 0 {
							sq.neighbors[check[1]] = selfID
							currSquare.neighbors[check[0]] = id
							sq.orientations = []int{o}
							currSquare.orientations = []int{orientation}
						}
					}
				}
			}
		}
		if len(currSquare.orientations) == 1 {
			break
		}
	}
	return []int{currSquare.neighbors["bottom"], currSquare.neighbors["right"]}
}

func findPosition(selfID int, squares map[int]*Square) []int {
	var next []int
	currSquare := squares[selfID]
	if currSquare.processed {
		return []int{}
	}
	currSquare.processed = true
	for _, orientation := range currSquare.orientations {
		for _, check := range borderChecks {
			// Don't overwrite neighbors
			if _, found := currSquare.neighbors[check[0]]; found {
				continue
			}
			for id, sq := range squares {
				if selfID == id {
					continue
				}
				for _, o := range sq.orientations {
					if _, found := sq.neighbors[check[1]]; !found && bytes.Compare(currSquare.edges[orientation][check[0]], sq.edges[o][check[1]]) == 0 {
						if bytes.Compare(currSquare.edges[orientation][check[0]], sq.edges[o][check[1]]) == 0 {
							sq.neighbors[check[1]] = selfID
							currSquare.neighbors[check[0]] = id
							next = append(next, id)
							sq.orientations = []int{o}
							currSquare.orientations = []int{orientation}
						}
					}
				}
			}
		}
		if len(currSquare.orientations) == 1 {
			break
		}
	}
	return next
}

func placeSquares(squares map[int]*Square) int {
	var corners []int
	for id := range squares {
		if isCorner(id, squares) {
			corners = append(corners, id)
		}
	}
	squares[corners[0]].processed = true
	// Force set topleft corner to make the whole thing deterministic
	queue := setTopLeftCorner(corners[0], squares)
	var top int
	for len(queue) > 0 {
		top, queue = queue[0], queue[1:]
		queue = append(queue, findPosition(top, squares)...)
	}
	return corners[0]
}

func stitchMap(topLeft int, squares map[int]*Square) Square {
	size := int(math.Sqrt(float64(len(squares)))) * (len(squares[topLeft].data) - 2)
	bigGiantMap := make([][]byte, size)
	for i := range bigGiantMap {
		bigGiantMap[i] = make([]byte, size)
	}
	i, j := 0, 0
	var currSquare *Square
	for rowStartID := topLeft; rowStartID != 0; rowStartID = squares[rowStartID].neighbors["bottom"] {
		j = 0
		for currID := rowStartID; currID != 0; currID = currSquare.neighbors["right"] {
			currSquare = squares[currID]
			orientation := currSquare.orientations[0]
			data := currSquare.variants[orientation]
			for m, line := range data[1 : len(data)-1] {
				for n, content := range line[1 : len(line)-1] {
					bigGiantMap[i+m][j+n] = content
				}
			}
			j += len(currSquare.data) - 2
		}
		i += len(currSquare.data) - 2
	}
	sq := Square{
		data:         bigGiantMap,
		edges:        make([]Edges, 8),
		variants:     make([][][]byte, 8),
		neighbors:    make(map[string]int),
		orientations: []int{0, 1, 2, 3, 4, 5, 6, 7},
	}
	sq.buildVariants()
	sq.buildEdges()
	return sq
}

func monsterAt(i, j int, bigGiantMap [][]byte) bool {
	for m, s := range seaMonster {
		for n, c := range s {
			if c == '#' && bigGiantMap[i+m][j+n] != '#' {
				return false
			}
		}
	}
	return true
}

func solve(sq Square) int {
	monsterSize := 0
	for _, m := range seaMonster {
		monsterSize += bytes.Count(m, []byte{'#'})
	}
	seaSize := 0
	for _, line := range sq.data {
		seaSize += bytes.Count(line, []byte{'#'})
	}
	for _, o := range sq.orientations {
		monstersFound := 0
		for i := 0; i < len(sq.variants[o])-len(seaMonster)+1; i++ {
			line := sq.variants[o][i]
			for j := 0; j < len(line)-len(seaMonster[0])+1; j++ {
				if monsterAt(i, j, sq.variants[o]) {
					monstersFound++
				}
			}
		}
		if monstersFound > 0 {
			return seaSize - monstersFound*monsterSize
		}
	}
	return -1
}

var seaMonster = [][]byte{
	[]byte("                  # "),
	[]byte("#    ##    ##    ###"),
	[]byte(" #  #  #  #  #  #   "),
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
	topLeft := placeSquares(squares)
	sq := stitchMap(topLeft, squares)
	fmt.Println(solve(sq))
}
