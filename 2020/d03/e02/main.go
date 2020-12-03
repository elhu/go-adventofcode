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

const treeSym = '#'

var vectors = []coords{
	coords{1, 1},
	coords{3, 1},
	coords{5, 1},
	coords{7, 1},
	coords{1, 2},
}

type coords struct {
	x, y int
}

func solve(trees [][]byte) int {
	total := 1
	for _, v := range vectors {
		res := 0
		pos := coords{0, 0}
		for pos.y+v.y < len(trees) {
			pos.x += v.x
			pos.x = pos.x % len(trees[0])
			pos.y += v.y
			if trees[pos.y][pos.x] == treeSym {
				res++
			}
		}
		total *= res
	}
	return total
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
