package find

import "strings"

type FindSubstringResult struct {
	StartIndexInclusive int
	EndIndexExclusive   int
}

func CustomFindSubstring(original, substring string) (FindSubstringResult, bool) {
	originalRunes := []rune(original)
	substringRunes := []rune(substring)

	idx := 0
	count := 0

	for idx < len(originalRunes) {
		if originalRunes[idx] == substringRunes[count] {
			count++
			idx++
		} else {
			idx -= count - 1
			count = 0
		}

		if count == len(substringRunes) {
			return FindSubstringResult{
				StartIndexInclusive: idx - count,
				EndIndexExclusive:   idx,
			}, true
		}
	}

	return FindSubstringResult{}, false
}

func BuiltinFindSubstring(original, substring string) (FindSubstringResult, bool) {
	idx := strings.Index(original, substring)
	if idx < 0 {
		return FindSubstringResult{}, false
	}

	return FindSubstringResult{
		StartIndexInclusive: idx,
		EndIndexExclusive:   idx + len(substring),
	}, true
}
