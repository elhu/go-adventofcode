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

type Node struct {
	value    byte
	children []byte
}

func computeAddresses(address, mask []byte) [][]byte {
	res := make([][]byte, 0)
	initial := make([]byte, len(mask))
	for i, c := range address {
		if mask[i] == '1' {
			initial[i] = '1'
		} else if mask[i] == '0' {
			initial[i] = c
		} else {
			initial[i] = 'X'
		}
	}
	res = append(res, initial)
	changed := 1
	for changed > 0 {
		changed = 0
		for i, r := range res {
			for j, c := range r {
				if c == 'X' {
					res[i][j] = '1'
					cpy := append([]byte{}, r...)
					cpy[j] = '0'
					res = append(res, cpy)
					changed++
				}
			}
		}
	}

	return res
}

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
			binAddress := []byte(strconv.FormatInt(address, 2))
			for _, addr := range computeAddresses(append([]byte(strings.Repeat("0", 36-len(binAddress))), binAddress...), mask) {
				a, err := strconv.ParseInt(string(addr), 2, 64)
				check(err)
				memory[a] = val
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
