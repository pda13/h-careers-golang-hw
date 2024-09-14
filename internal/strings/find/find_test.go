package find_test

import (
	"testing"

	"github.com/pda13/h-careers-golang-hw/internal/strings/find"
)

var testcases = []struct {
	original  string
	substring string
	result    find.FindSubstringResult
	contains  bool
}{
	{
		original:  "hello",
		substring: "ell",
		result:    find.FindSubstringResult{StartIndexInclusive: 1, EndIndexExclusive: 4},
		contains:  true,
	},
	{
		original:  "ABCDEFABCQM",
		substring: "ABCQ",
		result:    find.FindSubstringResult{StartIndexInclusive: 6, EndIndexExclusive: 10},
		contains:  true,
	},
	{
		original:  "ABCDEFABCQM",
		substring: "ABCQ",
		result:    find.FindSubstringResult{StartIndexInclusive: 6, EndIndexExclusive: 10},
		contains:  true,
	},
	// {
	// 	original:  "ПРИВЕТИКРИВЕТЫ",
	// 	substring: "РИВЕТЫ",
	// 	result:    find.FindSubstringResult{StartIndexInclusive: 8, EndIndexExclusive: 14},
	// 	contains:  true,
	// },
}

func TestCustomFindSubstring(t *testing.T) {
	for _, testcase := range testcases {
		t.Run(testcase.original, func(t *testing.T) {
			result, contains := find.CustomFindSubstring(testcase.original, testcase.substring)
			if result != testcase.result || contains != testcase.contains {
				t.Errorf(
					"finding %s in %s failed: expected (%v, %v), got (%v, %v)",
					testcase.substring,
					testcase.original,
					testcase.result,
					testcase.contains,
					result,
					contains,
				)
			}
		})
	}
}

func BenchmarkCustomFindSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testcase := range testcases {
			find.CustomFindSubstring(testcase.original, testcase.substring)
		}
	}
}

func TestBuiltinFindSubstring(t *testing.T) {
	for _, testcase := range testcases {
		t.Run(testcase.original, func(t *testing.T) {
			result, contains := find.BuiltinFindSubstring(testcase.original, testcase.substring)
			if result != testcase.result || contains != testcase.contains {
				t.Errorf(
					"finding %s in %s failed: expected (%v, %v), got (%v, %v)",
					testcase.substring,
					testcase.original,
					testcase.result,
					testcase.contains,
					result,
					contains,
				)
			}
		})
	}
}

func BenchmarkBuiltinFindSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testcase := range testcases {
			find.BuiltinFindSubstring(testcase.original, testcase.substring)
		}
	}
}
