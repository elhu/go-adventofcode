package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Node struct {
	name        string
	containedBy []*Node
}

func displayNodes(nodes map[string]*Node) {
	for _, n := range nodes {
		fmt.Println(fmt.Sprintf("%s is contained by %v", n.name, n.containedBy))
	}
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
			newNode = &Node{matches[1], make([]*Node, 0)}
		}
		nodes[newNode.name] = newNode
		contains := strings.TrimRight(matches[2], ".")
		parts := strings.Split(contains, ", ")
		for _, p := range parts {
			ms := partExp.FindStringSubmatch(p)
			name := ms[2]
			node, found := nodes[name]
			if !found {
				node = &Node{name, make([]*Node, 0)}
				nodes[name] = node
			}
			node.containedBy = append(node.containedBy, newNode)
		}
	}
	return nodes
}

func solve(nodes map[string]*Node, target string) int {
	res := make(map[string]struct{})
	curr := nodes[target]
	var queue []*Node
	queue = append(queue, curr.containedBy...)

	var popped *Node
	for len(queue) > 0 {
		popped, queue = queue[0], queue[1:]
		if _, visited := res[popped.name]; !visited {
			res[popped.name] = struct{}{}
			queue = append(queue, nodes[popped.name].containedBy...)
		}
	}

	return len(res)
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	nodes := buildGraph(lines)
	fmt.Println(solve(nodes, "shiny gold"))
}
