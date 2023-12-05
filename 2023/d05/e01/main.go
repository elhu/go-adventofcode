package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	check(err)
	return n
}

type Range struct {
	minDest   int
	minSource int
	length    int
}

type Map struct {
	ranges []Range
}

func parseMap(data []string) Map {
	m := Map{}
	for _, line := range data[1:] {
		r := Range{}
		fmt.Sscanf(line, "%d %d %d", &r.minDest, &r.minSource, &r.length)
		m.ranges = append(m.ranges, r)
	}
	return m
}

func parseSeeds(data string) []int {
	seeds := make([]int, 0)
	for _, seed := range strings.Fields(data)[1:] {
		seeds = append(seeds, atoi(seed))
	}
	return seeds
}

func min(arr []int) int {
	res := arr[0]
	for _, v := range arr[1:] {
		if v < res {
			res = v
		}
	}
	return res
}

func getSeedPosition(seed int, maps []Map) int {
	pos := seed
	for _, m := range maps {
		for _, r := range m.ranges {
			if pos >= r.minSource && pos < r.minSource+r.length {
				pos = r.minDest + (pos - r.minSource)
				break
			}
		}
	}
	return pos
}

func solve(seeds []int, maps []Map) int {
	seedPos := make([]int, 0)
	for _, seed := range seeds {
		p := getSeedPosition(seed, maps)
		seedPos = append(seedPos, p)
	}
	return min(seedPos)
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	mapsData := strings.Split(data, "\n\n")
	maps := make([]Map, 0)
	seeds := parseSeeds(mapsData[0])
	for _, mapData := range mapsData[1:] {
		maps = append(maps, parseMap(strings.Split(mapData, "\n")))
	}
	fmt.Println(solve(seeds, maps))
}
