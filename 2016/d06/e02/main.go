package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(fh *bufio.Reader, c chan string) {
	for {
		line, err := fh.ReadString('\n')
		c <- strings.Trim(line, " \n")
		if err == io.EOF {
			break
		}
	}
	close(c)
}

const numChars = 8

func recordFrequencies(frequencies []map[byte]int, line string) {
	for i, char := range line {
		frequencies[i][byte(char)]++
	}
}

func solve(frequencies []map[byte]int) []byte {
	res := make([]byte, len(frequencies))
	for i, freqs := range frequencies {
		min := int(^uint(0) >> 1)
		var candidate byte
		for k, v := range freqs {
			if v < min {
				min = v
				candidate = k
			}
		}
		res[i] = candidate
	}
	return res
}

func main() {
	fh, err := os.Open(os.Args[1])
	check(err)
	defer fh.Close()

	reader := bufio.NewReader(fh)
	c := make(chan string, 100)

	go readLines(reader, c)

	frequencies := make([]map[byte]int, numChars)
	for i := 0; i < numChars; i++ {
		frequencies[i] = make(map[byte]int)
	}
	for line := range c {
		if line == "" {
			break
		}
		recordFrequencies(frequencies, line)
	}
	fmt.Println(string(solve(frequencies)))
}
