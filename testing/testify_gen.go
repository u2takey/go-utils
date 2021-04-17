/*
* CODE GENERATED AUTOMATICALLY WITH github.com/stretchr/testify/_codegen
* THIS FILE MUST NOT BE EDITED BY HAND
 */

package testing

import (
	"net/http"
	"net/url"
	"time"

	"github.com/stretchr/testify/assert"
)

// Condition uses a Comparison to assert a complex condition.
func (a *Assertions) Condition(comp assert.Comparison, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Condition(comp, msgAndArgs...))
	return a
}

// Conditionf uses a Comparison to assert a complex condition.
func (a *Assertions) Conditionf(comp assert.Comparison, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Conditionf(comp, msg, args...))
	return a
}

// Contains asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
//    a.Contains("Hello World", "World")
//    a.Contains(["Hello", "World"], "World")
//    a.Contains({"Hello": "World"}, "Hello")
func (a *Assertions) Contains(s interface{}, contains interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Contains(s, contains, msgAndArgs...))
	return a
}

// Containsf asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
//    a.Containsf("Hello World", "World", "error message %s", "formatted")
//    a.Containsf(["Hello", "World"], "World", "error message %s", "formatted")
//    a.Containsf({"Hello": "World"}, "Hello", "error message %s", "formatted")
func (a *Assertions) Containsf(s interface{}, contains interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Containsf(s, contains, msg, args...))
	return a
}

// DirExists checks whether a directory exists in the given path. It also fails
// if the path is a file rather a directory or there is an error checking whether it exists.
func (a *Assertions) DirExists(path string, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.DirExists(path, msgAndArgs...))
	return a
}

// DirExistsf checks whether a directory exists in the given path. It also fails
// if the path is a file rather a directory or there is an error checking whether it exists.
func (a *Assertions) DirExistsf(path string, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.DirExistsf(path, msg, args...))
	return a
}

// ElementsMatch asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// a.ElementsMatch([1, 3, 2, 3], [1, 3, 3, 2])
func (a *Assertions) ElementsMatch(listA interface{}, listB interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.ElementsMatch(listA, listB, msgAndArgs...))
	return a
}

// ElementsMatchf asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// a.ElementsMatchf([1, 3, 2, 3], [1, 3, 3, 2], "error message %s", "formatted")
func (a *Assertions) ElementsMatchf(listA interface{}, listB interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.ElementsMatchf(listA, listB, msg, args...))
	return a
}

// Empty asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  a.Empty(obj)
func (a *Assertions) Empty(object interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Empty(object, msgAndArgs...))
	return a
}

// Emptyf asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  a.Emptyf(obj, "error message %s", "formatted")
func (a *Assertions) Emptyf(object interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Emptyf(object, msg, args...))
	return a
}

// Equal asserts that two objects are equal.
//
//    a.Equal(123, 123)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses). Function equality
// cannot be determined and will always fail.
func (a *Assertions) Equal(expected interface{}, actual interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Equal(expected, actual, msgAndArgs...))
	return a
}

// EqualError asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
//   actualObj, err := SomeFunction()
//   a.EqualError(err,  expectedErrorString)
func (a *Assertions) EqualError(theError error, errString string, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.EqualError(theError, errString, msgAndArgs...))
	return a
}

// EqualErrorf asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
//   actualObj, err := SomeFunction()
//   a.EqualErrorf(err,  expectedErrorString, "error message %s", "formatted")
func (a *Assertions) EqualErrorf(theError error, errString string, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.EqualErrorf(theError, errString, msg, args...))
	return a
}

// EqualValues asserts that two objects are equal or convertable to the same types
// and equal.
//
//    a.EqualValues(uint32(123), int32(123))
func (a *Assertions) EqualValues(expected interface{}, actual interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.EqualValues(expected, actual, msgAndArgs...))
	return a
}

