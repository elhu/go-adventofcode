package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var convert = map[string]int{
	"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
	"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
}

func rev(s string) string {
	res := ""
	for _, c := range s {
		res = string(c) + res
	}
	return res
}

func solve(strings []string) int {
	leftExp := regexp.MustCompile(`([1-9]|one|two|three|four|five|six|seven|eight|nine)`)
	rightExp := regexp.MustCompile(`(enin|thgie|neves|xis|evif|ruof|eerht|owt|eno|[1-9])`)
	res := 0
	for _, s := range strings {
		left := leftExp.FindString(s)
		right := rightExp.FindString(rev(s))
		res += convert[left]*10 + convert[rev(right)]
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	strings := strings.Split(data, "\n")
	fmt.Println(solve(strings))
}
