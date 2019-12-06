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
	objects []*orbitNode
}

func buildTree(orbitsData []string) map[string]*orbitNode {
	orbits := make(map[string]*orbitNode)
	orbits["COM"] = &orbitNode{"COM", make([]*orbitNode, 0)}
	for _, line := range orbitsData {
		parts := strings.Split(line, ")")
		object := parts[1]
		orbits[object] = &orbitNode{object, make([]*orbitNode, 0)}
	}
	for _, line := range orbitsData {
		parts := strings.Split(line, ")")
		center, object := parts[0], parts[1]
		orbits[center].objects = append(orbits[center].objects, orbits[object])
	}
	return orbits
}

func solve(orbit *orbitNode, depth int) int {
	count := 0
	for _, c := range orbit.objects {
		// fmt.Printf("Processing %s)%s (at depth %d)\n", orbit.center, c.center, count)
		count += solve(c, depth+1)
	}
	return depth + count
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	orbitsData := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	orbits := buildTree(orbitsData)
	fmt.Println(solve(orbits["COM"], 0))
}
