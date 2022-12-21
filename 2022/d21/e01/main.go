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
)

var operations = map[string]func(left, right string, monkeys map[string]*Monkey) int{
	"+": func(left, right string, monkeys map[string]*Monkey) int {
		return monkeys[left].number + monkeys[right].number
	},
	"-": func(left, right string, monkeys map[string]*Monkey) int {
		return monkeys[left].number - monkeys[right].number
	},
	"*": func(left, right string, monkeys map[string]*Monkey) int {
		return monkeys[left].number * monkeys[right].number
	},
	"/": func(left, right string, monkeys map[string]*Monkey) int {
		return monkeys[left].number / monkeys[right].number
	},
}

type Operation struct {
	left, right string
	operation   func(string, string, map[string]*Monkey) int
}

type Monkey struct {
	name      string
	kind      int
	number    int
	operation *Operation
}

func parseMonkey(line string) *Monkey {
	if strings.ContainsAny(line, "*+-/") {
		var name, left, operator, right string
		fmt.Sscanf(line, "%s %s %s %s", &name, &left, &operator, &right)
		operation := &Operation{left: left, right: right, operation: operations[operator]}
		return &Monkey{name: strings.TrimRight(name, ":"), kind: OPERATOR, operation: operation}
	}
	var name string
	var number int
	fmt.Sscanf(line, "%s %d", &name, &number)
	return &Monkey{name: strings.TrimRight(name, ":"), kind: NUMBER, number: number}
}

func solve(monkeys map[string]*Monkey) int {
	for monkeys["root"].kind != NUMBER {
		for _, monkey := range monkeys {
			if monkey.kind == OPERATOR {
				if monkeys[monkey.operation.left].kind == NUMBER && monkeys[monkey.operation.right].kind == NUMBER {
					monkey.kind = NUMBER
					monkey.number = monkey.operation.operation(monkey.operation.left, monkey.operation.right, monkeys)
				}
			}
		}
	}
	return monkeys["root"].number
}

func main() {
	data := files.ReadLines(os.Args[1])
	monkeys := make(map[string]*Monkey)
	for _, line := range data {
		m := parseMonkey(line)
		monkeys[m.name] = m
	}
	fmt.Println(solve(monkeys))
}
