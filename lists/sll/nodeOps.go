package sll

func (l *SLL) HeadNode() (*Node, error) {
	if l.length == 0 {
		return nil, ErrEmptyList
	}
	return l.head, nil
}

func (l *SLL) TailNode() (*Node, error) {
	if l.length == 0 {
		return nil, ErrEmptyList
	}
	return l.tail, nil
}

func (l *SLL) PrependNode(n *Node) {
	if l.length == 0 {
		l.head, l.tail = n, n
	} else {
		n.next = l.head
		l.head = n
	}
	l.length++
}

func (l *SLL) AppendNode(n *Node) {
	if l.length == 0 {
		l.head, l.tail = n, n
	} else {
		l.tail.next = n
		l.tail = n
	}
	l.length++
}

func (l *SLL) InsertNodeAt(idx int, n *Node) error {
	if idx > l.length {
		return ErrOutOfBounds
	}
	switch idx {
	case 0:
		l.PrependNode(n)
		return nil
	case l.length:
		l.AppendNode(n)
		return nil
	}
	current, prev := l.head, l.head
	for i := 0; i < idx; i++ {
		prev = current
		current = current.next
	}
	n.next = current
	prev.next = n
	l.length++
	return nil
}

func (l *SLL) GetNodeAt(idx int) (*Node, error) {
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
	return current, nil
}

func (l *SLL) RemoveNode(n *Node) error {
	switch l.length {
	case 0:
		return ErrEmptyList
	case 1:
		if n.val == l.head.val {
			l.head, l.tail = nil, nil
			l.length--
			return nil
		}
	}
	if n.val == l.head.val {
		l.head = l.head.next
		l.length--
		return nil
	}
	current, prev := l.head.next, l.head
	for i := 1; i < l.length-1; i++ {
		if current.val == n.val {
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
