package dll

import "errors"

type node struct {
	next *node
	prev *node
	val  interface{}
}

type DLL struct {
	length int
	head   *node
	tail   *node
}

func New() DLL {
	return DLL{
		length: 0,
	}
}

func (l *DLL) GetLength() int {
	return l.length
}

func (l *DLL) GetHead() (interface{}, error) {
	if l.length == 0 {
		return nil, errors.New("list has no nodes")
	}
	return l.head.val, nil
}

func (l *DLL) GetTail() (interface{}, error) {
	if l.length == 0 {
		return nil, errors.New("list has no nodes")
	}
	return l.tail.val, nil
}

func (l *DLL) Prepend(val interface{}) {
	n := node{val: val}
	if l.length == 0 {
		l.head, l.tail = &n, &n
	} else {
		n.next = l.head
		l.head.prev = &n
		l.head = &n
	}
	l.length++
}

func (l *DLL) Append(val interface{}) {
	n := node{val: val}
	if l.length == 0 {
		l.head, l.tail = &n, &n
	} else {
		l.tail.next = &n
		n.prev = l.tail
		l.tail = &n
	}
	l.length++
}
