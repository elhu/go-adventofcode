package main

import (
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

const vectorY = 1
const vectorX = 3
const treeSym = '#'

type coords struct {
	x, y int
}

func solve(trees [][]byte) int {
	res := 0
	pos := coords{0, 0}
	for pos.y+vectorY < len(trees) {
		pos.x += vectorX
		pos.x = pos.x % len(trees[0])
		pos.y += vectorY
		if trees[pos.y][pos.x] == treeSym {
			res++
		}
	}
	return res
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	var trees [][]byte
	for i, l := range lines {
		trees = append(trees, make([]byte, len(l)))
		for j, c := range []byte(l) {
			trees[i][j] = c
		}
	}
	fmt.Println(solve(trees))
}
