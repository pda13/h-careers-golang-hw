package intersection_test

import (
	"slices"
	"testing"

	"github.com/pda13/h-careers-golang-hw/internal/slices/intersection"
)

var testcases = []struct {
	first        []int
	second       []int
	intersection []int
}{
	{first: []int{}, second: []int{}, intersection: []int{}},
	{first: []int{1}, second: []int{2}, intersection: []int{}},
	{first: []int{1, 2, 3}, second: []int{4, 5, 6}, intersection: []int{}},
	{first: []int{1, 2, 3, 4, 5, 6}, second: []int{2, 4, 6, 8, 10}, intersection: []int{2, 4, 6}},
}

func TestFindSliceIntersectionUsingSet(t *testing.T) {
	for _, testcase := range testcases {
		t.Run("test", func(t *testing.T) {
			intersection := intersection.FindSliceIntersectionUsingSet(testcase.first, testcase.second)
			if !slices.Equal(intersection, testcase.intersection) {
				t.Errorf(
					"finding intersection of %v and %v failed: expected %v, got %v",
					testcase.first,
					testcase.second,
					testcase.intersection,
					intersection,
				)
			}
		})
	}
}
