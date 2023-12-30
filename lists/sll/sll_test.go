package sll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSLL_New(t *testing.T) {
	assert := assert.New(t)
	l := New()
	t.Run("Assert Empty SLL", func(t *testing.T) {
		t.Parallel()
		assert.Empty(l)
	})
	t.Run("Assert Length = 0", func(t *testing.T) {
		t.Parallel()
		assert.Equal(l.length, 0)
	})
}
