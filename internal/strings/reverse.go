package strings

import "strings"

func ReverseStringUsingRunes(original string) string {
	originalRunes := []rune(original)
	reversedRunes := make([]rune, len(originalRunes))

	for i := 0; i < len(originalRunes); i++ {
		reversedRunes[i] = originalRunes[len(originalRunes)-i-1]
	}

	return string(reversedRunes)
}

func ReverseStringUsingBuilder(original string) string {
	originalRunes := []rune(original)

	builder := strings.Builder{}
	builder.Grow(len(originalRunes))

	for i := len(originalRunes) - 1; i >= 0; i-- {
		builder.WriteRune(originalRunes[i])
	}

	return builder.String()
}
