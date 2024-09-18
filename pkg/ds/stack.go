package ds

import "errors"

type Stack struct {
	data []int
}

/*
	Add an item on top of the stack
*/
func (s *Stack) Push(i int) {
	s.data = append(s.data, i)
}

/*
	Remove an item from the top of the stack
*/
func (s *Stack) Pop() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("cannot pop from an empty stack")
	}

	popped := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return popped, nil
}

/*
	View an item from the top of the stack without removing it
*/
func (s *Stack) Peek() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("cannot peek from an empty stack")
	}
	return s.data[len(s.data)-1], nil
}

/*
	Return true if the stack is empty and false otherwise
*/
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}
