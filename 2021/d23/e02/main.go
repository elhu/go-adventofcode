package main

import (
	"adventofcode/utils/files"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type KeyHeap []MoveCost

func (h KeyHeap) Len() int {
	return len(h)
}

func (h KeyHeap) Less(i, j int) bool {
	return h[i].Cost < h[j].Cost
}

func (h KeyHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *KeyHeap) Push(x interface{}) {
	*h = append(*h, x.(MoveCost))
}

func (h *KeyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Edge struct {
	Node     *Node
	Distance int
	Path     []string
}

const (
	HALLWAY     = iota
	DESTINATION = iota
)

type Node struct {
	Kind      int // HALLWAY|DESTINATION
	Neighbors []*Edge
	Code      string
	Allowed   byte
}

type Amphipod struct {
	Kind     byte
	Cost     int
	Position *Node
}

type State []Amphipod

func connect(a, b *Node, path []string) {
	a.Neighbors = append(a.Neighbors, &Edge{Node: b, Distance: len(path) + 1, Path: path})
	b.Neighbors = append(b.Neighbors, &Edge{Node: a, Distance: len(path) + 1, Path: path})
}

//   #############
// H #...........#
// RA###B#C#B#D###
// RB  #A#D#C#A#
//     #########
func createGraph() map[string]*Node {
	nodes := make(map[string]*Node)
	// H (ignore the very end positions)
	hallways := make([]*Node, 0, 10)
	for i := 0; i <= 10; i++ {
		newNode := &Node{Kind: HALLWAY, Code: fmt.Sprintf("H%d", i)}
		hallways = append(hallways, newNode)
		nodes[newNode.Code] = newNode
	}
	// RA and RB
	var allowedMapping = map[int]byte{2: 'A', 4: 'B', 6: 'C', 8: 'D'}
	for i := 2; i < 10; i += 2 {
		newNodeA := &Node{Kind: DESTINATION, Code: fmt.Sprintf("RA%d", i), Allowed: allowedMapping[i]}
		newNodeB := &Node{Kind: DESTINATION, Code: fmt.Sprintf("RB%d", i), Allowed: allowedMapping[i]}
		newNodeC := &Node{Kind: DESTINATION, Code: fmt.Sprintf("RC%d", i), Allowed: allowedMapping[i]}
		newNodeD := &Node{Kind: DESTINATION, Code: fmt.Sprintf("RD%d", i), Allowed: allowedMapping[i]}
		nodes[newNodeA.Code] = newNodeA
		nodes[newNodeB.Code] = newNodeB
		nodes[newNodeC.Code] = newNodeC
		nodes[newNodeD.Code] = newNodeD
	}
	for _, move := range files.ReadLines("legalmoves.txt") {
		parts := strings.Split(move, " ")
		source := nodes[parts[0]]
		destination := nodes[parts[1]]
		path := parts[2:]
		connect(source, destination, path)
	}
	return nodes
}

func placeAmphipods(nodes map[string]*Node, data []string) []Amphipod {
	pods := make([]Amphipod, 0)
	var costs = map[byte]int{'A': 1, 'B': 10, 'C': 100, 'D': 1000}
	var depths = map[int]byte{2: 'A', 3: 'B', 4: 'C', 5: 'D'}
	for i := 3; i < 10; i += 2 {
		for j := 2; j < 6; j++ {
			a := Amphipod{Kind: data[j][i], Cost: costs[data[j][i]]}
			n := nodes[fmt.Sprintf("R%c%d", depths[j], i-1)]
			a.Position = n
			// n.Occupant = a
			pods = append(pods, a)
		}
	}
	return pods
}

func (s State) CopyState() State {
	newState := make(State, len(s))
	for i, a := range s {
		newState[i] = a
	}
	return newState
}

func (s State) Serialize() string {
	var res []string
	for _, a := range s {
		res = append(res, fmt.Sprintf("%s:%c", a.Position.Code, a.Kind))
	}
	sort.Sort(sort.StringSlice(res))
	return strings.Join(res, "|")
}

type MoveCost struct {
	Cost int
	Move State
}

func isFree(node *Node, state State) bool {
	for _, p := range state {
		if p.Position == node {
			return false
		}
	}
	return true
}

// func otherPodsStopped(state State) bool {
// 	for _, p := range state {
// 		if p.Position.Kind == HALLWAY && !p.Moving {
// 			return true
// 		}
// 	}
// 	return false
// }

func allowedDestination(pod Amphipod, to *Node) bool {
	// If destination is a hallway, it's always allowed
	if to.Kind == HALLWAY {
		return true
	}
	// If pod was in the wrong room, it can always move
	if pod.Position.Kind == DESTINATION && pod.Position.Allowed != pod.Kind {
		return true
	}
	// If pod was in hallway and enters its destination room, it's allowed
	return pod.Position.Kind == HALLWAY && pod.Kind == to.Allowed
}

func isReachable(path []string, state State) bool {
	for _, n := range path {
		for _, p := range state {
			if p.Position.Code == n {
				return false
			}
		}
	}
	return true
}

func potentialMoves(state State, nodes map[string]*Node) []MoveCost {
	var res []MoveCost
	var newState State
	// if a pod is moving, only this one can move
	for i, p := range state {
		for _, n := range p.Position.Neighbors {
			if isFree(n.Node, state) && isReachable(n.Path, state) {
				// Append new position with pod stopped if it's a destination or no other pod is stopped on the hallway
				if allowedDestination(p, n.Node) {
					newState = state.CopyState()
					newState[i].Position = n.Node
					res = append(res, MoveCost{Cost: n.Distance * p.Cost, Move: newState})
				}
			}
		}
	}

	return res
}

func solve(nodes map[string]*Node, pods []Amphipod, targetState string) int {
	var openStates KeyHeap
	heap.Init(&openStates)
	heap.Push(&openStates, MoveCost{Cost: 0, Move: pods})

	seenStates := make(map[string]int)
	from := make(map[string]string)
	var currentState State
	seenStates[State(pods).Serialize()] = 0

	for len(openStates) > 0 {
		currentMove := heap.Pop(&openStates).(MoveCost)
		currentState = currentMove.Move
		currentCost := currentMove.Cost
		currentStateKey := currentState.Serialize()

		if currentStateKey == targetState {
			fmt.Println("Found Solution!")
			break
		}

		for _, s := range potentialMoves(currentState, nodes) {
			moveKey := s.Move.Serialize()
			if knownCost, seen := seenStates[moveKey]; !seen || currentCost+s.Cost < knownCost {
				from[moveKey] = currentStateKey
				seenStates[moveKey] = currentCost + s.Cost
				heap.Push(&openStates, MoveCost{Cost: currentCost + s.Cost, Move: s.Move})
			}
		}
	}
	return seenStates[targetState]
}

// 12521 low
// 14354 high

func main() {
	s := time.Now()
	data := files.ReadLines(os.Args[1])
	nodes := createGraph()
	pods := placeAmphipods(nodes, data)
	target := placeAmphipods(nodes, files.ReadLines("endstate.txt"))
	fmt.Printf("Initial state: %s\n", State(pods).Serialize())
	fmt.Printf(" Target state: %s\n", State(target).Serialize())
	fmt.Println(solve(nodes, pods, State(target).Serialize()))
	fmt.Println(time.Since(s))
}
