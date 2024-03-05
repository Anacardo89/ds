package arrlst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrlstlNew(t *testing.T) {
	a := New(10)
	assert.Equal(t, nil, a.arr[0])
}

func TestArrlstPrepend(t *testing.T) {
	a := New(10)
	a.Prepend(3)
	t.Run("Prepend()", func(t *testing.T) {
		assert.Equal(t, 3, a.arr[0])
	})
	a.Prepend(9)
	t.Run("Prepend() - 2nd time", func(t *testing.T) {
		assert.Equal(t, 9, a.arr[0])
	})
}

func TestArrlstgrow(t *testing.T) {
	a := New(2)
	a.Prepend(3)
	a.Prepend(6)
	t.Run("grow() - Before growth", func(t *testing.T) {
		assert.Equal(t, 2, a.len)
		assert.Equal(t, 2, a.cap)
	})
	a.Prepend(9)
	t.Run("grow() - After growth", func(t *testing.T) {
		assert.Equal(t, 3, a.len)
		assert.Equal(t, 3, a.cap)
	})
}

func TestArrlstInsertAt(t *testing.T) {
	a := New(10)
	t.Run("InsertAt() - Err: index out of bounds", func(t *testing.T) {
		err := a.InsertAt(4, 2)
		assert.Equal(t, ErrIdxOutOfBounds, err)
	})
	a.Prepend(13)
	a.Prepend(11)
	a.Prepend(9)
	a.Prepend(7)
	a.Prepend(5)
	a.Prepend(3)
	a.Prepend(1)
	t.Run("InsertAt() - idx == len", func(t *testing.T) {
		_ = a.InsertAt(15, a.len)
		assert.Equal(t, 13, a.arr[6])
		assert.Equal(t, 15, a.arr[7])
		assert.Equal(t, nil, a.arr[8])
	})
	t.Run("InsertAt() - 0 < idx < len", func(t *testing.T) {
		_ = a.InsertAt(4, 2)
		assert.Equal(t, 3, a.arr[1])
		assert.Equal(t, 4, a.arr[2])
		assert.Equal(t, 5, a.arr[3])
		assert.Equal(t, 15, a.arr[8])
		assert.Equal(t, nil, a.arr[9])
	})
	t.Run("InsertAt() - idx == 0", func(t *testing.T) {
		_ = a.InsertAt(0, 0)
		assert.Equal(t, 0, a.arr[0])
		assert.Equal(t, 1, a.arr[1])
		assert.Equal(t, 13, a.arr[8])
		assert.Equal(t, 15, a.arr[9])
	})
}

func TestArrlstAppend(t *testing.T) {
	a := New(10)
	a.Append(3)
	t.Run("Append()", func(t *testing.T) {
		assert.Equal(t, 3, a.arr[a.len-1])
	})
	a.Append(9)
	t.Run("Append() - 2nd time", func(t *testing.T) {
		assert.Equal(t, 9, a.arr[a.len-1])
	})
}

func TestArrlstRemove(t *testing.T) {
	a := New(10)
	a.Append(1)
	a.Append(3)
	a.Append(5)
	a.Append(7)
	a.Append(9)
	a.Append(11)
	a.Append(13)
	t.Run("Remove() - Err: value not found", func(t *testing.T) {
		_, err := a.Remove(15)
		assert.Equal(t, ErrNotFound, err)
	})
	t.Run("Remove() - idx == len-1", func(t *testing.T) {
		v, _ := a.Remove(13)
		assert.Equal(t, 13, v)
		assert.Equal(t, nil, a.arr[6])
	})
	t.Run("Remove() - 0 < idx < len-1", func(t *testing.T) {
		v, _ := a.Remove(5)
		assert.Equal(t, 5, v)
		assert.Equal(t, 3, a.arr[1])
		assert.Equal(t, 7, a.arr[2])
		assert.Equal(t, nil, a.arr[5])
	})
	t.Run("Remove() - idx == 0", func(t *testing.T) {
		v, _ := a.Remove(1)
		assert.Equal(t, 1, v)
		assert.Equal(t, 3, a.arr[0])
		assert.Equal(t, 7, a.arr[1])
		assert.Equal(t, nil, a.arr[4])
	})
}

func TestArrlstGet(t *testing.T) {
	a := New(10)
	a.Append(1)
	a.Append(3)
	a.Append(5)
	a.Append(7)
	a.Append(9)
	a.Append(11)
	a.Append(13)
	t.Run("Get() - Err: index out of bounds", func(t *testing.T) {
		_, err := a.Get(7)
		assert.Equal(t, ErrIdxOutOfBounds, err)
	})
	t.Run("Get() - idx == len-1", func(t *testing.T) {
		v, _ := a.Get(6)
		assert.Equal(t, 13, v)
	})
	t.Run("Get() - 0 < idx < len-1", func(t *testing.T) {
		v, _ := a.Get(2)
		assert.Equal(t, 5, v)
	})
	t.Run("Get() - idx == 0", func(t *testing.T) {
		v, _ := a.Get(0)
		assert.Equal(t, 1, v)
	})
}

func TestArrlstRemoveAt(t *testing.T) {
	a := New(10)
	a.Append(1)
	a.Append(3)
	a.Append(5)
	a.Append(7)
	a.Append(9)
	a.Append(11)
	a.Append(13)
	t.Run("RemoveAt() - Err: index out of bounds", func(t *testing.T) {
		_, err := a.RemoveAt(7)
		assert.Equal(t, ErrIdxOutOfBounds, err)
	})
	t.Run("RemoveAt() - idx == len-1", func(t *testing.T) {
		v, _ := a.RemoveAt(6)
		assert.Equal(t, 13, v)
		assert.Equal(t, nil, a.arr[6])
	})
	t.Run("RemoveAt() - 0 < idx < len-1", func(t *testing.T) {
		v, _ := a.RemoveAt(2)
		assert.Equal(t, 5, v)
		assert.Equal(t, 3, a.arr[1])
		assert.Equal(t, 7, a.arr[2])
		assert.Equal(t, nil, a.arr[5])
	})
	t.Run("RemoveAt() - idx == 0", func(t *testing.T) {
		v, _ := a.RemoveAt(0)
		assert.Equal(t, 1, v)
		assert.Equal(t, 3, a.arr[0])
		assert.Equal(t, 7, a.arr[1])
		assert.Equal(t, nil, a.arr[4])
	})
}

func TestUsage(t *testing.T) {
	a := New(3)
	a.Append(5)
	a.Append(7)
	a.Append(9)
	v, _ := a.Get(2)
	assert.Equal(t, 9, v)
	v, _ = a.RemoveAt(1)
	assert.Equal(t, 7, v)
	assert.Equal(t, 2, a.len)
	a.Append(11)
	v, _ = a.RemoveAt(1)
	assert.Equal(t, 9, v)
	_, err := a.Remove(9)
	assert.Equal(t, ErrNotFound, err)
	v, _ = a.RemoveAt(0)
	assert.Equal(t, 5, v)
	v, _ = a.RemoveAt(0)
	assert.Equal(t, 11, v)
	assert.Equal(t, 0, a.len)
	a.Prepend(5)
	a.Prepend(7)
	a.Prepend(9)
	v, _ = a.Get(2)
	assert.Equal(t, 5, v)
	v, _ = a.Get(0)
	assert.Equal(t, 9, v)
	v, _ = a.Remove(9)
	assert.Equal(t, 9, v)
	assert.Equal(t, 2, a.len)
	v, _ = a.Get(0)
	assert.Equal(t, 7, v)
}
