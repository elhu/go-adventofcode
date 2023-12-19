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

type Rule struct {
	condition   func(Part) bool
	destination string
}

type Part struct {
	data map[string]int
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

func parseRule(str string) Rule {
	var r Rule
	matches := ruleExp.FindStringSubmatch(str)
	val := atoi(matches[3])
	r.condition = func(p Part) bool {
		if matches[2] == "<" {
			return p.data[matches[1]] < val
		} else {
			return p.data[matches[1]] > val
		}
	}
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
	w.rules = append(w.rules, Rule{condition: func(p Part) bool { return true }, destination: matches[3]})
	return w
}

func parsePart(line string) Part {
	var p Part
	var x, m, a, s int
	fmt.Sscanf(line, "{x=%d,m=%d,a=%d,s=%d}", &x, &m, &a, &s)
	p.data = make(map[string]int)
	p.data["x"] = x
	p.data["m"] = m
	p.data["a"] = a
	p.data["s"] = s
	return p
}

func solve(workflows map[string]Workflow, parts []Part) int {
	var rejected, accepted []Part
	workflows["R"] = Workflow{name: "R", rules: []Rule{{condition: func(p Part) bool {
		rejected = append(rejected, p)
		return false
	}}}}
	workflows["A"] = Workflow{name: "A", rules: []Rule{{condition: func(p Part) bool {
		accepted = append(accepted, p)
		return false
	}}}}
	for _, p := range parts {
		currentWorkflow := workflows["in"]
		for {
			matched := false
			for _, r := range currentWorkflow.rules {
				if r.condition(p) {
					matched = true
					currentWorkflow = workflows[r.destination]
					break
				}
			}
			if !matched {
				break
			}
		}
	}
	res := 0
	for _, a := range accepted {
		res += a.data["x"] + a.data["m"] + a.data["a"] + a.data["s"]
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
	var parts []Part
	for _, line := range strings.Split(sections[1], "\n") {
		parts = append(parts, parsePart(line))
	}
	fmt.Println(solve(workflows, parts))
}
