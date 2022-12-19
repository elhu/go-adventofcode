package main

import (
	"adventofcode/utils/files"
	"adventofcode/utils/sets/stringset"
	"fmt"
	"os"
)

type Robot struct {
	costs    map[string]int
	produces map[string]int
}

func parseBlueprint(line string) (int, map[string]Robot) {
	robots := make(map[string]Robot)
	robots["ore"] = Robot{produces: map[string]int{"ore": 1}, costs: map[string]int{}}
	robots["clay"] = Robot{produces: map[string]int{"clay": 1}, costs: map[string]int{}}
	robots["obsidian"] = Robot{produces: map[string]int{"obsidian": 1}, costs: map[string]int{}}
	robots["geode"] = Robot{produces: map[string]int{"geode": 1}, costs: map[string]int{}}

	var bpID, oreOreCost, clayOreCost, obsidianOreCost, obsidianClayCost, geodeOreCost, geodeObsidianCost int
	fmt.Sscanf(
		line,
		"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
		&bpID, &oreOreCost, &clayOreCost, &obsidianOreCost, &obsidianClayCost, &geodeOreCost, &geodeObsidianCost,
	)

	robots["ore"].costs["ore"] = oreOreCost
	robots["clay"].costs["ore"] = clayOreCost
	robots["obsidian"].costs["ore"] = obsidianOreCost
	robots["obsidian"].costs["clay"] = obsidianClayCost
	robots["geode"].costs["ore"] = geodeOreCost
	robots["geode"].costs["obsidian"] = geodeObsidianCost
	return bpID, robots
}

type State struct {
	robots      [4]int
	resources   [4]int
	timeLeft    int
	lastSkipped int
}

func toKey(state State) string {
	return fmt.Sprintf("%v", state)
}

func copyState(state State) State {
	var resources [4]int
	var robots [4]int
	resources = state.resources
	robots = state.robots
	return State{
		resources: resources,
		robots:    robots,
		timeLeft:  state.timeLeft - 1,
	}
}

var IDX = map[string]int{
	"ore":      0,
	"clay":     1,
	"obsidian": 2,
	"geode":    3,
}

func buildLimits(blueprint map[string]Robot) [4]int {
	var limits [4]int
	for _, r := range blueprint {
		for component, amount := range r.costs {
			if amount > limits[IDX[component]] {
				limits[IDX[component]] = amount
			}
		}
	}
	limits[IDX["geode"]] = 999999999999999
	return limits
}

func enoughResources(robot Robot, resources [4]int) bool {
	for k, v := range robot.costs {
		if resources[IDX[k]] < v {
			return false
		}
	}
	return true
}

func spendResources(robot Robot, resources [4]int) [4]int {
	nr := resources
	for k, v := range robot.costs {
		nr[IDX[k]] = resources[IDX[k]] - v
	}
	return nr
}

func gatherResources(state State) [4]int {
	nr := state.resources
	for i, count := range state.robots {
		nr[i] = state.resources[i] + count
	}
	return nr
}

const GEODE = "geode"

func maxPotential(state State) int {
	res := state.resources[IDX[GEODE]]
	for i := 0; i < state.timeLeft; i++ {
		res += state.robots[IDX[GEODE]] + i
	}
	return res
}

func solveBFS(bpID int, blueprint map[string]Robot) int {
	maxGeodes := 0
	mp := make(map[int]int)

	var resources [4]int
	var robots [4]int
	robots[IDX["ore"]] = 1
	limits := buildLimits(blueprint)

	queue := []State{{resources: resources, robots: robots, timeLeft: 32, lastSkipped: -1}}
	seen := stringset.New()
	var state State
	skipped := 0
	for len(queue) > 0 {
		state, queue = queue[0], queue[1:]

		if maxGeodes < state.resources[IDX[GEODE]] {
			maxGeodes = state.resources[IDX[GEODE]]
		}

		sk := toKey(state)
		if state.timeLeft == 0 || seen.HasMember(sk) {
			skipped += 1
			continue
		}
		seen.Add(sk)
		// If current branch has less potential than current branch at same state, skip it
		p := maxPotential(state)
		// Heuristic is buggy, add arbitrary tolerance to explore more branches
		if p+14 < mp[state.timeLeft]-1 {
			continue
		}
		if p > mp[state.timeLeft] {
			mp[state.timeLeft] = p
		}

		// Only push geode state if we can build a geode bot
		if enoughResources(blueprint[GEODE], state.resources) {
			newState := copyState(state)
			newState.resources = gatherResources(newState)
			newState.resources = spendResources(blueprint[GEODE], newState.resources)
			newState.robots[IDX[GEODE]] += 1
			queue = append(queue, newState)
		} else { // If not push every other possible robot build
			for rName, robot := range blueprint {
				if state.robots[IDX[rName]] < limits[IDX[rName]] && enoughResources(robot, state.resources) {
					newState := copyState(state)
					newState.resources = spendResources(blueprint[rName], newState.resources)
					newState.resources = gatherResources(newState)
					newState.robots[IDX[rName]] += 1
					queue = append(queue, newState)
				}
			}
			// Don't forget to push the case where we don't build a robot
			newState := copyState(state)
			newState.resources = gatherResources(newState)
			queue = append(queue, newState)
		}
	}
	return maxGeodes
}

func main() {
	data := files.ReadLines(os.Args[1])[0:3]
	res := 1
	for _, l := range data {
		id, blueprint := parseBlueprint(l)
		ql := solveBFS(id, blueprint)
		res *= ql
	}
	fmt.Println(res)
}
