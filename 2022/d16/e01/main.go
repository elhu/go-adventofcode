package main

import (
	"adventofcode/utils/files"
	"adventofcode/utils/sets/stringset"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	name      string
	value     int
	neighbors map[string]*Node
}

func parseGraph(data []string) map[string]*Node {
	nodes := make(map[string]*Node)
	for _, line := range data {
		node := &Node{neighbors: make(map[string]*Node)}
		var name string
		var value int
		fmt.Sscanf(line, "Valve %s has flow rate=%d", &name, &value)
		if n, exists := nodes[name]; exists {
			node = n
		}
		node.name = name
		node.value = value
		startIdx := strings.Index(line, "to valves ")
		if startIdx == -1 {
			startIdx = strings.Index(line, "to valve ") + len("to valve ")
		} else {
			startIdx += len("to valves ")
		}
		for _, neighborName := range strings.Split(line[startIdx:], ", ") {
			var neighbor *Node
			var exists bool
			if neighbor, exists = nodes[neighborName]; !exists {
				neighbor = &Node{name: neighborName, neighbors: make(map[string]*Node)}
				nodes[neighborName] = neighbor
			}
			neighbor.neighbors[node.name] = node
			node.neighbors[neighbor.name] = neighbor
		}

		nodes[node.name] = node
	}
	return nodes
}

func distance(n, m *Node) int {
	queue := []*Node{n}
	nextQueue := []*Node{}
	var head *Node
	distance := 0
	visited := stringset.New()
	visited.Add(n.name)

	for len(queue) > 0 {
		head, queue = queue[0], queue[1:]
		if head == m {
			return distance
		}

		for _, nb := range head.neighbors {
			if !visited.HasMember(nb.name) {
				visited.Add(nb.name)
				nextQueue = append(nextQueue, nb)
			}
		}
		if len(queue) == 0 {
			distance++
			queue = make([]*Node, len(nextQueue))
			copy(queue, nextQueue)
			nextQueue = nextQueue[:0]
		}
	}
	panic("WTF")
}

func buildDistanceMatrix(nodes map[string]*Node) map[string]map[string]int {
	distances := make(map[string]map[string]int)
	for i, n := range nodes {
		if n.value > 0 || n.name == "AA" {
			distances[i] = make(map[string]int)
			for j, m := range nodes {
				if m.value > 0 {
					if dist := distance(n, m); dist != 0 {
						// Adding 1 to account for the time needed to open the valve
						distances[i][j] = dist + 1
					}
				}
			}
		}
	}
	return distances
}

type QueueItem struct {
	visited       *stringset.StringSet
	lastVisited   string
	released      int
	timeRemaining int
}

func solve(nodes map[string]*Node) int {
	distances := buildDistanceMatrix(nodes)
	queue := []QueueItem{{visited: stringset.NewFromSlice([]string{"AA"}), lastVisited: "AA", released: 0, timeRemaining: 30}}
	max := 0
	var head QueueItem
	for len(queue) > 0 {
		head, queue = queue[0], queue[1:]
		for valve, distance := range distances[head.lastVisited] {
			timeRemaining := head.timeRemaining - distance
			if !head.visited.HasMember(valve) && timeRemaining > 0 {
				visited := stringset.NewFromSlice(head.visited.Members())
				visited.Add(valve)
				released := head.released + timeRemaining*nodes[valve].value
				if released > max {
					max = released
				}
				queue = append(queue, QueueItem{visited: visited, lastVisited: valve, released: released, timeRemaining: timeRemaining})
			}
		}
	}
	return max
}

func main() {
	data := files.ReadLines(os.Args[1])
	nodes := parseGraph(data)
	maxOpen := 0
	for _, n := range nodes {
		if n.value > 0 {
			maxOpen++
		}
	}
	fmt.Println(solve(nodes))
}
