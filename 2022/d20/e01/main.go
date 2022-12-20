package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
)

func atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return i
}

func index(n int, numbers []int) int {
	for i, v := range numbers {
		if v == n {
			return i
		}
	}
	return -1
}

type ListItem struct {
	value      int
	next, prev *ListItem
}

func buildList(data []string) []*ListItem {
	numbers := make([]*ListItem, len(data))
	for i, l := range data {
		numbers[i] = &ListItem{value: atoi(l)}
	}
	numbers[0].prev = numbers[len(numbers)-1]
	numbers[len(numbers)-1].next = numbers[0]
	for i := 1; i < len(numbers); i++ {
		numbers[i].prev = numbers[i-1]
		numbers[i-1].next = numbers[i]
	}
	return numbers
}

func findZeroItem(list []*ListItem) *ListItem {
	curr := list[0]
	for curr.value != 0 {
		curr = curr.next
	}
	return curr
}

func moveNumber(number *ListItem, listLength int) {
	curr := number.prev

	number.prev.next = number.next
	number.next.prev = number.prev

	i := number.value % (listLength - 1)
	for i < 0 {
		i++
		curr = curr.prev
	}
	for i > 0 {
		i--
		curr = curr.next
	}

	number.prev = curr
	number.next = curr.next
	number.prev.next = number
	number.next.prev = number
}

func mix(list []*ListItem) {
	listLength := len(list)
	for _, number := range list {
		moveNumber(number, listLength)
	}
}

func solve(data []string) int {
	list := buildList(data)
	mix(list)

	curr := findZeroItem(list)
	sum := 0
	for i := 0; i <= 3000; i++ {
		if i == 1000 || i == 2000 || i == 3000 {
			sum += curr.value
		}
		curr = curr.next
	}
	return sum
}

func main() {
	data := files.ReadLines(os.Args[1])
	fmt.Println(solve(data))
}

// 7083 too high
// 13343 too high
// 4183 wrong
