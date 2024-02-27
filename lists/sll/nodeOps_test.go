package sll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSllHeadNode(t *testing.T) {
	l := New()
	t.Run("HeadNode()  - Err: empty list", func(t *testing.T) {
		_, err := l.HeadNode()
		assert.Equal(t, err, ErrEmptyList)
	})
	n := &Node{val: 35}
	t.Run("HeadNode()", func(t *testing.T) {
		l.PrependNode(n)
		assert.Equal(t, l.head.val, 35)
	})
}

func TestSllTailNode(t *testing.T) {
	l := New()
	t.Run("TailNode()  - Err: empty list", func(t *testing.T) {
		_, err := l.HeadNode()
		assert.Equal(t, err, ErrEmptyList)
	})
	n := &Node{val: 35}
	t.Run("TailNode()", func(t *testing.T) {
		l.PrependNode(n)
		assert.Equal(t, l.tail.val, 35)
	})
}

func TestSllPrependNode(t *testing.T) {
	l := New()
	n := &Node{val: 21}
	t.Run("Prepend()  - Length 0", func(t *testing.T) {
		l.PrependNode(n)
		assert.Equal(t, l.head.val, 21)
	})
	n = &Node{val: 35}
	t.Run("Prepend() - Length 1", func(t *testing.T) {
		l.PrependNode(n)
		assert.Equal(t, l.head.val, 35)
		assert.Equal(t, l.head.next.val, 21)
	})
	n = &Node{val: 15}
	t.Run("Prepend() - Length > 1", func(t *testing.T) {
		l.PrependNode(n)
		assert.Equal(t, l.head.val, 15)
		assert.Equal(t, l.head.next.val, 35)
	})
}

func TestSllAppendNode(t *testing.T) {
	l := New()
	t.Run("Append() - Length 0", func(t *testing.T) {
		l.Append(57)
		assert.Equal(t, l.tail.val, 57)
	})
	t.Run("Append() - Length 1", func(t *testing.T) {
		l.Append(49)
		assert.Equal(t, l.tail.val, 49)
	})
	t.Run("Append() - Length > 1", func(t *testing.T) {
		l.Append(27)
		assert.Equal(t, l.tail.val, 27)
	})
}

func TestSllInsertNodeAt(t *testing.T) {
	l := New()
	n := &Node{val: 10}
	t.Run("InsertAt() - Err: out of bounds", func(t *testing.T) {
		err := l.InsertNodeAt(5, n)
		assert.Equal(t, err, ErrOutOfBounds)
	})
	n = &Node{val: 10}
	t.Run("InsertAt() - Length 0", func(t *testing.T) {
		_ = l.InsertNodeAt(0, n)
		assert.Equal(t, l.head.val, 10)
		assert.Equal(t, l.tail.val, 10)
	})
	n = &Node{val: 20}
	t.Run("InsertAt() - Index 0", func(t *testing.T) {
		_ = l.InsertNodeAt(0, n)
		assert.Equal(t, l.head.val, 20)
		assert.Equal(t, l.head.next.val, 10)
		assert.Equal(t, l.tail.val, 10)
	})
	n = &Node{val: 30}
	t.Run("InsertAt() - Index == Length", func(t *testing.T) {
		_ = l.InsertNodeAt(2, n)
		assert.Equal(t, l.tail.val, 30)
	})
	n = &Node{val: 40}
	t.Run("InsertAt() - 0 < Index < Length ", func(t *testing.T) {
		_ = l.InsertNodeAt(1, n)
		assert.Equal(t, l.head.val, 20)
		assert.Equal(t, l.head.next.val, 40)
		assert.Equal(t, l.head.next.next.val, 10)
		assert.Equal(t, l.tail.val, 30)
	})
}

func TestSllGetNodeAt(t *testing.T) {
	l := New()
	t.Run("GetAt() - Err: empty list", func(t *testing.T) {
		_, err := l.GetNodeAt(1)
		assert.Equal(t, err, ErrEmptyList)
	})
	l.Append(5)
	l.Append(10)
	l.Append(15)
	t.Run("GetAt() - Err: out of bounds", func(t *testing.T) {
		_, err := l.GetNodeAt(3)
		assert.Equal(t, err, ErrOutOfBounds)
	})
	t.Run("GetAt() - Index 0", func(t *testing.T) {
		n, _ := l.GetNodeAt(0)
		assert.Equal(t, n, l.head)
	})
	t.Run("GetAt() - Index == Length-1", func(t *testing.T) {
		n, _ := l.GetNodeAt(2)
		assert.Equal(t, n, l.tail)
	})
	t.Run("GetAt() - 0 < Index < Length", func(t *testing.T) {
		n, _ := l.GetNodeAt(1)
		assert.Equal(t, n, l.head.next)
	})
}

func TestSllRemoveNode(t *testing.T) {
	l := New()
	n := &Node{val: 5}
	t.Run("RemoveNode() - Err: empty list", func(t *testing.T) {
		err := l.RemoveNode(n)
		assert.Equal(t, err, ErrEmptyList)
	})
	l.Append(1)
	l.Append(2)
	t.Run("RemoveNode() - Err: value not in list", func(t *testing.T) {
		err := l.RemoveNode(n)
		assert.Equal(t, err, ErrNotInList)
	})
	n = &Node{val: 1}
	t.Run("RemoveNode() - Lenght > 1", func(t *testing.T) {
		l.RemoveNode(n)
		assert.Equal(t, l.head.val, 2)
		assert.Equal(t, l.head.next, (*Node)(nil))
		assert.Equal(t, l.tail.next, (*Node)(nil))
	})
	n = &Node{val: 2}
	t.Run("RemoveNode() - Length 1", func(t *testing.T) {
		l.RemoveNode(n)
		assert.Equal(t, l.head, (*Node)(nil))
		assert.Equal(t, l.tail, (*Node)(nil))
	})
}
