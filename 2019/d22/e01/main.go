package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func cut(deck []int, offset int) []int {
	if offset < 0 {
		return append(deck[len(deck)+offset:], deck[:len(deck)+offset]...)
	} else {
		return append(deck[offset:], deck[:offset]...)
	}
}

func deal(deck []int, increment int) []int {
	res := make([]int, len(deck))
	for i := range deck {
		res[(i*increment)%len(deck)] = deck[i]
	}
	return res
}

func dealNewStack(deck []int) []int {
	for i := 0; i < len(deck)/2; i++ {
		deck[i], deck[len(deck)-1-i] = deck[len(deck)-1-i], deck[i]
	}
	return deck
}

func solve(deck []int, lines [][]byte) int {
	for _, line := range lines {
		switch {
		case bytes.HasPrefix(line, []byte("cut")):
			offset := bytes.TrimSpace(line[3:])
			off, err := strconv.Atoi(string(offset))
			check(err)
			deck = cut(deck, off)
		case bytes.HasPrefix(line, []byte("deal with increment")):
			increment := bytes.TrimSpace(line[20:])
			inc, err := strconv.Atoi(string(increment))
			check(err)
			deck = deal(deck, inc)
		case bytes.HasPrefix(line, []byte("deal into new stack")):
			deck = dealNewStack(deck)
		}
		// for _, card := range deck {
		// 	fmt.Printf("%d ", card)
		// }
		// fmt.Println("")
	}
	for i, card := range deck {
		if card == 2019 {
			return i
		}
	}
	panic("WTF")
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	lines := bytes.Split(bytes.Trim(data, "\n"), []byte("\n"))
	deck := make([]int, 10007)
	// deck := make([]int, 10)
	for i := range deck {
		deck[i] = i
	}
	check(err)
	fmt.Println(solve(deck, lines))
}
