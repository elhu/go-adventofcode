package main

import (
	"adventofcode/utils/files"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func horizontalReflectionAt(row int, p [][]byte) bool {
	for offset := 0; offset < min(row, len(p)-row); offset++ {
		if bytes.Compare(p[row-offset-1], p[row+offset]) != 0 {
			return false
		}
	}
	return true
}

func verticalReflectionAt(col int, p [][]byte) bool {
	for offset := 0; offset < min(col, len(p[0])-col); offset++ {
		for i := 0; i < len(p); i++ {
			if p[i][col-offset-1] != p[i][col+offset] {
				return false
			}
		}
	}
	return true
}

func findReflections(p [][]byte) []int {
	var reflections []int
	for row := 1; row < len(p); row++ {
		if horizontalReflectionAt(row, p) {
			reflections = append(reflections, row*100)
		}
	}
	for col := 1; col < len(p[0]); col++ {
		if verticalReflectionAt(col, p) {
			reflections = append(reflections, col)
		}
	}
	return reflections
}

func smudgeAt(i, j int, p []string) [][]byte {
	var smudged [][]byte
	for _, l := range p {
		smudged = append(smudged, []byte(l))
	}
	if smudged[i][j] == '#' {
		smudged[i][j] = '.'
	} else {
		smudged[i][j] = '#'
	}
	return smudged
}

func smudge(p []string) [][][]byte {
	var smudges [][][]byte
	for i, l := range p {
		for j := range l {
			smudges = append(smudges, smudgeAt(i, j, p))
		}
	}
	return smudges
}

func pristine(p []string) [][]byte {
	var pristine [][]byte
	for _, l := range p {
		pristine = append(pristine, []byte(l))
	}
	return pristine
}

func exclude(a []int, b int) []int {
	var res []int
	for _, i := range a {
		if i != b {
			res = append(res, i)
		}
	}
	return res
}

func solve(patterns [][]string) int {
	res := 0
	for _, p := range patterns {
		pscore := findReflections(pristine(p))[0]
		for _, sp := range smudge(p) {
			if s := exclude(findReflections(sp), pscore); len(s) > 0 {
				res += s[0]
				break
			}
		}
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	pdata := strings.Split(data, "\n\n")
	var patterns [][]string
	for _, p := range pdata {
		patterns = append(patterns, strings.Split(p, "\n"))
	}
	fmt.Println(solve(patterns))
}
