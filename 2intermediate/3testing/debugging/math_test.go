package math

import "testing"

type testCase struct {
	a   int
	b   int
	sum int
}

func TestAdd(t *testing.T) {
	tc := testCase{1, 2, 1 + 2}
	expected := tc.sum
	actual := Add(tc.a, tc.b)

	if actual != expected {
		t.Error("Unexpected result for Add()")
	}
}
