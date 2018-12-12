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
	fmt.Println(string(state))
	for i := 0; i < len(state); i++ {
		res[center+i] = state[i]
	}
	return res
}

func parseNotes(rawNotes [][]byte) []note {
	res := make([]note, 0, len(rawNotes))
	for _, r := range rawNotes {
		n := note{pattern: r[0:5], result: r[9]}
		n.noop = n.pattern[2] == n.result
		res = append(res, n)
	}
	return res
}

func newState(sample []byte, notes []note) byte {
	for _, n := range notes {
		if bytes.Equal(sample, n.pattern) {
			return n.result
		}
	}
	panic("WTF")
}

func solve(initialState []byte, notes []note, genCount int) int {
	state := make([]byte, len(initialState))
	copy(state, initialState)
	fmt.Println(string(state))
	for i := 0; i < genCount; i++ {
		copy(state, initialState)
		for k := 2; k < potSize-2; k++ {
			state[k] = newState(initialState[k-2:k+3], notes)
		}
		state, initialState = initialState, state
	}
	res := 0
	for i, v := range initialState {
		if v == '#' {
			res += i - center
		}
	}
	return res
}

const genCount = 20

func main() {
	input, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input = bytes.TrimSuffix(input, []byte{'\n'})
	data := bytes.Split(input, []byte{'\n'})
	initialState := parseInitialState(data[0])
	notes := parseNotes(data[2:])
	fmt.Println(solve(initialState, notes, genCount))
}
