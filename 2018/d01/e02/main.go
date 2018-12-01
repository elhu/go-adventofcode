package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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
		c <- strings.TrimSuffix(line, "\n")
		if err == io.EOF {
			break
		}
	}
	close(c)
}

func buildChanges(c chan string) []int {
	res := make([]int, 0)
	for line := range c {
		if line == "" {
			break
		}
		offset, err := strconv.Atoi(line)
		check(err)
		res = append(res, offset)
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

	freq := 0
	freqHist := make(map[int]struct{})
	freqHist[0] = struct{}{}
	changes := buildChanges(c)
	for {
		for _, offset := range changes {
			freq += offset
			if _, seen := freqHist[freq]; seen {
				fmt.Println(freq)
				return
			}
			freqHist[freq] = struct{}{}
		}
	}
}
