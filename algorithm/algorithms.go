package algorithm

// AllOf is a function to check if all elements satisfy a condition.
func AllOf[T any](elements []T, pred func(T) bool) bool {
	for _, v := range elements {
		if !pred(v) {
			return false
		}
	}
	return true
}

func AnyOf[T any](elements []T, pred func(T) bool) bool {
	for _, v := range elements {
		if pred(v) {
			return true
		}
	}
	return false
}

func NoneOf[T any](elements []T, pred func(T) bool) bool {
	return !AnyOf(elements, pred)
}

func ForEach[T any](elements []T, f func(T)) {
	for _, v := range elements {
		f(v)
	}
}

func Count[T comparable](elements []T, value T) int64 {
	count := int64(0)
	for _, v := range elements {
		if v == value {
			count++
		}
	}
	return count
}
