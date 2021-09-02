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

func atoi(str string) int {
	res, err := strconv.Atoi(str)
	check(err)
	return res
}

type Node struct {
	x, y       int
	size, used int
}

func (n *Node) Size() int {
	return n.size
}

func (n *Node) Used() int {
	return n.used
}

func (n *Node) Avail() int {
	return n.size - n.used
}

func (n *Node) SameAs(a Node) bool {
	return n.x == a.x && n.y == a.y
}

var nodeExp = regexp.MustCompile(`/dev/grid/node-x(\d+)-y(\d+)\W+(\d+)T\W+(\d+)T\W+\d+T\W+\d+%`)

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	input = input[2:]
	var nodes [31][31]Node
	for _, l := range input {
		match := nodeExp.FindStringSubmatch(l)
		x := atoi(match[1])
		y := atoi(match[2])
		size := atoi(match[3])
		used := atoi(match[4])
		node := Node{x, y, size, used}
		nodes[y][x] = node
	}
	for y := range nodes {
		for _, n := range nodes[y] {
			if n.Size() == n.Avail() {
				fmt.Printf("[___] ")
			} else if n.Size() > 100 {
				fmt.Printf("------")
			} else {
				fmt.Printf("%d/%d ", n.Avail(), n.Size())
			}
		}
		fmt.Println(" ")
	}
	// Print data, solve manually
}
