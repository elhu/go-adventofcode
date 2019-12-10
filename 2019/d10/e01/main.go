package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(fh *bufio.Reader, c chan string) {
	for {
		line, err := fh.ReadString('\n')
		c <- strings.TrimSuffix(line, "\n")
		if err == io.EOF {
			break
		}
	}
	close(c)
}

type coords struct {
	x, y int
}

func coordsToKey(c coords) int {
	return c.y*10000 + c.x
}

func keyToCoords(i int) coords {
	return coords{i % 10000, i / 10000}
}

func parse(data []string) map[int]int {
	res := make(map[int]int)
	for i, line := range data {
		for j, c := range line {
			if c == '#' {
				res[coordsToKey(coords{j, i})] = 1
			}
		}
	}
	return res
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	if a < 0 {
		return -a
	}
	return a
}

func vector(a, b coords) coords {
	x := b.x - a.x
	y := b.y - a.y
	d := gcd(x, y)
	return coords{x / d, y / d}
}

func solve(asteroids map[int]int) int {
	max := 0
	var maxBase coords
	for base := range asteroids {
		for target, v := range asteroids {
			if v == 1 && target != base {
				bc := keyToCoords(base)
				tc := keyToCoords(target)
				vec := vector(bc, tc)
				inSight := true
				bc.x += vec.x
				bc.y += vec.y
				for bc.x != tc.x || bc.y != tc.y {
					if v, exists := asteroids[coordsToKey(bc)]; exists && v == 1 {
						if inSight {
							inSight = false
						} else {
							asteroids[target] = 0
						}
					}
					bc.x += vec.x
					bc.y += vec.y
				}
				if !inSight {
					asteroids[target] = 0
				} else {
				}
			}
		}
		// Remove base point from count
		count := -1
		for _, v := range asteroids {
			if v == 1 {
				count++
			}
		}
		// fmt.Printf("%v has %d in sight\n", keyToCoords(base), count)
		if count > max {
			max = count
			maxBase = keyToCoords(base)
		}
		// reset map
		for k := range asteroids {
			asteroids[k] = 1
		}
	}
	fmt.Printf("Base found: %v\n", maxBase)
	return max
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	fmt.Println(solve(parse(lines)))
	// fmt.Println(gcd(-2, 3))
}
