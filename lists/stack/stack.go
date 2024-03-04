package stack

import "errors"

var ErrEmptyStack = errors.New("cannot perform operation from empty stack")

type node struct {
	prev *node
	val  interface{}
}

type Stack struct {
	length int
	top    *node
}

func New() Stack {
	return Stack{
		length: 0,
	}
}

func (s *Stack) Length() int {
	return s.length
}

func (s *Stack) Push(val interface{}) {
	n := node{val: val}
	n.prev = s.top
	s.top = &n
	s.length++
}

func (s *Stack) Pop() (interface{}, error) {
	if s.length == 0 {
		return nil, ErrEmptyStack
	}
	val := s.top.val
	s.top = s.top.prev
	s.length--
	return val, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if s.length == 0 {
		return nil, ErrEmptyStack
	}
	return s.top.val, nil
}
