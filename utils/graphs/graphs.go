package graphs

import (
	"adventofcode/utils/pqueue"
	"container/heap"
	"errors"
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
	const maxInt = 2147483647

	visited := make(map[K]int)
	prev := make(map[K]K)
	for k := range g.Vertices {
		visited[k] = maxInt
	}
	visited[from] = 0

	var pq pqueue.PriorityQueue[K]
	heap.Init(&pq)
	heap.Push(&pq, &pqueue.Item[K]{Value: from, Priority: 0})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*pqueue.Item[K])
		if item.Value == to {
			path := []K{to}
			for prev[to] != from {
				path = append(path, prev[to])
				to = prev[to]
			}
			path = append(path, from)
			return path, nil
		}
		for _, edge := range g.Edges[item.Value] {
			if visited[edge.ToKey] > visited[edge.FromKey]+edge.Weight {
				visited[edge.ToKey] = visited[edge.FromKey] + edge.Weight
				prev[edge.ToKey] = edge.FromKey
				heap.Push(&pq, &pqueue.Item[K]{Value: edge.ToKey, Priority: visited[edge.ToKey]})
			}
		}
	}
	return nil, PathNotFound
}

func (g *Graph[K, T]) weightedShortestDistance(from K, to K) (int, error) {
	const maxInt = 2147483647

	visited := make(map[K]int)
	for k := range g.Vertices {
		visited[k] = maxInt
	}
	visited[from] = 0

	var pq pqueue.PriorityQueue[K]
	heap.Init(&pq)
	heap.Push(&pq, &pqueue.Item[K]{Value: from, Priority: 0})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*pqueue.Item[K])
		if item.Value == to {
			return visited[item.Value], nil
		}
		for _, edge := range g.Edges[item.Value] {
			if visited[edge.ToKey] > visited[edge.FromKey]+edge.Weight {
				visited[edge.ToKey] = visited[edge.FromKey] + edge.Weight
				heap.Push(&pq, &pqueue.Item[K]{Value: edge.ToKey, Priority: visited[edge.ToKey]})
			}
		}
	}
	return -1, PathNotFound
}
