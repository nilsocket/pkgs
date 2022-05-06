package stack

import "errors"

type Stack struct {
	data  []any
	index int
}

func New() *Stack {
	return &Stack{index: -1}
}

func (s *Stack) Top() any {
	return s.data[s.index]
}

func (s *Stack) Push(a any) {
	s.index++
	if s.index < len(s.data) {
		s.data[s.index] = a
	} else {
		s.data = append(s.data, a)
	}
}

var ErrNoElements = errors.New("no more elements in stack")

func (s *Stack) Pop() (any, error) {
	if s.index > -1 {
		s.index--
		return s.data[s.index+1], nil
	}
	return nil, ErrNoElements
}
