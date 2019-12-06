package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type orbitNode struct {
	center  string
	parent  *orbitNode
	objects []*orbitNode
}

func buildTree(orbitsData []string) map[string]*orbitNode {
	orbits := make(map[string]*orbitNode)
	orbits["COM"] = &orbitNode{"COM", nil, make([]*orbitNode, 0)}
	for _, line := range orbitsData {
		parts := strings.Split(line, ")")
		object := parts[1]
		orbits[object] = &orbitNode{object, nil, make([]*orbitNode, 0)}
	}
	for _, line := range orbitsData {
		parts := strings.Split(line, ")")
		center, object := parts[0], parts[1]
		orbits[object].parent = orbits[center]
		orbits[center].objects = append(orbits[center].objects, orbits[object])
	}
	return orbits
}

func isAncestor(a, b *orbitNode) (bool, int) {
	var c *orbitNode
	var hops int
	for c = a; c.parent != b; c = c.parent {
		hops++
		if c.parent == nil {
			break
		}
	}
	if c.parent != nil {
		return true, hops
	}
	return false, hops
}

func solve(origin *orbitNode, destination *orbitNode, orbits map[string]*orbitNode, hops int) int {
	if has, depth := isAncestor(destination, origin); has {
		return depth + hops
	}
	fmt.Printf("It is NOT!\n")
	return solve(origin.parent, destination, orbits, hops+1)
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	orbitsData := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	orbits := buildTree(orbitsData)
	fmt.Println(solve(orbits["YOU"].parent, orbits["SAN"], orbits, 0))
}
