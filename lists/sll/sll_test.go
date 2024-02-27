package sll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSllNew(t *testing.T) {
	l := New()
	assert.Empty(t, l)
}

func TestSllLength(t *testing.T) {
	l := New()
	t.Run("Length() - 0", func(t *testing.T) {
		length := l.Length()
		assert.Equal(t, length, 0)
	})
	l.Append(1)
	t.Run("Length()", func(t *testing.T) {
		length := l.Length()
		assert.Equal(t, length, 1)
	})
}

func TestSllHead(t *testing.T) {
	l := New()
	t.Run("Head() - Err: empty list", func(t *testing.T) {
		_, err := l.Head()
		assert.Equal(t, err, ErrEmptyList)
	})
	l.Append(1)
	t.Run("Head()", func(t *testing.T) {
		head, _ := l.Head()
		assert.Equal(t, head, 1)
	})
}

func TestSllTail(t *testing.T) {
	l := New()
	t.Run("Tail() - Err: empty list", func(t *testing.T) {
		_, err := l.Tail()
		assert.Equal(t, err, ErrEmptyList)
	})
	l.Append(1)
	t.Run("Tail()", func(t *testing.T) {
		tail, _ := l.Tail()
		assert.Equal(t, tail, 1)
	})
}

func TestSllPrepend(t *testing.T) {
	l := New()
	t.Run("Prepend() - Length 0", func(t *testing.T) {
		l.Prepend(57)
		assert.Equal(t, l.head.val, 57)
	})
	t.Run("Prepend - Length 1", func(t *testing.T) {
		l.Prepend(49)
		assert.Equal(t, l.head.val, 49)
		assert.Equal(t, l.head.next.val, 57)
	})
	t.Run("Prepend() - Length > 1", func(t *testing.T) {
		l.Prepend(13)
		assert.Equal(t, l.head.val, 13)
		assert.Equal(t, l.head.next.val, 49)
	})
}

func TestSllAppend(t *testing.T) {
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

func TestSllInsertAt(t *testing.T) {
	l := New()
	t.Run("InsertAt() - Err: out of bounds", func(t *testing.T) {
		err := l.InsertAt(5, 10)
		assert.Equal(t, err, ErrOutOfBounds)
	})
	t.Run("InsertAt() - Length 0", func(t *testing.T) {
		_ = l.InsertAt(0, 10)
		assert.Equal(t, l.head.val, 10)
		assert.Equal(t, l.tail.val, 10)
	})
	t.Run("InsertAt() - Index 0", func(t *testing.T) {
		_ = l.InsertAt(0, 20)
		assert.Equal(t, l.head.val, 20)
		assert.Equal(t, l.head.next.val, 10)
		assert.Equal(t, l.tail.val, 10)
	})
	t.Run("InsertAt() - Index == Length", func(t *testing.T) {
		_ = l.InsertAt(2, 30)
		assert.Equal(t, l.tail.val, 30)
	})
	t.Run("InsertAt() - 0 < Index < Length", func(t *testing.T) {
		_ = l.InsertAt(1, 40)
		assert.Equal(t, l.head.val, 20)
		assert.Equal(t, l.head.next.val, 40)
		assert.Equal(t, l.head.next.next.val, 10)
		assert.Equal(t, l.tail.val, 30)
	})
}

func TestSllGetAt(t *testing.T) {
	l := New()
	t.Run("GetAt() - Err: empty list", func(t *testing.T) {
		_, err := l.GetAt(1)
		assert.Equal(t, err, ErrEmptyList)
	})
	l.Append(5)
	l.Append(10)
	l.Append(15)
	t.Run("GetAt() - Err: out of bounds", func(t *testing.T) {
		_, err := l.GetAt(3)
		assert.Equal(t, err, ErrOutOfBounds)
	})
	t.Run("GetAt() - Index 0", func(t *testing.T) {
		val, _ := l.GetAt(0)
		assert.Equal(t, val, 5)
	})
	t.Run("GetAt() - Index == Length-1", func(t *testing.T) {
		val, _ := l.GetAt(2)
		assert.Equal(t, val, 15)
	})
	t.Run("GetAt() - 0 < Index < Length", func(t *testing.T) {
		val, _ := l.GetAt(1)
		assert.Equal(t, val, 10)
	})
}

func TestSllRemove(t *testing.T) {
	l := New()
	t.Run("Remove() - Err: empty list", func(t *testing.T) {
		err := l.Remove(5)
		assert.Equal(t, err, ErrEmptyList)
	})
	l.Append(1)
	l.Append(2)
	t.Run("Remove() - Err: value not in list", func(t *testing.T) {
		err := l.Remove(5)
		assert.Equal(t, err, ErrNotInList)
	})
	t.Run("Remove() - Lenght > 1", func(t *testing.T) {
		l.Remove(1)
		assert.Equal(t, l.head.val, 2)
		assert.Equal(t, l.head.next, (*Node)(nil))
		assert.Equal(t, l.tail.next, (*Node)(nil))
	})
	t.Run("Remove() - Length 1", func(t *testing.T) {
		l.Remove(2)
		assert.Equal(t, l.head, (*Node)(nil))
		assert.Equal(t, l.tail, (*Node)(nil))
	})
}

func TestSllRemoveAt(t *testing.T) {
	l := New()
	t.Run("RemoveAt() - Err: empty list", func(t *testing.T) {
		err := l.RemoveAt(1)
		assert.Equal(t, err, ErrEmptyList)
	})
	l.Append(5)
	t.Run("RemoveAt() - Err: Out of bounds", func(t *testing.T) {
		err := l.RemoveAt(1)
		assert.Equal(t, err, ErrOutOfBounds)
	})
	t.Run("RemoveAt() - Length 1", func(t *testing.T) {
		l.RemoveAt(0)
		assert.Equal(t, l.head, (*Node)(nil))
		assert.Equal(t, l.tail, (*Node)(nil))
	})
	l.Append(5)
	l.Append(10)
	t.Run("RemoveAt() - Index 0", func(t *testing.T) {
		_ = l.RemoveAt(0)
		assert.Equal(t, l.head.val, 10)
		assert.Equal(t, l.tail.val, 10)
	})
	l.Append(20)
	l.Append(30)
	t.Run("RemoveAt() - Index == Length-1", func(t *testing.T) {
		l.RemoveAt(2)
		assert.Equal(t, l.tail.val, 20)
		assert.Equal(t, l.tail.next, (*Node)(nil))
	})
	l.Append(30)
	t.Run("RemoveAt() - 0 < Index < Length", func(t *testing.T) {
		_ = l.RemoveAt(1)
		assert.Equal(t, l.head.val, 10)
		assert.Equal(t, l.head.next.val, 30)
		assert.Equal(t, l.tail.val, 30)
	})
}
