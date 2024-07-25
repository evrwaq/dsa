package data_structures_tests

import (
	"testing"

	ds "dsa/src/data_structures/graph_based/graph"
)

func TestNewGraph(t *testing.T) {
	graph := ds.NewGraph()
	if len(graph.GetVertices()) != 0 {
		t.Errorf("Expected 0 vertices, got %d", len(graph.GetVertices()))
	}
}

func TestAddVertex(t *testing.T) {
	graph := ds.NewGraph()
	graph.AddVertex("A")
	vertices := graph.GetVertices()
	if len(vertices) != 1 || vertices[0] != "A" {
		t.Errorf("Expected vertex A, got %v", vertices)
	}
}
