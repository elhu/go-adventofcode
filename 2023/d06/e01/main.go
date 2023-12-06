package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	check(err)
	return n
}

type Race struct {
	duration int
	record   int
}

func distanceTraveled(pressed, duration int) int {
	return (duration - pressed) * pressed
}

func solve(races []Race) int {
	res := 1
	for _, r := range races {
		low, high := 0, r.duration
		for ; low <= r.duration && distanceTraveled(low, r.duration) <= r.record; low++ {
		}
		for ; high >= 0 && distanceTraveled(high, r.duration) <= r.record; high-- {
		}
		res *= (high - low) + 1
	}
	return res
}

func main() {
	data := strings.Split(strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n"), "\n")
	durations := strings.Fields(data[0])[1:]
	records := strings.Fields(data[1])[1:]
	races := make([]Race, len(durations))
	for i := range durations {
		races[i] = Race{duration: atoi(durations[i]), record: atoi(records[i])}
	}
	fmt.Println(solve(races))
}
