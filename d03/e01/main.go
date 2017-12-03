package main

import "os"
import "fmt"
import "strconv"

const (
	right  = iota
	top    = iota
	left   = iota
	bottom = iota
)

func wtfIntAbs(value int) int {
	if value < 0 {
		return -value
	} else {
		return value
	}
}

func ringStart(index int) int {
	return (index-1)*(index)/2*8 + 2
}

func ringEnd(index int) int {
	return index*(index+1)/2*8 + 1
}

func findRing(target int) int {
	current := 0
	idx := 0
	for ; current < target; idx++ {
		current = ringEnd(idx)
	}
	return idx - 1
}

func solve(target int) int {
	ring := findRing(target)
	if ring == 0 {
		return 0
	}

	side := (target - ringStart(ring)) / (ring * 2)

	sideStart := ringStart(ring) + side*ring*2
	sideEnd := ringStart(ring) + (side+1)*ring*2 - 1

	sideMiddle := (sideEnd-sideStart)/2 + sideStart

	x, y := 0, 0

	switch side {
	case right:
		x = ring
		y = wtfIntAbs(sideMiddle - target)
	case top:
		x = wtfIntAbs(sideMiddle - target)
		y = ring
	case left:
		x = ring
		y = wtfIntAbs(sideMiddle - target)
	case bottom:
		x = wtfIntAbs(sideMiddle - target)
		y = ring
	}
	return x + y
}

func main() {
	target, _ := strconv.Atoi(os.Args[1])

	fmt.Println(solve(target))
}