// EqualValuesf asserts that two objects are equal or convertable to the same types
// and equal.
//
//    a.EqualValuesf(uint32(123, "error message %s", "formatted"), int32(123))
func (a *Assertions) EqualValuesf(expected interface{}, actual interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.EqualValuesf(expected, actual, msg, args...))
	return a
}

// Equalf asserts that two objects are equal.
//
//    a.Equalf(123, 123, "error message %s", "formatted")
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses). Function equality
// cannot be determined and will always fail.
func (a *Assertions) Equalf(expected interface{}, actual interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Equalf(expected, actual, msg, args...))
	return a
}

// Error asserts that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   if a.Error(err) {
// 	   assert.Equal(t, expectedError, err)
//   }
func (a *Assertions) Error(err error, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Error(err, msgAndArgs...))
	return a
}

// Errorf asserts that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   if a.Errorf(err, "error message %s", "formatted") {
// 	   assert.Equal(t, expectedErrorf, err)
//   }
func (a *Assertions) Errorf(err error, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Errorf(err, msg, args...))
	return a
}

// Eventually asserts that given condition will be met in waitFor time,
// periodically checking target function each tick.
//
//    a.Eventually(func() bool { return true; }, time.Second, 10*time.Millisecond)
func (a *Assertions) Eventually(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Eventually(condition, waitFor, tick, msgAndArgs...))
	return a
}

// Eventuallyf asserts that given condition will be met in waitFor time,
// periodically checking target function each tick.
//
//    a.Eventuallyf(func() bool { return true; }, time.Second, 10*time.Millisecond, "error message %s", "formatted")
func (a *Assertions) Eventuallyf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Eventuallyf(condition, waitFor, tick, msg, args...))
	return a
}

// Exactly asserts that two objects are equal in value and type.
//
//    a.Exactly(int32(123), int64(123))
func (a *Assertions) Exactly(expected interface{}, actual interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Exactly(expected, actual, msgAndArgs...))
	return a
}

// Exactlyf asserts that two objects are equal in value and type.
//
//    a.Exactlyf(int32(123, "error message %s", "formatted"), int64(123))
func (a *Assertions) Exactlyf(expected interface{}, actual interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Exactlyf(expected, actual, msg, args...))
	return a
}

// Fail reports a failure through
func (a *Assertions) Fail(failureMessage string, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Fail(failureMessage, msgAndArgs...))
	return a
}

// FailNow fails test
func (a *Assertions) FailNow(failureMessage string, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.FailNow(failureMessage, msgAndArgs...))
	return a
}

// FailNowf fails test
func (a *Assertions) FailNowf(failureMessage string, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.FailNowf(failureMessage, msg, args...))
	return a
}

// Failf reports a failure through
func (a *Assertions) Failf(failureMessage string, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Failf(failureMessage, msg, args...))
	return a
}

// False asserts that the specified value is false.
//
//    a.False(myBool)
func (a *Assertions) False(value bool, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.False(value, msgAndArgs...))
	return a
}

// Falsef asserts that the specified value is false.
//
//    a.Falsef(myBool, "error message %s", "formatted")
func (a *Assertions) Falsef(value bool, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Falsef(value, msg, args...))
	return a
}

// FileExists checks whether a file exists in the given path. It also fails if
// the path points to a directory or there is an error when trying to check the file.
func (a *Assertions) FileExists(path string, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.FileExists(path, msgAndArgs...))
	return a
}

// FileExistsf checks whether a file exists in the given path. It also fails if
// the path points to a directory or there is an error when trying to check the file.
func (a *Assertions) FileExistsf(path string, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.FileExistsf(path, msg, args...))
	return a
}

// Greater asserts that the first element is greater than the second
//
//    a.Greater(2, 1)
//    a.Greater(float64(2), float64(1))
//    a.Greater("b", "a")
func (a *Assertions) Greater(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Greater(e1, e2, msgAndArgs...))
	return a
}

