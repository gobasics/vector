package vector

import (
	"fmt"
	"strings"
)

// Vector is a float64 slice
type Vector []float64

// Add performs inplace elementwise addition of Vector b into b
func (a Vector) Add(b []float64) {
	la, lb := len(a), len(b)
	if la != lb {
		panic(fmt.Sprintf("can not calculate a.Add(b); len(a)==%d, len(b)==%d", la, lb))
	}
	for k := 0; k < la; k++ {
		a[k] += b[k]
	}
}

// Apply updates every element in Vector a with the results of applying function f
func (a Vector) Apply(f func(float64) float64) {
	for k := range a {
		a[k] = f(a[k])
	}
}

// Dot calculates the Dot Product of Vectors a and b
func (a Vector) Dot(b []float64) (sum float64) {
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
func (a Vector) Equal(b []float64) bool {
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

// Get returns a slice of floats that Vector a points to.
func (a Vector) Get() []float64 {
	return a
}

// Set copies the elements of slice b into Vector a.
func (a Vector) Set(b []float64) {
	copy(a, b)
}

// String returns a string representation of Vector a
func (a Vector) String() string {
	s := make([]string, len(a))
	for k := range a {
		s[k] = fmt.Sprintf("%g", a[k])
	}
	return fmt.Sprintf("[%s]", strings.Join(s, ", "))
}
