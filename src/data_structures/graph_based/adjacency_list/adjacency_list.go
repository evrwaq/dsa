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
