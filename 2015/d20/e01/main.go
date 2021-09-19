package main

import "fmt"

const input = 29000000

func main() {
	target := input / 10
	presents := make([]int, target)
	res := input
	for i := 1; i < target; i++ {
		for j := i; j < target; j += i {
			presents[j] += i
			if presents[j] >= target && j < res {
				res = j
			}
		}
	}
	fmt.Println(res)
}
