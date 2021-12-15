package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"container/heap"
	"fmt"
	"os"
	"time"
)

type KeyHeap [][2]int

func (h KeyHeap) Len() int {
	return len(h)
}

func (h KeyHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h KeyHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *KeyHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *KeyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Edge struct {
	Cost int
	Node *Node
}

type Node struct {
	Coords coords2d.Coords2d
	Value  int
	Key    int
	Edges  []*Edge
}

func coordsToKey(c coords2d.Coords2d) int {
	return c.Y*10000000 + c.X
}

func expandMap(data []string) [][]int {
	res := make([][]int, len(data)*5)
	for i := 0; i < len(data)*5; i++ {
		res[i] = make([]int, len(data)*5)
	}
	for yRep := 0; yRep < 5; yRep++ {
		for xRep := 0; xRep < 5; xRep++ {
			for i, line := range data {
				for j, cell := range line {
					val := int(cell-'0') + yRep + xRep
					if val > 9 {
						val -= 9
					}
					res[yRep*len(data)+i][xRep*len(data)+j] = val
				}
			}
		}
	}
	return res
}

func parseMap(data []string) map[int]*Node {
	expandedMap := expandMap(data)
	nodes := make(map[int]*Node)
	for i, l := range expandedMap {
		for j, c := range l {
			coord := coords2d.Coords2d{X: j, Y: i}
			key := coordsToKey(coord)
			nodes[key] = &Node{Coords: coord, Value: c, Key: key}
		}
	}
	for i, l := range expandedMap {
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
					node.Edges = append(node.Edges, &Edge{Cost: c, Node: currentNode})
				}
			}
		}
	}
	return nodes
}

const MAXINT = 2147483647

func solve(nodes map[int]*Node, startKey, endKey int) int {
	start := nodes[startKey]
	end := nodes[endKey]

	var open KeyHeap
	heap.Init(&open)
	heap.Push(&open, [2]int{startKey, start.Value})
	cameFrom := make(map[int]int)
	gScore := make(map[int]int)
	gScore[startKey] = 0

	for open.Len() != 0 {
		currKey := heap.Pop(&open).([2]int)[0]
		curr := nodes[currKey]
		evaluatedNodes++
		if curr == end {
			return gScore[endKey]
		}
		for _, e := range curr.Edges {
			eKey := e.Node.Key
			tempGScore := gScore[currKey] + e.Cost
			edgeGScore := MAXINT
			if s, found := gScore[eKey]; found {
				edgeGScore = s
			}
			if tempGScore < edgeGScore {
				gScore[eKey] = tempGScore
				cameFrom[eKey] = currKey
				heap.Push(&open, [2]int{eKey, tempGScore})
			}
		}
	}
	panic("WTF")
}

var evaluatedNodes = 0

func main() {
	start := time.Now()
	data := files.ReadLines(os.Args[1])
	nodes := parseMap(data)
	fmt.Printf("Map expansion & parsing took %s\n", time.Since(start))
	fmt.Println(solve(
		nodes,
		coordsToKey(coords2d.Coords2d{X: 0, Y: 0}),
		coordsToKey(coords2d.Coords2d{X: len(data)*5 - 1, Y: len(data)*5 - 1}),
	),
	)
	fmt.Printf("Took: %s, visited %d nodes", time.Since(start), evaluatedNodes)
}
