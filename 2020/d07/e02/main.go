package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type CountBag struct {
	count int
	bag   *Node
}

type Node struct {
	name     string
	contains []CountBag
}

func buildGraph(lines []string) map[string]*Node {
	lineExp := regexp.MustCompile(`^(\w+ \w+) bags contain (.+).$`)
	deadEndExp := regexp.MustCompile(`.+ contain no other bags.$`)
	partExp := regexp.MustCompile(`^(\d+) (\w+ \w+) bags?$`)

	nodes := make(map[string]*Node)

	for _, l := range lines {
		if deadEndExp.MatchString(l) {
			continue
		}
		matches := lineExp.FindStringSubmatch(l)
		newNode, found := nodes[matches[1]]
		if !found {
			newNode = &Node{matches[1], make([]CountBag, 0)}
		}
		nodes[newNode.name] = newNode
		contains := strings.TrimRight(matches[2], ".")
		parts := strings.Split(contains, ", ")
		for _, p := range parts {
			ms := partExp.FindStringSubmatch(p)
			count, err := strconv.Atoi(ms[1])
			check(err)
			name := ms[2]
			node, found := nodes[name]
			if !found {
				node = &Node{name, make([]CountBag, 0)}
				nodes[name] = node
			}
			newNode.contains = append(newNode.contains, CountBag{count, node})
		}
	}
	return nodes
}

func dfs(root *Node) int {
	res := 1
	for _, c := range root.contains {
		res += c.count * dfs(c.bag)
	}
	return res
}

func solve(nodes map[string]*Node, target string) int {
	return dfs(nodes[target]) - 1
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	nodes := buildGraph(lines)
	fmt.Println(solve(nodes, "shiny gold"))
}
