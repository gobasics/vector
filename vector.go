package vector

import (
	"fmt"
	"strings"
)

// Vector is a float64 slice
type Vector []float64

// Apply updates every element in Vector a with the results of applying function f
func (a Vector) Apply(f func(float64) float64) {
	for k := range a {
		a[k] = f(a[k])
	}
}

// Equal returns true if two Vectors have the same number of elements and
// the corresponding elements are equal.
func (a Vector) Equal(b Vector) bool {
	if len(a) != len(b) {
		return false
	}

	for k := range a {
		if !Equal(a[k], b[k]) {
			return false
		}
	}

	return true
}

// String returns a string representation of Vector a
func (a Vector) String() string {
	s := make([]string, len(a))
	for k := range a {
		s[k] = fmt.Sprintf("%g", a[k])
	}
	return fmt.Sprintf("[%s]", strings.Join(s, ", "))
}
