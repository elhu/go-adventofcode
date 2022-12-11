package main

import (
	"adventofcode/utils/files"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items     []int
	operation func(int) int
	testDivBy int
	testTrue  int
	testFalse int
}

func atoi(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

func parseMonkeys(monkeyData [][]byte) []*Monkey {
	monkeys := make([]*Monkey, 0, len(monkeyData))
	for _, md := range monkeyData {
		parts := strings.Split(string(md), "\n")
		startItems := make([]int, 0)
		for _, val := range strings.Split(parts[1][len("  Starting items: "):], ", ") {
			startItems = append(startItems, atoi(val))
		}
		var operator, operand string
		fmt.Sscanf(parts[2], "  Operation: new = old %s %s", &operator, &operand)
		operation := func(old int) int {
			if operator == "+" {
				if operand == "old" {
					return old + old
				} else {
					return old + atoi(operand)
				}
			} else {
				if operand == "old" {
					return old * old
				} else {
					return old * atoi(operand)
				}
			}
		}
		var divTest, testTrue, testFalse int
		fmt.Sscanf(parts[3], "  Test: divisible by %d", &divTest)
		fmt.Sscanf(parts[4], "    If true: throw to monkey %d", &testTrue)
		fmt.Sscanf(parts[5], "    If false: throw to monkey %d", &testFalse)
		monkeys = append(monkeys, &Monkey{
			items:     startItems,
			operation: operation,
			testDivBy: divTest,
			testTrue:  testTrue,
			testFalse: testFalse,
		})
	}
	return monkeys
}

func solve(monkeys []*Monkey, rounds int) int {
	processed := make([]int, len(monkeys))
	for r := 0; r < rounds; r++ {
		for mid, m := range monkeys {
			processed[mid] += len(m.items)
			for _, item := range m.items {
				worry := m.operation(item)
				worry /= 3
				if worry%m.testDivBy == 0 {
					monkeys[m.testTrue].items = append(monkeys[m.testTrue].items, worry)
				} else {
					monkeys[m.testFalse].items = append(monkeys[m.testFalse].items, worry)
				}
			}
			m.items = m.items[:0]
		}
	}
	sort.Ints(processed)
	return processed[len(processed)-1] * processed[len(processed)-2]
}

func main() {
	data := bytes.Split(files.ReadFile(os.Args[1]), []byte("\n\n"))
	monkeys := parseMonkeys(data)
	fmt.Println(solve(monkeys, 20))
}
