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

var pad = [][]string{
	// Use 0s as boundary values to detect out of bounds
	{"0", "0", "0", "0", "0", "0", "0"},
	{"0", "0", "0", "1", "0", "0", "0"},
	{"0", "0", "2", "3", "4", "0", "0"},
	{"0", "5", "6", "7", "8", "9", "0"},
	{"0", "0", "A", "B", "C", "0", "0"},
	{"0", "0", "0", "D", "0", "0", "0"},
	{"0", "0", "0", "0", "0", "0", "0"},
}

type position struct {
	posX int
	posY int
}

func (pos *position) move(dir byte) {
	newPos := *pos
	switch dir {
	case 'U':
		newPos.posY--
	case 'R':
		newPos.posX++
	case 'D':
		newPos.posY++
	case 'L':
		newPos.posX--
	}
	if pad[newPos.posY][newPos.posX] != "0" {
		pos.posX = newPos.posX
		pos.posY = newPos.posY
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

func solve(pos *position, line string) string {
	for _, dir := range line {
		pos.move(byte(dir))
	}
	return pad[pos.posY][pos.posX]
}

func main() {
	fh, err := os.Open(os.Args[1])
	check(err)
	defer fh.Close()

	reader := bufio.NewReader(fh)
	c := make(chan string, 100)

	go readLines(reader, c)

	pos := position{1, 3}
	var digits []string
	for line := range c {
		if line == "" {
			break
		}
		digits = append(digits, solve(&pos, line))
	}
	fmt.Println(strings.Join(digits, ""))
}
