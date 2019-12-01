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

func weightToFuel(w int) int {
	return w/3 - 2
}

func main() {
	fh, err := os.Open(os.Args[1])
	check(err)
	defer fh.Close()

	reader := bufio.NewReader(fh)

	c := make(chan string, 100)

	go readLines(reader, c)

	totalFuel := 0
	for line := range c {
		weight, err := strconv.Atoi(line)
		check(err)
		totalFuel += weightToFuel(weight)
	}
	fmt.Println(totalFuel)
}
