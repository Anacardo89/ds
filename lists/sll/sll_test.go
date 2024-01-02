package sll

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSLL_New(t *testing.T) {
	l := New()
	assert.Empty(t, l)
}

func TestSLL_GetLength(t *testing.T) {
	l := New()
	l.Append(1)

	t.Run("Gat Length 1", func(t *testing.T) {
		length := l.GetLength()
		assert.Equal(t, length, 1)
	})

	l.Append(1)

	t.Run("Gat Length 2", func(t *testing.T) {
		length := l.GetLength()
		assert.Equal(t, length, 2)
	})

	l.Remove(1)
	l.Remove(1)

	t.Run("Gat Length 0", func(t *testing.T) {
		length := l.GetLength()
		assert.Equal(t, length, 0)
	})
}

func TestSLL_GetHead(t *testing.T) {
	l := New()

	t.Run("Gat Head - No Node in List", func(t *testing.T) {
		_, err := l.GetHead()
		assert.Equal(t, err, errors.New("list has no nodes"))
	})

	l.Append(1)

	t.Run("Gat Head 1", func(t *testing.T) {
		head, _ := l.GetHead()
		assert.Equal(t, head, 1)
	})

	l.Prepend(2)

	t.Run("Gat Head 2", func(t *testing.T) {
		head, _ := l.GetHead()
		assert.Equal(t, head, 2)
	})

	l.Remove(2)

	t.Run("Gat Head After Remove", func(t *testing.T) {
		head, _ := l.GetHead()
		assert.Equal(t, head, 1)
	})
}

func TestSLL_GetTail(t *testing.T) {
	l := New()

	t.Run("Gat Tail - No Node in List", func(t *testing.T) {
		_, err := l.GetTail()
		assert.Equal(t, err, errors.New("list has no nodes"))
	})

	l.Append(1)

	t.Run("Gat Tail 1", func(t *testing.T) {
		tail, _ := l.GetTail()
		assert.Equal(t, tail, 1)
	})

	l.Append(2)

	t.Run("Gat Tail 2", func(t *testing.T) {
		tail, _ := l.GetTail()
		assert.Equal(t, tail, 2)
	})

	l.Remove(2)

	t.Run("Gat Tail After Remove", func(t *testing.T) {
		tail, _ := l.GetTail()
		assert.Equal(t, tail, 1)
	})
}

func TestSLL_Prepend(t *testing.T) {
	l := New()

	t.Run("Prepend 1", func(t *testing.T) {
		l.Prepend(1)
		assert.Equal(t, l.head.val, 1)
		assert.Equal(t, l.tail.val, 1)
	})

	t.Run("Prepend 2", func(t *testing.T) {
		l.Prepend(2)
		assert.Equal(t, l.head.val, 2)
		assert.Equal(t, l.tail.val, 1)
	})

	t.Run("Prepend 3", func(t *testing.T) {
		l.Prepend(3)
		assert.Equal(t, l.head.val, 3)
		assert.Equal(t, l.tail.val, 1)
	})
}

func TestSLL_Append(t *testing.T) {
	l := New()

	t.Run("Append 1", func(t *testing.T) {
		l.Append(1)
		assert.Equal(t, l.head.val, 1)
		assert.Equal(t, l.tail.val, 1)
	})

	t.Run("Append 2", func(t *testing.T) {
		l.Append(2)
		assert.Equal(t, l.head.val, 1)
		assert.Equal(t, l.tail.val, 2)
	})

	t.Run("Append 3", func(t *testing.T) {
		l.Append(3)
		assert.Equal(t, l.head.val, 1)
		assert.Equal(t, l.tail.val, 3)
	})
}

