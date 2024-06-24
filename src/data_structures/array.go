package data_structures

import "errors"

// Array represents a fixed-size array data structure.
type Array struct {
	elements []interface{} // A slice to store the elements of the array.
}

// NewArray creates a new array with the specified size.
// It initializes a slice with the given size to store the array elements.
func NewArray(size int) *Array {
	return &Array{
		elements: make([]interface{}, size), // Create a slice with the specified size.
	}
}

// Size returns the number of elements in the array.
// Since the size of the array is fixed, it simply returns the length of the slice.
func (array *Array) Size() int {
	return len(array.elements)
}

// Set assigns a value to the specified index in the array.
// It returns an error if the index is out of bounds.
func (array *Array) Set(index int, value interface{}) error {
	// Check if the index is within the valid range.
	if index < 0 || index >= len(array.elements) {
		return errors.New("index out of bounds") // Return an error if the index is invalid.
	}
	array.elements[index] = value // Set the value at the specified index.
	return nil
}

// Get retrieves the value at the specified index in the array.
// It returns an error if the index is out of bounds.
func (array *Array) Get(index int) (interface{}, error) {
	// Check if the index is within the valid range.
	if index < 0 || index >= len(array.elements) {
		return nil, errors.New("index out of bounds") // Return an error if the index is invalid.
	}
	return array.elements[index], nil // Return the value at the specified index.
}
