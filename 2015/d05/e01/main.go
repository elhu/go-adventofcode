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

func isNice(str string) bool {
	vowelCount := 0
	doubleLetter := false
	for i := 0; i < len(str); i++ {
		if str[i] == 'a' || str[i] == 'e' || str[i] == 'i' || str[i] == 'o' || str[i] == 'u' {
			vowelCount++
		}
		if i > 0 {
			if str[i] == str[i-1] {
				doubleLetter = true
			}
			if str[i-1:i+1] == "ab" || str[i-1:i+1] == "cd" || str[i-1:i+1] == "pq" || str[i-1:i+1] == "xy" {
				return false
			}
		}
	}
	return vowelCount >= 3 && doubleLetter
}

func solve(input []string) int {
	nice := 0
	for _, l := range input {
		if isNice(l) {
			nice++
		}
	}
	return nice
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	fmt.Println(solve(input))
}
