package main

import (
	"fmt"
	"sort"
	"strings"
)

type Floor struct {
	generators map[string]struct{}
	chips      map[string]struct{}
}

func getKeys(set map[string]struct{}) []string {
	var keys []string
	for k := range set {
		keys = append(keys, k)
	}
	return keys
}

func computeHash(floors [4]Floor) string {
	var hashes []string
	for _, f := range floors {
		chips := getKeys(f.chips)
		sort.Strings(chips)
		hashes = append(hashes, fmt.Sprintf("chips:%s", strings.Join(chips, "-")))

		generators := getKeys(f.generators)
		sort.Strings(generators)
		hashes = append(hashes, fmt.Sprintf("generators:%s", strings.Join(generators, "-")))
	}
	return strings.Join(hashes, "|")
}

func setupFloors() [4]Floor {
	var floors [4]Floor
	floors[0].generators = map[string]struct{}{
		"promethium": {},
	}
	floors[0].chips = map[string]struct{}{
		"promethium": {},
	}
	floors[1].generators = map[string]struct{}{
		"cobalt":    {},
		"curium":    {},
		"ruthenium": {},
		"plutonium": {},
	}
	floors[2].chips = map[string]struct{}{
		"cobalt":    {},
		"curium":    {},
		"ruthenium": {},
		"plutonium": {},
	}
	return floors
}

func isFinished(floors [4]Floor) bool {
	for _, chip := range []string{"promethium", "cobalt", "curium", "ruthenium", "plutonium"} {
		if _, found := floors[3].chips[chip]; !found {
			return false
		}
	}
	return true
}

func bfs(floors [4]Floor) int {
	var stepCount = 0
	seen := map[string]struct{}{
		computeHash(floors): {},
	}
	queue := make([][4]Floor, 0)
	newQueue := make([][4]Floor, 0)

	for {
		for _, q := range queue {
			seen[computeHash(q)] = struct{}{}
			if isFinished(q) {
				return stepCount
			}

		}
		queue = newQueue
		newQueue = nil
		stepCount++
	}

	if isFinished(floors) {
		return stepCount
	}
}

func main() {
	floors := setupFloors()
	bfs(floors)
}
