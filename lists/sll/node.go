package sll

import "errors"

var (
	ErrNotSet = errors.New("node not set")
)

type Node struct {
	next *Node
	val  interface{}
}

func (n *Node) Next() (*Node, error) {
	if n.next == (*Node)(nil) {
		return nil, ErrNotSet
	}
	return n.next, nil
}

func (n *Node) Val() interface{} {
	return n.val
}

func (n *Node) SetVal(val interface{}) {
	n.val = val
}
