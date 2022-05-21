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

package slices

import (
	"golang.org/x/exp/constraints"
)

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

func BinarySearch[T constraints.Ordered](values []T, x T) int {
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
