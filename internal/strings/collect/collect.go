package collect

import "strings"

func CollectStringFromRunesUsingConversion(runes []rune) string {
	result := ""
	for _, r := range runes {
		result += string(r)
	}

	// return string(runes)
	return result
}

func CollectStringFromRunesUsingBuilder(runes []rune) string {
	builder := strings.Builder{}
	builder.Grow(len(runes))

	for _, r := range runes {
		builder.WriteRune(r)
	}

	return builder.String()
}
