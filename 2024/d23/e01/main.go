package main

import (
	"adventofcode/utils/files"
	"adventofcode/utils/sets"
	"fmt"
	"os"
	"sort"
	"strings"
)

func solve(lines []string) int {
	computers := make(map[string]*sets.Set[string])
	for _, line := range lines {
		parts := strings.Split(line, "-")
		if _, found := computers[parts[0]]; !found {
			computers[parts[0]] = sets.New[string]()
		}
		if _, found := computers[parts[1]]; !found {
			computers[parts[1]] = sets.New[string]()
		}
		computers[parts[0]].Add(parts[1])
		computers[parts[1]].Add(parts[0])
	}
	matches := sets.New[string]()
	for k, v := range computers {
		for _, c := range v.Members() {
			for _, z := range computers[c].Members() {
				if computers[z].HasMember(k) && v.HasMember(z) {
					if strings.HasPrefix(k, "t") || strings.HasPrefix(c, "t") || strings.HasPrefix(z, "t") {
						s := []string{k, c, z}
						sort.Strings(s)
						matches.Add(strings.Join(s, ","))
					}
				}
			}
		}
	}
	return matches.Len()
}

// 11011 too high
func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	fmt.Println(solve(lines))
}
