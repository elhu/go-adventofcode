package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`(mul\(\d+,\d+\)|do\(\)|don't\(\))`)

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	ops := re.FindAllString(data, -1)
	res := 0
	mulEnabled := true
	for _, op := range ops {
		if op == "do()" {
			mulEnabled = true
		} else if op == "don't()" {
			mulEnabled = false
		} else if mulEnabled {
			var l, r int
			fmt.Sscanf(op, "mul(%d,%d)", &l, &r)
			res += l * r
		}
	}
	fmt.Println(res)
}
