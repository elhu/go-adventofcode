package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
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

// Use Chinese Remainder Theorem, or die trying
// https://blog.demofox.org/2015/09/12/solving-simultaneous-congruences-chinese-remainder-theorem/
// https://rosettacode.org/wiki/Chinese_remainder_theorem
func solve(buses []int) int {
	// Convert input in a set of remainders and modulos
	remMods := make([][2]int, 0)
	for i, val := range buses {
		if val == -42 {
			continue
		}
		remMods = append(remMods, [2]int{val - i, val})
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
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	buses := make([]int, 0)
	for _, t := range strings.Split(lines[1], ",") {
		if t != "x" {
			buses = append(buses, atoi(t))
		} else {
			buses = append(buses, -42)
		}
	}
	fmt.Println(solve(buses))
}
