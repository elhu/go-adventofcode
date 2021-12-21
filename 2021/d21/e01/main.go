package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
)

func solve(posOne, posTwo int) int {
	positions := []int{posOne, posTwo}
	scores := []int{0, 0}
	currentPlayer := 0
	dice := 0

	for rounds := 1; ; rounds++ {
		offset := 0
		for i := 0; i < 3; i++ {
			dice++
			if dice == 101 {
				dice = 1
			}
			offset += dice
		}
		positions[currentPlayer] += (offset % 10)
		if positions[currentPlayer] > 10 {
			positions[currentPlayer] -= 10
		}
		scores[currentPlayer] += positions[currentPlayer]
		if scores[currentPlayer] >= 1000 {
			return rounds * 3 * scores[(currentPlayer+1)%2]
		}
		currentPlayer = (currentPlayer + 1) % 2
	}
}

func main() {
	data := files.ReadLines(os.Args[1])
	var player, posOne, posTwo int
	fmt.Sscanf(data[0], "Player %d starting position: %d", &player, &posOne)
	fmt.Sscanf(data[1], "Player %d starting position: %d", &player, &posTwo)
	fmt.Println(solve(posOne, posTwo))
}
