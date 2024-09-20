package duplicates_test

import (
	"slices"
	"testing"

	"github.com/pda13/h-careers-golang-hw/internal/slices/duplicates"
)

var testcases = []struct {
	original []int
	unique   []int
}{
	{original: []int{}, unique: []int{}},
	{original: []int{1}, unique: []int{1}},
	{original: []int{1, 2, 3}, unique: []int{1, 2, 3}},
	{original: []int{1, 1, 2, 2, 3, 3, 3, 3, 4, 4}, unique: []int{1, 2, 3, 4}},
}

func TestRemoveDuplicatesUsingSet(t *testing.T) {
	for _, testcase := range testcases {
		t.Run("test", func(t *testing.T) {
			unique := duplicates.RemoveDuplicatesUsingSet(testcase.original)
			if !slices.Equal(unique, testcase.unique) {
				t.Errorf(
					"deleting duplicates from %v failed: expected %v, got %v",
					testcase.original,
					testcase.unique,
					unique,
				)
			}
		})
	}
}

func BenchmarkRemoveDuplicatesUsingSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testcase := range testcases {
			duplicates.RemoveDuplicatesUsingSet(testcase.original)
		}
	}
}

func TestRemoveDuplicatesUsingCompact(t *testing.T) {
	for _, testcase := range testcases {
		t.Run("test", func(t *testing.T) {
			unique := duplicates.RemoveDuplicatesUsingCompact(testcase.original)
			if !slices.Equal(unique, testcase.unique) {
				t.Errorf(
					"deleting duplicates from %v failed: expected %v, got %v",
					testcase.original,
					testcase.unique,
					unique,
				)
			}
		})
	}
}

func BenchmarkRemoveDuplicatesUsingCompact(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testcase := range testcases {
			duplicates.RemoveDuplicatesUsingCompact(testcase.original)
		}
	}
}
