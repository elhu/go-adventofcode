package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func atoi(str string) int64 {
	i, err := strconv.Atoi(str)
	check(err)
	return int64(i)
}

var memExp = regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)

func solve(input []string) int64 {
	memory := make(map[int64][]byte)
	var mask []byte
	for _, line := range input {
		if strings.HasPrefix(line, "mask = ") {
			mask = []byte(line[7:])
		} else {
			match := memExp.FindStringSubmatch(line)
			address := atoi(match[1])
			val := []byte(strconv.FormatInt(atoi(match[2]), 2))
			val = append([]byte(strings.Repeat("0", 36-len(val))), val...)
			memory[address] = append([]byte{}, mask...)
			for i, c := range val {
				if memory[address][i] == 'X' {
					memory[address][i] = c
				}
			}
		}
	}
	res := int64(0)
	for _, v := range memory {
		i, err := strconv.ParseInt(string(v), 2, 64)
		check(err)
		res += i
	}
	return res
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	fmt.Println(solve(lines))
}
