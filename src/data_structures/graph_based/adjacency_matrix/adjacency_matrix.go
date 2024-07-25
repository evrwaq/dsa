package data_structures

type AdjacencyMatrix struct {
	matrix  [][]bool
	indices map[string]int
	names   []string
}

func NewAdjacencyMatrix() *AdjacencyMatrix {
	return &AdjacencyMatrix{
		matrix:  make([][]bool, 0),
		indices: make(map[string]int),
		names:   make([]string, 0),
	}
}

func (am *AdjacencyMatrix) GetVertices() []string {
	return am.names
}

func (am *AdjacencyMatrix) AddVertex(vertex string) {
	if _, exists := am.indices[vertex]; !exists {
		am.indices[vertex] = len(am.names)
		am.names = append(am.names, vertex)
		for i := range am.matrix {
			am.matrix[i] = append(am.matrix[i], false)
		}
		am.matrix = append(am.matrix, make([]bool, len(am.names)))
	}
}
