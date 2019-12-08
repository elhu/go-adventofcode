package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func buildLayers(data []byte, w, h int) [][]int {
	layers := len(data) / (w * h)
	image := make([][]int, layers)
	for l := 0; l < layers; l++ {
		image[l] = make([]int, w*h)
		for i := 0; i < w*h; i++ {
			image[l][i] = int(data[l*w*h+i] - '0')
		}
	}
	return image
}

const maxInt = int(^uint(0) >> 1)

func solve(image [][]int) int {
	minZero := maxInt
	layer := 0
	for k, l := range image {
		zeros := 0
		for _, i := range l {
			if i == 0 {
				zeros++
			}
		}
		if zeros < minZero {
			minZero = zeros
			layer = k
		}
	}
	var ones, twos int
	for _, i := range image[layer] {
		if i == 1 {
			ones++
		} else if i == 2 {
			twos++
		}
	}
	return ones * twos
}

const imgWidth = 25
const imgHeight = 6

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	image := buildLayers(data, imgWidth, imgHeight)
	fmt.Println(solve(image))
}
