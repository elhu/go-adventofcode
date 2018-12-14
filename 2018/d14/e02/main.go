package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printState(posA, posB int, recipes []int) {
	for i, r := range recipes {
		if i == posA {
			fmt.Printf("(%d)", r)
		} else if i == posB {
			fmt.Printf("[%d]", r)
		} else {
			fmt.Printf(" %d ", r)
		}
	}
	fmt.Println("")
}

func equals(a, b []int) bool {
	if len(a) != len(b) {
		fmt.Println(a, b)
		panic("WTF")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func solve(posA, posB int, target []int, recipes []int) int {
	for {
		if len(recipes) >= len(target) {
			if equals(recipes[len(recipes)-len(target):], target) {
				return len(recipes) - len(target)
			}
		}
		if len(recipes) > len(target) {
			if equals(recipes[len(recipes)-len(target)-1:len(recipes)-1], target) {
				return len(recipes) - len(target) - 1
			}
		}
		newRecipeVal := recipes[posA] + recipes[posB]
		if newRecipeVal < 10 {
			recipes = append(recipes, newRecipeVal)
		} else {
			recipes = append(recipes, newRecipeVal/10, newRecipeVal%10)
		}
		posA = (posA + 1 + recipes[posA]) % len(recipes)
		posB = (posB + 1 + recipes[posB]) % len(recipes)
	}
}

func main() {
	input := strings.Split(os.Args[1], "")
	target := make([]int, len(input))
	for i, j := range input {
		target[i], _ = strconv.Atoi(j)
	}
	recipes := make([]int, 2, 1000000)
	recipes[0] = 3
	recipes[1] = 7
	fmt.Println(solve(0, 1, target, recipes))
}
