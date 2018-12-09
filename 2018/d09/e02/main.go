package main

import (
	"fmt"
	"os"
	"strconv"
)

type marble struct {
	prev, next *marble
	value      int
}

func printList(start, current *marble) {
	var c *marble
	for c = start; c.next != start; c = c.next {
		if c == current {
			fmt.Printf("(%2d)", c.value)
		} else {
			fmt.Printf(" %2d ", c.value)
		}
	}
	if c == current {
		fmt.Printf("(%2d)", c.value)
	} else {
		fmt.Printf(" %2d ", c.value)
	}
	fmt.Println("")
}

func solve(nbPlayer, nbMarble int, start *marble) []int {
	scores := make([]int, nbPlayer)
	curr := start
	for i := 1; i <= nbMarble; i++ {
		currentPlayer := (i - 1) % nbPlayer
		if i%23 == 0 {
			scores[currentPlayer] += i
			for j := 0; j < 7; j++ {
				curr = curr.prev
			}
			scores[currentPlayer] += curr.value
			nextCurr := curr.next
			curr.prev.next = curr.next
			curr.next.prev = curr.prev
			curr = nextCurr
		} else {
			newMarble := &marble{value: i, next: curr.next.next, prev: curr.next}
			newMarble.next.prev = newMarble
			newMarble.prev.next = newMarble
			curr = newMarble
		}
	}
	return scores
}

func main() {
	nbPlayer, _ := strconv.Atoi(os.Args[1])
	nbMarble, _ := strconv.Atoi(os.Args[2])

	start := &marble{value: 0}
	start.prev = start
	start.next = start
	scores := solve(nbPlayer, nbMarble, start)
	max := 0
	for _, score := range scores {
		if score > max {
			max = score
		}
	}
	fmt.Println(max)
}
