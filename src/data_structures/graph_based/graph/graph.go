package data_structures

type Graph struct {
	vertices map[string]struct{}
	edges    map[string]map[string]struct{}
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[string]struct{}),
		edges:    make(map[string]map[string]struct{}),
	}
}

func (g *Graph) AddVertex(vertex string) {
	if _, exists := g.vertices[vertex]; !exists {
		g.vertices[vertex] = struct{}{}
		g.edges[vertex] = make(map[string]struct{})
	}
}

func (g *Graph) AddEdge(from, to string) {
	if _, exists := g.vertices[from]; exists {
		if _, exists := g.vertices[to]; exists {
			g.edges[from][to] = struct{}{}
		}
	}
}

func (g *Graph) GetVertices() []string {
	vertices := make([]string, 0, len(g.vertices))
	for vertex := range g.vertices {
		vertices = append(vertices, vertex)
	}
	return vertices
}

func (g *Graph) GetEdges(vertex string) []string {
	if _, exists := g.edges[vertex]; exists {
		edges := make([]string, 0, len(g.edges[vertex]))
		for edge := range g.edges[vertex] {
			edges = append(edges, edge)
		}
		return edges
	}
	return nil
}

func (g *Graph) RemoveVertex(vertex string) {
	if _, exists := g.vertices[vertex]; exists {
		delete(g.vertices, vertex)
		delete(g.edges, vertex)
		for v := range g.edges {
			delete(g.edges[v], vertex)
		}
	}
}

func (g *Graph) RemoveEdge(from, to string) {
	if _, exists := g.edges[from]; exists {
		delete(g.edges[from], to)
	}
}
