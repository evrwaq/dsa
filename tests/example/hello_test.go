package example_test

import (
	"testing"

	"github.com/evrwaq/dsa/internal/example"
)

func TestHello(t *testing.T) {
	expected := "Hello, World!"
	result := example.Hello()

	if result != expected {
		t.Errorf("expected %q but got %q", expected, result)
	}
}

func TestPrintHello(t *testing.T) {
	example.PrintHello()
}
