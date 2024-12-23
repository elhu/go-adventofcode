package main

import (
	"adventofcode/utils/files"
	"adventofcode/utils/sets"
	"fmt"
	"os"
	"sort"
	"strings"
)

func solve(lines []string) string {
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
	matches := make(map[string]*sets.Set[string])
	for k := range computers {
		matches[k] = sets.NewFromSlice([]string{k})
	}
	// Now we have all 3-sized groups
	for mf := true; mf; {
		mf = false
		nm := make(map[string]*sets.Set[string])
		for _, members := range matches {
			for k, v := range computers {
				if v.Intersection(members).Len() == members.Len() {
					cp := copySet(members)
					cp.Add(k)
					pwd := cp.Members()
					sort.Strings(pwd)
					nm[strings.Join(pwd, ",")] = cp
					mf = true
				}
			}
		}
		if !mf {
			for k := range matches {
				return k
			}
		}
		matches = nm
	}
	panic("WTF")
}

func copySet(s *sets.Set[string]) *sets.Set[string] {
	res := sets.New[string]()
	s.Each(func(m string) {
		res.Add(m)
	})
	return res
}

// 11011 too high
func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	fmt.Println(solve(lines))
}
