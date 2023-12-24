package main

import (
	"adventofcode/utils/coords/coords3d"
	"adventofcode/utils/files"
	"fmt"
	"math"
	"os"
	"strings"
)

type Hail struct {
	pos        coords3d.Coords3d
	vec        coords3d.Coords3d
	slopeX     float64
	interceptX float64
	slopeZ     float64
	interceptZ float64
}

func parseHail(line string) Hail {
	var h Hail
	fmt.Sscanf(line, "%d, %d, %d @ %d, %d, %d", &h.pos.X, &h.pos.Y, &h.pos.Z, &h.vec.X, &h.vec.Y, &h.vec.Z)
	nextPos := coords3d.Add(h.pos, h.vec)
	h.slopeX = float64(nextPos.Y-h.pos.Y) / float64(nextPos.X-h.pos.X)
	h.interceptX = float64(h.pos.Y) - h.slopeX*float64(h.pos.X)
	h.slopeZ = float64(nextPos.Y-h.pos.Y) / float64(nextPos.Z-h.pos.Z)
	h.interceptZ = float64(h.pos.Y) - h.slopeZ*float64(h.pos.Z)
	return h
}

func intersects(a, b Hail) (bool, coords3d.Coords3d) {
	leftX := a.slopeX - b.slopeX
	rightX := b.interceptX - a.interceptX
	x := rightX / leftX
	y := a.slopeX*x + a.interceptX
	if x < math.Inf(1) && x > math.Inf(-1) && y < math.Inf(1) && y > math.Inf(-1) {
		za := (y - a.interceptZ) / a.slopeZ
		zb := (y - b.interceptZ) / b.slopeZ
		if za == zb || math.IsNaN(za) || math.IsNaN(zb) {
			z := za
			if math.IsNaN(za) {
				z = zb
			}
			return true, coords3d.Coords3d{X: int(x), Y: int(y), Z: int(z)}
		}
	}
	if a.slopeX == b.slopeX && a.interceptX == b.interceptX {
		leftZ := a.slopeZ - b.slopeZ
		rightZ := b.interceptZ - a.interceptZ
		z := rightZ / leftZ
		ya := a.slopeZ*z + a.interceptZ
		yb := b.slopeZ*z + b.interceptZ
		if ya == yb {
			x := (ya - a.interceptX) / a.slopeX
			return true, coords3d.Coords3d{X: int(math.Round(x)), Y: int(math.Round(ya)), Z: int(math.Round(z))}
		}
		if debug {
			fmt.Println("THE HARD STUFF")
			fmt.Println(a)
			fmt.Println(b)
			leftZ := a.slopeZ - b.slopeZ
			rightZ := b.interceptZ - a.interceptZ
			fmt.Println(leftX, rightX, x, y, leftZ, rightZ)
			z := rightZ / leftZ
			ya := a.slopeZ*z + a.interceptZ
			yb := b.slopeZ*z + b.interceptZ
			fmt.Println(ya, yb)
		}
	}
	return false, coords3d.Coords3d{}
}

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func preprocess(hailstones []Hail) ([]int, []int, []int) {
	validX := make(map[int]struct{})
	validY := make(map[int]struct{})
	validZ := make(map[int]struct{})
	for i := -AMPLITUDE; i <= AMPLITUDE; i++ {
		validX[i] = struct{}{}
		validY[i] = struct{}{}
		validZ[i] = struct{}{}
	}
	for i, a := range hailstones {
		for _, b := range hailstones[i+1:] {
			if a.pos.X > b.pos.X && a.vec.X > b.vec.X {
				min, max := minMax(a.vec.X, b.vec.X)
				for i := min; i <= max; i++ {
					delete(validX, i)
				}
			}
			if a.pos.Y > b.pos.Y && a.vec.Y > b.vec.Y {
				min, max := minMax(a.vec.Y, b.vec.Y)
				for i := min; i <= max; i++ {
					delete(validY, i)
				}
			}
			if a.pos.Z > b.pos.Z && a.vec.Z > b.vec.Z {
				min, max := minMax(a.vec.Z, b.vec.Z)
				for i := min; i <= max; i++ {
					delete(validZ, i)
				}
			}
		}
	}
	var resX, resY, resZ []int
	for i := -AMPLITUDE; i <= AMPLITUDE; i++ {
		if _, found := validX[i]; found {
			resX = append(resX, i)
		}
		if _, found := validY[i]; found {
			resY = append(resY, i)
		}
		if _, found := validZ[i]; found {
			resZ = append(resZ, i)
		}
	}
	return resX, resY, resZ
}

