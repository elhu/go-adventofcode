package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Coords3d struct {
	x, y, z int
}

type Moon struct {
	pos Coords3d
	vec Coords3d
}

func unsafeAtoi(s string) int {
	num, err := strconv.Atoi(s)
	check(err)
	return num
}

func parse(raw []string) []Moon {
	moons := make([]Moon, 0, len(raw))
	for _, line := range raw {
		match := moonExp.FindStringSubmatch(line)
		pos := Coords3d{unsafeAtoi(match[1]), unsafeAtoi(match[2]), unsafeAtoi(match[3])}
		moons = append(moons, Moon{pos, Coords3d{}})
	}
	return moons
}

func delta(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func addVector(a, b Coords3d) Coords3d {
	return Coords3d{a.x + b.x, a.y + b.y, a.z + b.z}
}

func dump(moons []Moon, step int) {
	fmt.Printf("After %3d steps:\n", step)
	for _, m := range moons {
		fmt.Printf("pos=<x=%4d, y=%4d, z=%4d>, vel=<x=%4d, y=%4d, z=%4d>\n", m.pos.x, m.pos.y, m.pos.z, m.vec.x, m.vec.y, m.vec.z)
	}
	fmt.Printf("\n")
}

func solve(moons []Moon) int {
	for step := 0; step < 1000; step++ {
		// dump(moons, step)
		gravityDelta := make([]Coords3d, 4)
		for i, m := range moons {
			for j := i + 1; j < len(moons); j++ {
				o := moons[j]
				dx := delta(m.pos.x, o.pos.x)
				dy := delta(m.pos.y, o.pos.y)
				dz := delta(m.pos.z, o.pos.z)
				gravityDelta[i].x -= dx
				gravityDelta[i].y -= dy
				gravityDelta[i].z -= dz
				gravityDelta[j].x += dx
				gravityDelta[j].y += dy
				gravityDelta[j].z += dz
			}
		}
		for i, d := range gravityDelta {
			moons[i].vec = addVector(moons[i].vec, d)
			moons[i].pos = addVector(moons[i].pos, moons[i].vec)
		}
	}
	// dump(moons, 100)
	return totalEnergy(moons)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func totalEnergy(moons []Moon) int {
	res := 0
	for _, moon := range moons {
		pot := abs(moon.pos.x) + abs(moon.pos.y) + abs(moon.pos.z)
		kin := abs(moon.vec.x) + abs(moon.vec.y) + abs(moon.vec.z)
		res += pot * kin
	}
	return res
}

var moonExp = regexp.MustCompile(`<x=(-?\d+), y=(-?\d+), z=(-?\d+)>`)

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	rawMoons := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	moons := parse(rawMoons)
	fmt.Println(solve(moons))
}
