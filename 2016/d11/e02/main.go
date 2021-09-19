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

type State struct {
	elevator int
	floors [4]Floor
}

func copyState(s State) State {
	floors := [4]Floor{}
	for i, f := range s.floors {
		chips := make(map[string]struct{})
		for c := range f.chips {
			chips[c] = struct{}{}
		}
		floors[i].chips = chips

		generators := make(map[string]struct{})
		for g := range f.generators {
			generators[g] = struct{}{}
		}
		floors[i].generators = generators
	}
	return State{s.elevator, floors}
}

func keys(h map[string]struct{}) (keys []string) {
	for k := range h {
		keys = append(keys, k)
	}
	return keys
}

func computeHash(state State) string {
	var hashes = []string{fmt.Sprintf("elevator:%d", state.elevator)}
	for fc, f := range state.floors {
		for c := range f.chips {
			for fg, o := range state.floors {
				if _, found :=  o.generators[c]; found {
					hashes = append(hashes,fmt.Sprintf("(%d,%d)", fc, fg))
				}
			}
		}
		sort.Strings(hashes)
	}
	return strings.Join(hashes, "|")
}

func setupState() State {
	var floors [4]Floor
	floors[0].generators = map[string]struct{}{
		"promethium": {},
		"elerium": {},
		"dilithium": {},
	}
	floors[0].chips = map[string]struct{}{
		"promethium": {},
		"elerium": {},
		"dilithium": {},
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
	return State{0, floors}
}

var elements = []string{"promethium", "cobalt", "curium", "ruthenium", "plutonium", "elerium", "dilithium"}

func isFinished(state State) bool {
	for _, e := range elements {
		if _, found := state.floors[3].chips[e]; !found {
			return false
		}
		if _, found := state.floors[3].generators[e]; !found {
			return false
		}
	}
	return true
}

func validState(state State) bool {
	for _, f := range state.floors {
		for c := range f.chips {
			if _, found := f.generators[c]; !found && len(f.generators) > 0 {
				return false
			}
		}
	}
	return true
}

func generateNewStates(state State, seen map[string]struct{}) []State {
	var candidates []State

	ef := state.elevator
	floor := state.floors[ef]
	var floorCandidates []int
	if ef > 0 {
		floorCandidates = append(floorCandidates, ef - 1)
	}
	if ef < len(state.floors) - 1 {
		floorCandidates = append(floorCandidates, ef + 1)
	}

	chips := keys(floor.chips)
	generators := keys(floor.generators)
	for _, newFloor := range(floorCandidates) {
		for i, c := range chips {
			// 1 chip only
			s := copyState(state)
			s.elevator = newFloor
			delete(s.floors[ef].chips, c)
			s.floors[newFloor].chips[c] = struct{}{}
			candidates = append(candidates, s)

			// 2 chips
			for j := i+1; j < len(chips); j++ {
				s := copyState(state)
				s.elevator = newFloor
				delete(s.floors[ef].chips, c)
				delete(s.floors[ef].chips, chips[j])
				s.floors[newFloor].chips[c] = struct{}{}
				s.floors[newFloor].chips[chips[j]] = struct{}{}
				candidates = append(candidates, s)
			}
			// 1 chip 1 generator
			for _, g := range generators {
				s := copyState(state)
				s.elevator = newFloor
				delete(s.floors[ef].chips, c)
				delete(s.floors[ef].generators, g)
				s.floors[newFloor].chips[c] = struct{}{}
				s.floors[newFloor].generators[g] = struct{}{}
				candidates = append(candidates, s)
			}
		}
		for i, g := range generators {
			// 1 generator
			s := copyState(state)
			s.elevator = newFloor
			delete(s.floors[ef].generators, g)
			s.floors[newFloor].generators[g] = struct{}{}
			candidates = append(candidates, s)

			// 2 generators
			for j := i + 1; j < len(generators);j++ {
				s := copyState(state)
				s.elevator = newFloor
				delete(s.floors[ef].generators, g)
				delete(s.floors[ef].generators, generators[j])
				s.floors[newFloor].generators[g] = struct{}{}
				s.floors[newFloor].generators[generators[j]] = struct{}{}
				candidates = append(candidates, s)
			}
		}
	}
	var res []State
	for _, c := range candidates {
		h := computeHash(c)
		if _, found := seen[h]; !found && validState(c) {
			seen[h] = struct{}{}
			res = append(res, c)
		}
	}
	return res
}

func bfs(state State) int {
	var stepCount = 0
	seen := map[string]struct{}{
		computeHash(state): {},
	}
	var queue = []State{state}
	newQueue := make([]State, 0)

	for len(queue) > 0 || len(newQueue) > 0 {
		for _, q := range queue {
			if isFinished(q) {
				return stepCount
			}
			newQueue = append(newQueue, generateNewStates(q, seen)...)
		}
		queue = newQueue
		newQueue = nil
		stepCount++
	}

	if isFinished(state) {
		return stepCount
	}
	panic("WTF")
}

func main() {
	state := setupState()
	fmt.Println(bfs(state))
}
