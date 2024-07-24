package data_structures

import (
	ds_errors "dsa/src/data_structures/errors"
	"errors"
)

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
		return nil, errors.New(ds_errors.QueueEmptyError)
	}
	value := queue.elements[0]
	queue.elements = queue.elements[1:]
	return value, nil
}

func (queue *Queue) Peek() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, errors.New(ds_errors.QueueEmptyError)
	}
	return queue.elements[0], nil
}
