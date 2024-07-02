package data_structures

import "errors"

// Stack represents a stack data structure that holds elements.
type Stack struct {
	elements []interface{} // A slice to store the stack elements.
}

// NewStack creates a new empty stack.
func NewStack() *Stack {
	return &Stack{elements: []interface{}{}} // Initialize the slice to store stack elements.
}

// Size returns the number of elements in the stack.
// It simply returns the length of the elements slice.
func (stack *Stack) Size() int {
	return len(stack.elements)
}

// IsEmpty checks if the stack is empty.
// It returns true if the elements slice has no elements, and false otherwise.
func (stack *Stack) IsEmpty() bool {
	return len(stack.elements) == 0
}

// Push adds a new element to the top of the stack.
// It appends the element to the end of the elements slice.
func (stack *Stack) Push(value interface{}) {
	stack.elements = append(stack.elements, value)
}

// Pop removes and returns the element at the top of the stack.
// It returns an error if the stack is empty.
func (stack *Stack) Pop() (interface{}, error) {
	// Check if the stack is empty.
	if stack.IsEmpty() {
		return nil, errors.New(StackEmptyError) // Return an error if the stack is empty.
	}
	// Get the value of the top element.
	value := stack.elements[len(stack.elements)-1]
	// Remove the top element by slicing off the last element.
	stack.elements = stack.elements[:len(stack.elements)-1]
	return value, nil // Return the value of the removed element.
}

// Peek returns the element at the top of the stack without removing it.
// It returns an error if the stack is empty.
func (stack *Stack) Peek() (interface{}, error) {
	// Check if the stack is empty.
	if stack.IsEmpty() {
		return nil, errors.New(StackEmptyError) // Return an error if the stack is empty.
	}
	// Return the value of the top element without removing it.
	return stack.elements[len(stack.elements)-1], nil
}
