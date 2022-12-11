package main

import (
	"fmt"
	"sort"
)

type Monkey struct {
	items     []int
	operation func(int) int
	testDivBy int
	testTrue  int
	testFalse int
}

var sampleMonkeys = []*Monkey{
	{
		items:     []int{79, 98},
		operation: func(old int) int { return old * 19 },
		testDivBy: 23,
		testTrue:  2,
		testFalse: 3,
	},
	{
		items:     []int{54, 65, 75, 74},
		operation: func(old int) int { return old + 6 },
		testDivBy: 19,
		testTrue:  2,
		testFalse: 0,
	},
	{
		items:     []int{79, 60, 97},
		operation: func(old int) int { return old * old },
		testDivBy: 13,
		testTrue:  1,
		testFalse: 3,
	},
	{
		items:     []int{74},
		operation: func(old int) int { return old + 3 },
		testDivBy: 17,
		testTrue:  0,
		testFalse: 1,
	},
}

var monkeys = []*Monkey{
	{
		items:     []int{93, 54, 69, 66, 71},
		operation: func(old int) int { return old * 3 },
		testDivBy: 7,
		testTrue:  7,
		testFalse: 1,
	},
	{
		items:     []int{89, 51, 80, 66},
		operation: func(old int) int { return old * 17 },
		testDivBy: 19,
		testTrue:  5,
		testFalse: 7,
	},
	{
		items:     []int{90, 92, 63, 91, 96, 63, 64},
		operation: func(old int) int { return old + 1 },
		testDivBy: 13,
		testTrue:  4,
		testFalse: 3,
	},
	{
		items:     []int{65, 77},
		operation: func(old int) int { return old + 2 },
		testDivBy: 3,
		testTrue:  4,
		testFalse: 6,
	},
	{
		items:     []int{76, 68, 94},
		operation: func(old int) int { return old * old },
		testDivBy: 2,
		testTrue:  0,
		testFalse: 6,
	},
	{
		items:     []int{86, 65, 66, 97, 73, 83},
		operation: func(old int) int { return old + 8 },
		testDivBy: 11,
		testTrue:  2,
		testFalse: 3,
	},
	{
		items:     []int{78},
		operation: func(old int) int { return old + 6 },
		testDivBy: 17,
		testTrue:  0,
		testFalse: 1,
	},
	{
		items:     []int{89, 57, 59, 61, 87, 55, 55, 88},
		operation: func(old int) int { return old + 7 },
		testDivBy: 5,
		testTrue:  2,
		testFalse: 5,
	},
}

func solve(monkeys []*Monkey, rounds int) int {
	processed := make([]int, len(monkeys))
	var prodDivBy = 1
	for _, m := range monkeys {
		prodDivBy *= m.testDivBy
	}
	for r := 0; r < rounds; r++ {
		for mid, m := range monkeys {
			processed[mid] += len(m.items)
			for _, item := range m.items {
				worry := m.operation(item)
				if worry%m.testDivBy == 0 {
					monkeys[m.testTrue].items = append(monkeys[m.testTrue].items, worry%prodDivBy)
				} else {
					monkeys[m.testFalse].items = append(monkeys[m.testFalse].items, worry%prodDivBy)
				}
			}
			m.items = m.items[:0]
		}
	}
	sort.Ints(processed)
	return processed[len(processed)-1] * processed[len(processed)-2]
}

func main() {
	fmt.Println(solve(monkeys, 10000))
}
