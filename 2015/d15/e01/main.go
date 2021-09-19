package main

import "fmt"

// Frosting: capacity 4, durability -2, flavor 0, texture 0, calories 5
// Candy: capacity 0, durability 5, flavor -1, texture 0, calories 8
// Butterscotch: capacity -1, durability 0, flavor 5, texture 0, calories 6
// Sugar: capacity 0, durability 0, flavor -2, texture 2, calories 1

var data = [][]int{
	{4, -2, 0, 0, 5},
	{0, 5, -1, 0, 8},
	{-1, 0, 5, 0, 6},
	{0, 0, -2, 2, 1},
}

func main() {
	max := 0
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100-i; j++ {
			for k := 0; k <= 100-i-j; k++ {
				l := 100 - i - j - k
				cap := data[0][0]*i + data[1][0]*j + data[2][0]*k + data[3][0]*l
				dur := data[0][1]*i + data[1][1]*j + data[2][1]*k + data[3][1]*l
				fla := data[0][2]*i + data[1][2]*j + data[2][2]*k + data[3][2]*l
				tex := data[0][3]*i + data[1][3]*j + data[2][3]*k + data[3][3]*l

				if cap <= 0 || dur <= 0 || fla <= 0 || tex <= 0 {
					continue
				}
				score := cap * dur * fla * tex
				if score > max {
					max = score
				}
			}
		}
	}
	fmt.Println(max)
}
