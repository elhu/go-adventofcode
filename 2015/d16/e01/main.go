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
	res, err := strconv.Atoi(str)
	check(err)
	return res
}

func parse(input []string) map[int]map[string]int {
	sues := make(map[int]map[string]int)
	for i, l := range input {
		sues[i+1] = make(map[string]int)
		parts := strings.SplitN(l, ": ", 2)
		attributes := parts[1]
		parts = strings.Split(attributes, ", ")
		for _, p := range parts {
			attr := strings.Split(p, ": ")
			sues[i+1][attr[0]] = atoi(attr[1])
		}
	}
	return sues
}

var facts = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func solve(sues map[int]map[string]int, facts map[string]int) int {
	for k, v := range facts {
		for i, sue := range sues {
			if val, found := sue[k]; found && val != v {
				delete(sues, i)
			}
		}
	}
	for k := range sues {
		return k
	}
	panic("wtf")
}

func main() {
	rawData, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(rawData), "\n"), "\n")
	sues := parse(input)
	fmt.Println(solve(sues, facts))
	// fmt.Println(sues)
}
