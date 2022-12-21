package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

const (
	NUMBER   = iota
	OPERATOR = iota
	HUMAN    = iota
)

var operations = map[string]func(left, right *Monkey) int{
	"+": func(left, right *Monkey) int {
		return left.number + right.number
	},
	"-": func(left, right *Monkey) int {
		return left.number - right.number
	},
	"*": func(left, right *Monkey) int {
		return left.number * right.number
	},
	"/": func(left, right *Monkey) int {
		return left.number / right.number
	},
}

type Operation struct {
	left, right *Monkey
	operation   func(*Monkey, *Monkey) int
}

type Monkey struct {
	name      string
	kind      int
	number    int
	operation *Operation
}

func parseMonkey(line string) *Monkey {
	if strings.Contains(line, "humn:") {
		return &Monkey{name: "humn", kind: HUMAN}
	} else if strings.ContainsAny(line, "*+-/") {
		var name, left, operator, right string
		fmt.Sscanf(line, "%s %s %s %s", &name, &left, &operator, &right)
		return &Monkey{name: strings.TrimRight(name, ":"), kind: OPERATOR}
	} else {
		var name string
		var number int
		fmt.Sscanf(line, "%s %d", &name, &number)
		return &Monkey{name: strings.TrimRight(name, ":"), kind: NUMBER, number: number}
	}
}

func solve(monkeys map[string]*Monkey) int {
	left := monkeys["root"].operation.left
	right := monkeys["root"].operation.right
	fmt.Println(left, right)
	rl := resolve(left, true)
	rr := resolve(right, true)
	fmt.Println(rl, left.number)
	fmt.Println(rr, right.number)
	var target, variable *Monkey
	if rl {
		target = left
		variable = right
	} else {
		target = right
		variable = left
	}
	humn := monkeys["humn"]
	humn.kind = NUMBER
	// Hacked it by printing the value for variable at 1000000 intervals
	// Pick i that has a value just below target, iterate from there
	for i := 3059361000000; ; i++ {
		humn.number = i
		resolve(variable, false)
		if variable.number == target.number {
			return i
		}
	}
}

func resolve(monkey *Monkey, set bool) bool {
	if monkey.kind == NUMBER {
		return true
	} else if monkey.kind == OPERATOR {
		successLeft := resolve(monkey.operation.left, set)
		successRight := resolve(monkey.operation.right, set)
		if successLeft && successRight {
			if set {
				monkey.kind = NUMBER
			}
			monkey.number = monkey.operation.operation(monkey.operation.left, monkey.operation.right)
			return true
		}
		return false
	} else {
		return false
	}
}

func main() {
	data := files.ReadLines(os.Args[1])
	monkeys := make(map[string]*Monkey)
	// build all monkeys
	for _, line := range data {
		m := parseMonkey(line)
		monkeys[m.name] = m
	}
	// build all operations
	for _, line := range data {
		if strings.ContainsAny(line, "*+-/") {
			var name, left, operator, right string
			fmt.Sscanf(line, "%s %s %s %s", &name, &left, &operator, &right)
			operation := &Operation{left: monkeys[left], right: monkeys[right], operation: operations[operator]}
			monkeys[strings.TrimRight(name, ":")].operation = operation
		}
	}
	fmt.Println(solve(monkeys))
}
