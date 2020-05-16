package vector

import (
	"strconv"
	"testing"
)

func TestApply(t *testing.T) {
	for k, v := range []struct {
		a    Vector
		f    func(float64) float64
		want Vector
	}{
		{Vector{1, 2, 3}, func(v float64) float64 { return v * v }, Vector{1, 4, 9}},
		{Vector{1, 2, 3}, func(v float64) float64 { return v + v }, Vector{2, 4, 6}},
		{Vector{1, 2, 3}, func(v float64) float64 { return v * -1 }, Vector{-1, -2, -3}},
	} {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			want := v.want.String()
			v.a.Apply(v.f)
			got := v.a.String()
			if want != got {
				t.Errorf("want %s, got %s", want, got)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	for k, v := range []struct {
		a, b Vector
		want bool
	}{
		{Vector{0}, Vector{0}, true},
		{Vector{1}, Vector{1}, true},
		{Vector{1, 2}, Vector{1, 2}, true},
		{Vector{1, 2, 3}, Vector{1, 2, 3}, true},
		{Vector{1, 2}, Vector{1, 2, 3}, false},
		{Vector{1, 2, 4}, Vector{1, 2, 3}, false},
	} {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			got := v.a.Equal(v.b)
			if v.want != got {
				t.Errorf("want %t, got %t(%s,%s)", v.want, got, v.a, v.b)
			}
		})
	}
}

func TestString(t *testing.T) {
	for k, v := range []struct {
		a    Vector
		want string
	}{
		{Vector{1, 2, 3}, "[1, 2, 3]"},
		{Vector{4, 5, 6}, "[4, 5, 6]"},
		{Vector{7, 8, 9}, "[7, 8, 9]"},
	} {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			got := v.a.String()
			if v.want != got {
				t.Errorf("want %s, got %s", v.want, got)
			}
		})
	}
}