// GreaterOrEqual asserts that the first element is greater than or equal to the second
//
//    a.GreaterOrEqual(2, 1)
//    a.GreaterOrEqual(2, 2)
//    a.GreaterOrEqual("b", "a")
//    a.GreaterOrEqual("b", "b")
func (a *Assertions) GreaterOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.GreaterOrEqual(e1, e2, msgAndArgs...))
	return a
}

// GreaterOrEqualf asserts that the first element is greater than or equal to the second
//
//    a.GreaterOrEqualf(2, 1, "error message %s", "formatted")
//    a.GreaterOrEqualf(2, 2, "error message %s", "formatted")
//    a.GreaterOrEqualf("b", "a", "error message %s", "formatted")
//    a.GreaterOrEqualf("b", "b", "error message %s", "formatted")
func (a *Assertions) GreaterOrEqualf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.GreaterOrEqualf(e1, e2, msg, args...))
	return a
}

// Greaterf asserts that the first element is greater than the second
//
//    a.Greaterf(2, 1, "error message %s", "formatted")
//    a.Greaterf(float64(2, "error message %s", "formatted"), float64(1))
//    a.Greaterf("b", "a", "error message %s", "formatted")
func (a *Assertions) Greaterf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Greaterf(e1, e2, msg, args...))
	return a
}

// HTTPBodyContains asserts that a specified handler returns a
// body that contains a string.
//
//  a.HTTPBodyContains(myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPBodyContains(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.HTTPBodyContains(handler, method, url, values, str, msgAndArgs...))
	return a
}

// HTTPBodyContainsf asserts that a specified handler returns a
// body that contains a string.
//
//  a.HTTPBodyContainsf(myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky", "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPBodyContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.HTTPBodyContainsf(handler, method, url, values, str, msg, args...))
	return a
}

// HTTPBodyNotContains asserts that a specified handler returns a
// body that does not contain a string.
//
//  a.HTTPBodyNotContains(myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPBodyNotContains(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.HTTPBodyNotContains(handler, method, url, values, str, msgAndArgs...))
	return a
}

// HTTPBodyNotContainsf asserts that a specified handler returns a
// body that does not contain a string.
//
//  a.HTTPBodyNotContainsf(myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky", "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPBodyNotContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.HTTPBodyNotContainsf(handler, method, url, values, str, msg, args...))
	return a
}

// HTTPError asserts that a specified handler returns an error status code.
//
//  a.HTTPError(myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPError(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.HTTPError(handler, method, url, values, msgAndArgs...))
	return a
}

// HTTPErrorf asserts that a specified handler returns an error status code.
//
//  a.HTTPErrorf(myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true, "error message %s", "formatted") or not (false).
func (a *Assertions) HTTPErrorf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.HTTPErrorf(handler, method, url, values, msg, args...))
	return a
}

// HTTPRedirect asserts that a specified handler returns a redirect status code.
//
//  a.HTTPRedirect(myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPRedirect(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.HTTPRedirect(handler, method, url, values, msgAndArgs...))
	return a
}

// HTTPRedirectf asserts that a specified handler returns a redirect status code.
//
//  a.HTTPRedirectf(myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true, "error message %s", "formatted") or not (false).
func (a *Assertions) HTTPRedirectf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.HTTPRedirectf(handler, method, url, values, msg, args...))
	return a
}

// HTTPSuccess asserts that a specified handler returns a success status code.
//
//  a.HTTPSuccess(myHandler, "POST", "http://www.google.com", nil)
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPSuccess(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.HTTPSuccess(handler, method, url, values, msgAndArgs...))
	return a
}

// HTTPSuccessf asserts that a specified handler returns a success status code.
//
//  a.HTTPSuccessf(myHandler, "POST", "http://www.google.com", nil, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPSuccessf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.HTTPSuccessf(handler, method, url, values, msg, args...))
	return a
}

