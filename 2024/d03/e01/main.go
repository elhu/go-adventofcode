package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`mul\(\d+,\d+\)`)

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	muls := re.FindAllString(data, -1)
	res := 0
	for _, mul := range muls {
		var l, r int
		fmt.Sscanf(mul, "mul(%d,%d)", &l, &r)
		res += l * r
	}
	fmt.Println(res)
}
