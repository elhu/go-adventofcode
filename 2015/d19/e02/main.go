package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseReplacements(input []string) map[string]string {
	res := make(map[string]string)
	for _, l := range input {
		parts := strings.Split(l, " => ")
		res[parts[1]] = parts[0]
	}
	return res
}

func allIndexes(str string, search string) []int {
	var res []int
	for i := 0; i < len(str); i++ {
		idx := strings.Index(str[i:], search)
		if idx == -1 {
			break
		}
		res = append(res, i+idx)
		i += idx
	}
	return res
}

func shuffle(keys []string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(keys), func(i, j int) { keys[i], keys[j] = keys[j], keys[i] })
}

func solve(molecule string, replacements map[string]string) int {
	steps := 0
	var keys []string
	for k := range replacements {
		keys = append(keys, k)
	}
	currMolecule := molecule
	for currMolecule != "e" {
		// Try different sequences of keys until one works
		shuffle(keys)
		replaced := false
		for _, to := range keys {
			from := replacements[to]
			if strings.Index(currMolecule, to) != -1 {
				currMolecule = strings.Replace(currMolecule, to, from, 1)
				steps++
				replaced = true
				break
			}
		}
		if !replaced {
			steps = 0
			currMolecule = molecule
			fmt.Println("No possible replacement")
		}
	}
	return steps
}

func main() {
	rawData, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(rawData), "\n"), "\n")
	molecule := input[len(input)-1]
	replacements := parseReplacements(input[:len(input)-2])
	fmt.Println(solve(molecule, replacements))
}
