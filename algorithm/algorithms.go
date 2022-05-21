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

// AnyOf is a function to check that any element has a true evaluation
// of the predicate pre.
func AnyOf[T any](elements []T, pred func(T) bool) bool {
	for _, v := range elements {
		if pred(v) {
			return true
		}
	}
	return false
}

// NoneOf is a function that ensures that all items in elements are not
// truthy when evaluated by the predicate pred.
func NoneOf[T any](elements []T, pred func(T) bool) bool {
	return !AnyOf(elements, pred)
}

// ForEach applies function f to every element in elements.
// The function will not mutate any elements and ignores the return
// value of the function f.
func ForEach[T any](elements []T, f func(T)) {
	for _, v := range elements {
		f(v)
	}
}

// Count counts all occurrences of a given value.
func Count[T comparable](elements []T, value T) int64 {
	var count int64 = 0
	for _, v := range elements {
		if v == value {
			count++
		}
	}
	return count
}

// CountIf counts all values who have a truthy evaluation of predicate pred.
func CountIf[T any](elements []T, pred func(T) bool) int64 {
	var c int64 = 0
	for _, v := range elements {
		if pred(v) {
			c++
		}
	}
	return c
}

// GroupBy groups all the unique items in elements slice into a key value
// map that contains the unique key, and the number of occurrence.
func GroupBy[T comparable](elements []T) map[T]int64 {
	m := make(map[T]int64)
	for _, v := range elements {
		m[v] += 1
	}
	return m
}

// UniqueCopy returns a copy of unique items in elements.
func UniqueCopy[T comparable](elements []T) []T {
	m := map[T]struct{}{}
	out := make([]T, 0, len(elements))
	for _, v := range elements {
		if _, ok := m[v]; !ok {
			out = append(out, v)
			m[v] = struct{}{}
		}
	}
	return out
}
