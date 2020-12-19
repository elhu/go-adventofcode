package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var baseRuleExp = regexp.MustCompile(`(\d+): "(\w)"`)

type Rule struct {
	options [][]string
	cache   []string
}

func combineArrays(input [][]string) []string {
	res := []string{""}
	for _, arr := range input {
		prod := []string{}
		for _, i := range arr {
			for _, r := range res {
				prod = append(prod, r+i)
			}
		}
		res = prod
	}
	return res
}

func buildRuleGraph(rulesRef map[string]*Rule, root string) {
	if len(rulesRef[root].cache) == 0 {
		for _, option := range rulesRef[root].options {
			var cache []string
			for _, name := range option {
				buildRuleGraph(rulesRef, name)
			}
			var nestedOptions [][]string
			for _, opt := range option {
				nestedOptions = append(nestedOptions, rulesRef[opt].cache)
			}
			cache = append(cache, combineArrays(nestedOptions)...)
			rulesRef[root].cache = append(rulesRef[root].cache, cache...)
		}
	}
}

func parseRules(input []string) map[string]*Rule {
	rulesRef := make(map[string]*Rule)

	for _, l := range input {
		if matches := baseRuleExp.FindStringSubmatch(l); len(matches) > 0 {
			rulesRef[matches[1]] = &Rule{cache: []string{matches[2]}}
		} else {
			parts := strings.Split(l, ": ")
			rule := Rule{cache: []string{}}
			for _, opt := range strings.Split(parts[1], " | ") {
				rule.options = append(rule.options, strings.Split(opt, " "))
			}
			rulesRef[parts[0]] = &rule
		}
	}
	buildRuleGraph(rulesRef, "0")
	return rulesRef
}

func solve(rule *Rule, messages []string) int {
	res := 0
	for _, m := range messages {
		for _, c := range rule.cache {
			if m == c {
				res++
				break
			}
		}
	}
	return res
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n\n")
	rules := parseRules(strings.Split(input[0], "\n"))
	fmt.Println(solve(rules["0"], strings.Split(input[1], "\n")))
}