// Implements asserts that an object is implemented by the specified interface.
//
//    a.Implements((*MyInterface)(nil), new(MyObject))
func (a *Assertions) Implements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Implements(interfaceObject, object, msgAndArgs...))
	return a
}

// Implementsf asserts that an object is implemented by the specified interface.
//
//    a.Implementsf((*MyInterface, "error message %s", "formatted")(nil), new(MyObject))
func (a *Assertions) Implementsf(interfaceObject interface{}, object interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Implementsf(interfaceObject, object, msg, args...))
	return a
}

// InDelta asserts that the two numerals are within delta of each other.
//
// 	 a.InDelta(math.Pi, 22/7.0, 0.01)
func (a *Assertions) InDelta(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.InDelta(expected, actual, delta, msgAndArgs...))
	return a
}

// InDeltaMapValues is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.
func (a *Assertions) InDeltaMapValues(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.InDeltaMapValues(expected, actual, delta, msgAndArgs...))
	return a
}

// InDeltaMapValuesf is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.
func (a *Assertions) InDeltaMapValuesf(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.InDeltaMapValuesf(expected, actual, delta, msg, args...))
	return a
}

// InDeltaSlice is the same as InDelta, except it compares two slices.
func (a *Assertions) InDeltaSlice(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.InDeltaSlice(expected, actual, delta, msgAndArgs...))
	return a
}

// InDeltaSlicef is the same as InDelta, except it compares two slices.
func (a *Assertions) InDeltaSlicef(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.InDeltaSlicef(expected, actual, delta, msg, args...))
	return a
}

// InDeltaf asserts that the two numerals are within delta of each other.
//
// 	 a.InDeltaf(math.Pi, 22/7.0, 0.01, "error message %s", "formatted")
func (a *Assertions) InDeltaf(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.InDeltaf(expected, actual, delta, msg, args...))
	return a
}

// InEpsilon asserts that expected and actual have a relative error less than epsilon
func (a *Assertions) InEpsilon(expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.InEpsilon(expected, actual, epsilon, msgAndArgs...))
	return a
}

// InEpsilonSlice is the same as InEpsilon, except it compares each value from two slices.
func (a *Assertions) InEpsilonSlice(expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.InEpsilonSlice(expected, actual, epsilon, msgAndArgs...))
	return a
}

// InEpsilonSlicef is the same as InEpsilon, except it compares each value from two slices.
func (a *Assertions) InEpsilonSlicef(expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.InEpsilonSlicef(expected, actual, epsilon, msg, args...))
	return a
}

// InEpsilonf asserts that expected and actual have a relative error less than epsilon
func (a *Assertions) InEpsilonf(expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.InEpsilonf(expected, actual, epsilon, msg, args...))
	return a
}

// IsType asserts that the specified objects are of the same type.
func (a *Assertions) IsType(expectedType interface{}, object interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.IsType(expectedType, object, msgAndArgs...))
	return a
}

// IsTypef asserts that the specified objects are of the same type.
func (a *Assertions) IsTypef(expectedType interface{}, object interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.IsTypef(expectedType, object, msg, args...))
	return a
}

// JSONEq asserts that two JSON strings are equivalent.
//
//  a.JSONEq(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
func (a *Assertions) JSONEq(expected string, actual string, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.JSONEq(expected, actual, msgAndArgs...))
	return a
}

// JSONEqf asserts that two JSON strings are equivalent.
//
//  a.JSONEqf(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`, "error message %s", "formatted")
func (a *Assertions) JSONEqf(expected string, actual string, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.JSONEqf(expected, actual, msg, args...))
	return a
}

// Len asserts that the specified object has specific length.
// Len also fails if the object has a type that len() not accept.
//
//    a.Len(mySlice, 3)
func (a *Assertions) Len(object interface{}, length int, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Len(object, length, msgAndArgs...))
	return a
}

