package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func atoi(str string) int {
	n, err := strconv.Atoi(str)
	check(err)
	return n
}

// Perm calls f with each permutation of a.
func Perm(a []string, f func([]string)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []string, f func([]string), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func parse(input []string) map[string]map[string]int {
	res := make(map[string]map[string]int)
	for _, l := range input {
		var verb, source, target string
		var value int
		_, err := fmt.Sscanf(l, "%s would %s %d happiness units by sitting next to %s", &source, &verb, &value, &target)
		check(err)
		target = target[:len(target)-1]
		if verb == "lose" {
			value = -value
		}
		if _, found := res[source]; !found {
			res[source] = make(map[string]int)
		}
		res[source][target] = value
	}
	return res
}

func solve(data map[string]map[string]int) int {
	var names []string
	for k := range data {
		names = append(names, k)
	}
	names = append(names, "self")
	max := 0
	Perm(names, func(orderedNames []string) {
		curr := 0
		for i := 1; i < len(orderedNames); i++ {
			left := orderedNames[i-1]
			right := orderedNames[i]
			curr += data[left][right]
			curr += data[right][left]
		}
		left := orderedNames[0]
		right := orderedNames[len(orderedNames)-1]
		curr += data[left][right]
		curr += data[right][left]
		if curr > max {
			max = curr
		}
	})
	return max
}

func main() {
	rawData, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(rawData), "\n"), "\n")
	data := parse(input)
	fmt.Println(solve(data))
}
