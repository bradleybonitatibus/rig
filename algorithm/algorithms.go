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