func shiftHail(hailstones []Hail, shift coords3d.Coords3d) []Hail {
	var res []Hail
	for _, oh := range hailstones {
		h := Hail{
			pos: oh.pos,
			vec: coords3d.Coords3d{X: oh.vec.X - shift.X, Y: oh.vec.Y - shift.Y, Z: oh.vec.Z - shift.Z},
		}
		nextPos := coords3d.Add(h.pos, h.vec)
		h.slopeX = float64(nextPos.Y-h.pos.Y) / float64(nextPos.X-h.pos.X)
		h.interceptX = float64(h.pos.Y) - h.slopeX*float64(h.pos.X)
		h.slopeZ = float64(nextPos.Y-h.pos.Y) / float64(nextPos.Z-h.pos.Z)
		h.interceptZ = float64(h.pos.Y) - h.slopeZ*float64(h.pos.Z)
		res = append(res, h)
	}
	if debug {
		fmt.Println(res)
	}
	return res
}

func eq(a, b coords3d.Coords3d) bool {
	if a == b {
		return true
	}
	return a.X == b.X && a.Y == b.Y && a.Z == 0 || b.Z == 0
}

var debug = false

func solve(hailstones []Hail) int {
	validX, validY, validZ := preprocess(hailstones)
	fmt.Println(len(validX), len(validY), len(validZ))
	for _, x := range validX {
		for _, y := range validY {
			for _, z := range validZ {
				debug = false
				if x == -3 && y == 1 && z == 2 {
					debug = true
				}
				sh := shiftHail(hailstones, coords3d.Coords3d{X: x, Y: y, Z: z})
				if debug {
					fmt.Println(sh[0])
					fmt.Println(sh[1])
				}
				pf, p := intersects(sh[0], sh[1])
				if !pf {
					if debug {
						fmt.Println(pf, p)
						fmt.Println("FAIL", x, y, z)
					}
					continue
				}
				fmt.Println(pf)
				var pairs [][2]Hail
				for i, a := range sh[1:] {
					for _, b := range sh[i+2:] {
						pairs = append(pairs, [2]Hail{a, b})
					}
				}
				success := true
				for _, pair := range pairs {
					rf, r := intersects(pair[0], pair[1])
					if !rf || !eq(r, p) {
						if debug {
							fmt.Printf("%v %v rf: %v eq: %v\n", pair[0], pair[1], rf, eq(r, p))
							fmt.Println("FAIL2", rf, r)
						}
						success = false
						break
					}
				}
				if success {
					fmt.Println("SUCCESS", p.X, p.Y, p.Z)
					return p.X + p.Y + p.Z
				}
			}
		}
	}
	return 0
}

const (
	// MIN = 200000000000000
	// MAX = 400000000000000
	// Assume that the velocities will be in the -500..500 range
	AMPLITUDE = 00
	// AMPLITUDE = 25
)

func solve2(hailStones []Hail) int {
	fmt.Println("Plug the following equations into https://www.wolframalpha.com/input?i=system+equation+calculator")
	for i := 0; i < 4; i++ {
		h := hailStones[i]
		str := fmt.Sprintf("(X-%d)/(A-%d)=(Y-%d)/(B-%d)=(Z-%d)/(C-%d)", h.pos.X, h.vec.X, h.pos.Y, h.vec.Y, h.pos.Z, h.vec.Z)
		str = strings.ReplaceAll(str, "--", "+")
		fmt.Println(str)
	}
	// A = 44 and B = 305 and C = 75 and X = 234382970331570 and Y = 100887864960615 and Z = 231102671115832
	return 234382970331570 + 100887864960615 + 231102671115832
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	var hailStones []Hail
	for _, line := range lines {
		hailStones = append(hailStones, parseHail(line))
	}
	// fmt.Println(solve(hailStones))
	fmt.Println(solve2(hailStones))
}
