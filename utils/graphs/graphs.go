package graphs

import (
	"errors"

	"github.com/rdleal/go-priorityq/kpq"
)

var (
	EdgeWeightNotSpecified = errors.New("Edge weight not specified")
	PathNotFound           = errors.New("Path not found")
	VertexNotFound         = errors.New("Vertex not found")
	NotImplemented         = errors.New("Not implemented")
)

type Edge[K comparable, T any] struct {
	From, To       T
	FromKey, ToKey K
	Weight         int
}

type Graph[K comparable, T any] struct {
	Edges    map[K](map[K]Edge[K, T])
	Vertices map[K]T
	Weighted bool
}

func NewGraph[K comparable, T any]() *Graph[K, T] {
	return &Graph[K, T]{
		Edges:    make(map[K](map[K]Edge[K, T])),
		Vertices: make(map[K]T),
		Weighted: false,
	}
}

func NewWeightedGraph[K comparable, T any]() *Graph[K, T] {
	return &Graph[K, T]{
		Edges:    make(map[K](map[K]Edge[K, T])),
		Vertices: make(map[K]T),
		Weighted: true,
	}
}

func (g *Graph[K, T]) AddVertex(key K, value T) {
	g.Vertices[key] = value
}

func (g *Graph[K, T]) GetVertex(k K) (T, error) {
	v, found := g.Vertices[k]
	if !found {
		return v, VertexNotFound
	}
	return v, nil
}

func (g *Graph[K, T]) AddEdge(from K, to K, weight ...int) error {
	if g.Weighted && len(weight) != 1 {
		return EdgeWeightNotSpecified
	}

	fromVertex, err := g.GetVertex(from)
	if err != nil {
		return err
	}
	toVertex, err := g.GetVertex(to)
	if err != nil {
		return err
	}

	if _, found := g.Edges[from]; !found {
		g.Edges[from] = make(map[K]Edge[K, T])
	}
	w := 1
	if g.Weighted {
		w = weight[0]
	}
	g.Edges[from][to] = Edge[K, T]{fromVertex, toVertex, from, to, w}
	return nil
}

func (g *Graph[K, T]) RemoveEdge(from K, to K) error {
	_, err := g.GetVertex(from)
	if err != nil {
		return err
	}
	_, err = g.GetVertex(to)
	if err != nil {
		return err
	}

	delete(g.Edges[from], to)
	return nil
}

func (g *Graph[K, T]) ShortestPath(from K, to K) ([]K, error) {
	if g.Weighted {
		return g.weightedShortestPath(from, to)
	} else {
		return g.unWeightedShortestPath(from, to)
	}
}

func (g *Graph[K, T]) AllShortestPaths(from K, to K) ([][]K, error) {
	if g.Weighted {
		return g.weightedAllShortestPaths(from, to)
	} else {
		return nil, NotImplemented
	}
}

func (g *Graph[K, T]) ShortestDistance(from K, to K) (int, error) {
	if g.Weighted {
		return g.weightedShortestDistance(from, to)
	} else {
		return -1, NotImplemented
	}
}

func (g *Graph[K, T]) unWeightedShortestPath(from K, to K) ([]K, error) {
	type state struct {
		current K
		visited []K
	}
	toVisit := []state{state{from, []K{from}}}
	visited := make(map[K]struct{})
	visited[from] = struct{}{}
	var head state

	for len(toVisit) > 0 {
		head, toVisit = toVisit[0], toVisit[1:]
		if head.current == to {
			return head.visited, nil
		}
		for to := range g.Edges[head.current] {
			if _, found := visited[to]; !found {
				toVisit = append(toVisit, state{to, append(head.visited, to)})
				visited[to] = struct{}{}
			}
		}
	}
	return nil, PathNotFound
}

func (g *Graph[K, T]) weightedShortestPath(from K, to K) ([]K, error) {
	paths, err := g.weightedAllShortestPaths(from, to)
	if err != nil {
		return nil, err
	}
	return paths[0], nil
}

func (g *Graph[K, T]) weightedAllShortestPaths(from K, to K) ([][]K, error) {
	const maxDist = 99999999999999999
	cmp := func(a, b int) bool { return a < b }
	pq := kpq.NewKeyedPriorityQueue[K](cmp)
	prev := make(map[K][]K)

	distances := make(map[K]int)
	for v := range g.Vertices {
		distances[v] = maxDist
		prev[v] = []K{}
		pq.Push(v, maxDist)
	}
	distances[from] = 0
	pq.Update(from, 0)

	for pq.Len() > 0 {
		curr, _, _ := pq.Pop()
		for _, edge := range g.Edges[curr] {
			newDist := distances[curr] + edge.Weight
			if newDist <= distances[edge.ToKey] {
				distances[edge.ToKey] = newDist
				prev[edge.ToKey] = append(prev[edge.ToKey], curr)
				pq.Update(edge.ToKey, newDist)
			}
		}
	}
	if distances[to] == maxDist {
		return nil, PathNotFound
	}
	return backtrackPaths(prev, to), nil
}

type n[K comparable] struct {
	curr K
	path []K
}

func backtrackPaths[K comparable](prev map[K][]K, to K) [][]K {
	queue := []n[K]{{to, []K{to}}}
	paths := make([][]K, 0)
	var curr n[K]
	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		if len(prev[curr.curr]) == 0 {
			paths = append(paths, curr.path)
		}
		for _, e := range prev[curr.curr] {
			np := make([]K, len(curr.path))
			copy(np, curr.path)
			np = append(np, e)
			queue = append(queue, n[K]{e, np})
		}
	}
	return paths
}

func (g *Graph[K, T]) weightedShortestDistance(from K, to K) (int, error) {
	const maxDist = 99999999999999999
	cmp := func(a, b int) bool { return a < b }
	pq := kpq.NewKeyedPriorityQueue[K](cmp)

	distances := make(map[K]int)
	for v := range g.Vertices {
		distances[v] = maxDist
		pq.Push(v, maxDist)
	}
	pq.Update(from, 0)
	distances[from] = 0

	for pq.Len() > 0 {
		curr, dist, _ := pq.Pop()
		if curr == to {
			return dist, nil
		}
		for _, edge := range g.Edges[curr] {
			newDist := distances[curr] + edge.Weight
			if newDist <= distances[edge.ToKey] {
				distances[edge.ToKey] = newDist
				pq.Update(edge.ToKey, newDist)
			}
		}
	}
	return -1, PathNotFound
}
