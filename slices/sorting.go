package slices

// MergeSort is a generic implementation of the merge sort algorithm.
func MergeSort[T Ordered](items []T) []T {
	num := len(items)

	if num == 1 {
		return items
	}
	mid := int(num / 2)

	left := make([]T, mid)
	right := make([]T, num-mid)

	for i := 0; i < num; i++ {
		if i < mid {
			left[i] = items[i]
		} else {
			right[i-mid] = items[i]
		}
	}
	return merge(MergeSort(left), MergeSort(right))
}

func merge[T Ordered](left, right []T) []T {
	result := make([]T, len(left)+len(right))
	i := 0

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return result
}
