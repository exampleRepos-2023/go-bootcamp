package math

import "testing"

func TestAbs(t *testing.T) {
	pairs := [][]int{{1, -1}, {0, 0}, {-1, 1}}
	for _, pair := range pairs {
		a, b := pair[0], pair[1]
		if Add(a, b) != a+b {
			t.Errorf("%d + %d != %d\n", a, b, a+b)
		}
	}
}

func FuzzAbs(f *testing.F) {
	f.Add(1, -1)
	f.Add(0, 0)
	f.Add(-1, 1)
	f.Fuzz(func(t *testing.T, a, b int) {
		if Add(a, b) != a+b {
			t.Errorf("%d + %d != %d\n", a, b, a+b)
		}
	})
}
