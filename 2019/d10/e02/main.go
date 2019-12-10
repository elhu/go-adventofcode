package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"sort"
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

const coordFactor = 10000000000

func coordsToKey(c coords) int {
	return c.y*coordFactor + c.x
}

func keyToCoords(i int) coords {
	return coords{i % coordFactor, i / coordFactor}
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

func getAsteroidsInSight(asteroids map[int]int) []coords {
	base := coordsToKey(baseCoords)
	result := make([]coords, 0)
	for target, v := range asteroids {
		if v == 1 && target != base {
			bc := keyToCoords(base)
			tc := keyToCoords(target)
			vec := vector(bc, tc)
			inSight := true
			bc.x += vec.x
			bc.y += vec.y
			for bc.x != tc.x || bc.y != tc.y {
				if _, exists := asteroids[coordsToKey(bc)]; exists {
					inSight = false
				}
				bc.x += vec.x
				bc.y += vec.y
			}
			if inSight {
				result = append(result, bc)
			}
		}
	}
	return result
}

func angle(c coords) float64 {
	res := math.Atan2(float64(baseCoords.y-c.y), float64(baseCoords.x-c.x)) - math.Pi/2
	if res < 0 {
		res += 2 * math.Pi
	}
	return res
}

var baseCoords = coords{23, 29}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	asteroidsInSight := getAsteroidsInSight(parse(lines))
	sort.SliceStable(asteroidsInSight, func(i, j int) bool {
		return angle(asteroidsInSight[i]) < angle(asteroidsInSight[j])
	})
	bet := asteroidsInSight[199]
	fmt.Println(bet.x*100 + bet.y)
}
