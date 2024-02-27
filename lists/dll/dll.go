package dll

import "errors"

var (
	ErrEmptyList   = errors.New("cannot perform operation in empty list")
	ErrOutOfBounds = errors.New("index exceeds length of list")
	ErrNotInList   = errors.New("value not in list")
)

type DLL struct {
	length int
	head   *Node
	tail   *Node
}

func New() DLL {
	return DLL{
		length: 0,
	}
}

func (l *DLL) Length() int {
	return l.length
}

func (l *DLL) Head() (interface{}, error) {
	if l.length == 0 {
		return nil, ErrEmptyList
	}
	return l.head.val, nil
}

func (l *DLL) Tail() (interface{}, error) {
	if l.length == 0 {
		return nil, ErrEmptyList
	}
	return l.tail.val, nil
}

func (l *DLL) Prepend(val interface{}) {
	n := Node{val: val}
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
	n := Node{val: val}
	if l.length == 0 {
		l.head, l.tail = &n, &n
	} else {
		l.tail.next = &n
		n.prev = l.tail
		l.tail = &n
	}
	l.length++
}

func (l *DLL) InsertAt(idx int, val interface{}) error {
	if idx > l.length {
		return ErrOutOfBounds
	}
	if l.length == 0 {
		l.Prepend(val)
		return nil
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
	n.next, n.prev = current, prev
	prev.next, current.prev = &n, &n
	l.length++
	return nil
}

func (l *DLL) GetAt(idx int) (interface{}, error) {
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

func (l *DLL) Remove(val interface{}) error {
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
	switch val {
	case l.head.val:
		l.head = l.head.next
		l.head.prev = nil
		l.length--
		return nil
	case l.tail.val:
		l.tail = l.tail.prev
		l.tail.next = nil
		l.length--
		return nil
	}
	current, prev := l.head.next, l.head
	for i := 1; i < l.length-1; i++ {
		if current.val == val {
			prev.next = current.next
			current.next.prev = prev
			current = nil
			l.length--
			return nil
		}
		prev = current
		current = current.next
	}
	return ErrNotInList
}

func (l *DLL) RemoveAt(idx int) error {
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
	switch idx {
	case 0:
		l.head = l.head.next
		l.head.prev = nil
		l.length--
		return nil
	case l.length - 1:
		l.tail = l.tail.prev
		l.tail.next = nil
		l.length--
		return nil
	}
	current, prev := l.head, l.head
	for i := 0; i < idx; i++ {
		prev = current
		current = current.next
	}
	prev.next = current.next
	current.next.prev = prev
	current = nil
	l.length--
	return nil
}
