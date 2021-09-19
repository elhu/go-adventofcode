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

func parseReplacements(input []string) map[string][]string {
	res := make(map[string][]string)
	for _, l := range input {
		parts := strings.Split(l, " => ")
		res[parts[0]] = append(res[parts[0]], parts[1])
	}
	// fmt.Println(res)
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

func solve(molecule string, replacements map[string][]string) int {
	generated := make(map[string]struct{})
	for k, v := range replacements {
		for _, newMol := range v {
			for _, idx := range allIndexes(molecule, k) {
				newMolecule := fmt.Sprintf("%s%s%s", molecule[:idx], newMol, molecule[idx+len(k):])
				generated[newMolecule] = struct{}{}
			}
		}
	}
	return len(generated)
}

func main() {
	rawData, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(rawData), "\n"), "\n")
	molecule := input[len(input)-1]
	replacements := parseReplacements(input[:len(input)-2])
	fmt.Println(solve(molecule, replacements))
}
