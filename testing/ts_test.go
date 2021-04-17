package testing

import "testing"

type TestA struct {
	A string
}

func (a TestA) FuncB() string {
	return "b"
}

func TestTsWithExamples(t *testing.T) {
	ts := New(t)

	ts.Case("show case usage 1", map[string]interface{}{"A": "b"}).PropEqual("A", "b")
	ts.Case("show case usage 2", struct{ A string }{A: "b"}).PropEqual("A", "b")
	ts.Case("show case usage 3", TestA{}).PropEqual("FuncB", "b")

	ts.Call(func(a string) *TestA { return &TestA{A: a} }, "b").PropEqual("A", "b")
	ts.Call(func(a string) (*TestA, error) { return &TestA{A: a}, nil }, "b").PropEqual("A", "b").With(1).Equal(nil)
}
