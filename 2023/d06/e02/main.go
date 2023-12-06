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

func solve(race Race) int {
	low, high := 0, race.duration
	for ; low <= race.duration && distanceTraveled(low, race.duration) <= race.record; low++ {
	}
	for ; high >= 0 && distanceTraveled(high, race.duration) <= race.record; high-- {
	}
	return (high - low) + 1
}

func main() {
	data := strings.Split(strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n"), "\n")
	duration := strings.Join(strings.Fields(data[0])[1:], "")
	record := strings.Join(strings.Fields(data[1])[1:], "")
	fmt.Println(solve(Race{atoi(duration), atoi(record)}))
}
