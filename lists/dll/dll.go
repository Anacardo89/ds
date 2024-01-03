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
		return nil, errors.New("empty list")
	}
	return l.head.val, nil
}

func (l *DLL) GetTail() (interface{}, error) {
	if l.length == 0 {
		return nil, errors.New("empty list")
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

func (l *DLL) Remove(val interface{}) (interface{}, error) {
	if l.length == 0 {
		return val, errors.New("empty list")
	}
	if l.length == 1 {
		l.head = nil
		l.tail = nil
		l.length--
		return val, nil
	}
	if val == l.head.val {
		l.head = l.head.next
		l.head.prev = nil
		l.length--
		return val, nil
	} else if val == l.tail.val {
		l.tail = l.tail.prev
		l.tail.next = nil
		l.length--
		return val, nil
	}
	current := l.head
	var prev *node
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
	prev.next = current.next
	current = nil
	l.length--
	return val, nil
}

func (l *DLL) InsertAt(idx int, val interface{}) error {
	if idx > l.length {
		return errors.New("index exceeds length")
	}
	n := node{val: val}
	if l.length == 0 {
		l.head = &n
		l.tail = &n
		l.length++
		return nil
	}
	if idx == 0 {
		n.next = l.head
		l.head.prev = &n
		l.head = &n
		l.length++
		return nil
	}
	if idx == l.length {
		l.tail.next = &n
		n.prev = l.tail
		l.tail = &n
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
	current.prev = &n
	l.length++
	return nil
}

func (l *DLL) RemoveAt(idx int) error {
	if l.length == 0 {
		return errors.New("empty list")
	}
	if idx >= l.length {
		return errors.New("index exceeds length")
	}
	if l.length == 1 {
		l.length--
		l.head = nil
		l.tail = nil
		return nil
	}
	if idx == 0 {
		l.head = l.head.next
		l.head.prev = nil
		l.length--
		return nil
	}
	if idx == l.length-1 {
		l.tail = l.tail.prev
		l.tail.next = nil
		l.length--
		return nil
	}
	current := l.head
	var prev *node
	for i := 0; i < idx; i++ {
		prev = current
		current = current.next
	}
	prev.next = current.next
	current = nil
	l.length--
	return nil
}

func (l *DLL) GetAt(idx int) (interface{}, error) {
	if l.length == 0 {
		return nil, errors.New("empty list")
	}
	if idx >= l.length {
		return nil, errors.New("index exceeds length")
	}
	current := l.head
	for i := 0; i < idx; i++ {
		current = current.next
	}
	return current.val, nil
}
