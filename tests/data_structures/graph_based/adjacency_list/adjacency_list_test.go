package data_structures_tests

import (
	"testing"

	ds "dsa/src/data_structures/graph_based/adjacency_list"
)

func TestNewAdjacencyList(t *testing.T) {
	al := ds.NewAdjacencyList()
	if len(al.GetVertices()) != 0 {
		t.Errorf("Expected 0 vertices, got %d", len(al.GetVertices()))
	}
}

func TestAddVertex(t *testing.T) {
	al := ds.NewAdjacencyList()
	al.AddVertex("A")
	vertices := al.GetVertices()
	if len(vertices) != 1 || vertices[0] != "A" {
		t.Errorf("Expected vertex A, got %v", vertices)
	}
}
