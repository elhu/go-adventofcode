package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
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

var roomExp = regexp.MustCompile(`(?P<id>\d+) @ (?P<x>\d+),(?P<y>\d+): (?P<w>\d+)x(?P<h>\d+)`)

func processLine(line string, fabric [][]int) {
	match := roomExp.FindStringSubmatch(line)
	x, _ := strconv.Atoi(match[2])
	y, _ := strconv.Atoi(match[3])
	w, _ := strconv.Atoi(match[4])
	h, _ := strconv.Atoi(match[5])

	for i := y; i < y+h; i++ {
		for j := x; j < x+w; j++ {
			fabric[i][j]++
		}
	}
}
func countOverlap(fabric [][]int) int {
	res := 0
	for i := 0; i < len(fabric); i++ {
		for j := 0; j < len(fabric[0]); j++ {
			if fabric[i][j] > 1 {
				res++
			}
		}
	}
	return res
}

func main() {
	fh, err := os.Open(os.Args[1])
	check(err)
	defer fh.Close()

	reader := bufio.NewReader(fh)

	fabric := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		fabric[i] = make([]int, 1000)
	}

	c := make(chan string, 100)

	go readLines(reader, c)

	for line := range c {
		processLine(line, fabric)
	}
	fmt.Println(countOverlap(fabric))
}
