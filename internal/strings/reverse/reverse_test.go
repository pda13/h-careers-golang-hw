package reverse_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/pda13/h-careers-golang-hw/internal/strings/reverse"
)

var testcases = []struct {
	original string
	reversed string
}{
	{original: "hello", reversed: "olleh"},
	{original: "o", reversed: "o"},
	{original: "", reversed: ""},
	{original: "привет, мир!", reversed: "!рим ,тевирп"},
}

func TestReverseStringUsingRunes(t *testing.T) {
	for _, testcase := range testcases {
		t.Run(testcase.original, func(t *testing.T) {
			reversed := reverse.ReverseStringUsingRunes(testcase.original)
			if reversed != testcase.reversed {
				t.Errorf(
					"reversing %s failed: expected %s, got %s",
					testcase.original,
					testcase.reversed,
					reversed,
				)
			}
		})
	}
}

func BenchmarkReverseStringUsingRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = reverse.ReverseStringUsingRunes(gofakeit.Username())
	}
}

func TestReverseStringUsingBuilder(t *testing.T) {
	for _, testcase := range testcases {
		t.Run(testcase.original, func(t *testing.T) {
			reversed := reverse.ReverseStringUsingBuilder(testcase.original)
			if reversed != testcase.reversed {
				t.Errorf(
					"reversing %s failed: expected %s, got %s",
					testcase.original,
					testcase.reversed,
					reversed,
				)
			}
		})
	}
}

func BenchmarkReverseStringUsingBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = reverse.ReverseStringUsingBuilder(gofakeit.Username())
	}
}
