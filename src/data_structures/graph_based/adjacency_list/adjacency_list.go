package data_structures

type AdjacencyList struct {
	vertices map[string][]string
}

func NewAdjacencyList() *AdjacencyList {
	return &AdjacencyList{
		vertices: make(map[string][]string),
	}
}

func (al *AdjacencyList) GetVertices() []string {
	vertices := make([]string, 0, len(al.vertices))
	for vertex := range al.vertices {
		vertices = append(vertices, vertex)
	}
	return vertices
}

func (al *AdjacencyList) AddVertex(vertex string) {
	if _, exists := al.vertices[vertex]; !exists {
		al.vertices[vertex] = []string{}
	}
}

func (al *AdjacencyList) AddEdge(from, to string) {
	if _, exists := al.vertices[from]; exists {
		if _, exists := al.vertices[to]; exists {
			al.vertices[from] = append(al.vertices[from], to)
		}
	}
}

func (al *AdjacencyList) GetEdges(vertex string) []string {
	if edges, exists := al.vertices[vertex]; exists {
		return edges
	}
	return nil
}

func (al *AdjacencyList) RemoveVertex(vertex string) {
	delete(al.vertices, vertex)
	for v := range al.vertices {
		edges := al.vertices[v]
		for i := len(edges) - 1; i >= 0; i-- {
			if edges[i] == vertex {
				al.vertices[v] = append(edges[:i], edges[i+1:]...)
			}
		}
	}
}

func (al *AdjacencyList) RemoveEdge(from, to string) {
	if edges, exists := al.vertices[from]; exists {
		for i := len(edges) - 1; i >= 0; i-- {
			if edges[i] == to {
				al.vertices[from] = append(edges[:i], edges[i+1:]...)
			}
		}
	}
}