func TestSLL_Remove(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)

	t.Run("Remove First", func(t *testing.T) {
		l.Remove(1)
		assert.Equal(t, l.head.val, 2)
		assert.Equal(t, l.tail.val, 3)
	})

	t.Run("Remove Unexistent", func(t *testing.T) {
		_, err := l.Remove(5)
		assert.Equal(t, err, errors.New("value not in list"))
	})

	t.Run("Remove Last", func(t *testing.T) {
		l.Remove(3)
		assert.Equal(t, l.head.val, 2)
		assert.Equal(t, l.tail.val, 2)
	})

	t.Run("Remove Only", func(t *testing.T) {
		l.Remove(2)
		assert.Equal(t, l.head, (*node)(nil))
		assert.Equal(t, l.tail, (*node)(nil))
	})

	t.Run("Remove From Empty List", func(t *testing.T) {
		_, err := l.Remove(5)
		assert.Equal(t, err, errors.New("empty list"))
	})
}

func TestSLL_InsertAt(t *testing.T) {
	l := New()

	t.Run("Insert outside the list", func(t *testing.T) {
		err := l.InsertAt(5, 10)
		assert.Equal(t, err, errors.New("index exceeds length"))
	})

	t.Run("Insert first Value", func(t *testing.T) {
		_ = l.InsertAt(0, 10)
		assert.Equal(t, l.head.val, 10)
		assert.Equal(t, l.tail.val, 10)
	})

	t.Run("Insert first Index", func(t *testing.T) {
		_ = l.InsertAt(0, 20)
		assert.Equal(t, l.head.val, 20)
		assert.Equal(t, l.tail.val, 10)
	})

	t.Run("Insert last Index", func(t *testing.T) {
		_ = l.InsertAt(2, 30)
		assert.Equal(t, l.head.val, 20)
		assert.Equal(t, l.tail.val, 30)
	})

	t.Run("Insert middle Index", func(t *testing.T) {
		_ = l.InsertAt(1, 40)
		assert.Equal(t, l.head.val, 20)
		assert.Equal(t, l.tail.val, 30)
	})
}

func TestSLL_RemoveAt(t *testing.T) {
	l := New()

	t.Run("Remove from empty list", func(t *testing.T) {
		err := l.RemoveAt(1)
		assert.Equal(t, err, errors.New("empty list"))
	})

	l.Append(5)

	t.Run("Remove outside the list", func(t *testing.T) {
		err := l.RemoveAt(1)
		assert.Equal(t, err, errors.New("index exceeds length"))
	})

	t.Run("Remove only value", func(t *testing.T) {
		_ = l.RemoveAt(0)
		assert.Equal(t, l.head, (*node)(nil))
		assert.Equal(t, l.tail, (*node)(nil))
	})

	l.Append(5)
	l.Append(10)

	t.Run("Remove first Index", func(t *testing.T) {
		_ = l.RemoveAt(0)
		assert.Equal(t, l.head.val, 10)
		assert.Equal(t, l.tail.val, 10)
	})

	l.Append(20)
	l.Append(30)

	t.Run("Remove last Index", func(t *testing.T) {
		_ = l.RemoveAt(2)
		assert.Equal(t, l.head.val, 10)
		assert.Equal(t, l.tail.val, 20)
	})

	l.Append(30)

	t.Run("Remove middle Index", func(t *testing.T) {
		_ = l.RemoveAt(1)
		assert.Equal(t, l.head.val, 10)
		assert.Equal(t, l.tail.val, 30)
	})
}

func TestSLL_GetAt(t *testing.T) {
	l := New()

	t.Run("Get from empty list", func(t *testing.T) {
		_, err := l.GetAt(1)
		assert.Equal(t, err, errors.New("empty list"))
	})

	l.Append(5)
	l.Append(10)
	l.Append(15)

	t.Run("Get outside the list", func(t *testing.T) {
		_, err := l.GetAt(3)
		assert.Equal(t, err, errors.New("index exceeds length"))
	})

	t.Run("Get first Index", func(t *testing.T) {
		val, _ := l.GetAt(0)
		assert.Equal(t, val, 5)
	})

	t.Run("Remove last Index", func(t *testing.T) {
		val, _ := l.GetAt(2)
		assert.Equal(t, val, 15)
	})

	t.Run("Remove middle Index", func(t *testing.T) {
		val, _ := l.GetAt(1)
		assert.Equal(t, val, 10)
	})
}
