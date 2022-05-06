package stack

type Stack struct {
	data  []any
	index int
}

// New Stack
func New() *Stack {
	return &Stack{index: -1}
}

// Top returns top element
func (s *Stack) Top() any {
	return s.data[s.index]
}

// Push given element `a` into stack
func (s *Stack) Push(a any) {
	s.index++
	if s.index < len(s.data) {
		s.data[s.index] = a
	} else {
		s.data = append(s.data, a)
	}
}

// Pop element from stack,
// return `nil` when stack is empty.
func (s *Stack) Pop() any {
	if s.index > -1 {
		s.index--
		return s.data[s.index+1]
	}
	return nil
}
