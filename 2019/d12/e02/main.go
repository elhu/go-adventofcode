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

func getFrequencies(moons []Moon) (int, int, int) {
	targetX := hashX(moons)
	targetY := hashY(moons)
	targetZ := hashZ(moons)
	freqX, freqY, freqZ := -1, -1, -1
	for step := 1; ; step++ {
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
		if freqX == -1 && hashX(moons) == targetX {
			freqX = step
		}
		if freqY == -1 && hashY(moons) == targetY {
			freqY = step
		}
		if freqZ == -1 && hashZ(moons) == targetZ {
			freqZ = step
		}
		if freqX != -1 && freqY != -1 && freqZ != -1 {
			return freqX, freqY, freqZ
		}
	}
}

func hashX(ms []Moon) string {
	res := make([]string, len(ms))
	for i := 0; i < len(ms); i++ {
		res[i] = fmt.Sprintf("%d:%d", ms[i].pos.x, ms[i].vec.x)
	}
	return strings.Join(res, "|")
}

func hashY(ms []Moon) string {
	res := make([]string, len(ms))
	for i := 0; i < len(ms); i++ {
		res[i] = fmt.Sprintf("%d:%d", ms[i].pos.y, ms[i].vec.y)
	}
	return strings.Join(res, "|")
}

func hashZ(ms []Moon) string {
	res := make([]string, len(ms))
	for i := 0; i < len(ms); i++ {
		res[i] = fmt.Sprintf("%d:%d", ms[i].pos.z, ms[i].vec.z)
	}
	return strings.Join(res, "|")
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func solve(moons []Moon) int {
	xFreq, yFreq, zFreq := getFrequencies(moons)
	return lcm(xFreq, yFreq, zFreq)
}

var moonExp = regexp.MustCompile(`<x=(-?\d+), y=(-?\d+), z=(-?\d+)>`)

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	rawMoons := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	moons := parse(rawMoons)
	fmt.Println(solve(moons))
}
