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

func (g *Graph) GetVertices() []string {
	vertices := make([]string, 0, len(g.vertices))
	for vertex := range g.vertices {
		vertices = append(vertices, vertex)
	}
	return vertices
}
