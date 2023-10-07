package math

import "testing"

func BenchmarkAbs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Abs(i)
	}
}

func BenchmarkAbsSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AbsSlow(i)
	}
}
