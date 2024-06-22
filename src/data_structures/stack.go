package data_structures

type Stack struct {
	elements []interface{}
}

func NewStack() *Stack {
	return &Stack{elements: []interface{}{}}
}

func (stack *Stack) Size() int {
	return len(stack.elements)
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.elements) == 0
}

func (stack *Stack) Push(value interface{}) {
	stack.elements = append(stack.elements, value)
}
