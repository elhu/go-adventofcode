package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Name      string
	Neighbors []*Node
}

func connect(a, b *Node) {
	a.Neighbors = append(a.Neighbors, b)
	b.Neighbors = append(b.Neighbors, a)
}

func buildPath(source, destination string, visited map[string]string) {
	name, found := visited[destination]
	var path []string
	for found {
		if name == source {
			break
		}
		path = append(path, name)
		name, found = visited[name]
	}
	fmt.Printf("%s %s %s\n", source, destination, strings.Join(path, " "))
}

func bfs(source, destination *Node) {
	var queue = []*Node{source}
	var current *Node
	visited := make(map[string]string)
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		if current == destination {
			buildPath(source.Name, destination.Name, visited)
			break
		}

		for _, n := range current.Neighbors {
			if _, found := visited[n.Name]; !found {
				visited[n.Name] = current.Name
				queue = append(queue, n)
			}
		}
	}
}

func main() {
	r2s := []*Node{&Node{Name: "RA2"}, &Node{Name: "RB2"}, &Node{Name: "RC2"}, &Node{Name: "RD2"}}
	r4s := []*Node{&Node{Name: "RA4"}, &Node{Name: "RB4"}, &Node{Name: "RC4"}, &Node{Name: "RD4"}}
	r6s := []*Node{&Node{Name: "RA6"}, &Node{Name: "RB6"}, &Node{Name: "RC6"}, &Node{Name: "RD6"}}
	r8s := []*Node{&Node{Name: "RA8"}, &Node{Name: "RB8"}, &Node{Name: "RC8"}, &Node{Name: "RD8"}}

	rs := [][]*Node{r2s, r4s, r6s, r8s}
	for _, r := range rs {
		for i := 1; i < len(r); i++ {
			connect(r[i], r[i-1])
		}
	}

	var hs []*Node
	for i := 0; i <= 10; i++ {
		hs = append(hs, &Node{Name: fmt.Sprintf("H%d", i)})
	}
	for i := 1; i <= 10; i++ {
		connect(hs[i], hs[i-1])
	}

	connect(hs[2], r2s[0])
	connect(hs[4], r4s[0])
	connect(hs[6], r6s[0])
	connect(hs[8], r8s[0])

	// for i, r := range rs {
	// 	for _, or := range rs[i+1:] {
	// 		for _, n := range r {
	// 			for _, on := range or {
	// 				bfs(n, on)
	// 			}
	// 		}
	// 	}
	// }
	allowedHs := []*Node{hs[0], hs[1], hs[3], hs[5], hs[7], hs[9], hs[10]}
	for _, r := range rs {
		for _, n := range r {
			for _, h := range allowedHs {
				bfs(n, h)
			}
		}
	}
}