// Lenf asserts that the specified object has specific length.
// Lenf also fails if the object has a type that len() not accept.
//
//    a.Lenf(mySlice, 3, "error message %s", "formatted")
func (a *Assertions) Lenf(object interface{}, length int, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Lenf(object, length, msg, args...))
	return a
}

// Less asserts that the first element is less than the second
//
//    a.Less(1, 2)
//    a.Less(float64(1), float64(2))
//    a.Less("a", "b")
func (a *Assertions) Less(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Less(e1, e2, msgAndArgs...))
	return a
}

// LessOrEqual asserts that the first element is less than or equal to the second
//
//    a.LessOrEqual(1, 2)
//    a.LessOrEqual(2, 2)
//    a.LessOrEqual("a", "b")
//    a.LessOrEqual("b", "b")
func (a *Assertions) LessOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.LessOrEqual(e1, e2, msgAndArgs...))
	return a
}

// LessOrEqualf asserts that the first element is less than or equal to the second
//
//    a.LessOrEqualf(1, 2, "error message %s", "formatted")
//    a.LessOrEqualf(2, 2, "error message %s", "formatted")
//    a.LessOrEqualf("a", "b", "error message %s", "formatted")
//    a.LessOrEqualf("b", "b", "error message %s", "formatted")
func (a *Assertions) LessOrEqualf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.LessOrEqualf(e1, e2, msg, args...))
	return a
}

// Lessf asserts that the first element is less than the second
//
//    a.Lessf(1, 2, "error message %s", "formatted")
//    a.Lessf(float64(1, "error message %s", "formatted"), float64(2))
//    a.Lessf("a", "b", "error message %s", "formatted")
func (a *Assertions) Lessf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Lessf(e1, e2, msg, args...))
	return a
}

// Never asserts that the given condition doesn't satisfy in waitFor time,
// periodically checking the target function each tick.
//
//    a.Never(func() bool { return false; }, time.Second, 10*time.Millisecond)
func (a *Assertions) Never(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Never(condition, waitFor, tick, msgAndArgs...))
	return a
}

// Neverf asserts that the given condition doesn't satisfy in waitFor time,
// periodically checking the target function each tick.
//
//    a.Neverf(func() bool { return false; }, time.Second, 10*time.Millisecond, "error message %s", "formatted")
func (a *Assertions) Neverf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Neverf(condition, waitFor, tick, msg, args...))
	return a
}

// Nil asserts that the specified object is nil.
//
//    a.Nil(err)
func (a *Assertions) Nil(object interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Nil(object, msgAndArgs...))
	return a
}

// Nilf asserts that the specified object is nil.
//
//    a.Nilf(err, "error message %s", "formatted")
func (a *Assertions) Nilf(object interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Nilf(object, msg, args...))
	return a
}

// NoDirExists checks whether a directory does not exist in the given path.
// It fails if the path points to an existing _directory_ only.
func (a *Assertions) NoDirExists(path string, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.NoDirExists(path, msgAndArgs...))
	return a
}

// NoDirExistsf checks whether a directory does not exist in the given path.
// It fails if the path points to an existing _directory_ only.
func (a *Assertions) NoDirExistsf(path string, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.NoDirExistsf(path, msg, args...))
	return a
}

// NoError asserts that a function returned no error (i.e. `nil`).
//
//   actualObj, err := SomeFunction()
//   if a.NoError(err) {
// 	   assert.Equal(t, expectedObj, actualObj)
//   }
func (a *Assertions) NoError(err error, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.NoError(err, msgAndArgs...))
	return a
}

// NoErrorf asserts that a function returned no error (i.e. `nil`).
//
//   actualObj, err := SomeFunction()
//   if a.NoErrorf(err, "error message %s", "formatted") {
// 	   assert.Equal(t, expectedObj, actualObj)
//   }
func (a *Assertions) NoErrorf(err error, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.NoErrorf(err, msg, args...))
	return a
}

