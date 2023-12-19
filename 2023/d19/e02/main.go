package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Workflow struct {
	name  string
	rules []Rule
}

type Range struct {
	left, right int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func intersect(a, b Range) Range {
	return Range{max(a.left, b.left), min(a.right, b.right)}
}

func invert(a Range) Range {
	// [1:1500] -> [1500:4001]
	// [1500:4001] -> [1:1500]
	if a.left == minRange {
		return Range{a.right, maxRange}
	} else {
		return Range{minRange, a.left}
	}
}

type Rule struct {
	attr         string
	matchedRange Range
	destination  string
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func atoi(str string) int {
	v, err := strconv.Atoi(str)
	check(err)
	return v
}

var ruleExp = regexp.MustCompile(`([xmas])([<>])(\d+):(\w+)`)

const minRange, maxRange = 1, 4001

func parseRule(str string) Rule {
	var r Rule
	matches := ruleExp.FindStringSubmatch(str)
	val := atoi(matches[3])
	if matches[2] == "<" {
		r.matchedRange = Range{left: minRange, right: val}
	} else {
		r.matchedRange = Range{left: val + 1, right: maxRange}
	}
	r.attr = matches[1]
	r.destination = matches[4]
	return r
}

var worfklowExp = regexp.MustCompile(`(\w+){(.+),(\w+)}`)

func parseWorkflow(line string) Workflow {
	var w Workflow
	matches := worfklowExp.FindStringSubmatch(line)
	w.name = matches[1]
	for _, r := range strings.Split(matches[2], ",") {
		w.rules = append(w.rules, parseRule(r))
	}
	w.rules = append(w.rules, Rule{matchedRange: Range{left: minRange, right: maxRange}, destination: matches[3]})
	return w
}

type State struct {
	curr      Workflow
	path      []Workflow
	indexPath []int
}

func bfs(workflows map[string]Workflow) ([][]Workflow, [][]int) {
	var queue []State
	queue = append(queue, State{curr: workflows["in"]})
	var head State
	var validPaths [][]Workflow
	var validIndexPaths [][]int
	for len(queue) > 0 {
		head, queue = queue[0], queue[1:]
		if head.curr.name == "A" {
			validPaths = append(validPaths, head.path)
			validIndexPaths = append(validIndexPaths, head.indexPath)
		}
		for i, r := range head.curr.rules {
			newIndexPath := make([]int, len(head.indexPath))
			copy(newIndexPath, head.indexPath)
			newPath := make([]Workflow, len(head.path))
			copy(newPath, head.path)
			queue = append(queue, State{curr: workflows[r.destination], path: append(newPath, head.curr), indexPath: append(newIndexPath, i)})
		}
	}
	return validPaths, validIndexPaths
}

func resolvePath(path []Workflow, indexPath []int) int {
	attrRanges := map[string]Range{"x": {minRange, maxRange}, "m": {minRange, maxRange}, "a": {minRange, maxRange}, "s": {minRange, maxRange}}
	for i, w := range path {
		for r := 0; r < indexPath[i]; r++ {
			rule := w.rules[r]
			attrRanges[rule.attr] = intersect(attrRanges[rule.attr], invert(rule.matchedRange))
		}
		rule := w.rules[indexPath[i]]
		attrRanges[rule.attr] = intersect(attrRanges[rule.attr], rule.matchedRange)
	}
	res := 1
	for _, k := range []string{"x", "m", "a", "s"} {
		v := attrRanges[k]
		res *= v.right - v.left
	}
	return res
}

func solve(workflows map[string]Workflow) int {
	workflows["A"] = Workflow{name: "A"}
	paths, indexPaths := bfs(workflows)
	res := 0
	for i, p := range paths {
		res += resolvePath(p, indexPaths[i])
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	sections := strings.Split(data, "\n\n")
	workflows := make(map[string]Workflow)
	for _, line := range strings.Split(sections[0], "\n") {
		w := parseWorkflow(line)
		workflows[w.name] = w
	}
	fmt.Println(solve(workflows))
}
