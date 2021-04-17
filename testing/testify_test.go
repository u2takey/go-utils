package testing

import (
	"errors"
	"testing"
)

func TestAssert(t *testing.T) {
	assert := NewAssert(t)
	a, b, e := Fun("a", "b", nil)
	assert.Equal(a, "a").Equal(b, "b").Nil(e)

	a, b, e = Fun("a", "b", errors.New("ss"))
	assert.NotNil(e).Equal(b, "b").Equal(a, "a")
}

func Fun(a, b string, err error) (string, string, error) {
	return a, b, err
}