// NoFileExists checks whether a file does not exist in a given path. It fails
// if the path points to an existing _file_ only.
func (a *Assertions) NoFileExists(path string, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.NoFileExists(path, msgAndArgs...))
	return a
}

// NoFileExistsf checks whether a file does not exist in a given path. It fails
// if the path points to an existing _file_ only.
func (a *Assertions) NoFileExistsf(path string, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.NoFileExistsf(path, msg, args...))
	return a
}

// NotContains asserts that the specified string, list(array, slice...) or map does NOT contain the
// specified substring or element.
//
//    a.NotContains("Hello World", "Earth")
//    a.NotContains(["Hello", "World"], "Earth")
//    a.NotContains({"Hello": "World"}, "Earth")
func (a *Assertions) NotContains(s interface{}, contains interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.NotContains(s, contains, msgAndArgs...))
	return a
}

// NotContainsf asserts that the specified string, list(array, slice...) or map does NOT contain the
// specified substring or element.
//
//    a.NotContainsf("Hello World", "Earth", "error message %s", "formatted")
//    a.NotContainsf(["Hello", "World"], "Earth", "error message %s", "formatted")
//    a.NotContainsf({"Hello": "World"}, "Earth", "error message %s", "formatted")
func (a *Assertions) NotContainsf(s interface{}, contains interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.NotContainsf(s, contains, msg, args...))
	return a
}

// NotEmpty asserts that the specified object is NOT empty.  I.e. not nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  if a.NotEmpty(obj) {
//    assert.Equal(t, "two", obj[1])
//  }
func (a *Assertions) NotEmpty(object interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.NotEmpty(object, msgAndArgs...))
	return a
}

// NotEmptyf asserts that the specified object is NOT empty.  I.e. not nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  if a.NotEmptyf(obj, "error message %s", "formatted") {
//    assert.Equal(t, "two", obj[1])
//  }
func (a *Assertions) NotEmptyf(object interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.NotEmptyf(object, msg, args...))
	return a
}

// NotEqual asserts that the specified values are NOT equal.
//
//    a.NotEqual(obj1, obj2)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
func (a *Assertions) NotEqual(expected interface{}, actual interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.NotEqual(expected, actual, msgAndArgs...))
	return a
}

// NotEqualf asserts that the specified values are NOT equal.
//
//    a.NotEqualf(obj1, obj2, "error message %s", "formatted")
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
func (a *Assertions) NotEqualf(expected interface{}, actual interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.NotEqualf(expected, actual, msg, args...))
	return a
}

// NotNil asserts that the specified object is not nil.
//
//    a.NotNil(err)
func (a *Assertions) NotNil(object interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.NotNil(object, msgAndArgs...))
	return a
}

// NotNilf asserts that the specified object is not nil.
//
//    a.NotNilf(err, "error message %s", "formatted")
func (a *Assertions) NotNilf(object interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.NotNilf(object, msg, args...))
	return a
}

// NotPanics asserts that the code inside the specified PanicTestFunc does NOT panic.
//
//   a.NotPanics(func(){ RemainCalm() })
func (a *Assertions) NotPanics(f assert.PanicTestFunc, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.NotPanics(f, msgAndArgs...))
	return a
}

// NotPanicsf asserts that the code inside the specified PanicTestFunc does NOT panic.
//
//   a.NotPanicsf(func(){ RemainCalm() }, "error message %s", "formatted")
func (a *Assertions) NotPanicsf(f assert.PanicTestFunc, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.NotPanicsf(f, msg, args...))
	return a
}

// NotRegexp asserts that a specified regexp does not match a string.
//
//  a.NotRegexp(regexp.MustCompile("starts"), "it's starting")
//  a.NotRegexp("^start", "it's not starting")
func (a *Assertions) NotRegexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.NotRegexp(rx, str, msgAndArgs...))
	return a
}

