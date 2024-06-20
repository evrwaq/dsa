package data_structures

import "errors"

type Array struct {
	elements []interface{}
}

func NewArray(size int) *Array {
	return &Array{
		elements: make([]interface{}, size),
	}
}

func (array *Array) Set(index int, value interface{}) error {
	if index < 0 || index >= len(array.elements) {
		return errors.New("index out of bounds")
	}
	array.elements[index] = value
	return nil
}

func (array *Array) Get(index int) (interface{}, error) {
	if index < 0 || index >= len(array.elements) {
		return nil, errors.New("index out of bounds")
	}
	return array.elements[index], nil
}

func (array *Array) Size() int {
	return len(array.elements)
}
