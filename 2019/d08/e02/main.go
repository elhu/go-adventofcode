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

func flattenImage(image [][]int) []int {
	res := make([]int, len(image[0]))
	for i := 0; i < len(image[0]); i++ {
		// Set pixel to transparent to start with
		res[i] = 2
		for l := 0; l < len(image); l++ {
			if image[l][i] != 2 {
				res[i] = image[l][i]
				break
			}
		}
	}
	return res
}

func print(layer []int) {
	for i := 0; i < imgHeight; i++ {
		for j := 0; j < imgWidth; j++ {
			if layer[i*imgWidth+j] != 1 {
				fmt.Printf(" ")
			} else {
				fmt.Printf("%d", layer[i*imgWidth+j])
			}
		}
		fmt.Printf("\n")
	}
}

func solve(image [][]int) {
	flat := flattenImage(image)
	print(flat)
}

const imgWidth = 25
const imgHeight = 6

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	image := buildLayers(data, imgWidth, imgHeight)
	// for _, l := range image {
	// 	print(l)
	// }
	solve(image)
}
