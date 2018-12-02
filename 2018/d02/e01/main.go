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

func countFreqs(line []byte) map[byte]int {
	res := make(map[byte]int)
	for _, c := range line {
		res[c]++
	}
	return res
}

func hasFreq(freqs map[byte]int, freq int) bool {
	for _, v := range freqs {
		if v == freq {
			return true
		}
	}
	return false
}

func solve(lines [][]byte) int {
	twos, threes := 0, 0
	for _, line := range lines {
		freqs := countFreqs(line)
		if hasFreq(freqs, 2) {
			twos++
		}
		if hasFreq(freqs, 3) {
			threes++
		}
	}
	return twos * threes
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	lines := bytes.Split(data, []byte{'\n'})
	fmt.Println(solve(lines))
}
