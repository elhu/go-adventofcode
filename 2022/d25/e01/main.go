package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var digits = map[rune]int{
	'=': -2,
	'-': -1,
	'0': 0,
	'1': 1,
	'2': 2,
}

func pow(m, n int) int {

	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func toDecimal(str string) int {
	res := 0
	for i, c := range str {
		res += digits[c] * pow(len(str)-i-1, 5)
	}
	return res
}

func toSnafu(num int) string {
	var places []int
	place := 1

	for num > 0 {
		remainder := num % 5
		if remainder == 3 {
			remainder = -2
		} else if remainder == 4 {
			remainder = -1
		}
		places = append(places, remainder)
		num -= remainder
		num /= 5
		place++
	}
	var res []string
	for i := len(places) - 1; i >= 0; i-- {
		if places[i] == -2 {
			res = append(res, "=")
		} else if places[i] == -1 {
			res = append(res, "-")
		} else {
			res = append(res, strconv.Itoa(places[i]))
		}
	}
	return strings.Join(res, "")
}

func solve(data []string) string {
	sum := 0
	for _, l := range data {
		sum += toDecimal(l)
	}
	return toSnafu(sum)
}

func main() {
	data := files.ReadLines(os.Args[1])
	fmt.Println(solve(data))
}
