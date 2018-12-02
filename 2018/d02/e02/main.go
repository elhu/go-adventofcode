package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func distance(a, b []byte) int {
	res := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			res++
		}
	}
	return res
}

func formatResponse(a, b []byte) []byte {
	res := make([]byte, 0)
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			res = append(res, a[i])
		}
	}
	return res
}

func solve(lines [][]byte) []byte {
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			if distance(lines[i], lines[j]) == 1 {
				return formatResponse(lines[i], lines[j])
			}
		}
	}
	return []byte("Not Found")
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	lines := bytes.Split(data, []byte{'\n'})
	fmt.Println(string(solve(lines)))
}
