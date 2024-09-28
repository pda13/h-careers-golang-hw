package duplicates

import (
	"cmp"
	"slices"
)

func RemoveDuplicatesUsingSet[T comparable](original []T) []T {
	uniqueElements := make([]T, 0)
	encounteredElements := make(map[T]struct{}) // map[T]bool - так не делать

	for _, elem := range original {
		if _, exists := encounteredElements[elem]; !exists {
			uniqueElements = append(uniqueElements, elem)
		}

		encounteredElements[elem] = struct{}{}
	}

	return uniqueElements
}

func RemoveDuplicatesUsingCompact[T cmp.Ordered](original []T) []T {
	uniqueElements := make([]T, len(original))
	copy(uniqueElements, original)

	slices.Sort(uniqueElements)
	return slices.Compact(uniqueElements)
}
