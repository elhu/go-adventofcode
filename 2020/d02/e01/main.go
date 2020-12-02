package main

import (
	"fmt"
	"io/ioutil"
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

func unsafeAtoi(s string) int {
	num, err := strconv.Atoi(s)
	check(err)
	return num
}

type passwordAndPolicy struct {
	min, max int
	char     byte
	password []byte
}

func countLetters(str []byte) map[byte]int {
	res := make(map[byte]int)
	for _, c := range str {
		res[c]++
	}
	return res
}

func solve(passwords []passwordAndPolicy) int {
	res := 0
	for _, p := range passwords {
		lc := countLetters(p.password)
		if lc[p.char] >= p.min && lc[p.char] <= p.max {
			res++
		}
	}
	return res
}

var passwordExp = regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	var input []passwordAndPolicy
	for _, l := range lines {
		match := passwordExp.FindStringSubmatch(l)
		pass := passwordAndPolicy{unsafeAtoi(match[1]), unsafeAtoi(match[2]), match[3][0], []byte(match[4])}
		input = append(input, pass)
	}
	fmt.Println(solve(input))
}
