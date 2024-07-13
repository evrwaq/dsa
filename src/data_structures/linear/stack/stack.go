package data_structures

import (
	ds_errors "dsa/src/data_structures/errors"
	"errors"
)

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

func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New(ds_errors.StackEmptyError)
	}
	value := stack.elements[len(stack.elements)-1]
	stack.elements = stack.elements[:len(stack.elements)-1]
	return value, nil
}

func (stack *Stack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New(ds_errors.StackEmptyError)
	}
	return stack.elements[len(stack.elements)-1], nil
}
