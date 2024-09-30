package ds

import "errors"

type Stack struct {
	Data []int
}

/*
	Add an item on top of the stack
*/
func (s *Stack) Push(i int) {
	s.Data = append(s.Data, i)
}

/*
	Remove an item from the top of the stack
*/
func (s *Stack) Pop() (int, error) {
	if len(s.Data) == 0 {
		return 0, errors.New("cannot pop from an empty stack")
	}

	popped := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]

	return popped, nil
}

/*
	View an item from the top of the stack without removing it
*/
func (s *Stack) Peek() (int, error) {
	if len(s.Data) == 0 {
		return 0, errors.New("cannot peek from an empty stack")
	}
	return s.Data[len(s.Data)-1], nil
}

/*
	Return true if the stack is empty and false otherwise
*/
func (s *Stack) IsEmpty() bool {
	return len(s.Data) == 0
}
