package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	check(err)
	return i
}

func parseTicket(str string) []int {
	var ticket []int
	for _, n := range strings.Split(str, ",") {
		ticket = append(ticket, atoi(n))
	}
	return ticket
}

var ruleExp = regexp.MustCompile(`([\w\s]+): (\d+)-(\d+) or (\d+)-(\d+)`)

func parseRules(input []string) map[string][][]int {
	rules := make(map[string][][]int)
	for _, line := range input {
		matches := ruleExp.FindStringSubmatch(line)
		rules[matches[1]] = [][]int{
			[]int{atoi(matches[2]), atoi(matches[3])},
			[]int{atoi(matches[4]), atoi(matches[5])},
		}
	}
	return rules
}

func isValidField(field int, rules map[string][][]int) bool {
	for _, rule := range rules {
		if (field >= rule[0][0] && field <= rule[0][1]) || (field >= rule[1][0] && field <= rule[1][1]) {
			return true
		}
	}
	return false
}

func filterValidTickets(rules map[string][][]int, tickets [][]int) [][]int {
	var validTickets [][]int
	for _, t := range tickets {
		valid := true
		for _, f := range t {
			valid = valid && isValidField(f, rules)
		}
		if valid {
			validTickets = append(validTickets, t)
		}
	}
	return validTickets
}

func ruleMatchesPos(rule [][]int, pos int, tickets [][]int) bool {
	for _, t := range tickets {
		if !((t[pos] >= rule[0][0] && t[pos] <= rule[0][1]) || (t[pos] >= rule[1][0] && t[pos] <= rule[1][1])) {
			return false
		}
	}
	return true
}

func index(list []int, n int) int {
	for i, v := range list {
		if v == n {
			return i
		}
	}
	return -1
}

func solve(rules map[string][][]int, tickets [][]int) int {
	// Generate list of potential positions for each rule
	rulePos := make(map[string][]int)
	for name, rule := range rules {
		for pos := 0; pos < len(tickets[0]); pos++ {
			if ruleMatchesPos(rule, pos, tickets) {
				rulePos[name] = append(rulePos[name], pos)
			}
		}
	}
	// Find rules with only one potential position, remove that position from every other rules
	// Repeat until every rule has only one potential position
	for reduced := 1; reduced > 0; {
		reduced = 0
		for name, potentialPos := range rulePos {
			if len(potentialPos) == 1 {
				for k, v := range rulePos {
					if idx := index(v, potentialPos[0]); k != name && idx != -1 {
						reduced++
						rulePos[k][idx] = rulePos[k][len(rulePos[k])-1]
						rulePos[k] = rulePos[k][:len(rulePos[k])-1]
					}
				}
			}
		}
	}
	res := 1
	for k, v := range rulePos {
		if strings.HasPrefix(k, "departure") {
			res *= tickets[0][v[0]]
		}
	}
	return res
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	var myTicketIndex int
	for i, line := range input {
		if line == "your ticket:" {
			myTicketIndex = i
			break
		}
	}
	rules := parseRules(input[:myTicketIndex-1])
	myTicket := parseTicket(input[myTicketIndex+1])
	var otherTickets [][]int
	for _, line := range input[myTicketIndex+4:] {
		otherTickets = append(otherTickets, parseTicket(line))
	}
	fmt.Println(solve(rules, append([][]int{myTicket}, filterValidTickets(rules, otherTickets)...)))
}
