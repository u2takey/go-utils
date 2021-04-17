package assert

import (
	"errors"
	"testing"
)

func TestAssert(t *testing.T) {

	tests := []struct {
		name string
		err  error
		msg  string
	}{
		{
			"test1",
			errors.New("xxx"),
			"xxxx",
		},
	}
	//var a *struct{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//Assert(tt.err, tt.msg)
			//AssertNotNil(a, "a is nil")
		})
	}
}
