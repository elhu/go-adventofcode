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

func getStartPosition(seed int, maps []Map) int {
	for i := len(maps) - 1; i >= 0; i-- {
		m := maps[i]
		for _, r := range m.ranges {
			if seed >= r.minDest && seed < r.minDest+r.length {
				seed = r.minSource + (seed - r.minDest)
				break
			}
		}
	}
	return seed
}

func solve(seeds []int, maps []Map) int {
	for i := 0; ; i++ {
		startPos := getStartPosition(i, maps)
		for s := 0; s < len(seeds); s += 2 {
			if startPos >= seeds[s] && startPos < seeds[s]+seeds[s+1] {
				return i
			}
		}
	}
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
