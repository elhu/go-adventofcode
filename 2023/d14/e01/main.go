package main

import (
	"adventofcode/utils/files"
	"bytes"
	"fmt"
	"os"
)

func computeScore(lines [][]byte) int {
	score := 0
	for i := 0; i < len(lines); i++ {
		score += bytes.Count(lines[i], []byte("O")) * (len(lines) - i - 1)
	}
	return score
}

func buildFixedMap(lines [][]byte) [][]byte {
	var res [][]byte
	for i, l := range lines {
		res = append(res, bytes.Repeat([]byte("."), len(l)))
		for j, c := range l {
			if c == '#' {
				res[i][j] = c
			}
		}
	}
	return res
}

func tiltNorth(lines [][]byte) [][]byte {
	fm := buildFixedMap(lines)
	for j := 0; j < len(lines[0])-1; j++ {
		for i := 0; i < len(lines[0])-1; i++ {
			if lines[i][j] == '#' {
				rr := 0
				for k := i + 1; lines[k][j] != '#'; k++ {
					if lines[k][j] == 'O' {
						rr++
					}
				}
				for k := 0; k < rr; k++ {
					fm[i+k+1][j] = 'O'
				}
			}
		}
	}
	return fm
}

func solve(lines [][]byte) int {
	lines = tiltNorth(lines)
	return computeScore(lines)
}

func pad(lines [][]byte, val byte) [][]byte {
	padded := make([][]byte, len(lines)+2)
	padded[0] = bytes.Repeat([]byte{val}, len(lines[0])+2)
	for i, line := range lines {
		padded[i+1] = append([]byte{val}, line...)
		padded[i+1] = append(padded[i+1], val)
	}
	padded[len(lines)+1] = bytes.Repeat([]byte{val}, len(lines[0])+2)
	return padded
}

func main() {
	data := bytes.TrimRight(files.ReadFile(os.Args[1]), "\n")
	lines := bytes.Split(data, []byte("\n"))
	lines = pad(lines, '#')
	fmt.Println(solve(lines))
}
