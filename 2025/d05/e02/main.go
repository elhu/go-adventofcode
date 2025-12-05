package main

import (
	"adventofcode/utils/files"
	"adventofcode/utils/sets"
	"fmt"
	"os"
	"strings"
)

func intersect(a, b [2]int) bool {
	return a[0] <= b[1] && b[0] <= a[1]
}

func union(a, b [2]int) (bool, [2]int) {
	if intersect(a, b) {
		return true, [2]int{min(a[0], b[0]), max(a[1], b[1])}
	}
	return false, [2]int{}
}

func solve(ranges [][2]int) int {
	rangeSet := sets.New[[2]int]()
	for _, r := range ranges {
		rangeSet.Add(r)
	}
	merged := true
	for merged {
		merged = false
		rangeMembers := rangeSet.Members()
		for i := 0; i < rangeSet.Len()-1; i++ {
			for j := i + 1; j < rangeSet.Len(); j++ {
				if merged, r := union(rangeMembers[i], rangeMembers[j]); merged {
					rangeSet.Remove(rangeMembers[i])
					rangeSet.Remove(rangeMembers[j])
					rangeSet.Add(r)
					merged = true
					break
				}
			}
		}
	}
	res := 0
	for _, r := range rangeSet.Members() {
		res += r[1] - r[0] + 1
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	parts := strings.Split(data, "\n\n")
	var ranges [][2]int
	for _, p := range strings.Split(parts[0], "\n") {
		bounds := strings.Split(p, "-")
		var r [2]int
		fmt.Sscanf(bounds[0], "%d", &r[0])
		fmt.Sscanf(bounds[1], "%d", &r[1])
		ranges = append(ranges, r)
	}
	fmt.Println(solve(ranges))
}
