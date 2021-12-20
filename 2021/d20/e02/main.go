package main

import (
	"adventofcode/utils/files"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func pad(data [][]byte, size int, padChar string) [][]byte {
	res := make([][]byte, len(data)+2*size)
	for i := range res {
		res[i] = make([]byte, len(data[0])+2*size)
		res[i] = bytes.Repeat([]byte(padChar), len(data[0])+2*size)
		if i >= size && i < len(data)+size {
			copy(res[i][size:], data[i-size])
		}
	}
	return res
}

func printImage(image [][]byte) {
	for _, l := range image {
		fmt.Println(string(l))
	}
	fmt.Println("--")
}

func decodeIndex(image [][]byte, y, x int) int {
	data := make([]byte, 0, 9)
	data = append(data, image[y-1][x-1:x+2]...)
	data = append(data, image[y][x-1:x+2]...)
	data = append(data, image[y+1][x-1:x+2]...)

	data = bytes.ReplaceAll(data, []byte{'.'}, []byte{'0'})
	data = bytes.ReplaceAll(data, []byte{'#'}, []byte{'1'})
	n, err := strconv.ParseInt(string(data), 2, 64)
	if err != nil {
		panic(fmt.Errorf("Couldn't parse binary number %s\n", data))
	}
	return int(n)
}

func copyImage(image [][]byte) [][]byte {
	cpy := make([][]byte, len(image))
	for i, l := range image {
		cpy[i] = make([]byte, len(l))
		copy(cpy[i], l)
	}
	return cpy
}

func score(image [][]byte) int {
	res := 0
	for _, l := range image {
		res += bytes.Count(l, []byte{'#'})
	}
	return res
}

func solve(image [][]byte, algo string, rounds int) int {
	fillers := []byte{algo[0], algo[len(algo)-1]}
	fillerIdx := 0
	for r := 0; r < rounds; r++ {
		tmp := copyImage(image)
		for i := 1; i < len(image)-1; i++ {
			for j := 1; j < len(image[i])-1; j++ {
				idx := decodeIndex(image, i, j)
				tmp[i][j] = algo[idx]
			}
		}
		currentFiller := fillers[fillerIdx]
		tmp[0] = bytes.Repeat([]byte{currentFiller}, len(tmp[0]))
		tmp[len(tmp)-1] = bytes.Repeat([]byte{currentFiller}, len(tmp[0]))
		for i := range tmp {
			tmp[i][0] = currentFiller
			tmp[i][len(tmp[i])-1] = currentFiller
		}
		image = pad(tmp, 1, string(currentFiller))
		fillerIdx = (fillerIdx + 1) % 2
	}
	return score(image)
}

func main() {
	data := files.ReadLinesWithSeparator(os.Args[1], "\n\n")
	algo := data[0]
	image := pad(bytes.Split([]byte(data[1]), []byte{'\n'}), 2, ".")
	fmt.Println(solve(image, algo, 50))
}
