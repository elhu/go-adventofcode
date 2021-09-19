package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func computeNextPassword(password []byte) {
	rank := len(password) - 1
	password[rank] = 'a' + (password[rank]+1-'a')%26
	for password[rank] == 'a' {
		rank--
		password[rank] = 'a' + (password[rank]+1-'a')%26
	}
}

func valid(password []byte) bool {
	for _, c := range password {
		if c == 'i' || c == 'o' || c == 'l' {
			return false
		}
	}
	straight := false
	for i := 2; i < len(password); i++ {
		if password[i] == password[i-1]+1 && password[i] == password[i-2]+2 {
			straight = true
		}
	}
	var pairs []int
	for i := 1; i < len(password); i++ {
		if password[i] == password[i-1] {
			pairs = append(pairs, i-1)
		}
	}
	repeatPairs := false
	if len(pairs) > 2 {
		repeatPairs = true
	} else if len(pairs) == 2 {
		repeatPairs = pairs[0]+1 < pairs[1]
	}
	return straight && repeatPairs
}

func solve(input string) []byte {
	password := []byte(input)
	computeNextPassword(password)
	for !valid(password) {
		computeNextPassword(password)
	}
	return password
}

func main() {
	fmt.Println(string(solve(os.Args[1])))
}
