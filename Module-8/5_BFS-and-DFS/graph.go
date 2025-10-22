package main

import "sort"

type Vertex struct {
	key string
}
type Graph struct {
	Vertices      map[string]*Vertex
	AdjacencyList map[string][]*Vertex
	Directed      bool
}

func (g *Graph) AddVertex(key string) {
	if _, exists := g.Vertices[key]; !exists {
		g.Vertices[key] = &Vertex{key: key}
		g.AdjacencyList[key] = []*Vertex{}
	}
}

func (g *Graph) AddEdge(from, to string) {
	g.AddVertex(from)
	g.AddVertex(to)
	g.AdjacencyList[from] = append(g.AdjacencyList[from], g.Vertices[to])

	// if undirectional, it needs to go be in both arrays
	if !g.Directed {
		g.AdjacencyList[to] = append(g.AdjacencyList[to], g.Vertices[from])
	}
}

func NewGraph(directed bool) Graph {
	return Graph{
		Vertices:      map[string]*Vertex{},
		AdjacencyList: map[string][]*Vertex{},
		Directed:      directed,
	}
}

// GetVerticesSorted returns vertices in alphabetical order cus the exercise assumes they are sorted this way so we kinda have to do this to check if the result matches.
// due to the nature of maps, we cant just sort them in the graph, we have to return them as a slice when we want to iterate over them
func (g *Graph) GetVerticesSorted() []*Vertex {
	keys := make([]string, 0, len(g.Vertices))
	for key := range g.Vertices {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	vertices := make([]*Vertex, len(keys))
	for i, key := range keys {
		vertices[i] = g.Vertices[key]
	}
	return vertices
}

// SortAdjacencyLists sorts all adjacency lists alphabetically. As these are slices, we can just sort them in the graph with no problem :pray:
func (g *Graph) SortAdjacencyLists() {
	for key := range g.AdjacencyList {
		sort.Slice(g.AdjacencyList[key], func(i, j int) bool {
			return g.AdjacencyList[key][i].key < g.AdjacencyList[key][j].key
		})
	}
}
