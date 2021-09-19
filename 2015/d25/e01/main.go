package main

import "fmt"

func solve(y, x int) int {
	last := 20151125
	maxI, i := 2, 2
	j := 1
	for {
		last = (last * 252533) % 33554393
		if i == y && j == x {
			return last
		}
		if i <= 1 {
			maxI++
			i = maxI
			j = 1
			continue
		}
		j++
		i--
	}
}

func main() {
	fmt.Println(solve(3010, 3019))
}
