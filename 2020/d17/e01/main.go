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

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	check(err)
	return i
}

type Coord3d struct {
	x, y, z int
}

func toKey(c Coord3d) string {
	return fmt.Sprintf("%d:%d:%d", c.x, c.y, c.z)
}

func surroundingKeys(c Coord3d) []string {
	var res []string
	for i := c.x - 1; i <= c.x+1; i++ {
		for j := c.y - 1; j <= c.y+1; j++ {
			for k := c.z - 1; k <= c.z+1; k++ {
				if i != c.x || j != c.y || k != c.z {
					res = append(res, fmt.Sprintf("%d:%d:%d", i, j, k))
				}
			}
		}
	}
	return res
}

func fromKey(str string) Coord3d {
	parts := strings.Split(str, ":")
	return Coord3d{atoi(parts[0]), atoi(parts[1]), atoi(parts[2])}
}

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func minMax(space map[string]struct{}) (int, int, int, int, int, int) {
	var minX, minY, minZ int = MaxInt, MaxInt, MaxInt
	var maxX, maxY, maxZ int = MinInt, MinInt, MinInt
	for k := range space {
		coords := fromKey(k)
		if coords.x < minX {
			minX = coords.x
		}
		if coords.x > maxX {
			maxX = coords.x
		}
		if coords.y < minY {
			minY = coords.y
		}
		if coords.y > maxY {
			maxY = coords.y
		}
		if coords.z < minZ {
			minZ = coords.z
		}
		if coords.z > maxZ {
			maxZ = coords.z
		}
	}
	return minX, maxX, minY, maxY, minZ, maxZ
}

func playTurn(space map[string]struct{}) map[string]struct{} {
	newSpace := make(map[string]struct{})
	minX, maxX, minY, maxY, minZ, maxZ := minMax(space)
	for x := minX - 1; x <= maxX+1; x++ {
		for y := minY - 1; y <= maxY+1; y++ {
			for z := minZ - 1; z <= maxZ+1; z++ {
				currPos := Coord3d{x, y, z}
				currKey := toKey(currPos)
				neighborCount := 0
				for _, k := range surroundingKeys(currPos) {
					if _, exists := space[k]; exists {
						neighborCount++
					}
				}

				_, active := space[currKey]
				if active && (neighborCount == 2 || neighborCount == 3) {
					newSpace[currKey] = struct{}{}
				} else if !active && neighborCount == 3 {
					newSpace[currKey] = struct{}{}
				}
			}
		}
	}
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
				space[toKey(Coord3d{x, y, 0})] = struct{}{}
			}
		}
	}
	fmt.Println(solve(space, 6))
}
