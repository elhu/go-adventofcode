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

func comb(data []int) [][]int {
	var res [][]int
	for i := 1; i < (1 << len(data)); i++ {
		var subRes []int
		for k := range data {
			if (i>>k)&1 == 1 {
				subRes = append(subRes, data[k])
			}
		}
		res = append(res, subRes)
	}
	return res
}

func solve(containers []int) int {
	res := 0
	for _, c := range comb(containers) {
		sum := 0
		for _, e := range c {
			sum += e
		}
		if sum == target {
			res++
		}
	}
	return res
}

const target = 150

func main() {
	rawData, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(rawData), "\n"), "\n")
	var containers []int
	for _, l := range input {
		containers = append(containers, atoi(l))
	}
	fmt.Println(solve(containers))
}
