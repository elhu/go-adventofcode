package main

import (
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

const MAXINT = 2147483647

func toKey(x, y int) int {
	return y*1000000 + x
}

func fromKey(k int) (int, int) {
	y := k / 1000000
	return k - y*1000000, y
}

func solve(data [][]int) int {
	var open KeyHeap
	startKey := toKey(0, 0)
	heap.Init(&open)
	heap.Push(&open, [2]int{startKey, data[0][0]})
	gScore := make(map[int]int)
	gScore[startKey] = 0

	for open.Len() != 0 {
		currKey := heap.Pop(&open).([2]int)[0]
		x, y := fromKey(currKey)
		evaluatedNodes++
		if x == len(data)-1 && y == len(data)-1 {
			return gScore[currKey]
		}
		neighbors := [][2]int{{x, y - 1}, {x, y + 1}, {x - 1, y}, {x + 1, y}}

		for _, n := range neighbors {
			i, j := n[0], n[1]
			if i < 0 || i >= len(data) || j < 0 || j >= len(data) {
				continue
			}
			eKey := toKey(i, j)
			tempGScore := gScore[currKey] + data[i][j]
			edgeGScore := MAXINT
			if s, found := gScore[eKey]; found {
				edgeGScore = s
			}
			if tempGScore < edgeGScore {
				gScore[eKey] = tempGScore
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
	expandedMap := expandMap(data)
	fmt.Println(solve(expandedMap))
	fmt.Printf("Took: %s, visited %d nodes", time.Since(start), evaluatedNodes)
}
