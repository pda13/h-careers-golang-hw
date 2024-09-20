package concat_test

import (
	"slices"
	"testing"

	"github.com/pda13/h-careers-golang-hw/internal/slices/concat"
)

var testcases = []struct {
	slices        [][]int
	concatenation []int
}{
	{slices: [][]int{{1, 2, 3}, {4}, {5, 6, 7, 8}, {9, 10}}, concatenation: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{slices: [][]int{{1, 2, 3}, {4, 5, 5, 5}, {5, 6, 7, 8}, {}}, concatenation: []int{1, 2, 3, 4, 5, 5, 5, 5, 6, 7, 8}},
	{slices: [][]int{{1}, {2}, {3}, {4}, {5}}, concatenation: []int{1, 2, 3, 4, 5}},
}

func TestConcatSlicesUsingCopy(t *testing.T) {
	for _, testcase := range testcases {
		t.Run("test", func(t *testing.T) {
			concatenation := concat.ConcatSlicesUsingCopy(testcase.slices...)
			if !slices.Equal(concatenation, testcase.concatenation) {
				t.Errorf(
					"concatenation of %v failed: expected %v, got %v",
					testcase.slices,
					testcase.concatenation,
					concatenation,
				)
			}
		})
	}
}

func BenchmarkConcatSlicesUsingCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testcase := range testcases {
			concat.ConcatSlicesUsingCopy(testcase.slices)
		}
	}
}

func TestConcatSlicesUsingAppend(t *testing.T) {
	for _, testcase := range testcases {
		t.Run("test", func(t *testing.T) {
			concatenation := concat.ConcatSlicesUsingAppend(testcase.slices...)
			if !slices.Equal(concatenation, testcase.concatenation) {
				t.Errorf(
					"concatenation of %v failed: expected %v, got %v",
					testcase.slices,
					testcase.concatenation,
					concatenation,
				)
			}
		})
	}
}

func BenchmarkConcatSlicesUsingAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testcase := range testcases {
			concat.ConcatSlicesUsingAppend(testcase.slices)
		}
	}
}
