package vector

import "math"

// Equal checks if two float64 scalars are equal.
func Equal(x, y float64) bool {
	return math.Abs(x-y) <= 1e-9
}
