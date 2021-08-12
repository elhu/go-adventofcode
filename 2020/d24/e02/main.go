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

var neighboringVectors = [][]int{
	[]int{1, -1, 0},
	[]int{0, -1, 1},
	[]int{-1, 0, 1},
	[]int{-1, 1, 0},
	[]int{0, 1, -1},
	[]int{1, 0, -1},
}

func surroundingTiles(coords string) []string {
	var x, y, z int
	fmt.Sscanf(coords, "%d:%d:%d", &x, &y, &z)
	res := make([]string, len(neighboringVectors))
	for i, v := range neighboringVectors {
		res[i] = fmt.Sprintf("%d:%d:%d", x+v[0], y+v[1], z+v[2])
	}
	return res
}

func turn(tiles map[string]struct{}) map[string]struct{} {
	toCheck := make(map[string]struct{})
	// Put all the tiles to check in a set
	for tile := range tiles {
		for _, adj := range surroundingTiles(tile) {
			toCheck[adj] = struct{}{}
		}
	}

	// For each tile, count the neighbors, apply rule
	newTiles := make(map[string]struct{})
	for tile := range toCheck {
		blackCount := 0
		for _, adj := range surroundingTiles(tile) {
			if _, found := tiles[adj]; found {
				blackCount++
			}
		}
		_, tileIsBlack := tiles[tile]
		if !tileIsBlack && blackCount == 2 {
			newTiles[tile] = struct{}{}
		} else if tileIsBlack && (blackCount == 0 || blackCount > 2) {
			delete(newTiles, tile)
		} else if tileIsBlack {
			newTiles[tile] = struct{}{}
		}
	}
	return newTiles
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
	for i := 1; i <= 100; i++ {
		flippedTiles = turn(flippedTiles)
	}
	fmt.Println(len(flippedTiles))
}
