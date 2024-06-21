package data_structures

import "errors"

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
		return nil, errors.New("index out of bounds")
	}
	current := linkedList.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.value, nil
}
