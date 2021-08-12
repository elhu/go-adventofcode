package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	check(err)
	return i
}

type CoordNd []int

func toKey(c CoordNd) string {
	var res []byte
	for _, n := range c {
		res = strconv.AppendInt(res, int64(n), 10)
		res = append(res, ':')
	}
	return string(res[:len(res)-1])
}

func surroundingKeys(c CoordNd, pos int, res []string) []string {
	currentCoord := make(CoordNd, len(c))
	copy(currentCoord, c)
	for i := c[pos] - 1; i <= c[pos]+1; i++ {
		currentCoord[pos] = i
		if pos == len(c)-1 {
			res = append(res, toKey(currentCoord))
		} else {
			res = surroundingKeys(currentCoord, pos+1, res)
		}
	}
	return res
}

func fromKey(str string) CoordNd {
	parts := strings.Split(str, ":")
	res := make(CoordNd, len(parts))
	for i, p := range parts {
		res[i] = atoi(p)
	}
	return res
}

var wg sync.WaitGroup

func playTurn(space map[string]struct{}) map[string]struct{} {
	newSpace := make(map[string]struct{})
	var mu sync.Mutex
	allKeys := make(map[string]struct{})
	for c := range space {
		for _, k := range surroundingKeys(fromKey(c), 0, make([]string, 0)) {
			allKeys[k] = struct{}{}
		}
	}
	for currKey := range allKeys {
		wg.Add(1)
		go func(currKey string) {
			defer wg.Done()
			neighborCount := 0
			currPos := fromKey(currKey)
			for _, k := range surroundingKeys(currPos, 0, make([]string, 0)) {
				if _, exists := space[k]; currKey != k && exists {
					neighborCount++
				}
			}

			_, active := space[currKey]
			if (!active && neighborCount == 3) || (active && (neighborCount == 2 || neighborCount == 3)) {
				mu.Lock()
				newSpace[currKey] = struct{}{}
				mu.Unlock()
			}
		}(currKey)
	}
	wg.Wait()
	return newSpace
}

func solve(space map[string]struct{}, cycles int) int {
	for i := 0; i < cycles; i++ {
		space = playTurn(space)
	}
	return len(space)
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	space := make(map[string]struct{})
	for y, l := range input {
		for x, c := range l {
			if c == '#' {
				space[toKey(CoordNd{x, y, 0, 0, 0, 0})] = struct{}{}
			}
		}
	}
	fmt.Println(solve(space, 6))
}
