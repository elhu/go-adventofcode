package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	check(err)
	return i
}

func findEndPos(values []byte) int {
	parensOpen := 1
	for i, c := range values[1:] {
		if c == '(' {
			parensOpen++
		}
		if c == ')' {
			parensOpen--
		}
		if parensOpen == 0 {
			return i + 1
		}
	}
	panic(fmt.Sprintf("Cannot find matching closing parenthesis in %s\n", values))
}

func calculate(values []byte) int {
	result := 0
	operator := byte('+')
	for i := 0; i < len(values); i++ {
		c := values[i]
		if c >= '0' && c <= '9' {
			if operator == '+' {
				result += int(c - '0')
			} else if operator == '*' {
				result *= int(c - '0')
			}
		} else if c == '*' || c == '+' {
			operator = c
		} else if c == '(' {
			endPos := i + findEndPos(values[i:])
			sub := calculate(values[i+1 : endPos])
			if operator == '+' {
				result += sub
			} else if operator == '*' {
				result *= sub
			}
			i = endPos
		}
	}
	return result
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	total := 0
	for _, l := range input {
		str := strings.ReplaceAll(l, " ", "")
		res := calculate([]byte(str))
		total += res
	}
	fmt.Println(total)
}
