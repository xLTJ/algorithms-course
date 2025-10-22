package main

import "fmt"

type VertexColor int

const (
	WHITE VertexColor = iota
	GRAY
	BLACK
)

type SearchState struct {
	color         map[*Vertex]VertexColor
	parent        map[*Vertex]*Vertex
	distance      map[*Vertex]int // bfs
	currentTime   int             // dfs, keep track of time
	discoveryTime map[*Vertex]int // dfs
	finishTime    map[*Vertex]int // dfs
}

func newSearchState(graph *Graph) SearchState {
	verticesAmount := len(graph.Vertices)
	newState := SearchState{
		color:         make(map[*Vertex]VertexColor, verticesAmount),
		parent:        make(map[*Vertex]*Vertex, verticesAmount),
		distance:      make(map[*Vertex]int, verticesAmount),
		currentTime:   0,
		discoveryTime: make(map[*Vertex]int, verticesAmount),
		finishTime:    make(map[*Vertex]int, verticesAmount),
	}

	for _, vertex := range graph.Vertices {
		newState.color[vertex] = WHITE
		newState.parent[vertex] = nil
		newState.distance[vertex] = -1 // we just use -1 for infinity
		newState.discoveryTime[vertex] = -1
		newState.finishTime[vertex] = -1
	}

	return newState
}

func BFS(graph *Graph, start *Vertex) SearchState {
	// storing color and all that stuff in a separate struct instead of making the algorithm tightly coupled with the graph struct
	state := newSearchState(graph)
	state.color[start] = GRAY
	state.distance[start] = 0
	state.parent[start] = nil

	// for queue im just using a basic slice cus that works fine in this scenario
	queue := []*Vertex{start}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:] // dequeue

		for _, v := range graph.AdjacencyList[u.key] {
			if state.color[v] == WHITE {
				state.color[v] = GRAY
				state.distance[v] = state.distance[u] + 1
				state.parent[v] = u
				queue = append(queue, v)
			}
		}

		state.color[u] = BLACK
	}

	return state
}

func DFS(graph *Graph) SearchState {
	state := newSearchState(graph)

	sortedVertices := graph.GetVerticesSorted()
	for _, vertex := range sortedVertices {
		if state.color[vertex] == WHITE {
			dfsVisit(graph, vertex, &state)
		}
	}
	return state
}

func dfsVisit(graph *Graph, vertex *Vertex, state *SearchState) {
	state.currentTime++
	state.discoveryTime[vertex] = state.currentTime
	state.color[vertex] = GRAY

	for _, v := range graph.AdjacencyList[vertex.key] {
		if state.color[v] == WHITE {
			state.parent[v] = vertex
			dfsVisit(graph, v, state)
		}
	}

	state.currentTime++
	state.finishTime[vertex] = state.currentTime
	state.color[vertex] = BLACK
}

func main() {
	fmt.Println("============ BFS ============ ")

	bfsTestGraph := NewGraph(false)
	// (bla bla bad code use a map instea- I DONT CARE STFU)
	bfsTestGraph.AddEdge("S", "B")
	bfsTestGraph.AddEdge("B", "D")
	bfsTestGraph.AddEdge("D", "A")
	bfsTestGraph.AddEdge("S", "E")
	bfsTestGraph.AddEdge("B", "E")
	bfsTestGraph.AddEdge("E", "C")
	bfsTestGraph.AddEdge("E", "D")
	bfsTestGraph.AddEdge("E", "F")
	bfsTestGraph.AddEdge("C", "F")
	bfsTestGraph.AddEdge("D", "F")
	bfsTestGraph.AddEdge("A", "F")

	finishedBfsState := BFS(&bfsTestGraph, bfsTestGraph.Vertices["S"])

	for _, vertex := range bfsTestGraph.Vertices {
		fmt.Println("\nKey:", vertex.key)
		fmt.Println("Distance:", finishedBfsState.distance[vertex])
		if parent := finishedBfsState.parent[vertex]; parent != nil {
			fmt.Println("Parent:", finishedBfsState.parent[vertex].key)
		} else {
			fmt.Println("Parent:", "Is Root")
		}
	}

	fmt.Println("============ DFS ============ ")
	dfsTestGraph := NewGraph(true)
	dfsTestGraph.AddEdge("q", "s")
	dfsTestGraph.AddEdge("q", "t")
	dfsTestGraph.AddEdge("q", "w")
	dfsTestGraph.AddEdge("r", "y")
	dfsTestGraph.AddEdge("r", "u")
	dfsTestGraph.AddEdge("s", "v")
	dfsTestGraph.AddEdge("t", "x")
	dfsTestGraph.AddEdge("t", "y")
	dfsTestGraph.AddEdge("v", "w")
	dfsTestGraph.AddEdge("w", "s")
	dfsTestGraph.AddEdge("x", "z")
	dfsTestGraph.AddEdge("y", "q")
	dfsTestGraph.AddEdge("u", "y")
	dfsTestGraph.AddEdge("z", "x")

	dfsTestGraph.SortAdjacencyLists()
	finishedDfsState := DFS(&dfsTestGraph)

	for _, vertex := range dfsTestGraph.Vertices {
		fmt.Println("\nKey:", vertex.key)
		fmt.Println("Discovery Time:", finishedDfsState.discoveryTime[vertex])
		fmt.Println("Finish Time:", finishedDfsState.finishTime[vertex])
		if parent := finishedDfsState.parent[vertex]; parent != nil {
			fmt.Println("Parent:", finishedDfsState.parent[vertex].key)
		} else {
			fmt.Println("Parent:", "Is Root")
		}
	}
}
