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

func kindOf(counts []int) int {
	if counts[0] == 5 {
		return fiveOfKind
	} else if counts[0] == 4 {
		return fourOfKind
	} else if counts[0] == 3 && counts[1] == 2 {
		return fullHouse
	} else if counts[0] == 3 {
		return threeOfKind
	} else if counts[0] == 2 && counts[1] == 2 {
		return twoPair
	} else if counts[0] == 2 {
		return onePair
	}
	return highCard
}

func parseHand(value string) Hand {
	counters := make(map[byte]int)
	for _, b := range value {
		counters[byte(b)]++
	}
	counts := make([]int, 0)
	for _, v := range counters {
		counts = append(counts, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	return Hand{value: value, kind: kindOf(counts)}
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

const order = "23456789TJQKA"
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
