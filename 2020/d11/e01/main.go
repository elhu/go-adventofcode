package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isEqual(fpa, fpb [][]byte) bool {
	for i := range fpa {
		if bytes.Compare(fpa[i], fpb[i]) != 0 {
			return false
		}
	}
	return true
}

var surroundingsPos = [][]int{
	[]int{-1, -1},
	[]int{0, -1},
	[]int{1, -1},
	[]int{-1, 0},
	[]int{1, 0},
	[]int{-1, 1},
	[]int{0, 1},
	[]int{1, 1},
}

func playTurn(fp [][]byte) [][]byte {
	newState := make([][]byte, len(fp))
	for i := range fp {
		newState[i] = []byte(strings.Repeat("?", len(fp[0])))
	}
	for i := 0; i < len(fp); i++ {
		for j := 0; j < len(fp[i]); j++ {
			c := fp[i][j]
			surroundings := make(map[byte]int)
			for _, pos := range surroundingsPos {
				if i+pos[0] >= 0 && i+pos[0] < len(fp) &&
					j+pos[1] >= 0 && j+pos[1] < len(fp[i]) {
					surroundings[fp[i+pos[0]][j+pos[1]]]++
				}
			}
			if c == 'L' && surroundings['#'] == 0 {
				newState[i][j] = '#'
			} else if c == '#' && surroundings['#'] >= 4 {
				newState[i][j] = 'L'
			} else {
				newState[i][j] = c
			}
		}
	}
	return newState
}

func solve(fp [][]byte) int {
	for {
		newState := playTurn(fp)
		if isEqual(fp, newState) {
			break
		}
		fp = newState
	}
	res := 0
	for _, l := range fp {
		for _, c := range l {
			if c == '#' {
				res++
			}
		}
	}
	return res
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	floorPlan := make([][]byte, len(lines))
	for i, l := range lines {
		floorPlan[i] = []byte(l)
	}
	fmt.Println(solve(floorPlan))
}
