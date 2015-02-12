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

func Test_MinAegs(t *testing.T) {

	if 0 != MinArgs(0) {
		t.Errorf("MinArgs(%v) == %d, want %d", []int{0}, MinArgs(0), 0)
	}

	if 1 != MinArgs(3, 2, 1) {
		t.Errorf("MinArgs(%v) == %d, want %d", []int{3, 2, 1}, MinArgs(3, 2, 1), 1)
	}

	if -5 != MinArgs(-5, 0, 1) {
		t.Errorf("MinArgs(%v) == %d, want %d", []int{-5, 0, 1}, MinArgs(-5, 0, 1), -5)
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
