package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

var signals = map[int][]byte{
	0: {'a', 'b', 'c', 'e', 'f', 'g'},
	1: {'c', 'f'},
	2: {'a', 'c', 'd', 'e', 'g'},
	3: {'a', 'c', 'd', 'f', 'g'},
	4: {'b', 'c', 'd', 'f'},
	5: {'a', 'b', 'd', 'f', 'g'},
	6: {'a', 'b', 'd', 'e', 'f', 'g'},
	7: {'a', 'c', 'f'},
	8: {'a', 'b', 'c', 'd', 'e', 'f', 'g'},
	9: {'a', 'b', 'c', 'd', 'f', 'g'},
}

func main() {
	data := files.ReadLines(os.Args[1])
	res := 0
	for _, line := range data {
		parts := strings.Fields(line)
		for _, out := range parts[11:] {
			if len(out) == len(signals[1]) || len(out) == len(signals[4]) ||
				len(out) == len(signals[7]) || len(out) == len(signals[8]) {
				res++
			}
		}
	}
	fmt.Println(res)
}
