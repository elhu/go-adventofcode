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

func allowedDestination(pod Amphipod, to *Node) bool {
	// If destination is a hallway, it's always allowed
	if to.Kind == HALLWAY {
		return true
	}
	// If pod was in the wrong room, it can move if the target is the right kind of room
	if pod.Position.Kind == DESTINATION && pod.Position.Allowed != pod.Kind {
		return to.Allowed == pod.Kind
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

func podAt(pos string, state State) (Amphipod, bool) {
	for _, p := range state {
		if p.Position.Code == pos {
			return p, true
		}
	}
	return state[0], false
}

var rows = []string{"RA", "RB", "RC", "RD"}

func containsWrongPods(node *Node, state State) bool {
	if node.Kind == HALLWAY {
		return false
	}
	col := node.Code[2]
	for _, r := range rows {
		if a, b := podAt(fmt.Sprintf("%s%c", col, r), state); b && a.Kind != node.Allowed {
			return true
		}
	}
	return false
}

func isOptimallyPlaced(p Amphipod, state State) bool {
	if p.Position.Kind == HALLWAY || p.Position.Allowed != p.Kind {
		return false
	}
	// If anything below is not in the right place, allow move
	row, col := p.Position.Code[0:2], p.Position.Code[2]
	found := false
	for _, r := range rows {
		if found == true {
			if a, b := podAt(fmt.Sprintf("%s%c", r, col), state); b && a.Kind != p.Kind {
				return false
			}
		}
		if r == row {
			found = true
		}
	}
	return true
}

func potentialMoves(state State, nodes map[string]*Node) []MoveCost {
	var res []MoveCost
	var newState State
	for i, p := range state {
		// If the pod is in the right room and there are no wrong pods below it
		// don't move
		if isOptimallyPlaced(p, state) {
			continue
		}
		for _, n := range p.Position.Neighbors {
			// Check if destination is legit (moving to hallway or the right room)
			if allowedDestination(p, n.Node) {
				// Check if destination isn't already occupied
				if isFree(n.Node, state) {
					// Check if there are no pods between source and dest
					if isReachable(n.Path, state) {
						// Check if destination doesn't contain pods that don't belong to that room
						if !containsWrongPods(n.Node, state) {
							newState = state.CopyState()
							newState[i].Position = n.Node
							res = append(res, MoveCost{Cost: n.Distance * p.Cost, Move: newState})
						}
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
		currentStateKey := currentState.Serialize()
		currentCost := seenStates[currentStateKey]

		if currentStateKey == targetState {
			fmt.Println("Found Solution!")
			break
		}

		for _, s := range potentialMoves(currentState, nodes) {
			moveKey := s.Move.Serialize()
			totalCost := currentCost + s.Cost
			if knownCost, seen := seenStates[moveKey]; !seen || totalCost < knownCost {
				from[moveKey] = currentStateKey
				seenStates[moveKey] = totalCost
				heap.Push(&openStates, MoveCost{Cost: totalCost, Move: s.Move})
			}
		}
	}
	return seenStates[targetState]
}

// 46108 low

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
