package data_structures

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
