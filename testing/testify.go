package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 对 https://github.com/stretchr/testify 的简单包装，支持链式调用
// _codegen  -output-package=testing -template=testify_gen.go.tmpl -include-format-funcs
type Assertions struct {
	A *assert.Assertions
	R []bool
}

func NewAssert(t *testing.T) *Assertions {
	return &Assertions{A: assert.New(t)}
}
