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

func validTriangle(data string) bool {
	sizes := strings.Fields(data)
	var err error
	var a, b, c int
	a, err = strconv.Atoi(sizes[0])
	check(err)
	b, err = strconv.Atoi(sizes[1])
	check(err)
	c, err = strconv.Atoi(sizes[2])
	check(err)

	return a+b > c && a+c > b && b+c > a
}

func main() {
	fh, err := os.Open(os.Args[1])
	check(err)
	defer fh.Close()

	reader := bufio.NewReader(fh)
	c := make(chan string, 100)

	go readLines(reader, c)

	valid := 0
	for line := range c {
		if line == "" {
			break
		}
		if validTriangle(line) {
			valid++
		}
	}
	fmt.Println(valid)
}
