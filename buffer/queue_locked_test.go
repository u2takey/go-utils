package buffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLockQueue(t *testing.T) {
	a := NewLockQueue(1)
	for i := 0; i < 10; i++ {
		a.Put(i)
	}
	read := 10
	for {
		b, ok := a.Get()
		if !ok {
			break
		}
		read--
		assert.Equal(t, read, b)
	}
}
