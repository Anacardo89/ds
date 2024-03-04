package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackNew(t *testing.T) {
	s := New()
	assert.Empty(t, s)
}

func TestStackLength(t *testing.T) {
	s := New()
	t.Run("Length() - 0", func(t *testing.T) {
		length := s.Length()
		assert.Equal(t, length, 0)
	})
	s.Push(1)
	t.Run("Length()", func(t *testing.T) {
		length := s.Length()
		assert.Equal(t, length, 1)
	})
}

func TestStackPush(t *testing.T) {
	s := New()
	t.Run("Push() - Length 0", func(t *testing.T) {
		s.Push(57)
		assert.Equal(t, s.top.val, 57)
	})
	t.Run("Push() - Length 1", func(t *testing.T) {
		s.Push(49)
		assert.Equal(t, s.top.val, 49)
	})
	t.Run("Enqueue() - Length > 1", func(t *testing.T) {
		s.Push(27)
		assert.Equal(t, s.top.val, 27)
	})
}

func TestStackPop(t *testing.T) {
	q := New()
	t.Run("Pop() - Err: empty stack", func(t *testing.T) {
		_, err := q.Pop()
		assert.Equal(t, err, ErrEmptyStack)
	})
	q.Push(57)
	q.Push(43)
	q.Push(29)
	t.Run("Pop()", func(t *testing.T) {
		val, _ := q.Pop()
		assert.Equal(t, val, 29)
	})
	t.Run("Pop() - 2nd time", func(t *testing.T) {
		val, _ := q.Pop()
		assert.Equal(t, val, 43)
	})
}

func TestStackPeek(t *testing.T) {
	s := New()
	t.Run("Peek() - Err: empty stack", func(t *testing.T) {
		_, err := s.Peek()
		assert.Equal(t, err, ErrEmptyStack)
	})
	s.Push(57)
	s.Push(43)
	s.Push(29)
	t.Run("Peek()", func(t *testing.T) {
		val, _ := s.Peek()
		assert.Equal(t, val, 29)
	})
	t.Run("Peek() - 2nd time", func(t *testing.T) {
		val, _ := s.Peek()
		assert.Equal(t, val, 29)
	})
}
