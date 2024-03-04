package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueNew(t *testing.T) {
	q := New()
	assert.Empty(t, q)
}

func TestQueueLength(t *testing.T) {
	q := New()
	t.Run("Length() - 0", func(t *testing.T) {
		length := q.Length()
		assert.Equal(t, length, 0)
	})
	q.Enqueue(1)
	t.Run("Length()", func(t *testing.T) {
		length := q.Length()
		assert.Equal(t, length, 1)
	})
}

func TestQueueEnqueue(t *testing.T) {
	q := New()
	t.Run("Enqueue() - Length 0", func(t *testing.T) {
		q.Enqueue(57)
		assert.Equal(t, q.tail.val, 57)
	})
	t.Run("Enqueue() - Length 1", func(t *testing.T) {
		q.Enqueue(49)
		assert.Equal(t, q.tail.val, 49)
	})
	t.Run("Enqueue() - Length > 1", func(t *testing.T) {
		q.Enqueue(27)
		assert.Equal(t, q.tail.val, 27)
	})
}

func TestQueueDequeue(t *testing.T) {
	q := New()
	t.Run("Dequeue() - Err: empty list", func(t *testing.T) {
		_, err := q.Dequeue()
		assert.Equal(t, err, ErrEmptyQueue)
	})
	q.Enqueue(57)
	q.Enqueue(43)
	q.Enqueue(29)
	t.Run("Dequeue() - Length > 1", func(t *testing.T) {
		val, _ := q.Dequeue()
		assert.Equal(t, val, 57)
	})
	t.Run("Enqueue() - Length 1", func(t *testing.T) {
		val, _ := q.Dequeue()
		assert.Equal(t, val, 43)
	})
	t.Run("Enqueue() - Length 0", func(t *testing.T) {
		val, _ := q.Dequeue()
		assert.Equal(t, val, 29)
	})
}

func TestQueuePeek(t *testing.T) {
	q := New()
	q.Enqueue(57)
	q.Enqueue(43)
	q.Enqueue(29)
	t.Run("Peek()", func(t *testing.T) {
		val := q.Peek()
		assert.Equal(t, val, 57)
	})
	t.Run("Peek() - 2nd time", func(t *testing.T) {
		val := q.Peek()
		assert.Equal(t, val, 57)
	})
}
