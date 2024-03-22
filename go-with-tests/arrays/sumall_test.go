package arrays

import (
	"slices"
	"testing"
)

func TestSumAll(t *testing.T) {
	// PASSED: ???
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func BenchmarkSumAll(t *testing.B) {
	for i := 0; i < t.N; i++ {
		SumAll([]int{1, 2}, []int{0, 9})
	}
}
