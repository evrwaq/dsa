package data_structures

import "errors"

// Node represents a node in the linked list.
type Node struct {
	value interface{} // The value stored in the node.
	next  *Node       // Pointer to the next node in the list.
}

// LinkedList represents a singly linked list.
type LinkedList struct {
	head *Node // Pointer to the first node in the list.
	size int   // The number of nodes in the list.
}

// NewLinkedList creates a new empty linked list.
func NewLinkedList() *LinkedList {
	return &LinkedList{nil, 0}
}

// Size returns the number of nodes in the list.
// Since the linked list can dynamically change size, it simply returns the number of nodes.
func (linkedList *LinkedList) Size() int {
	return linkedList.size
}

// IsEmpty checks if the list is empty.
// It returns true if the list has no nodes, and false otherwise.
func (linkedList *LinkedList) IsEmpty() bool {
	return linkedList.size == 0
}

// Add appends a new node with the given value to the end of the list.
func (linkedList *LinkedList) Add(value interface{}) {
	newNode := &Node{value, nil} // Create a new node with the given value.
	if linkedList.head == nil {
		linkedList.head = newNode // If the list is empty, set the new node as the head.
	} else {
		current := linkedList.head
		for current.next != nil {
			current = current.next // Traverse to the last node.
		}
		current.next = newNode // Append the new node at the end.
	}
	linkedList.size++
}

// Get retrieves the value at the specified index in the list.
// It returns an error if the index is out of bounds.
func (linkedList *LinkedList) Get(index int) (interface{}, error) {
	// Check if the index is within the valid range.
	if index < 0 || index >= linkedList.size {
		return nil, errors.New("index out of bounds") // Return an error if the index is invalid.
	}
	current := linkedList.head
	for i := 0; i < index; i++ {
		current = current.next // Traverse to the node at the specified index.
	}
	return current.value, nil // Return the value at the specified index.
}

// Remove deletes the node at the specified index in the list and returns its value.
// It returns an error if the index is out of bounds.
func (linkedList *LinkedList) Remove(index int) (interface{}, error) {
	// Check if the index is within the valid range.
	if index < 0 || index >= linkedList.size {
		return nil, errors.New("index out of bounds") // Return an error if the index is invalid.
	}
	var removedValue interface{}
	if index == 0 {
		removedValue = linkedList.head.value   // Store the value of the head node.
		linkedList.head = linkedList.head.next // Update the head to the next node.
	} else {
		current := linkedList.head
		for i := 0; i < index-1; i++ {
			current = current.next // Traverse to the node before the specified index.
		}
		removedValue = current.next.value // Store the value of the node to be removed.
		current.next = current.next.next  // Update the next pointer to skip the removed node.
	}
	linkedList.size--
	return removedValue, nil // Return the value of the removed node.
}
