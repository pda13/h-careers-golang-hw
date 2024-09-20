package intersection

func FindSliceIntersectionUsingSet[T comparable](first, second []T) []T {
	firstUniqueElements := make(map[T]struct{})
	for _, elem := range first {
		firstUniqueElements[elem] = struct{}{}
	}

	intersection := make([]T, 0)
	for _, elem := range second {
		if _, exists := firstUniqueElements[elem]; exists {
			intersection = append(intersection, elem)
		}
	}

	return intersection
}
