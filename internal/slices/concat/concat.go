package concat

func ConcatSlicesUsingCopy[T any](slices ...[]T) []T {
	totalElementsCount := 0
	for _, slice := range slices {
		totalElementsCount += len(slice)
	}

	result := make([]T, totalElementsCount)
	start := 0
	end := 0

	for _, slice := range slices {
		end += len(slice)
		copy(result[start:end], slice)
		start += len(slice)
	}

	return result
}

func ConcatSlicesUsingAppend[T any](slices ...[]T) []T {
	totalElementsCount := 0
	for _, slice := range slices {
		totalElementsCount += len(slice)
	}

	result := make([]T, 0, totalElementsCount)

	for _, slice := range slices {
		result = append(result, slice...)
	}

	return result
}
