package data_structures

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
