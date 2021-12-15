package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"adventofcode/utils/sets/intset"
	"fmt"
	"os"
)

type Edge struct {
	Cost int
	Node *Node
}

type Node struct {
	Coords coords2d.Coords2d
	Value  int
	Edges  []*Edge
}

func coordsToKey(c coords2d.Coords2d) int {
	return c.Y*10000000 + c.X
}

func parseMap(data []string) map[int]*Node {
	nodes := make(map[int]*Node)
	for i, l := range data {
		for j, c := range l {
			coord := coords2d.Coords2d{X: j, Y: i}
			nodes[coordsToKey(coord)] = &Node{Coords: coord, Value: int(c - '0')}
		}
	}
	for i, l := range data {
		for j, c := range l {
			currentNode := nodes[coordsToKey(coords2d.Coords2d{X: j, Y: i})]
			neighbors := []coords2d.Coords2d{
				{X: j, Y: i - 1},
				{X: j, Y: i + 1},
				{X: j - 1, Y: i},
				{X: j + 1, Y: i},
			}
			for _, coord := range neighbors {
				if node, exists := nodes[coordsToKey(coord)]; exists {
					node.Edges = append(node.Edges, &Edge{Cost: int(c - '0'), Node: currentNode})
				}
			}
		}
	}
	return nodes
}

const MAXINT = 2147483647

func heuristic(n, end *Node) int {
	return coords2d.Distance(n.Coords, end.Coords)
}

func minFScore(open *intset.IntSet, fScore map[int]int) int {
	minScore := MAXINT
	minKey := -1
	open.Each(func(key int) {
		var score = MAXINT
		if s, found := fScore[key]; found {
			score = s
		}
		if score <= minScore {
			minScore = score
			minKey = key
		}
	})
	if minKey == -1 {
		panic("WTF")
	}
	return minKey
}

func remove(open []*Node, pos int) []*Node {
	open[pos] = open[len(open)-1]
	return open[:len(open)-1]
}

func reconstructPath(endKey int, cameFrom map[int]int, nodes map[int]*Node) {
	k := endKey
	path := []coords2d.Coords2d{nodes[k].Coords}
	for {
		f, found := cameFrom[k]
		if !found {
			break
		}
		path = append(path, nodes[f].Coords)
		k = f
	}
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Printf("%d,%d => %d\n", path[i].X, path[i].Y, nodes[coordsToKey(path[i])].Value)
	}
}

func solve(nodes map[int]*Node, startKey, endKey int) int {
	start := nodes[startKey]
	end := nodes[endKey]

	open := intset.NewFromSlice([]int{startKey})
	cameFrom := make(map[int]int)
	gScore := make(map[int]int)
	gScore[startKey] = 0
	fScore := make(map[int]int)
	fScore[startKey] = heuristic(start, end)

	for open.Len() != 0 {
		currKey := minFScore(open, fScore)
		curr := nodes[currKey]
		if curr == end {
			// reconstructPath(endKey, cameFrom, nodes)
			return gScore[endKey]
		}
		open.Remove(currKey)
		for _, e := range curr.Edges {
			eKey := coordsToKey(e.Node.Coords)
			tempGScore := gScore[currKey] + e.Cost
			edgeGScore := MAXINT
			if s, found := gScore[eKey]; found {
				edgeGScore = s
			}
			if tempGScore < edgeGScore {
				gScore[eKey] = tempGScore
				fScore[eKey] = tempGScore + heuristic(e.Node, end)
				cameFrom[eKey] = currKey
				open.Add(eKey)
			}
		}
	}
	panic("WTF")
}

func main() {
	data := files.ReadLines(os.Args[1])
	nodes := parseMap(data)
	fmt.Println(solve(
		nodes,
		coordsToKey(coords2d.Coords2d{X: 0, Y: 0}),
		coordsToKey(coords2d.Coords2d{X: len(data) - 1, Y: len(data) - 1}),
	),
	)
}
