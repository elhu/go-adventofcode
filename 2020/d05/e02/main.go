package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func computeSeatID(str string) int64 {
	str = strings.ReplaceAll(str, "F", "0")
	str = strings.ReplaceAll(str, "B", "1")
	str = strings.ReplaceAll(str, "L", "0")
	str = strings.ReplaceAll(str, "R", "1")
	str = fmt.Sprintf("0b%s", str)
	i, err := strconv.ParseInt(str, 0, 64)
	check(err)
	return i
}

func solve(seats []int64) int64 {
	sort.Slice(seats, func(i, j int) bool { return seats[i] < seats[j] })
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
	var seats []int64
	for _, l := range lines {
		id := computeSeatID(l)
		seats = append(seats, id)
	}
	fmt.Println(solve(seats))
}
