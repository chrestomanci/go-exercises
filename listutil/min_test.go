package listutil

import (
	"testing"
)

// NB: Not working. Syntax errors.
func TestXYZ(t *testing.T) {
	cases := []struct {
		in []int, want int
	}{
		{ [0], 0 },
		{ [3,2,1], 1 },
		{ [-5,0,1], -1},
    };

	for _, c := range cases {
		got := Min(c.in)
		if got != c.want {
			t.Errorf("Min(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
