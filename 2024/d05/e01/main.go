package main

import (
	"adventofcode/utils/files"
	"adventofcode/utils/sets"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func parsePageLists(data []string) [][]int {
	var res [][]int
	for _, line := range data {
		pages := strings.Split(line, ",")
		var pagesInt []int
		for _, page := range pages {
			pagesInt = append(pagesInt, atoi(page))
		}
		res = append(res, pagesInt)
	}
	return res
}

func parseOrders(data []string) map[int]*sets.Set[int] {
	ltgt := make(map[int]*sets.Set[int])

	for _, line := range data {
		var left, right int
		fmt.Sscanf(line, "%d|%d", &left, &right)
		if _, found := ltgt[left]; !found {
			ltgt[left] = sets.New[int]()
		}
		ltgt[left].Add(right)
	}
	return ltgt
}

func isOrdered(ltgt map[int]*sets.Set[int], list []int) bool {
	for i := 1; i < len(list); i++ {
		if ltgt[list[i-1]] == nil {
			return false
		}
		if !ltgt[list[i-1]].HasMember(list[i]) {
			return false
		}
	}
	return true
}

func solve(ltgt map[int]*sets.Set[int], lists [][]int) int {
	res := 0
	for _, list := range lists {
		if isOrdered(ltgt, list) {
			res += list[len(list)/2]
		}
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	parts := strings.Split(data, "\n\n")
	lists := parsePageLists(strings.Split(parts[1], "\n"))
	ltgt := parseOrders(strings.Split(parts[0], "\n"))
	fmt.Println(solve(ltgt, lists))
}
