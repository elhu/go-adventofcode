package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func pathToCoord(path string) string {
	coords := make(map[string]int)
	for i := 0; i < len(path); i++ {
		dir := path[i : i+1]
		if dir == "s" || dir == "n" {
			dir = path[i : i+2]
			i++
		}
		switch dir {
		case "e":
			coords["x"]++
			coords["y"]--
		case "se":
			coords["y"]--
			coords["z"]++
		case "sw":
			coords["x"]--
			coords["z"]++
		case "w":
			coords["x"]--
			coords["y"]++
		case "nw":
			coords["y"]++
			coords["z"]--
		case "ne":
			coords["x"]++
			coords["z"]--
		default:
			panic(fmt.Sprintf("Parsing Error, found %s", dir))
		}
	}
	return fmt.Sprintf("%d:%d:%d", coords["x"], coords["y"], coords["z"])
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	flippedTiles := make(map[string]struct{})
	for _, l := range input {
		coords := pathToCoord(l)
		if _, found := flippedTiles[coords]; found {
			delete(flippedTiles, coords)
		} else {
			flippedTiles[coords] = struct{}{}
		}
	}
	fmt.Println(len(flippedTiles))
}
