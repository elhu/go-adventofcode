package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	check(err)
	return n
}

type Hand struct {
	value string
	kind  int
}

type HandBid struct {
	hand Hand
	bid  int
}

var kindMap = map[[2]int]int{
	{5, 0}: fiveOfKind,
	{4, 1}: fourOfKind,
	{3, 2}: fullHouse,
	{3, 1}: threeOfKind,
	{2, 2}: twoPair,
	{2, 1}: onePair,
}

func parseHand(value string) Hand {
	counters := make(map[byte]int)
	for _, b := range value {
		counters[byte(b)]++
	}
	toAdd := counters['J']
	counters['J'] = 0
	counts := make([]int, 6)
	i := 0
	for _, v := range counters {
		counts[i] = v
		i++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	counts[0] += toAdd

	return Hand{value: value, kind: kindMap[[2]int{counts[0], counts[1]}]}
}

func compare(a, b Hand) int {
	if a.kind != b.kind {
		return a.kind - b.kind
	}
	for i := 0; i < 5; i++ {
		left := strings.Index(order, string(a.value[i]))
		right := strings.Index(order, string(b.value[i]))
		if left != right {
			return left - right
		}
	}
	return 0
}

const order = "J23456789TQKA"
const (
	highCard    = iota
	onePair     = iota
	twoPair     = iota
	threeOfKind = iota
	fullHouse   = iota
	fourOfKind  = iota
	fiveOfKind  = iota
)

func solve(hbs []HandBid) int {
	sort.Slice(hbs, func(i, j int) bool {
		return compare(hbs[i].hand, hbs[j].hand) < 0
	})
	res := 0
	for i, hb := range hbs {
		res += hb.bid * (i + 1)
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	hbs := make([]HandBid, 0)
	for _, line := range strings.Split(data, "\n") {
		parts := strings.Split(line, " ")
		hbs = append(hbs, HandBid{parseHand(parts[0]), atoi(parts[1])})
	}
	fmt.Println(solve(hbs))
}
