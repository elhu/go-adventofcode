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

func supportsTLS(line string) bool {
	abba := false
	insideBrackets := false
	for i := 0; i <= len(line)-4; i++ {
		if line[i] == '[' {
			insideBrackets = true
		} else if line[i] == ']' {
			insideBrackets = false
		} else {
			if line[i] == line[i+3] && line[i+1] == line[i+2] && line[i] != line[i+1] {
				if insideBrackets {
					return false
				}
				abba = true
			}
		}
	}
	return abba
}

func main() {
	fh, err := os.Open(os.Args[1])
	check(err)
	defer fh.Close()

	reader := bufio.NewReader(fh)
	c := make(chan string, 100)

	go readLines(reader, c)

	res := 0
	for line := range c {
		if line == "" {
			break
		}
		if supportsTLS(line) {
			res++
		}
	}
	fmt.Println(res)
}
