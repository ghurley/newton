package newton

import (
	"testing"
)

/*
We're testing functions that return floating point numbers so the usual caveats
apply. We're expressing the expected results in source as decimal strings but
conversion into IEEE 754 may be slightly different.

These tests are in fact backwards in that I wrote them after the libraries
were working because it was too annoying to figure out a good decimal
representation that parses into the correct IEEE 754 without running
the calculation! Even then, it was much trial and error to get the right number
of decimal digits that caused the parsing to get the right results and it
may be that I just got lucky. Perhaps a better way would be to avail myself
of the math.Float64frombits et al. functions but the purpose here is to
allow refactoring and optimizations that keep calculating correctly and what's
here supports that.
*/

func TestNewton(t *testing.T) {
	p := 1 - .5i

	p64 := complex64(p)
	expected64 := complex64(0.82666665 - 0.1199999973i)
	if got := newton64(p64); got != expected64 {
		t.Errorf("Expected: %1.28f Got: %1.28f", expected64, got)
	}

	p128 := p
	expected128 := 0.82666666666666666 - 0.12000000000000005i
	if got := newton128(p128); got != expected128 {
		t.Errorf("Expected: %1.28f Got: %1.28f", expected128, got)
	}
}
