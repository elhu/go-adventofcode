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

func findLeft(values []byte, pos int) int {
	if values[pos-1] >= '0' && values[pos-1] <= '9' {
		return pos - 1
	} else if values[pos-1] == ')' {
		parensClosed := 1
		for i := pos - 2; ; i-- {
			if values[i] == ')' {
				parensClosed++
			} else if values[i] == '(' {
				parensClosed--
			}
			if parensClosed == 0 {
				return i
			}
		}
	}
	panic("wtf left?")
}

func findRight(values []byte, pos int) int {
	if values[pos+1] >= '0' && values[pos+1] <= '9' {
		return pos + 1
	} else if values[pos+1] == '(' {
		parensOpen := 1
		for i := pos + 2; ; i++ {
			if values[i] == '(' {
				parensOpen++
			} else if values[i] == ')' {
				parensOpen--
			}
			if parensOpen == 0 {
				return i
			}
		}
	}
	panic("wtf right?")
}

func transform(res []byte) []byte {
	for i := 0; i < len(res); i++ {
		c := res[i]
		if c == '+' {
			res = append(res, []byte{0, 0}...) // make room at the end of the slice
			lft := findLeft(res, i)
			copy(res[lft+1:], res[lft:])
			res[lft] = '('
			i++
			rgt := findRight(res, i)
			copy(res[rgt+2:], res[rgt+1:])
			res[rgt+1] = ')'
		}
	}
	return res
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
		res := calculate(transform([]byte(str)))
		total += res
	}
	fmt.Println(total)
}
