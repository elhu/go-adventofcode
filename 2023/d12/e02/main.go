package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func atoi(str string) int {
	n, err := strconv.Atoi(str)
	check(err)
	return n
}

type State struct {
	lineIndex      int
	reportIndex    int
	workingSprings int
}

func toKey(s State) string {
	return fmt.Sprintf("%d:%d:%d", s.lineIndex, s.reportIndex, s.workingSprings)
}

func dfs(cache map[string]int, line string, report []int, s State) int {
	if val, seen := cache[toKey(s)]; seen {
		return val
	}
	if s.lineIndex == len(line) {
		if (s.reportIndex == len(report) && s.workingSprings == 0) ||
			s.reportIndex == len(report)-1 && s.workingSprings == report[s.reportIndex] {
			return 1
		}
		return 0
	}
	res := 0
	if line[s.lineIndex] != '#' {
		if s.workingSprings == 0 {
			res += dfs(cache, line, report, State{s.lineIndex + 1, s.reportIndex, s.workingSprings})
		} else if s.workingSprings == report[s.reportIndex] {
			res += dfs(cache, line, report, State{s.lineIndex + 1, s.reportIndex + 1, 0})
		}
	}
	if line[s.lineIndex] != '.' && s.reportIndex < len(report) && s.workingSprings < report[s.reportIndex] {
		res += dfs(cache, line, report, State{s.lineIndex + 1, s.reportIndex, s.workingSprings + 1})
	}
	cache[toKey(s)] = res
	return res
}

func unfold(str string, sep string) string {
	var parts []string
	for i := 0; i < 5; i++ {
		parts = append(parts, str)
	}
	return strings.Join(parts, sep)
}

func solve(lines []string) int {
	res := 0
	for _, line := range lines {
		parts := strings.Fields(line)
		var report []int
		for _, r := range strings.Split(unfold(parts[1], ","), ",") {
			report = append(report, atoi(r))
		}
		res += dfs(make(map[string]int), unfold(parts[0], "?"), report, State{0, 0, 0})
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	fmt.Println(solve(lines))
}