// NotRegexpf asserts that a specified regexp does not match a string.
//
//  a.NotRegexpf(regexp.MustCompile("starts", "error message %s", "formatted"), "it's starting")
//  a.NotRegexpf("^start", "it's not starting", "error message %s", "formatted")
func (a *Assertions) NotRegexpf(rx interface{}, str interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.NotRegexpf(rx, str, msg, args...))
	return a
}

// NotSame asserts that two pointers do not reference the same object.
//
//    a.NotSame(ptr1, ptr2)
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func (a *Assertions) NotSame(expected interface{}, actual interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.NotSame(expected, actual, msgAndArgs...))
	return a
}

// NotSamef asserts that two pointers do not reference the same object.
//
//    a.NotSamef(ptr1, ptr2, "error message %s", "formatted")
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func (a *Assertions) NotSamef(expected interface{}, actual interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.NotSamef(expected, actual, msg, args...))
	return a
}

// NotSubset asserts that the specified list(array, slice...) contains not all
// elements given in the specified subset(array, slice...).
//
//    a.NotSubset([1, 3, 4], [1, 2], "But [1, 3, 4] does not contain [1, 2]")
func (a *Assertions) NotSubset(list interface{}, subset interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.NotSubset(list, subset, msgAndArgs...))
	return a
}

// NotSubsetf asserts that the specified list(array, slice...) contains not all
// elements given in the specified subset(array, slice...).
//
//    a.NotSubsetf([1, 3, 4], [1, 2], "But [1, 3, 4] does not contain [1, 2]", "error message %s", "formatted")
func (a *Assertions) NotSubsetf(list interface{}, subset interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.NotSubsetf(list, subset, msg, args...))
	return a
}

// NotZero asserts that i is not the zero value for its type.
func (a *Assertions) NotZero(i interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.NotZero(i, msgAndArgs...))
	return a
}

// NotZerof asserts that i is not the zero value for its type.
func (a *Assertions) NotZerof(i interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.NotZerof(i, msg, args...))
	return a
}

// Panics asserts that the code inside the specified PanicTestFunc panics.
//
//   a.Panics(func(){ GoCrazy() })
func (a *Assertions) Panics(f assert.PanicTestFunc, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Panics(f, msgAndArgs...))
	return a
}

// PanicsWithError asserts that the code inside the specified PanicTestFunc
// panics, and that the recovered panic value is an error that satisfies the
// EqualError comparison.
//
//   a.PanicsWithError("crazy error", func(){ GoCrazy() })
func (a *Assertions) PanicsWithError(errString string, f assert.PanicTestFunc, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.PanicsWithError(errString, f, msgAndArgs...))
	return a
}

// PanicsWithErrorf asserts that the code inside the specified PanicTestFunc
// panics, and that the recovered panic value is an error that satisfies the
// EqualError comparison.
//
//   a.PanicsWithErrorf("crazy error", func(){ GoCrazy() }, "error message %s", "formatted")
func (a *Assertions) PanicsWithErrorf(errString string, f assert.PanicTestFunc, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.PanicsWithErrorf(errString, f, msg, args...))
	return a
}

// PanicsWithValue asserts that the code inside the specified PanicTestFunc panics, and that
// the recovered panic value equals the expected panic value.
//
//   a.PanicsWithValue("crazy error", func(){ GoCrazy() })
func (a *Assertions) PanicsWithValue(expected interface{}, f assert.PanicTestFunc, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.PanicsWithValue(expected, f, msgAndArgs...))
	return a
}

// PanicsWithValuef asserts that the code inside the specified PanicTestFunc panics, and that
// the recovered panic value equals the expected panic value.
//
//   a.PanicsWithValuef("crazy error", func(){ GoCrazy() }, "error message %s", "formatted")
func (a *Assertions) PanicsWithValuef(expected interface{}, f assert.PanicTestFunc, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.PanicsWithValuef(expected, f, msg, args...))
	return a
}

