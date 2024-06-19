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

func (a *Array) Set(index int, value interface{}) error {
	if index < 0 || index >= len(a.elements) {
		return errors.New("index out of bounds")
	}
	a.elements[index] = value
	return nil
}

func (a *Array) Get(index int) (interface{}, error) {
	if index < 0 || index >= len(a.elements) {
		return nil, errors.New("index out of bounds")
	}
	return a.elements[index], nil
}

func (a *Array) Size() int {
	return len(a.elements)
}
