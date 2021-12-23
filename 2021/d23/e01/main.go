package main

import (
	"adventofcode/utils/files"
<<<<<<< HEAD
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

=======
	"fmt"
	"os"
)

>>>>>>> 0092f18 (Limit number of legal moves)
type Edge struct {
	Node     *Node
	Distance int
}

const (
	HALLWAY     = iota
	DESTINATION = iota
)

type Node struct {
	Kind      int // HALLWAY|DESTINATION
	Neighbors []*Edge
	Code      string
<<<<<<< HEAD
	Allowed   byte
=======
>>>>>>> 0092f18 (Limit number of legal moves)
	// Occupant  *Amphipod
}

type Amphipod struct {
	Kind     byte
	Cost     int
	Position *Node
	Moving   bool
}

<<<<<<< HEAD
// type State struct {
// 	Pods []Amphipod
// 	From []State
// 	Cost int
// }

=======
>>>>>>> 0092f18 (Limit number of legal moves)
type State []Amphipod

func connect(a, b *Node, dist int) {
	a.Neighbors = append(a.Neighbors, &Edge{Node: b, Distance: dist})
	b.Neighbors = append(b.Neighbors, &Edge{Node: a, Distance: dist})
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
<<<<<<< HEAD
	for i := 0; i <= 10; i++ {
=======
	for i := 1; i < 10; i++ {
>>>>>>> 0092f18 (Limit number of legal moves)
		newNode := &Node{Kind: HALLWAY, Code: fmt.Sprintf("H%d", i)}
		hallways = append(hallways, newNode)
		nodes[newNode.Code] = newNode
	}
	// RA and RB
	ra := make([]*Node, 0, 4)
	rb := make([]*Node, 0, 4)
<<<<<<< HEAD
	var allowedMapping = map[int]byte{2: 'A', 4: 'B', 6: 'C', 8: 'D'}
	for i := 2; i < 10; i += 2 {
		newNodeA := &Node{Kind: DESTINATION, Code: fmt.Sprintf("RA%d", i), Allowed: allowedMapping[i]}
		newNodeB := &Node{Kind: DESTINATION, Code: fmt.Sprintf("RB%d", i), Allowed: allowedMapping[i]}
=======
	for i := 2; i < 10; i += 2 {
		newNodeA := &Node{Kind: DESTINATION, Code: fmt.Sprintf("RA%d", i)}
		newNodeB := &Node{Kind: DESTINATION, Code: fmt.Sprintf("RB%d", i)}
>>>>>>> 0092f18 (Limit number of legal moves)
		ra = append(ra, newNodeA)
		rb = append(rb, newNodeB)
		nodes[newNodeA.Code] = newNodeA
		nodes[newNodeB.Code] = newNodeB
	}
	// Connect it all
	connect(nodes["RB2"], nodes["RA2"], 1)
	connect(nodes["RB4"], nodes["RA4"], 1)
	connect(nodes["RB6"], nodes["RA6"], 1)
	connect(nodes["RB8"], nodes["RA8"], 1)
<<<<<<< HEAD

	connect(nodes["H0"], nodes["H1"], 1)
=======
>>>>>>> 0092f18 (Limit number of legal moves)
	connect(nodes["H1"], nodes["RA2"], 2)
	connect(nodes["H1"], nodes["H3"], 2)
	connect(nodes["H3"], nodes["RA2"], 2)
	connect(nodes["H3"], nodes["RA4"], 2)
	connect(nodes["H3"], nodes["H5"], 2)
	connect(nodes["H5"], nodes["RA4"], 2)
	connect(nodes["H5"], nodes["RA6"], 2)
	connect(nodes["H5"], nodes["H7"], 2)
	connect(nodes["H7"], nodes["RA6"], 2)
	connect(nodes["H7"], nodes["RA8"], 2)
	connect(nodes["H7"], nodes["H9"], 2)
	connect(nodes["H9"], nodes["RA8"], 2)
<<<<<<< HEAD
	connect(nodes["H9"], nodes["H10"], 1)
=======
>>>>>>> 0092f18 (Limit number of legal moves)

	return nodes
}

func placeAmphipods(nodes map[string]*Node, data []string) []Amphipod {
	pods := make([]Amphipod, 0)
	var costs = map[byte]int{'A': 1, 'B': 10, 'C': 100, 'D': 1000}
	for i := 3; i < 10; i += 2 {
		for j := 2; j < 4; j++ {
<<<<<<< HEAD
			a := Amphipod{Kind: data[j][i], Cost: costs[data[j][i]]}
=======
			a := Amphipod{Kind: data[j][i], Cost: costs[data[j][i+1]]}
>>>>>>> 0092f18 (Limit number of legal moves)
			prefix := "RB"
			if j == 2 {
				prefix = "RA"
			}
			n := nodes[fmt.Sprintf("%s%d", prefix, i-1)]
			a.Position = n
			// n.Occupant = a
			pods = append(pods, a)
		}
	}
	return pods
}

<<<<<<< HEAD
func (s State) CopyState() State {
=======
func copyState(s State) State {
>>>>>>> 0092f18 (Limit number of legal moves)
	newState := make(State, len(s))
	for i, a := range s {
		newState[i] = a
	}
	return newState
}

<<<<<<< HEAD
func (s State) Serialize() string {
	var res []string
	for _, a := range s {
		res = append(res, fmt.Sprintf("%c%s%v", a.Kind, a.Position.Code, a.Moving))
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

func otherPodsStopped(state State) bool {
	for _, p := range state {
		if p.Position.Kind == HALLWAY && !p.Moving {
			return true
		}
	}
	return false
}

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

func potentialMoves(state State, nodes map[string]*Node) []MoveCost {
	var res []MoveCost
	var newState State
	// if a pod is moving, only this one can move
	for i, p := range state {
		if p.Moving {
			for _, n := range p.Position.Neighbors {
				if isFree(n.Node, state) {
					// Append new position with pod stopped if it's a destination or no other pod is stopped on the hallway
					if (n.Node.Kind == DESTINATION && allowedDestination(p, n.Node)) || !otherPodsStopped(state) {
						newState = state.CopyState()
						newState[i].Position = n.Node
						newState[i].Moving = false
						res = append(res, MoveCost{Cost: n.Distance * p.Cost, Move: newState})
					}
					// If new position is hallway, also append new state with pod moving
					if n.Node.Kind == HALLWAY {
						newState = state.CopyState()
						newState[i].Position = n.Node
						newState[i].Moving = true
						res = append(res, MoveCost{Cost: n.Distance * p.Cost, Move: newState})
					}
				}
			}
			return res
		}
	}
	// if no pod is moving, move them all in all directions
	for i, p := range state {
		for _, n := range p.Position.Neighbors {
			if isFree(n.Node, state) {
				if !allowedDestination(p, p.Position) || allowedDestination(p, n.Node) {
					newState = state.CopyState()
					newState[i].Position = n.Node
					newState[i].Moving = false
					res = append(res, MoveCost{Cost: n.Distance * p.Cost, Move: newState})
					if n.Node.Kind == HALLWAY {
						newState = state.CopyState()
						newState[i].Position = n.Node
						newState[i].Moving = true
						res = append(res, MoveCost{Cost: n.Distance * p.Cost, Move: newState})
					}
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
=======
func solve(nodes map[string]*Node, pods []Amphipod) int {
	openStates := []State{pods}
	var currentState State
	for len(openStates) > 0 {
		currentState, openStates = openStates[0], openStates[1:]
		ostate := copyState(currentState)
		currentState[0].Kind = 'Z'
		for i, c := range currentState {
			fmt.Printf("%d %c %s\n", i, c.Kind, c.Position.Code)
		}
		for i, c := range ostate {
			fmt.Printf("%d %c %s\n", i, c.Kind, c.Position.Code)
		}
	}
	return 0
}

func main() {
	data := files.ReadLines(os.Args[1])
	nodes := createGraph()
	pods := placeAmphipods(nodes, data)
	fmt.Println(solve(nodes, pods))
>>>>>>> 0092f18 (Limit number of legal moves)
}
