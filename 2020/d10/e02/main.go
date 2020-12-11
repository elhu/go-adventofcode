package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	check(err)
	return i
}

/*
Sort the input
Keep a list of the number of branches at each adapter
Traverse the adapters, starting from the end.
At the end, the number of branches is 1 (the device's adapter)
For each adapter (adapter A):
* check the next 3 ones in the chain (there's the only ones that could connect), named adapter B:
* If the adapter A connects adapter B (i.e its value is within +3 of the adapter A):
	* Add the number of branches of adapter B to adapter A

Keeping going until you get to the first adapter (the charging outlet)
*/

func solve(adapters []int) int {
	sort.Slice(adapters, func(i, j int) bool { return adapters[i] < adapters[j] })
	// Pad the adapters with fake ones to avoid having to check boundary conditions
	adapters = append(adapters, []int{adapters[len(adapters)-1] + 3, adapters[len(adapters)-1] + 6}...)
	branches := make([]int, len(adapters))
	branches[len(branches)-1] = 1
	branches[len(branches)-2] = 1
	branches[len(branches)-3] = 1
	for i := len(adapters) - 4; i >= 0; i-- {
		for _, dist := range []int{1, 2, 3} {
			if adapters[i+dist]-adapters[i] <= 3 {
				branches[i] += branches[i+dist]
			}
		}
	}
	return branches[0]
}
func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	var adapters []int
	adapters = append(adapters, 0)
	maxJolt := 0
	for _, l := range lines {
		i := atoi(l)
		if i > maxJolt {
			maxJolt = i
		}
		adapters = append(adapters, i)
	}
	adapters = append(adapters, maxJolt+3)
	fmt.Println(solve(adapters))
}
