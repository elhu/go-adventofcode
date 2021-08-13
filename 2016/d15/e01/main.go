package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Disc struct {
	number    int
	positions int
	startPos  int
}

// Solution for AoC 2020 day 13 part 2
func solve(disks []Disc) int {
	// (x + 1 + 11) % 13 = 0
	// (x + 2 + 0) % 5 = 0
	// (x + 3 + 11) % 17 = 0
	// (x + 4 + 0) % 3 = 0
	// (x + 5 + 2) % 7 = 0
	// (x + 6 + 17) % 19 = 0

	remMods := make([][2]int, 0)
	for _, d := range disks {
		remMods = append(remMods, [2]int{-d.number - d.startPos, d.positions})
	}
	// For each congruence, compute coefficient
	coefficients := make([]int, len(remMods))
	for i := range remMods {
		coefficients[i] = 1
		for j, rm := range remMods {
			if i != j {
				coefficients[i] *= rm[1]
			}
		}
	}
	results := make([]int, len(coefficients))
	// For each coefficient, compute sub result
	for i, c := range coefficients {
		rm := remMods[i]
		var s, x big.Int
		// GCD here avoids having to solve the modulo the hard way, apparently
		x.GCD(nil, &s, big.NewInt(int64(rm[1])), big.NewInt(int64(c)))
		results[i] = rm[0] * int(s.Int64()) * c
	}
	res := 0
	mod := 1
	// Add up all the subresults, result is sum modulo the product of all the remainders
	for i, r := range results {
		res += r
		mod *= remMods[i][1]
	}
	var x big.Int
	// For some reason, needs euclidian modulo ¯\_(ツ)_/¯
	return int(x.Mod(big.NewInt(int64(res)), big.NewInt(int64(mod))).Int64())
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	var discs []Disc
	for _, l := range input {
		var disc Disc
		_, err := fmt.Sscanf(l, "Disc #%d has %d positions; at time=0, it is at position %d.", &disc.number, &disc.positions, &disc.startPos)
		check(err)
		discs = append(discs, disc)
	}
	fmt.Println(solve(discs))
}
