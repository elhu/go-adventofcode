package main

import (
	"fmt"
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

func computeRow(str string) int {
	low := 0.0
	high := 127.0
	for _, c := range str {
		if c == 'F' {
			high = math.Floor(high - (high-low)/2)
		} else if c == 'B' {
			low = math.Ceil(low + (high-low)/2)
		}
	}
	return int(low)
}

func computeColumn(str string) int {
	low := 0.0
	high := 7.0
	for _, c := range str {
		if c == 'L' {
			high = math.Floor(high - (high-low)/2)
		} else if c == 'R' {
			low = math.Ceil(low + (high-low)/2)
		}
	}
	return int(low)
}

func computeSeatID(ref string) int {
	r := computeRow(ref[:7])
	c := computeColumn(ref[7:])
	return r*8 + c
}

func solve(seats []int) int {
	sort.Sort(sort.IntSlice(seats))
	for i := 1; i < len(seats); i++ {
		if seats[i]-1 != seats[i-1] {
			return seats[i] - 1
		}
	}
	return -1
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	var seats []int
	for _, l := range lines {
		id := computeSeatID(l)
		seats = append(seats, id)
	}
	fmt.Println(solve(seats))
}
