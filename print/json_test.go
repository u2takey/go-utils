package print

import (
	"fmt"
	"testing"
)

func TestMarshalStringIgnoreError(t *testing.T) {
	a := struct {
		A string
		B []string
	}{
		A: "abcdf",
		B: []string{"abcdf", "cdf", "", "abcdf", "adf", "abf", "bcdf", "a", "cdf", "", "abcdf", "adf", "abf", "bcdf", "a"},
	}
	fmt.Println(MarshalStringIgnoreError(a))

	a = struct {
		A string
		B []string
	}{
		A: "abcdf",
		B: []string{"abcdf", "cdf", "", "abcdf", "adf", "abf", "bcdf", "a", "cdf"},
	}
	fmt.Println(MarshalStringIgnoreError(a))

	b := struct {
		A string
		B map[string]string
	}{
		A: "abcdf",
		B: map[string]string{"1": "1", "2": "2", "3": "3"},
	}
	fmt.Println(MarshalStringIgnoreError(b))

	b = struct {
		A string
		B map[string]string
	}{
		A: "abcdf",
		B: map[string]string{},
	}
	fmt.Println(MarshalStringIgnoreError(b))

	b = struct {
		A string
		B map[string]string
	}{
		A: "abcdf",
		B: map[string]string{"1": "1", "2": "2", "3": "3", "4": "4", "5": "5", "6": "6", "7": "7", "8": "8", "9": "9", "10": "10", "11": "11"},
	}
	fmt.Println(MarshalStringIgnoreError(b))

	c := struct {
		A string
		B map[string][]string
	}{
		A: "abcdf",
		B: map[string][]string{"1": {"abcdf", "cdf", "", "abcdf", "adf", "abf", "bcdf", "a", "cdf", "", "abcdf", "adf", "abf", "bcdf", "a"}},
	}
	fmt.Println(MarshalStringIgnoreError(c))
}
