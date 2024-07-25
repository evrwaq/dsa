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
