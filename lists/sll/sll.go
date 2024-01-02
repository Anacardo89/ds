package sll

import "errors"

type node struct {
	next *node
	val  interface{}
}

type SLL struct {
	length int
	head   *node
	tail   *node
}

func New() SLL {
	return SLL{
		length: 0,
	}
}

func (l *SLL) GetLength() int {
	return l.length
}

func (l *SLL) GetHead() (interface{}, error) {
	if l.length == 0 {
		return nil, errors.New("list has no nodes")
	}
	return l.head.val, nil
}

func (l *SLL) GetTail() (interface{}, error) {
	if l.length == 0 {
		return nil, errors.New("list has no nodes")
	}
	return l.tail.val, nil
}

func (l *SLL) Prepend(val interface{}) {
	n := node{val: val}
	if l.length == 0 {
		l.head, l.tail = &n, &n
	} else {
		n.next = l.head
		l.head = &n
	}
	l.length++
}

func (l *SLL) Append(val interface{}) {
	n := node{val: val}
	if l.length == 0 {
		l.head, l.tail = &n, &n
	} else {
		l.tail.next = &n
		l.tail = &n
	}
	l.length++
}

func (l *SLL) Remove(val interface{}) (interface{}, error) {
	current := l.head
	var prev *node
	if l.length == 0 {
		return val, errors.New("empty list")
	}
	for i := 0; i < l.length; i++ {
		if current.val == val {
			break
		}
		prev = current
		if current.next != nil {
			current = current.next
		} else {
			return val, errors.New("value not in list")
		}
	}
	if l.length == 1 {
		l.length--
		l.head = nil
		l.tail = nil
		return val, nil
	}
	if current == l.head {
		l.head = current.next
		current = nil
	} else if current == l.tail {
		l.tail = prev
		prev.next = nil
	} else {
		prev.next = current.next
		current = nil
	}
	l.length--
	return val, nil
}

func (l *SLL) InsertAt(idx int, val interface{}) error {
	if idx > l.length {
		return errors.New("index exceeds length")
	}
	n := node{val: val}
	if idx == 0 {
		n.next = l.head
		l.head = &n
		l.length++
		return nil
	}
	current, prev := l.head, l.head
	for i := 0; i < idx; i++ {
		prev = current
		current = current.next
	}
	n.next = current
	prev.next = &n
	l.length++
	return nil
}

func (l *SLL) RemoveAt(idx int) error {
	if idx > l.length-1 {
		return errors.New("index exceeds length")
	}
	current, prev := l.head, l.head
	for i := 0; i < idx; i++ {
		prev = current
		current = current.next
	}
	prev.next = current.next
	current.next = nil
	l.length--
	return nil
}

func (l *SLL) GetAt(idx int) (interface{}, error) {
	if idx > l.length-1 {
		return nil, errors.New("index exceeds length")
	}
	current := l.head
	for i := 0; i < idx; i++ {
		current = current.next
	}
	return current.val, nil
}