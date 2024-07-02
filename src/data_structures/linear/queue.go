package data_structures

import "errors"

// Queue represents a queue data structure that holds elements.
type Queue struct {
	elements []interface{} // A slice to store the queue elements.
}

// NewQueue creates a new empty queue.
func NewQueue() *Queue {
	return &Queue{elements: []interface{}{}} // Initialize the slice to store queue elements.
}

// Size returns the number of elements in the queue.
// It simply returns the length of the elements slice.
func (queue *Queue) Size() int {
	return len(queue.elements)
}

// IsEmpty checks if the queue is empty.
// It returns true if the elements slice has no elements, and false otherwise.
func (queue *Queue) IsEmpty() bool {
	return len(queue.elements) == 0
}

// Enqueue adds a new element to the end of the queue.
// It appends the element to the end of the elements slice.
func (queue *Queue) Enqueue(value interface{}) {
	queue.elements = append(queue.elements, value)
}

// Dequeue removes and returns the element at the front of the queue.
// It returns an error if the queue is empty.
func (queue *Queue) Dequeue() (interface{}, error) {
	// Check if the queue is empty.
	if queue.IsEmpty() {
		return nil, errors.New(QueueEmptyError) // Return an error if the queue is empty.
	}
	// Get the value of the front element.
	value := queue.elements[0]
	// Remove the front element by slicing off the first element.
	queue.elements = queue.elements[1:]
	return value, nil // Return the value of the removed element.
}

// Peek returns the element at the front of the queue without removing it.
// It returns an error if the queue is empty.
func (queue *Queue) Peek() (interface{}, error) {
	// Check if the queue is empty.
	if queue.IsEmpty() {
		return nil, errors.New(QueueEmptyError) // Return an error if the queue is empty.
	}
	// Return the value of the front element without removing it.
	return queue.elements[0], nil
}
