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

func (am *AdjacencyMatrix) AddEdge(from, to string) {
	fromIndex, fromExists := am.indices[from]
	toIndex, toExists := am.indices[to]
	if fromExists && toExists {
		am.matrix[fromIndex][toIndex] = true
	}
}

func (am *AdjacencyMatrix) GetEdges(vertex string) []string {
	if index, exists := am.indices[vertex]; exists {
		edges := []string{}
		for i, connected := range am.matrix[index] {
			if connected {
				edges = append(edges, am.names[i])
			}
		}
		return edges
	}
	return nil
}

func (am *AdjacencyMatrix) RemoveVertex(vertex string) {
	if index, exists := am.indices[vertex]; exists {
		delete(am.indices, vertex)
		am.names = append(am.names[:index], am.names[index+1:]...)

		for i := range am.matrix {
			am.matrix[i] = append(am.matrix[i][:index], am.matrix[i][index+1:]...)
		}
		am.matrix = append(am.matrix[:index], am.matrix[index+1:]...)

		for i := index; i < len(am.names); i++ {
			am.indices[am.names[i]] = i
		}
	}
}

func (am *AdjacencyMatrix) RemoveEdge(from, to string) {
	fromIndex, fromExists := am.indices[from]
	toIndex, toExists := am.indices[to]
	if fromExists && toExists {
		am.matrix[fromIndex][toIndex] = false
	}
}

func (am *AdjacencyMatrix) AreVerticesConnected(from, to string) bool {
	fromIndex, fromExists := am.indices[from]
	toIndex, toExists := am.indices[to]
	if fromExists && toExists {
		return am.matrix[fromIndex][toIndex]
	}
	return false
}
