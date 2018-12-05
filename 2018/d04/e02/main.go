package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var beginExp = regexp.MustCompile(`\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}\] Guard #(\d+) begins shift`)
var sleepExp = regexp.MustCompile(`\[\d{4}-\d{2}-\d{2} \d{2}:(\d{2})\] falls asleep`)
var wakeExp = regexp.MustCompile(`\[\d{4}-\d{2}-\d{2} \d{2}:(\d{2})\] wakes up`)

func solve(schedule map[string][]int) int {
	maxMinutesAsleep := 0
	guard := ""
	sleepiestMinute := 0
	for g, s := range schedule {
		for i, sleepTimes := range s {
			if sleepTimes > maxMinutesAsleep {
				maxMinutesAsleep = sleepTimes
				sleepiestMinute = i
				guard = g
			}
		}
	}
	guardInt, _ := strconv.Atoi(guard)
	return sleepiestMinute * guardInt
}

func extractGuard(line string) string {
	if match := beginExp.FindStringSubmatch(line); match != nil {
		return match[1]
	}
	return ""
}

func processCycle(schedule []int, sleep, wake string) {
	matchSleep := sleepExp.FindStringSubmatch(sleep)
	minuteStart, _ := strconv.Atoi(matchSleep[1])
	matchWake := wakeExp.FindStringSubmatch(wake)
	minuteEnd, _ := strconv.Atoi(matchWake[1])
	for i := minuteStart; i < minuteEnd; i++ {
		schedule[i]++
	}
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	lines := strings.Split(string(data), "\n")
	sort.Strings(lines)
	schedule := make(map[string][]int)
	guard := ""
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		if newGuard := extractGuard(lines[i]); newGuard != "" {
			guard = newGuard
			if _, seen := schedule[guard]; !seen {
				schedule[guard] = make([]int, 60)
			}
		} else {
			processCycle(schedule[guard], lines[i], lines[i+1])
			i++
		}
	}
	fmt.Println(solve(schedule))
}
