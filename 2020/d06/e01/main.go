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
	data = strings.ReplaceAll(data, "\n", "")
	unique := make(map[rune]struct{})
	for _, c := range data {
		unique[c] = struct{}{}
	}
	return len(unique)
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
