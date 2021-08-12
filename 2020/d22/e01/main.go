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

func play(player1, player2 []int) int {
	var top1, top2 int
	for len(player1) > 0 && len(player2) > 0 {
		top1, player1 = player1[0], player1[1:]
		top2, player2 = player2[0], player2[1:]
		if top1 > top2 {
			player1 = append(player1, top1, top2)
		} else {
			player2 = append(player2, top2, top1)
		}
	}
	if len(player1) > 0 {
		return calcScore(player1)
	}
	return calcScore(player2)
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n\n")
	player1 := parseDeck(input[0])
	player2 := parseDeck(input[1])
	fmt.Println(play(player1, player2))
}
