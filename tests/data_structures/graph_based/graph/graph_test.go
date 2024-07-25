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
