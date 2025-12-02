package main

import (
	"adventofcode/utils/files"
	"adventofcode/utils/sets"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parts(a string, count int) []string {
	var res []string
	size := len(a) / count
	for i := 0; i < len(a); i += size {
		res = append(res, a[i:i+size])
	}
	return res
}

func solve(ranges [][2]int) int {
	res := 0
	for _, r := range ranges {
		invalidIDs := sets.New[int]()
		for z := r[0]; z <= r[1]; z++ {
			a := strconv.Itoa(z)
			for i := 2; i <= len(a); i++ {
				if len(a)%i == 0 {
					parts := parts(a, i)
					equal := true
					for j := 1; j < len(parts); j++ {
						if parts[j] != parts[0] {
							equal = false
							break
						}
					}
					if equal {
						invalidIDs.Add(z)
					}
				}
			}
		}
		for _, k := range invalidIDs.Members() {
			res += k
		}
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	rawRanges := strings.Split(data, ",")
	ranges := make([][2]int, 0, len(rawRanges))
	for _, r := range rawRanges {
		bounds := strings.Split(r, "-")
		lower, err := strconv.Atoi(bounds[0])
		if err != nil {
			panic(err)
		}
		upper, err := strconv.Atoi(bounds[1])
		if err != nil {
			panic(err)
		}
		ranges = append(ranges, [2]int{lower, upper})
	}
	fmt.Println(solve(ranges))
}
