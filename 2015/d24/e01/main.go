package main

import (
	"fmt"
	"io/ioutil"
	"math/bits"
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
	i, e := strconv.Atoi(str)
	check(e)
	return i
}

func Combinations(set []int, n int) (subsets [][]int) {
	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		var subset []int

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// add subset to subsets
		subsets = append(subsets, subset)
	}
	return subsets
}

func wqe(packages []int) (int, int) {
	qe := 1
	weight := 0
	for _, p := range packages {
		qe *= p
		weight += p
	}
	return weight, qe
}

const MAX = 9999999999999999

func solve(targetWeight int, packages []int) int {
	bestQe := MAX
	for packSize := 5; packSize < len(packages); packSize++ {
		for _, c := range Combinations(packages, packSize) {
			w, qe := wqe(c)
			if w == targetWeight && qe < bestQe {
				bestQe = qe
			}
		}
		if bestQe != MAX {
			return bestQe
		}
	}
	return 0
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	var packages []int
	var totalWeight int
	for _, l := range input {
		weight := atoi(l)
		packages = append(packages, weight)
		totalWeight += weight
	}
	fmt.Println(solve(totalWeight/3, packages))
}
