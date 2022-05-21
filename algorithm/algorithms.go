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
	var count int64 = 0
	for _, v := range elements {
		if v == value {
			count++
		}
	}
	return count
}

func CountIf[T any](elements []T, pred func(T) bool) int64 {
	var c int64 = 0
	for _, v := range elements {
		if pred(v) {
			c++
		}
	}
	return c
}

func GroupBy[T comparable](elements []T) map[T]int64 {
	m := make(map[T]int64)
	for _, v := range elements {
		m[v]+=1
	}
	return m
}
