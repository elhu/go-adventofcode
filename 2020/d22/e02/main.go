package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	check(err)
	return i
}

func parseDeck(input string) []int {
	res := []int{}
	for _, c := range strings.Split(input, "\n")[1:] {
		res = append(res, atoi(c))
	}
	return res
}

func calcScore(deck []int) int {
	res := 0
	for i := 1; i <= len(deck); i++ {
		res += deck[len(deck)-i] * i
	}
	return res
}

type Set map[string]struct{}

func inSet(k string, s Set) bool {
	_, found := s[k]
	return found
}

func deckToKey(deck []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(deck), " ", ",", -1), "[]")
}

func play(player1, player2 []int, depth int) (string, []int) {
	previouslyPlayed1 := make(Set)
	previouslyPlayed2 := make(Set)
	var top1, top2 int
	round := 1
	for len(player1) > 0 && len(player2) > 0 {
		// Previously played rule
		key1 := deckToKey(player1)
		key2 := deckToKey(player2)
		if inSet(key1, previouslyPlayed1) || inSet(key2, previouslyPlayed2) {
			return "player1", player1
		}
		previouslyPlayed1[key1] = struct{}{}
		previouslyPlayed2[key2] = struct{}{}

		top1, player1 = player1[0], player1[1:]
		top2, player2 = player2[0], player2[1:]
		// Recursive rule
		if len(player1) >= top1 && len(player2) >= top2 {
			cpy1 := make([]int, top1)
			cpy2 := make([]int, top2)
			copy(cpy1, player1[:top1])
			copy(cpy2, player2[:top2])
			if winner, _ := play(cpy1, cpy2, depth+1); winner == "player1" {
				player1 = append(player1, top1, top2)
			} else {
				player2 = append(player2, top2, top1)
			}
		} else { // Default rule rule
			if top1 > top2 {
				player1 = append(player1, top1, top2)
			} else {
				player2 = append(player2, top2, top1)
			}
		}
		round++
	}
	if len(player1) > 0 {
		return "player1", player1
	}
	// fmt.Printf("The winner of game %d is player 2!\n", depth)
	return "player2", player2
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n\n")
	player1 := parseDeck(input[0])
	player2 := parseDeck(input[1])
	winner, deck := play(player1, player2, 1)
	fmt.Printf("%s won with %d\n", winner, calcScore(deck))
}
