package graphs

import "errors"

var (
	VertexNotFound = errors.New("Vertex not found")
)

type Edge[K comparable, T any] struct {
	From, To       T
	FromKey, ToKey K
}

type Graph[K comparable, T any] struct {
	Edges    map[K](map[K]Edge[K, T])
	Vertices map[K]T
}

func NewGraph[K comparable, T any]() *Graph[K, T] {
	return &Graph[K, T]{
		Edges:    make(map[K](map[K]Edge[K, T])),
		Vertices: make(map[K]T),
	}
}

func (g *Graph[K, T]) AddVertex(key K, value T) {
	g.Vertices[key] = value
}

func (g *Graph[K, T]) getVertex(k K) (T, error) {
	v, found := g.Vertices[k]
	if !found {
		return v, VertexNotFound
	}
	return v, nil
}

func (g *Graph[K, T]) AddEdge(from K, to K) error {
	fromVertex, err := g.getVertex(from)
	if err != nil {
		return err
	}
	toVertex, err := g.getVertex(to)
	if err != nil {
		return err
	}

	if _, found := g.Edges[from]; !found {
		g.Edges[from] = make(map[K]Edge[K, T])
	}
	g.Edges[from][to] = Edge[K, T]{fromVertex, toVertex, from, to}
	return nil
}

func (g *Graph[K, T]) RemoveEdge(from K, to K) error {
	_, err := g.getVertex(from)
	if err != nil {
		return err
	}
	_, err = g.getVertex(to)
	if err != nil {
		return err
	}

	delete(g.Edges[from], to)
	return nil
}

func (g *Graph[K, T]) ShortestPath(from K, to K) ([]K, error) {
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
	return nil, errors.New("No path found")
}
