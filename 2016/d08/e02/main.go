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
		c <- strings.Trim(line, " \n")
		if err == io.EOF {
			break
		}
	}
	close(c)
}

var rectExp = regexp.MustCompile(`rect (\d+)x(\d+)`)
var rotateRowExp = regexp.MustCompile(`rotate row y=(\d+) by (\d+)`)
var rotateColumnExp = regexp.MustCompile(`rotate column x=(\d+) by (\d+)`)

const (
	screenX      = 50
	screenY      = 6
	rect         = iota
	rotateRow    = iota
	rotateColumn = iota
)

type screen struct {
	pixels [][]bool
}

type instruction struct {
	label, x, y int
}

func initScreen() *screen {
	s := &screen{}
	s.pixels = make([][]bool, screenY)
	for i := 0; i < screenY; i++ {
		s.pixels[i] = make([]bool, screenX)
	}
	return s
}

func (s *screen) rect(x, y int) {
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			s.pixels[i][j] = true
		}
	}
}

func (s *screen) rotateRow(y, offset int) {
	offset %= screenX
	for z := 0; z < offset; z++ {
		tmp := s.pixels[y][screenX-1]
		for i := screenX - 2; i >= 0; i-- {
			s.pixels[y][i+1] = s.pixels[y][i]
		}
		s.pixels[y][0] = tmp
	}
}

func (s *screen) rotateColumn(x, offset int) {
	offset %= screenY
	for z := 0; z < offset; z++ {
		tmp := s.pixels[screenY-1][x]
		for i := screenY - 2; i >= 0; i-- {
			s.pixels[i+1][x] = s.pixels[i][x]
		}
		s.pixels[0][x] = tmp
	}
}

func (s *screen) countPixels() int {
	res := 0
	for i := 0; i < screenY; i++ {
		for j := 0; j < screenX; j++ {
			if s.pixels[i][j] {
				res++
			}
		}
	}
	return res
}

func (s *screen) print() {
	for i := 0; i < screenY; i++ {
		for j := 0; j < screenX; j++ {
			if s.pixels[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Println("--")
}

func parseLine(line string) instruction {
	res := instruction{}
	if match := rectExp.FindStringSubmatch(line); match != nil {
		res.label = rect
		res.x, _ = strconv.Atoi(match[1])
		res.y, _ = strconv.Atoi(match[2])
	} else if match := rotateColumnExp.FindStringSubmatch(line); match != nil {
		res.label = rotateColumn
		res.x, _ = strconv.Atoi(match[1])
		res.y, _ = strconv.Atoi(match[2])
	} else if match := rotateRowExp.FindStringSubmatch(line); match != nil {
		res.label = rotateRow
		res.x, _ = strconv.Atoi(match[1])
		res.y, _ = strconv.Atoi(match[2])
	} else {
		panic("Unable to parse line " + line)
	}
	return res
}

func main() {
	s := initScreen()

	fh, err := os.Open(os.Args[1])
	check(err)
	defer fh.Close()

	reader := bufio.NewReader(fh)
	c := make(chan string, 100)

	go readLines(reader, c)
	for line := range c {
		if line == "" {
			break
		}
		inst := parseLine(line)
		switch inst.label {
		case rect:
			s.rect(inst.x, inst.y)
		case rotateRow:
			s.rotateRow(inst.x, inst.y)
		case rotateColumn:
			s.rotateColumn(inst.x, inst.y)
		}
	}
	s.print()
	fmt.Println(s.countPixels())
}
