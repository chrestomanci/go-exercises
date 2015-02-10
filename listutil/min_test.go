package listutil

import (
	"testing"
)

type case_ent struct {
	in   []int
	want int
}

func Test_Min(t *testing.T) {
	cases := []case_ent{
		{[]int{0}, 0},
		{[]int{3, 2, 1}, 1},
		{[]int{-5, 0, 1}, -5},
	}

	for _, c := range cases {
		got := Min(c.in)
		if got != c.want {
			t.Errorf("Min(%v) == %d, want %d", c.in, got, c.want)
		}
	}
}

func Test_Max(t *testing.T) {
	cases := []case_ent{
		{[]int{0}, 0},
		{[]int{3, 2, 1}, 3},
		{[]int{-5, 0, 1}, 1},
	}

	for _, c := range cases {
		got := Max(c.in)
		if got != c.want {
			t.Errorf("Min(%v) == %d, want %d", c.in, got, c.want)
		}
	}
}

