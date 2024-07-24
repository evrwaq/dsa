package data_structures

import (
	ds_errors "dsa/src/data_structures/errors"
	"errors"
)

type Array struct {
	elements []interface{}
}

func NewArray(size int) *Array {
	return &Array{
		elements: make([]interface{}, size),
	}
}

func (array *Array) Size() int {
	return len(array.elements)
}

func (array *Array) Set(index int, value interface{}) error {
	if index < 0 || index >= len(array.elements) {
		return errors.New(ds_errors.IndexOutOfBoundsError)
	}
	array.elements[index] = value
	return nil
}

func (array *Array) Get(index int) (interface{}, error) {
	if index < 0 || index >= len(array.elements) {
		return nil, errors.New(ds_errors.IndexOutOfBoundsError)
	}
	return array.elements[index], nil
}
