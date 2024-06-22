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

func (queue *Queue) Enqueue(value interface{}) {
	queue.elements = append(queue.elements, value)
}
