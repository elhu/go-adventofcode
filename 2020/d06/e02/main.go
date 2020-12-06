package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func processGroup(data string) int {
	people := strings.Count(data, "\n") + 1
	data = strings.ReplaceAll(data, "\n", "")
	counts := make(map[rune]int)
	for _, c := range data {
		counts[c]++
	}
	res := 0
	for _, v := range counts {
		if v == people {
			res++
		}
	}
	return res
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	groups := strings.Split(strings.TrimRight(string(data), "\n"), "\n\n")
	res := 0
	for _, g := range groups {
		res += processGroup(g)
	}
	fmt.Println(res)
}
