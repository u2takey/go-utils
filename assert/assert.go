package assert

import (
	"fmt"
	"reflect"
)

// Assert if err!=nil panic with error and msg if not niu.
func Assert(err error, format string, a ...interface{}) {
	if err == nil {
		return
	}
	s := fmt.Sprintf("check error failed, error: %s", err)
	if len(format) > 0 {
		s += fmt.Sprintf(", message: "+format, a...)
	}
	panic(s)
}

func AssertNotNil(c interface{}, format string, a ...interface{}) {
	if c == nil || (reflect.ValueOf(c).Kind() == reflect.Ptr && reflect.ValueOf(c).IsNil()) {
		panic(fmt.Sprintf("unexpected nil object"+format, a...))
	}
}

func AssertAllNotNil(errorMessage string, cs ...interface{}) {
	for c := range cs {
		AssertNotNil(c, errorMessage)
	}
}
