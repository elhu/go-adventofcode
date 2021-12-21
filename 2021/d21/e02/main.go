package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
)

func possibleRolls() map[int]int {
	res := make(map[int]int)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				res[i+j+k]++
			}
		}
	}
	return res
}

func solve(posOne, posTwo int) int {
	// map[position][score]count
	// offset positions by 1 to simplify modulo
	currentPosScore := map[int]map[int]int{(posOne - 1): {0: 1}}
	nextPosScore := map[int]map[int]int{(posTwo - 1): {0: 1}}
	pr := possibleRolls()
	var currentWon, nextWon int
	currentUniverseCount, nextUniverseCount := 1, 1
	// Play until every universe has reached the high score
	for len(currentPosScore) > 0 && len(nextPosScore) > 0 {
		for toggle := 0; toggle <= 1; toggle++ {
			tmpScorePos := make(map[int]map[int]int)
			for r, c := range pr {
				for pos, scoresCounts := range currentPosScore {
					newPos := (pos + r) % 10
					for score, count := range scoresCounts {
						newScore := score + newPos + 1 // fix position offset
						if newScore >= 21 {
							currentWon += count * c * nextUniverseCount
						} else {
							currentUniverseCount += count * c
							if _, found := tmpScorePos[newPos]; !found {
								tmpScorePos[newPos] = make(map[int]int)
							}
							tmpScorePos[newPos][newScore] += count * c
						}
					}
				}
			}
			nextPosScore, currentPosScore = tmpScorePos, nextPosScore
			nextWon, currentWon = currentWon, nextWon
			nextUniverseCount, currentUniverseCount = currentUniverseCount, 0
		}
	}
	if nextWon > currentWon {
		return nextWon
	}
	return currentWon
}

func main() {
	data := files.ReadLines(os.Args[1])
	var player, posOne, posTwo int
	fmt.Sscanf(data[0], "Player %d starting position: %d", &player, &posOne)
	fmt.Sscanf(data[1], "Player %d starting position: %d", &player, &posTwo)
	fmt.Println(solve(posOne, posTwo))
}
