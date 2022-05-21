/*
Copyright 2022 Bradley Bonitatibus

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package algorithm

// Ordered is the interface containing all the types that implement ordering.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}

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

// LinearSearch iterates through all the values and checks if the value of
// x is within the slice of values. If the value exists, it will return
// the index at which the value is contained.
func LinearSearch[T comparable](values []T, x T) int {
	for i, v := range values {
		if x == v {
			return i
		}
	}
	return -1
}

// BinarySearch searches a slices for a value using binary search.
// This function will return the index of the element, or -1 if the value
// is not found.
func BinarySearch[T Ordered](values []T, x T) int {
	if len(values) == 0 {
		return -1
	}

	low := 0
	high := len(values) - 1

	for low <= high {
		mid := (low + high) / 2
		if values[mid] == x {
			return mid
		}
		if values[mid] > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

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
