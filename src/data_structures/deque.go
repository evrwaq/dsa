package data_structures

import "errors"

type Deque struct {
	elements []interface{}
}

func NewDeque() *Deque {
	return &Deque{elements: []interface{}{}}
}

func (deque *Deque) Size() int {
	return len(deque.elements)
}

func (deque *Deque) IsEmpty() bool {
	return len(deque.elements) == 0
}

func (deque *Deque) AddFront(value interface{}) {
	deque.elements = append([]interface{}{value}, deque.elements...)
}

func (deque *Deque) AddBack(value interface{}) {
	deque.elements = append(deque.elements, value)
}

func (deque *Deque) RemoveFront() (interface{}, error) {
	if deque.IsEmpty() {
		return nil, errors.New("deque is empty")
	}
	value := deque.elements[0]
	deque.elements = deque.elements[1:]
	return value, nil
}

func (deque *Deque) RemoveBack() (interface{}, error) {
	if deque.IsEmpty() {
		return nil, errors.New("deque is empty")
	}
	value := deque.elements[len(deque.elements)-1]
	deque.elements = deque.elements[:len(deque.elements)-1]
	return value, nil
}
