package collect_test

import (
	"testing"

	"github.com/pda13/h-careers-golang-hw/internal/strings/collect"
)

var testcases = []struct {
	runes  []rune
	result string
}{
	{runes: []rune{104, 101, 108, 108, 111}, result: "hello"},
	{runes: []rune{1087, 1088, 1080, 1074, 1077, 1090, 44, 32, 1084, 1080, 1088}, result: "привет, мир"},
}

func TestCollectStringFromRunesUsingConversion(t *testing.T) {
	for _, testcase := range testcases {
		t.Run(testcase.result, func(t *testing.T) {
			result := collect.CollectStringFromRunesUsingConversion(testcase.runes)
			if result != testcase.result {
				t.Errorf(
					"collecting %v failed: expected %s, got %s",
					testcase.runes,
					testcase.result,
					result,
				)
			}
		})
	}
}

func BenchmarkCollectStringFromRunesUsingConversion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testcase := range testcases {
			collect.CollectStringFromRunesUsingConversion(testcase.runes)
		}
	}
}

func TestCollectStringFromRunesUsingBuilder(t *testing.T) {
	for _, testcase := range testcases {
		t.Run(testcase.result, func(t *testing.T) {
			result := collect.CollectStringFromRunesUsingBuilder(testcase.runes)
			if result != testcase.result {
				t.Errorf(
					"collecting %v failed: expected %s, got %s",
					testcase.runes,
					testcase.result,
					result,
				)
			}
		})
	}
}

func BenchmarkCollectStringFromRunesUsingBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testcase := range testcases {
			collect.CollectStringFromRunesUsingBuilder(testcase.runes)
		}
	}
}
