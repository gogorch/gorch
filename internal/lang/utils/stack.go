package utils

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value any
		prev  *node
	}
)

// Create a new stack
func NewStack() *Stack {
	return &Stack{nil, 0}
}

// Return the number of items in the stack
func (s *Stack) Len() int {
	return s.length
}

// View the top item on the stack
func (s *Stack) Peek() any {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

// Pop the top item of the stack and return it
func (s *Stack) Pop() any {
	if s.length == 0 {
		return nil
	}

	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

// Push a value onto the top of the stack
func (s *Stack) Push(value any) {
	n := &node{value, s.top}
	s.top = n
	s.length++
}
