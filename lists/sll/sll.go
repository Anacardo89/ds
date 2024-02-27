package sll

import "errors"

var (
	ErrEmptyList   = errors.New("cannot perform operation in empty list")
	ErrOutOfBounds = errors.New("index exceeds length of list")
	ErrNotInList   = errors.New("value not in list")
)

type SLL struct {
	length int
	head   *Node
	tail   *Node
}

func New() SLL {
	return SLL{
		length: 0,
	}
}

func (l *SLL) Length() int {
	return l.length
}

func (l *SLL) Head() (interface{}, error) {
	if l.length == 0 {
		return nil, ErrEmptyList
	}
	return l.head.val, nil
}

func (l *SLL) Tail() (interface{}, error) {
	if l.length == 0 {
		return nil, ErrEmptyList
	}
	return l.tail.val, nil
}

func (l *SLL) Prepend(val interface{}) {
	n := Node{val: val}
	if l.length == 0 {
		l.head, l.tail = &n, &n
	} else {
		n.next = l.head
		l.head = &n
	}
	l.length++
}

func (l *SLL) Append(val interface{}) {
	n := Node{val: val}
	if l.length == 0 {
		l.head, l.tail = &n, &n
	} else {
		l.tail.next = &n
		l.tail = &n
	}
	l.length++
}

func (l *SLL) InsertAt(idx int, val interface{}) error {
	if idx > l.length {
		return ErrOutOfBounds
	}
	switch idx {
	case 0:
		l.Prepend(val)
		return nil
	case l.length:
		l.Append(val)
		return nil
	}
	current, prev := l.head, l.head
	for i := 0; i < idx; i++ {
		prev = current
		current = current.next
	}
	n := Node{val: val}
	n.next = current
	prev.next = &n
	l.length++
	return nil
}

func (l *SLL) GetAt(idx int) (interface{}, error) {
	if l.length == 0 {
		return nil, ErrEmptyList
	}
	if idx >= l.length {
		return nil, ErrOutOfBounds
	}
	current := l.head
	for i := 0; i < idx; i++ {
		current = current.next
	}
	return current.val, nil
}

func (l *SLL) Remove(val interface{}) error {
	switch l.length {
	case 0:
		return ErrEmptyList
	case 1:
		if val == l.head.val {
			l.head, l.tail = nil, nil
			l.length--
			return nil
		}
	}
	if val == l.head.val {
		l.head = l.head.next
		l.length--
		return nil
	}
	current, prev := l.head.next, l.head
	for i := 1; i < l.length-1; i++ {
		if current.val == val {
			prev.next = current.next
			current = nil
			l.length--
			return nil
		}
		prev = current
		current = current.next
	}
	return ErrNotInList
}

func (l *SLL) RemoveAt(idx int) error {
	if l.length == 0 {
		return ErrEmptyList
	}
	if idx >= l.length {
		return ErrOutOfBounds
	}
	if l.length == 1 {
		l.head, l.tail = nil, nil
		l.length--
		return nil
	}
	if idx == 0 {
		l.head = l.head.next
		l.length--
		return nil
	}
	current, prev := l.head, l.head
	for i := 0; i < idx; i++ {
		if current.next == nil {
			break
		}
		prev = current
		current = current.next
	}
	if idx == l.length-1 {
		l.tail = prev
	}
	prev.next = current.next
	current = nil
	l.length--
	return nil
}
