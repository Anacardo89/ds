package dll

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDLL_New(t *testing.T) {
	l := New()
	assert.Empty(t, l)
}

func TestDLL_GetLength(t *testing.T) {
	l := New()
	l.Append(1)

	t.Run("Get Length 1", func(t *testing.T) {
		length := l.GetLength()
		assert.Equal(t, length, 1)
	})

	l.Append(1)

	t.Run("Get Length 2", func(t *testing.T) {
		length := l.GetLength()
		assert.Equal(t, length, 2)
	})

	l.Remove(1)
	l.Remove(1)

	t.Run("Get Length 0", func(t *testing.T) {
		length := l.GetLength()
		assert.Equal(t, length, 0)
	})
}

func TestDLL_GetHead(t *testing.T) {
	l := New()

	t.Run("Get Head - No Node in List", func(t *testing.T) {
		_, err := l.GetHead()
		assert.Equal(t, err, errors.New("empty list"))
	})

	l.Append(1)

	t.Run("Get Head 1", func(t *testing.T) {
		head, _ := l.GetHead()
		assert.Equal(t, head, 1)
	})

	l.Prepend(2)

	t.Run("Get Head 2", func(t *testing.T) {
		head, _ := l.GetHead()
		assert.Equal(t, head, 2)
	})

	l.Remove(2)

	t.Run("Get Head After Remove", func(t *testing.T) {
		head, _ := l.GetHead()
		assert.Equal(t, head, 1)
	})
}

func TestDLL_GetTail(t *testing.T) {
	l := New()

	t.Run("Get Tail - No Node in List", func(t *testing.T) {
		_, err := l.GetTail()
		assert.Equal(t, err, errors.New("empty list"))
	})

	l.Append(1)

	t.Run("Get Tail 1", func(t *testing.T) {
		tail, _ := l.GetTail()
		assert.Equal(t, tail, 1)
	})

	l.Append(2)

	t.Run("Get Tail 2", func(t *testing.T) {
		tail, _ := l.GetTail()
		assert.Equal(t, tail, 2)
	})

	l.Remove(2)

	t.Run("Get Tail After Remove", func(t *testing.T) {
		tail, _ := l.GetTail()
		assert.Equal(t, tail, 1)
	})
}

func TestDLL_Prepend(t *testing.T) {
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

func TestDLL_Append(t *testing.T) {
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

func TestDLL_Remove(t *testing.T) {
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
		assert.Equal(t, l.head, (*Node)(nil))
		assert.Equal(t, l.tail, (*Node)(nil))
	})

	t.Run("Remove From Empty List", func(t *testing.T) {
		_, err := l.Remove(5)
		assert.Equal(t, err, errors.New("empty list"))
	})
}

func TestDLL_InsertAt(t *testing.T) {
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

func TestDLL_RemoveAt(t *testing.T) {
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
		assert.Equal(t, l.head, (*Node)(nil))
		assert.Equal(t, l.tail, (*Node)(nil))
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

func TestDLL_GetAt(t *testing.T) {
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

	t.Run("Get last Index", func(t *testing.T) {
		val, _ := l.GetAt(2)
		assert.Equal(t, val, 15)
	})

	t.Run("Get middle Index", func(t *testing.T) {
		val, _ := l.GetAt(1)
		assert.Equal(t, val, 10)
	})
}

func TestDLL_WalkTo(t *testing.T) {
	idx := 0
	l := New()

	t.Run("Walk in empty list", func(t *testing.T) {
		idx = 1
		_, err := l.WalkTo(idx)
		assert.Equal(t, err, errors.New("empty list"))
	})

	l.Append(5)
	l.Append(10)
	l.Append(15)

	t.Run("Walk outside the list", func(t *testing.T) {
		idx = 3
		_, err := l.WalkTo(idx)
		assert.Equal(t, err, errors.New("index exceeds length"))
	})

	t.Run("Walk to first Index", func(t *testing.T) {
		idx = 0
		Node, _ := l.WalkTo(idx)
		assert.Equal(t, Node, l.head)
	})

	t.Run("Walk to last Index", func(t *testing.T) {
		idx = 2
		Node, _ := l.WalkTo(idx)
		assert.Equal(t, Node, l.tail)
	})

	t.Run("Walk to middle Index", func(t *testing.T) {
		idx = 1
		Node, _ := l.WalkTo(idx)
		current := l.head
		for i := 0; i < idx; i++ {
			current = current.next
		}
		assert.Equal(t, Node, current)
	})
}
