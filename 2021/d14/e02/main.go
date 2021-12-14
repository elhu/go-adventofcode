package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
)

func parsePairs(data []string) map[string]byte {
	res := make(map[string]byte)
	var pair string
	var insert byte
	for _, d := range data {
		fmt.Sscanf(d, "%s -> %c", &pair, &insert)
		res[pair] = insert
	}
	return res
}

func parseTemplate(ts string) map[string]int {
	res := make(map[string]int)
	for i := 0; i < len(ts)-1; i++ {
		res[ts[i:i+2]]++
	}
	return res
}

func computeScore(template map[string]int) int {
	counts := make(map[string]int)
	for k, v := range template {
		counts[k[0:1]] += v
		counts[k[1:]] += v
	}
	for k, v := range counts {
		counts[k] = (v + 1) / 2
	}
	min := 999999999999999
	max := 0
	for _, v := range counts {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max - min
}

func solve(template map[string]int, pairs map[string]byte) int {
	for i := 0; i < turns; i++ {
		newTemplate := make(map[string]int)
		for k, v := range template {
			if insert, found := pairs[k]; found {
				k1 := fmt.Sprintf("%c%c", k[0], insert)
				k2 := fmt.Sprintf("%c%c", insert, k[1])
				newTemplate[k1] += v
				newTemplate[k2] += v
			} else {
				panic("WTF")
			}
			template = newTemplate
		}
	}
	return computeScore(template)
}

const turns = 40

func main() {
	data := files.ReadLines(os.Args[1])
	template := parseTemplate(data[0])
	pairs := parsePairs(data[2:])
	fmt.Println(solve(template, pairs))
}
