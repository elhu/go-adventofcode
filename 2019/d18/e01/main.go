package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/sets/byteset"
	"adventofcode/utils/sets/intset"
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type State struct {
	collected *byteset.ByteSet
	distance  int
	currNode  *Node
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    *State // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

// CHANGED IT FROM EXAMPLE
func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func listKeys(lines [][]byte) *byteset.ByteSet {
	keys := byteset.New()
	for _, line := range lines {
		for _, c := range line {
			if c >= 'a' && c <= 'z' {
				keys.Add(c)
			}
		}
	}
	return keys
}

func bsCopy(a *byteset.ByteSet) *byteset.ByteSet {
	b := byteset.New()
	for _, m := range a.Members() {
		b.Add(m)
	}
	return b
}

func solve(lines [][]byte, toCollect *byteset.ByteSet) int {
	graph := buildGraph(lines)
	pq := make(PriorityQueue, 1)
	pq[0] = &Item{
		value:    &State{collected: byteset.New(), distance: 0, currNode: graph['@']},
		priority: 0,
		index:    0,
	}
	heap.Init(&pq)
	for pq.Len() > 0 {
		state := pq.Pop().(*Item).value
		if state.collected.Len() == toCollect.Len() {
			return state.distance
		}

		for key, dist := range reachableKeys(state.currNode, state.collected, 0, make(map[byte]int)) {
			newCollected := bsCopy(state.collected)
			newCollected.Add(key)
			pq.Push(&Item{
				value:    &State{collected: newCollected, distance: state.distance + dist, currNode: graph[key]},
				priority: state.distance + dist,
			})
		}
	}

	return 0
}

func toLower(b byte) byte {
	return bytes.ToLower([]byte{b})[0]
}

func reachableKeys(node *Node, collected *byteset.ByteSet, distance int, visited map[byte]int) map[byte]int {
	keys := make(map[byte]int)
	for _, e := range node.edges {
		if dist, exists := visited[e.node.name]; !exists || dist > distance {
			visited[e.node.name] = distance
			if e.node.kind == key && !collected.HasMember(e.node.name) {
				keys[e.node.name] = distance + e.distance
			}
			if e.node.kind == door && collected.HasMember(toLower(e.node.name)) {
				for key, dist := range reachableKeys(e.node, collected, distance+e.distance, visited) {
					keys[key] = dist
				}
			}
		}
	}

	return keys
}

type Edge struct {
	distance int
	node     *Node
}

type Node struct {
	kind  int
	name  byte
	edges map[byte]Edge
	pos   coords2d.Coords2d
}

const (
	start = iota
	key   = iota
	door  = iota
)

func buildGraph(lines [][]byte) map[byte]*Node {
	nodes := make(map[byte]*Node)
	for y, line := range lines {
		for x, c := range line {
			if c == '@' {
				nodes[c] = &Node{kind: start, name: c, edges: make(map[byte]Edge), pos: coords2d.Coords2d{X: x, Y: y}}
			} else if c >= 'a' && c <= 'z' {
				nodes[c] = &Node{kind: key, name: c, edges: make(map[byte]Edge), pos: coords2d.Coords2d{X: x, Y: y}}

			} else if c >= 'A' && c <= 'Z' {
				nodes[c] = &Node{kind: door, name: c, edges: make(map[byte]Edge), pos: coords2d.Coords2d{X: x, Y: y}}
			}
		}
	}
	for _, node := range nodes {
		buildEdges(nodes, lines, node)
	}
	return nodes
}

func printGraph(graph map[byte]*Node) {
	for name, node := range graph {
		var strEdges []string
		for _, e := range node.edges {
			strEdges = append(strEdges, fmt.Sprintf("%c(%d)", e.node.name, e.distance))
		}
		fmt.Printf("%c => %s\n", name, strings.Join(strEdges, "|"))
	}
}

func coordsToKey(c coords2d.Coords2d) int {
	return c.Y*10000 + c.X
}

func coordsFromKey(i int) coords2d.Coords2d {
	return coords2d.Coords2d{X: i % 10000, Y: i / 10000}
}

func adjacentPos(lines [][]byte, pos coords2d.Coords2d) []coords2d.Coords2d {
	var adj []coords2d.Coords2d
	if pos.X > 0 && lines[pos.Y][pos.X-1] != '#' {
		adj = append(adj, coords2d.Coords2d{X: pos.X - 1, Y: pos.Y})
	}
	if pos.X < len(lines[pos.Y])-1 && lines[pos.Y][pos.X+1] != '#' {
		adj = append(adj, coords2d.Coords2d{X: pos.X + 1, Y: pos.Y})
	}
	if pos.Y > 0 && lines[pos.Y-1][pos.X] != '#' {
		adj = append(adj, coords2d.Coords2d{X: pos.X, Y: pos.Y - 1})
	}
	if pos.Y < len(lines)-1 && lines[pos.Y+1][pos.X] != '#' {
		adj = append(adj, coords2d.Coords2d{X: pos.X, Y: pos.Y + 1})
	}
	var n []byte
	for _, a := range adj {
		n = append(n, lines[a.Y][a.X])
	}
	return adj
}

func buildEdges(nodes map[byte]*Node, lines [][]byte, startNode *Node) {
	visited := intset.New()
	visited.Add(coordsToKey(startNode.pos))

	queue := adjacentPos(lines, startNode.pos)
	var curr coords2d.Coords2d
	var newQueue []coords2d.Coords2d
	distance := 1

	for len(queue) != 0 {
		curr, queue = queue[0], queue[1:]
		// fmt.Printf("visiting %v\n", curr)
		key := coordsToKey(curr)
		if !visited.HasMember(key) {
			visited.Add(key)
			if lines[curr.Y][curr.X] == '.' {
				// fmt.Printf("Empty - enqueue next move (%v)\n", adjacentPos(lines, curr))
				newQueue = append(newQueue, adjacentPos(lines, curr)...)
			} else if lines[curr.Y][curr.X] == '@' {
				// fmt.Println("Adding entrance edge")
				startNode.edges[lines[curr.Y][curr.X]] = Edge{distance: distance, node: nodes[lines[curr.Y][curr.X]]}
				newQueue = append(newQueue, adjacentPos(lines, curr)...)
			} else if lines[curr.Y][curr.X] >= 'a' && lines[curr.Y][curr.X] <= 'z' {
				// fmt.Println("Adding key edge")
				startNode.edges[lines[curr.Y][curr.X]] = Edge{distance: distance, node: nodes[lines[curr.Y][curr.X]]}
				newQueue = append(newQueue, adjacentPos(lines, curr)...)
			} else if lines[curr.Y][curr.X] >= 'A' && lines[curr.Y][curr.X] <= 'Z' {
				// fmt.Println("Adding door edge")
				startNode.edges[lines[curr.Y][curr.X]] = Edge{distance: distance, node: nodes[lines[curr.Y][curr.X]]}
				// } else {
				// 	panic("WTF")
			}
		}

		if len(queue) == 0 {
			// fmt.Println("Flipping queues")
			queue = make([]coords2d.Coords2d, len(newQueue))
			copy(queue, newQueue)
			newQueue = nil
			distance++
		}
	}
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	check(err)
	lines := bytes.Split(data, []byte("\n"))
	toCollect := listKeys(lines)
	fmt.Println(solve(lines, toCollect))
}
