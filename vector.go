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

// Dot calculates the Dot Product of Vectors a and b
func (a Vector) Dot(b Vector) (sum float64) {
	la, lb := len(a), len(b)
	if la != lb {
		panic(fmt.Sprintf("can not calculate Dot Product; len(a)==%d, len(b)==%d", la, lb))
	}
	for k := 0; k < la; k++ {
		sum += a[k] * b[k]
	}
	return sum
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