// Panicsf asserts that the code inside the specified PanicTestFunc panics.
//
//   a.Panicsf(func(){ GoCrazy() }, "error message %s", "formatted")
func (a *Assertions) Panicsf(f assert.PanicTestFunc, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Panicsf(f, msg, args...))
	return a
}

// Regexp asserts that a specified regexp matches a string.
//
//  a.Regexp(regexp.MustCompile("start"), "it's starting")
//  a.Regexp("start...$", "it's not starting")
func (a *Assertions) Regexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Regexp(rx, str, msgAndArgs...))
	return a
}

// Regexpf asserts that a specified regexp matches a string.
//
//  a.Regexpf(regexp.MustCompile("start", "error message %s", "formatted"), "it's starting")
//  a.Regexpf("start...$", "it's not starting", "error message %s", "formatted")
func (a *Assertions) Regexpf(rx interface{}, str interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Regexpf(rx, str, msg, args...))
	return a
}

// Same asserts that two pointers reference the same object.
//
//    a.Same(ptr1, ptr2)
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func (a *Assertions) Same(expected interface{}, actual interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Same(expected, actual, msgAndArgs...))
	return a
}

// Samef asserts that two pointers reference the same object.
//
//    a.Samef(ptr1, ptr2, "error message %s", "formatted")
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func (a *Assertions) Samef(expected interface{}, actual interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Samef(expected, actual, msg, args...))
	return a
}

// Subset asserts that the specified list(array, slice...) contains all
// elements given in the specified subset(array, slice...).
//
//    a.Subset([1, 2, 3], [1, 2], "But [1, 2, 3] does contain [1, 2]")
func (a *Assertions) Subset(list interface{}, subset interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Subset(list, subset, msgAndArgs...))
	return a
}

// Subsetf asserts that the specified list(array, slice...) contains all
// elements given in the specified subset(array, slice...).
//
//    a.Subsetf([1, 2, 3], [1, 2], "But [1, 2, 3] does contain [1, 2]", "error message %s", "formatted")
func (a *Assertions) Subsetf(list interface{}, subset interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Subsetf(list, subset, msg, args...))
	return a
}

// True asserts that the specified value is true.
//
//    a.True(myBool)
func (a *Assertions) True(value bool, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.True(value, msgAndArgs...))
	return a
}

// Truef asserts that the specified value is true.
//
//    a.Truef(myBool, "error message %s", "formatted")
func (a *Assertions) Truef(value bool, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Truef(value, msg, args...))
	return a
}

// WithinDuration asserts that the two times are within duration delta of each other.
//
//   a.WithinDuration(time.Now(), time.Now(), 10*time.Second)
func (a *Assertions) WithinDuration(expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.WithinDuration(expected, actual, delta, msgAndArgs...))
	return a
}

// WithinDurationf asserts that the two times are within duration delta of each other.
//
//   a.WithinDurationf(time.Now(), time.Now(), 10*time.Second, "error message %s", "formatted")
func (a *Assertions) WithinDurationf(expected time.Time, actual time.Time, delta time.Duration, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.WithinDurationf(expected, actual, delta, msg, args...))
	return a
}

// YAMLEq asserts that two YAML strings are equivalent.
func (a *Assertions) YAMLEq(expected string, actual string, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.YAMLEq(expected, actual, msgAndArgs...))
	return a
}

// YAMLEqf asserts that two YAML strings are equivalent.
func (a *Assertions) YAMLEqf(expected string, actual string, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.YAMLEqf(expected, actual, msg, args...))
	return a
}

// Zero asserts that i is the zero value for its type.
func (a *Assertions) Zero(i interface{}, msgAndArgs ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Zero(i, msgAndArgs...))
	return a
}

// Zerof asserts that i is the zero value for its type.
func (a *Assertions) Zerof(i interface{}, msg string, args ...interface{}) *Assertions {
	a.R = append(a.R, a.A.Zerof(i, msg, args...))
	return a
}
