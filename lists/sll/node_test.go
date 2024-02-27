package sll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeNext(t *testing.T) {
	l := New()
	l.Append(10)
	current := l.head
	t.Run("Next() - Err: node not set", func(t *testing.T) {
		_, err := current.Next()
		assert.Equal(t, err, ErrNotSet)
	})
	l.Append(20)
	t.Run("Next()", func(t *testing.T) {
		next, _ := current.Next()
		assert.Equal(t, current.next.val, next.val)
	})
}

func TestNodeVal(t *testing.T) {
	l := New()
	l.Append(10)
	current := l.head
	t.Run("Val()", func(t *testing.T) {
		val := current.Val()
		assert.Equal(t, val, 10)
	})
}

func TestNodeSetVal(t *testing.T) {
	l := New()
	l.Append(10)
	current := l.head
	t.Run("Get Val", func(t *testing.T) {
		current.SetVal(30)
		assert.Equal(t, current.val, 30)
	})
}
