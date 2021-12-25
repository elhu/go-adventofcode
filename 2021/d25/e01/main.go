package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
)

func moveEast(data [][]byte, copy [][]byte) int {
	moved := 0
	for i := 0; i < len(data); i++ {
		l := data[i]
		for j := 0; j < len(l); j++ {
			if data[i][j] == '>' {
				if data[i][(j+1)%len(l)] == '.' {
					copy[i][(j+1)%len(l)] = '>'
					copy[i][j] = '.'
					moved++
				}
			}
		}
	}
	return moved
}

func moveSouth(data [][]byte, copy [][]byte) int {
	moved := 0
	for i := len(data) - 1; i >= 0; i-- {
		l := data[i]
		for j := 0; j < len(l); j++ {
			if data[i][j] == 'v' {
				if data[(i+1)%len(data)][j] == '.' {
					copy[(i+1)%len(data)][j] = 'v'
					copy[i][j] = '.'
					moved++
				}
			}
		}
	}
	return moved
}

func copyMap(data [][]byte) [][]byte {
	res := make([][]byte, len(data))
	for i, c := range data {
		res[i] = make([]byte, len(c))
		copy(res[i], c)
	}
	return res
}

func printMap(data [][]byte) {
	for _, l := range data {
		fmt.Printf("%s\n", l)
	}
	fmt.Println("--")
}

func solve(data [][]byte) int {
	for i := 0; ; i++ {
		moved := 0
		copy := copyMap(data)
		moved += moveEast(data, copy)
		data = copy
		copy = copyMap(data)
		moved += moveSouth(data, copy)
		data = copy
		if moved == 0 {
			return i + 1
		}
	}
}

func main() {
	data := files.ReadLines(os.Args[1])
	input := make([][]byte, len(data))
	for i, l := range data {
		input[i] = []byte(l)
	}
	fmt.Println(solve([][]byte(input)))
}
