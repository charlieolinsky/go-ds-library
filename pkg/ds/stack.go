package ds

import "errors"

type Stack[T any] struct {
	Data []T
}

/*
	Add an item on top of the stack
*/
func (s *Stack[T]) Push(x T) {
	s.Data = append(s.Data, x)
}

/*
	Remove an item from the top of the stack
*/
func (s *Stack[T]) Pop() (T, error) {
	if len(s.Data) == 0 {
		return *new(T), errors.New("cannot pop from an empty stack")
	}

	popped := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]

	return popped, nil
}

/*
	View an item from the top of the stack without removing it
*/
func (s *Stack[T]) Peek() (T, error) {
	if len(s.Data) == 0 {
		return *new(T), errors.New("cannot peek from an empty stack")
	}
	return s.Data[len(s.Data)-1], nil
}

/*
	Return true if the stack is empty and false otherwise
*/
func (s *Stack[T]) IsEmpty() bool {
	return len(s.Data) == 0
}
