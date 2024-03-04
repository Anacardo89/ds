package queue

import "errors"

var ErrEmptyQueue = errors.New("cannot dequeue an empty queue")

type node struct {
	next *node
	val  interface{}
}

type Queue struct {
	length int
	head   *node
	tail   *node
}

func New() Queue {
	return Queue{
		length: 0,
	}
}

func (q *Queue) Length() int {
	return q.length
}

func (q *Queue) Enqueue(val interface{}) {
	n := node{val: val}
	if q.length == 0 {
		q.head, q.tail = &n, &n
	} else {
		q.tail.next = &n
		q.tail = &n
	}
	q.length++
}

func (q *Queue) Dequeue() (interface{}, error) {
	var val interface{}
	switch q.length {
	case 0:
		return nil, ErrEmptyQueue
	case 1:
		val = q.head.val
		q.head, q.tail = nil, nil
	default:
		val = q.head.val
		q.head = q.head.next
	}
	q.length--
	return val, nil
}

func (q *Queue) Peek() interface{} {
	return q.head.val
}
