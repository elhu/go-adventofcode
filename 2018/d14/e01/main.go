package main

import (
	"fmt"
	"os"
	"strconv"
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

func solve(posA, posB, target int, recipes []int) []int {
	for len(recipes) < target {
		newRecipeVal := recipes[posA] + recipes[posB]
		if newRecipeVal < 10 {
			recipes = append(recipes, newRecipeVal)
		} else {
			recipes = append(recipes, newRecipeVal/10, newRecipeVal%10)
		}
		posA = (posA + 1 + recipes[posA]) % len(recipes)
		posB = (posB + 1 + recipes[posB]) % len(recipes)
		// printState(posA, posB, recipes)
	}
	return recipes[target-10 : target]
}

func main() {
	input, _ := strconv.Atoi(os.Args[1])
	recipes := make([]int, 2, input+10)
	recipes[0] = 3
	recipes[1] = 7
	for _, r := range solve(0, 1, input+10, recipes) {
		fmt.Printf("%d", r)
	}
	fmt.Println("")
}
