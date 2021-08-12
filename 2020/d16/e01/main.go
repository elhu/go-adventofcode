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

var ruleExp = regexp.MustCompile(`(\w+): (\d+)-(\d+) or (\d+)-(\d+)`)

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

func solve(rules map[string][][]int, tickets [][]int) int {
	invalidSum := 0
	for _, t := range tickets {
		for _, f := range t {
			if !isValidField(f, rules) {
				invalidSum += f
			}
		}
	}
	return invalidSum
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
	var otherTickets [][]int
	for _, line := range input[myTicketIndex+4:] {
		otherTickets = append(otherTickets, parseTicket(line))
	}
	fmt.Println(solve(rules, otherTickets))
}
