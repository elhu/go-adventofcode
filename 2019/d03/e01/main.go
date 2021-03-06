package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type coord struct {
	x, y int
}

func recordPos(pos coord, path map[string]struct{}) {
	key := fmt.Sprintf("%d:%d", pos.x, pos.y)
	path[key] = struct{}{}
}

func wireToPath(wire []string) map[string]struct{} {
	path := make(map[string]struct{})
	pos := coord{0, 0}
	for _, elem := range wire {
		val, err := strconv.Atoi(elem[1:])
		check(err)
		for i := 0; i < val; i++ {
			switch elem[0] {
			case 'R':
				pos.x++
			case 'L':
				pos.x--
			case 'U':
				pos.y++
			case 'D':
				pos.y--
			}
			recordPos(pos, path)
		}
	}
	return path
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func keyToDistance(key string) int {
	parts := strings.Split(key, ":")
	x, err := strconv.Atoi(parts[0])
	check(err)
	y, err := strconv.Atoi(parts[1])
	check(err)

	return abs(x) + abs(y)
}

func intersect(a, b map[string]struct{}) map[string]struct{} {
	res := make(map[string]struct{})
	for k := range a {
		if _, exists := b[k]; exists {
			res[k] = struct{}{}
		}
	}
	return res
}

func solve(wireA, wireB []string) int {
	pathA := wireToPath(wireA)
	pathB := wireToPath(wireB)
	// hack max value
	path := intersect(pathA, pathB)
	min := 999999999999999999
	for k := range path {
		if dist := keyToDistance(k); dist < min {
			min = dist
		}
	}
	return min
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	rawLines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	wireA := strings.Split(rawLines[0], ",")
	wireB := strings.Split(rawLines[1], ",")

	fmt.Println(solve(wireA, wireB))
}
