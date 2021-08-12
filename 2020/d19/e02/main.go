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
	return rulesRef
}

func buildExpressions(rulesRef map[string]*Rule) *regexp.Regexp {
	// 0: 8 11
	// 8: 42 | 42 8
	// 11: 42 31 | 42 11 31
	// compute subgraphs for 31 and 42 (only deps of 8 and 11, which are the only deps of 0)
	buildRuleGraph(rulesRef, "31")
	buildRuleGraph(rulesRef, "42")
	exp31 := fmt.Sprintf("(%s)", strings.Join(rulesRef["31"].cache, "|"))
	exp42 := fmt.Sprintf("(%s)", strings.Join(rulesRef["42"].cache, "|"))
	// 8 is just 42 repeated 1-n times
	exp8 := fmt.Sprintf("(%s)+", exp42)
	// 11 is 42 repeated n times followed by 31 repeated the same number of time
	// build a list for each n from 0 to 10 (the messages aren't that long)
	var expandedExp11 []string
	for i := 1; i < 10; i++ {
		expandedExp11 = append(expandedExp11, fmt.Sprintf("%s{%d}%s{%d}", exp42, i, exp31, i))
	}
	// 0 is 8 followed by 11
	exp0 := fmt.Sprintf("^%s(%s)$", exp8, strings.Join(expandedExp11, "|"))
	return regexp.MustCompile(exp0)
}

func solve(exp *regexp.Regexp, messages []string) int {
	res := 0
	for _, m := range messages {
		if exp.MatchString(m) {
			res++
		}
	}
	return res
}

func main() {
	bData, err := ioutil.ReadFile(os.Args[1])
	check(err)
	data := string(bData)
	data = strings.Replace(data, "8: 42", "8: 42 | 42 8", 1)
	data = strings.Replace(data, "11: 42 31", "11: 42 31 | 42 11 31", 1)
	input := strings.Split(strings.TrimRight(data, "\n"), "\n\n")
	rules := parseRules(strings.Split(input[0], "\n"))
	exp := buildExpressions(rules)
	fmt.Println(solve(exp, strings.Split(input[1], "\n")))
}
