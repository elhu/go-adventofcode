package main

import "fmt"

const input = 29000000

func main() {
	target := input
	presents := make([]int, target)
	res := input
	for i := 1; i < target; i++ {
		visited := 0
		for j := i; j < target && visited < 50; j += i {
			visited++
			presents[j] += i * 11
			if presents[j] >= target && j < res {
				res = j
			}
		}
	}
	fmt.Println(res)
}
