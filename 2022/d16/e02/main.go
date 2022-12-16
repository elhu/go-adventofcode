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
	visited       int
	lastVisited   string
	released      int
	timeRemaining int
}

func allPairs(list []string) [][2]string {
	pairs := make([][2]string, 0)
	for i, m := range list {
		for j, n := range list {
			if i != j {
				pairs = append(pairs, [2]string{m, n})
			}
		}
	}
	return pairs
}

func bitPos(nodes map[string]map[string]int) map[string]uint {
	positions := make(map[string]uint)
	pos := uint(1)
	for k := range nodes {
		positions[k] = pos
		pos++
	}
	positions["AA"] = 0
	return positions
}

func setBit(n int, pos uint) int {
	n |= (1 << pos)
	return n
}

func hasBit(n int, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}

func countSetBits(n int) int {
	count := 0
	for n != 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

func solve(nodes map[string]*Node) int {
	distances := buildDistanceMatrix(nodes)
	bp := bitPos(distances)
	visited := setBit(0, bp["AA"])

	queue := []QueueItem{{visited: visited, lastVisited: "AA", released: 0, timeRemaining: 26}}
	bestReleased := make(map[int]int)

	var head QueueItem
	for len(queue) > 0 {
		head, queue = queue[0], queue[1:]
		for valve, distance := range distances[head.lastVisited] {
			timeRemaining := head.timeRemaining - distance
			if !hasBit(head.visited, bp[valve]) && timeRemaining > 0 {
				visited = setBit(head.visited, bp[valve])
				released := head.released + timeRemaining*nodes[valve].value
				if bestReleased[visited] < released {
					bestReleased[visited] = released
				}
				queue = append(queue, QueueItem{visited: visited, lastVisited: valve, released: released, timeRemaining: timeRemaining})
			}
		}
	}

	max := 0
	for k1, v1 := range bestReleased {
		for k2, v2 := range bestReleased {
			if k1&k2 == 1 {
				if v1+v2 > max {
					max = v1 + v2
				}
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
