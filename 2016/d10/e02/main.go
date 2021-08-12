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

type bot struct {
	id         string
	values     []int
	lowTarget  *bot
	highTarget *bot
}

func (b *bot) assignValue(val int) {
	b.values = append(b.values, val)
	if len(b.values) == 2 {
		if b.lowTarget != nil {
			b.lowTarget.assignValue(min(b.values))
		}
		if b.highTarget != nil {
			b.highTarget.assignValue(max(b.values))
		}
		b.values = b.values[:]
	} else if len(b.values) > 2 {
		panic("WTF")
	}
}

func fetchBot(bots map[string](*bot), name string) *bot {
	if b, exists := bots[name]; exists {
		return b
	}
	b := &bot{name, make([]int, 0), nil, nil}
	bots[name] = b
	return b
}

func min(vals []int) int {
	min := vals[0]
	for _, v := range vals[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func max(vals []int) int {
	max := vals[0]
	for _, v := range vals[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

var assignValueExp = regexp.MustCompile(`value (\d+) goes to bot (\d+)`)
var giveValuesExp = regexp.MustCompile(`bot (\d+) gives low to (bot|output) (\d+) and high to (bot|output) (\d+)`)

const targetLow = 17
const targetHigh = 61

func process(line string, bots map[string](*bot)) {
	if match := assignValueExp.FindStringSubmatch(line); match != nil {
		b := fetchBot(bots, match[2])
		val, _ := strconv.Atoi(match[1])
		b.assignValue(val)
	} else if match := giveValuesExp.FindStringSubmatch(line); match != nil {
		b := fetchBot(bots, match[1])
		target := match[3]
		if match[2] == "output" {
			target = fmt.Sprintf("output %s", match[3])
		}
		r := fetchBot(bots, target)
		b.lowTarget = r

		target = match[5]
		if match[4] == "output" {
			target = fmt.Sprintf("output %s", match[5])
		}
		r = fetchBot(bots, target)
		b.highTarget = r
	} else {
		panic("Unexpected instruction: " + line)
	}
}

func main() {
	fh, err := os.Open(os.Args[1])
	check(err)
	defer fh.Close()

	reader := bufio.NewReader(fh)
	c := make(chan string, 100)

	bots := make(map[string](*bot))
	go readLines(reader, c)
	for line := range c {
		if line == "" {
			break
		}
		process(line, bots)
	}
	println(bots["output 0"].values[0] * bots["output 1"].values[0] * bots["output 2"].values[0])
}
