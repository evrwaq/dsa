package data_structures

import "errors"

type Queue struct {
	elements []interface{}
}

func NewQueue() *Queue {
	return &Queue{elements: []interface{}{}}
}

func (queue *Queue) Size() int {
	return len(queue.elements)
}

func (queue *Queue) IsEmpty() bool {
	return len(queue.elements) == 0
}

func (queue *Queue) Enqueue(value interface{}) {
	queue.elements = append(queue.elements, value)
}

func (queue *Queue) Dequeue() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	value := queue.elements[0]
	queue.elements = queue.elements[1:]
	return value, nil
}
