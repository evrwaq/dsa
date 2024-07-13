package data_structures

import (
	ds_errors "dsa/src/data_structures/errors"
	"errors"
)

type Node struct {
	value interface{}
	next  *Node
}

type LinkedList struct {
	head *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{nil, 0}
}

func (linkedList *LinkedList) Size() int {
	return linkedList.size
}

func (linkedList *LinkedList) IsEmpty() bool {
	return linkedList.size == 0
}

func (linkedList *LinkedList) Add(value interface{}) {
	newNode := &Node{value, nil}
	if linkedList.head == nil {
		linkedList.head = newNode
	} else {
		current := linkedList.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	linkedList.size++
}

func (linkedList *LinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= linkedList.size {
		return nil, errors.New(ds_errors.IndexOutOfBoundsError)
	}
	current := linkedList.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.value, nil
}

func (linkedList *LinkedList) Remove(index int) (interface{}, error) {
	if index < 0 || index >= linkedList.size {
		return nil, errors.New(ds_errors.IndexOutOfBoundsError)
	}
	var removedValue interface{}
	if index == 0 {
		removedValue = linkedList.head.value
		linkedList.head = linkedList.head.next
	} else {
		current := linkedList.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		removedValue = current.next.value
		current.next = current.next.next
	}
	linkedList.size--
	return removedValue, nil
}
