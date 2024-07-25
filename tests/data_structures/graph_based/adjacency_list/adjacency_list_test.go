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

func TestAddEdge(t *testing.T) {
	al := ds.NewAdjacencyList()
	al.AddVertex("A")
	al.AddVertex("B")
	al.AddEdge("A", "B")

	edges := al.GetEdges("A")
	if len(edges) != 1 || edges[0] != "B" {
		t.Errorf("Expected edge from A to B, got %v", edges)
	}
}

func TestGetEdgesNonExistentVertex(t *testing.T) {
	al := ds.NewAdjacencyList()
	edges := al.GetEdges("NonExistent")
	if edges != nil {
		t.Errorf("Expected nil, got %v", edges)
	}
}

func TestRemoveVertex(t *testing.T) {
	al := ds.NewAdjacencyList()
	al.AddVertex("A")
	al.AddVertex("B")
	al.AddEdge("A", "B")
	al.RemoveVertex("A")

	if len(al.GetVertices()) != 1 {
		t.Errorf("Expected 1 vertex, got %d", len(al.GetVertices()))
	}

	edges := al.GetEdges("A")
	if edges != nil {
		t.Errorf("Expected nil, got %v", edges)
	}

	al.AddVertex("A")
	al.AddVertex("B")
	al.AddVertex("C")
	al.AddEdge("A", "B")
	al.AddEdge("B", "C")
	al.AddEdge("A", "C")

	al.RemoveVertex("B")

	edgesA := al.GetEdges("A")
	if len(edgesA) != 1 || edgesA[0] != "C" {
		t.Errorf("Expected edges of A to be [C], got %v", edgesA)
	}

	edgesC := al.GetEdges("C")
	if len(edgesC) != 0 {
		t.Errorf("Expected edges of C to be [], got %v", edgesC)
	}
}

func TestRemoveEdge(t *testing.T) {
	al := ds.NewAdjacencyList()
	al.AddVertex("A")
	al.AddVertex("B")
	al.AddEdge("A", "B")
	al.RemoveEdge("A", "B")

	edges := al.GetEdges("A")
	if len(edges) != 0 {
		t.Errorf("Expected 0 edges, got %d", len(edges))
	}
}

func TestBFS(t *testing.T) {
	al := ds.NewAdjacencyList()
	al.AddVertex("A")
	al.AddVertex("B")
	al.AddVertex("C")
	al.AddVertex("D")
	al.AddEdge("A", "B")
	al.AddEdge("A", "C")
	al.AddEdge("B", "D")
	al.AddEdge("C", "D")

	expectedOrder := []string{"A", "B", "C", "D"}
	order := al.BFS("A")

	if !equal(order, expectedOrder) {
		t.Errorf("Expected BFS order %v, got %v", expectedOrder, order)
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestDFS(t *testing.T) {
	al := ds.NewAdjacencyList()
	al.AddVertex("A")
	al.AddVertex("B")
	al.AddVertex("C")
	al.AddVertex("D")
	al.AddEdge("A", "B")
	al.AddEdge("A", "C")
	al.AddEdge("B", "D")
	al.AddEdge("C", "D")

	expectedOrder := []string{"A", "B", "D", "C"}
	order := al.DFS("A")

	if !equal(order, expectedOrder) {
		t.Errorf("Expected DFS order %v, got %v", expectedOrder, order)
	}
}
