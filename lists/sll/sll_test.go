package sll

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSLL_New(t *testing.T) {
	l := New()
	t.Parallel()
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

}

func TestSLL_RemoveAt(t *testing.T) {

}

func TestSLL_GetAt(t *testing.T) {

}
