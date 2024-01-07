package main

import (
	"adventofcode/utils/files"
	set "adventofcode/utils/sets"
	"fmt"
	"os"
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

type card struct {
	id   int
	wins int
	memo int
}

func parseLine(line string) *card {
	parts := strings.Split(line, ": ")
	id := atoi(strings.Fields(parts[0])[1])
	numbers := parts[1]
	parts = strings.Split(numbers, " | ")
	left, right := set.New[int](), set.New[int]()
	for _, num := range strings.Fields(parts[0]) {
		left.Add(atoi(num))
	}
	for _, num := range strings.Fields(parts[1]) {
		right.Add(atoi(num))
	}
	return &card{id: id, wins: left.Intersection(right).Len(), memo: -1}
}

func computeWins(cards map[int]*card, card *card) int {
	if card.memo != -1 {
		return card.memo
	}
	res := 1
	for i := card.id + 1; i <= card.id+card.wins; i++ {
		res += computeWins(cards, cards[i])
	}
	card.memo = res
	return res
}

func solve(lines []string) int {
	cards := make(map[int]*card)
	for _, line := range lines {
		card := parseLine(line)
		cards[card.id] = card
	}
	var res int
	for _, card := range cards {
		res += computeWins(cards, card)
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	fmt.Println(solve(lines))
}
