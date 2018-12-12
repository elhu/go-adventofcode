package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type note struct {
	pattern []byte
	result  byte
	noop    bool
}

const potSize = 1001
const center = potSize / 2

func parseInitialState(state []byte) []byte {
	res := make([]byte, potSize)
	for i := 0; i < potSize; i++ {
		res[i] = '.'
	}
	state = state[15:]
	for i := 0; i < len(state); i++ {
		res[center+i] = state[i]
	}
	return res
}

func parseNotes(rawNotes [][]byte) []note {
	res := make([]note, 0, len(rawNotes))
	for _, r := range rawNotes {
		if r[9] == '#' {
			n := note{pattern: r[0:5], result: r[9]}
			n.noop = n.pattern[2] == n.result
			res = append(res, n)
		}
	}
	return res
}

func newState(sample []byte, notes []note) byte {
	for _, n := range notes {
		if bytes.Equal(sample, n.pattern) {
			return '#'
		}
	}
	return '.'
}

func computeScore(state []byte) int {
	res := 0
	for i, v := range state {
		if v == '#' {
			res += i - center
		}
	}
	return res
}

func solve(initialState []byte, notes []note, genCount int) int {
	state := make([]byte, len(initialState))
	prevScore := 0
	prevPlantCount := 0
	similarityCounter := 0
	for i := 0; i < genCount; i++ {
		copy(state, initialState)
		for k := 2; k < potSize-2; k++ {
			state[k] = newState(initialState[k-2:k+3], notes)
		}
		score := computeScore(state)
		plantCount := bytes.Count(state, []byte{'#'})
		// Count how many times we have the same plant count in a row
		if plantCount == prevPlantCount {
			similarityCounter++
		} else {
			similarityCounter = 0
		}
		// If we get the same plant count more than 10 times in a row, then we can assume it won't change
		// Once that happens, from observation we know that the score increases a fixed amount for each new gen
		if similarityCounter >= 10 {
			scoreDiff := score - prevScore
			return score + (genCount-i)*scoreDiff
		}
		prevPlantCount = plantCount
		prevScore = score
		state, initialState = initialState, state
	}
	return computeScore(initialState)
}

const genCount = 50000000000

func main() {
	input, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input = bytes.TrimSuffix(input, []byte{'\n'})
	data := bytes.Split(input, []byte{'\n'})
	initialState := parseInitialState(data[0])
	notes := parseNotes(data[2:])
	fmt.Println(solve(initialState, notes, genCount))
}
