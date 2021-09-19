package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

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

type Edge struct {
	distance int
	node     *Node
}

type Node struct {
	name  string
	edges map[string]Edge
}

func parse(input []string) map[string]*Node {
	res := make(map[string]*Node)
	for _, l := range input {
		parts := strings.Fields(l)
		from, to, distance := parts[0], parts[2], atoi(parts[4])
		if _, found := res[from]; !found {
			res[from] = &Node{name: from, edges: make(map[string]Edge)}
		}
		if _, found := res[to]; !found {
			res[to] = &Node{name: to, edges: make(map[string]Edge)}
		}
		res[from].edges[to] = Edge{distance, res[to]}
		res[to].edges[from] = Edge{distance, res[from]}
	}
	return res
}

func solve(nodes map[string]*Node) int {
	var names []string
	for k := range nodes {
		names = append(names, k)
	}
	longest := 0
	Perm(names, func(path []string) {
		distance := 0
		for i := 1; i < len(path); i++ {
			from, to := path[i-1], path[i]
			edge, found := nodes[from].edges[to]
			if found {
				distance += edge.distance
			} else {
				return
			}
		}
		if distance > longest {
			longest = distance
		}
	})
	return longest
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	nodes := parse(input)
	fmt.Println(solve(nodes))
}
