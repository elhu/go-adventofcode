package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func litNeighbors(x, y int, lights [][]byte) int {
	res := 0
	neighbors := [8][2]int{
		{x - 1, y - 1}, {x - 1, y}, {x - 1, y + 1},
		{x, y - 1}, {x, y + 1},
		{x + 1, y - 1}, {x + 1, y}, {x + 1, y + 1},
	}
	for _, c := range neighbors {
		j, i := c[0], c[1]
		if i >= 0 && i < len(lights) && j >= 0 && j < len(lights[i]) {
			if lights[i][j] == '#' {
				res++
			}
		}
	}
	return res
}

func step(lights [][]byte) [][]byte {
	newSetup := make([][]byte, len(lights))
	for i := range lights {
		newSetup[i] = make([]byte, len(lights[i]))
	}

	for i := range lights {
		for j := range lights[i] {
			ln := litNeighbors(j, i, lights)
			if lights[i][j] == '#' {
				if ln == 2 || ln == 3 {
					newSetup[i][j] = '#'
				} else {
					newSetup[i][j] = '.'
				}
			} else {
				if ln == 3 {
					newSetup[i][j] = '#'
				} else {
					newSetup[i][j] = '.'
				}
			}
		}
	}
	return newSetup
}

func countLights(lights [][]byte) int {
	res := 0
	for _, l := range lights {
		for _, c := range l {
			if c == '#' {
				res++
			}
		}
	}
	return res
}

func solve(lights [][]byte) int {
	for i := 0; i < 100; i++ {
		lights = step(lights)
	}
	return countLights(lights)
}

func main() {
	rawData, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := bytes.Split(bytes.TrimRight(rawData, "\n"), []byte("\n"))
	fmt.Println(solve(input))
}
