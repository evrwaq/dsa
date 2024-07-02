package data_structures

import "errors"

// Deque represents a double-ended queue data structure that holds elements.
type Deque struct {
	elements []interface{} // A slice to store the deque elements.
}

// NewDeque creates a new empty deque.
func NewDeque() *Deque {
	return &Deque{elements: []interface{}{}} // Initialize the slice to store deque elements.
}

// Size returns the number of elements in the deque.
// It simply returns the length of the elements slice.
func (deque *Deque) Size() int {
	return len(deque.elements)
}

// IsEmpty checks if the deque is empty.
// It returns true if the elements slice has no elements, and false otherwise.
func (deque *Deque) IsEmpty() bool {
	return len(deque.elements) == 0
}

// AddFront adds a new element to the front of the deque.
// It prepends the element to the start of the elements slice.
func (deque *Deque) AddFront(value interface{}) {
	deque.elements = append([]interface{}{value}, deque.elements...)
}

// AddBack adds a new element to the end of the deque.
// It appends the element to the end of the elements slice.
func (deque *Deque) AddBack(value interface{}) {
	deque.elements = append(deque.elements, value)
}

// RemoveFront removes and returns the element at the front of the deque.
// It returns an error if the deque is empty.
func (deque *Deque) RemoveFront() (interface{}, error) {
	// Check if the deque is empty.
	if deque.IsEmpty() {
		return nil, errors.New(DequeEmptyError) // Return an error if the deque is empty.
	}
	// Get the value of the front element.
	value := deque.elements[0]
	// Remove the front element by slicing off the first element.
	deque.elements = deque.elements[1:]
	return value, nil // Return the value of the removed element.
}

// RemoveBack removes and returns the element at the end of the deque.
// It returns an error if the deque is empty.
func (deque *Deque) RemoveBack() (interface{}, error) {
	// Check if the deque is empty.
	if deque.IsEmpty() {
		return nil, errors.New(DequeEmptyError) // Return an error if the deque is empty.
	}
	// Get the value of the back element.
	value := deque.elements[len(deque.elements)-1]
	// Remove the back element by slicing off the last element.
	deque.elements = deque.elements[:len(deque.elements)-1]
	return value, nil // Return the value of the removed element.
}

// PeekFront returns the element at the front of the deque without removing it.
// It returns an error if the deque is empty.
func (deque *Deque) PeekFront() (interface{}, error) {
	// Check if the deque is empty.
	if deque.IsEmpty() {
		return nil, errors.New(DequeEmptyError) // Return an error if the deque is empty.
	}
	// Return the value of the front element without removing it.
	return deque.elements[0], nil
}

// PeekBack returns the element at the end of the deque without removing it.
// It returns an error if the deque is empty.
func (deque *Deque) PeekBack() (interface{}, error) {
	// Check if the deque is empty.
	if deque.IsEmpty() {
		return nil, errors.New(DequeEmptyError) // Return an error if the deque is empty.
	}
	// Return the value of the back element without removing it.
	return deque.elements[len(deque.elements)-1], nil
}
