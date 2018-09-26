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

func intersects(a, b map[string]struct{}) bool {
	for k := range a {
		_, seen := b[k]
		if seen {
			return true
		}
	}
	return false
}

func supportsSSL(line string) bool {
	aba := make(map[string]struct{})
	bab := make(map[string]struct{})
	insideBrackets := false
	for i := 0; i <= len(line)-3; i++ {
		if line[i] == '[' {
			insideBrackets = true
		} else if line[i] == ']' {
			insideBrackets = false
		} else {
			if line[i] == line[i+2] && line[i] != line[i+1] {
				if insideBrackets {
					aba[line[i:i+3]] = struct{}{}
				} else {
					bab[string([]byte{line[i+1], line[i], line[i+1]})] = struct{}{}
				}
			}
		}
	}
	return intersects(aba, bab)
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
		if supportsSSL(line) {
			res++
		}
	}
	fmt.Println(res)
}
