package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

type game struct {
	id    int
	hands []map[string]int
}

func parseGame(rawGame string) game {
	game := game{}
	parts := strings.Split(rawGame, ": ")
	fmt.Sscanf(parts[0], "Game %d", &game.id)

	hands := strings.Split(parts[1], ";")
	for _, hand := range hands {
		handMap := make(map[string]int)
		cubes := strings.Split(hand, ", ")
		for _, cube := range cubes {
			var count int
			var color string
			fmt.Sscanf(cube, "%d %s", &count, &color)
			handMap[color] = count
		}
		game.hands = append(game.hands, handMap)
	}
	return game
}

var limits map[string]int = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func solve(games []game) int {
	res := 0
	for _, game := range games {
		possible := true
		for _, hand := range game.hands {
			for k, v := range limits {
				if hand[k] > v {
					possible = false
					break
				}
			}
		}
		if possible {
			res += game.id
		}
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	rawGames := strings.Split(data, "\n")
	games := make([]game, 0)
	for _, rawGame := range rawGames {
		games = append(games, parseGame(rawGame))
	}
	fmt.Println(solve(games))
}
