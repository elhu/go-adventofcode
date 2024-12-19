package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func solve(designs []string, patterns []string) int {
	es := fmt.Sprintf("^(%s)+$", strings.Join(patterns, "|"))
	exp := regexp.MustCompile(es)
	res := 0
	for _, design := range designs {
		if exp.MatchString(design) {
			res++
		}
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	parts := strings.Split(data, "\n\n")

	patterns := strings.Split(parts[0], ", ")
	designs := strings.Split(parts[1], "\n")
	fmt.Println(solve(designs, patterns))
}
