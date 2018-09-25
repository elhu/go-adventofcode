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
		c <- strings.Trim(line, " \n")
		if err == io.EOF {
			break
		}
	}
	close(c)
}

func validTriangle(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
}

func parseLine(line string) [3]int {
	parts := strings.Fields(line)
	var a, b, c int
	var err error
	a, err = strconv.Atoi(parts[0])
	check(err)
	b, err = strconv.Atoi(parts[1])
	check(err)
	c, err = strconv.Atoi(parts[2])
	check(err)
	return [3]int{a, b, c}
}

func main() {
	fh, err := os.Open(os.Args[1])
	check(err)
	defer fh.Close()

	reader := bufio.NewReader(fh)
	c := make(chan string, 100)

	go readLines(reader, c)

	valid := 0
	data := [3][3]int{}
	i := 0
	for line := range c {
		if line == "" {
			break
		}
		data[i] = parseLine(line)
		if i == 2 {
			for j := 0; j < 3; j++ {
				if validTriangle(data[0][j], data[1][j], data[2][j]) {
					valid++
				}
			}
		}
		i++
		i %= 3
	}
	fmt.Println(valid)
}
