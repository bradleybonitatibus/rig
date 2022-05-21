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

import (
	"math"
	"sort"
	"strings"
	"testing"
)

func TestAllOf(t *testing.T) {
	evens := []int32{2, 4, 6, 8, 10}
	isEven := func(x int32) bool {
		return x%2 == 0
	}

	if !AllOf(evens, isEven) {
		t.Error("expected slice of even numbers to be all even")
	}

	allCaps := []string{
		"HELLO",
		"WORLD",
		"GOODNIGHT",
	}
	isAllCaps := func(s string) bool {
		return strings.ToUpper(s) == s
	}

	if !AllOf(allCaps, isAllCaps) {
		t.Error("expected slice of all capital words to be truthy")
	}

	type test struct {
		k string
		v int32
	}

	structsAllHave5LengthPropertyK := []test{
		{
			k: "11111",
			v: 1234,
		},
		{
			k: "22222",
			v: 12345,
		},
		{
			k: "33333",
			v: 89823,
		},
	}
	has5LengthKeys := func(x test) bool {
		return len(x.k) == 5
	}

	if !AllOf(structsAllHave5LengthPropertyK, has5LengthKeys) {
		t.Error("structs did not all have key of length 5")
	}

	notAllEven := []int32{2, 3, 4, 6, 8}

	if AllOf(notAllEven, isEven) {
		t.Error("expected not all even to be false when provided non-even integers")
	}
}

func TestAnyOf(t *testing.T) {
	oddsWithOneEven := []int64{1, 3, 4, 5, 7}
	isEven := func(x int64) bool {
		return x%2 == 0
	}

	if !AnyOf(oddsWithOneEven, isEven) {
		t.Error("expected AnyOf to be truthy with odd slice with a single even value")
	}

	allLower := []string{
		"hello",
		"goodnight",
		"sleep tight",
	}
	isAllUpper := func(s string) bool {
		return strings.ToUpper(s) == s
	}

	if AnyOf(allLower, isAllUpper) {
		t.Error("expected allLower to not have any upper case values")
	}
}

func TestNoneOf(t *testing.T) {
	evens := []int32{2, 4, 6, 8, 10, 12, 14}
	isOdd := func(x int32) bool {
		return x%2 != 0
	}

	if !NoneOf(evens, isOdd) {
		t.Error("expected NoneOf to be truthy for a slice of all even numbers")
	}

	type test struct {
		k string
		v int
	}

	v := []test{
		{
			k: "2123",
			v: 15,
		},
		{
			k: "111",
			v: 12,
		},
		{
			k: "123123123123123123123",
			v: 11,
		},
	}
	keyOfLength3 := func(x test) bool {
		return len(x.k) == 3
	}

	if NoneOf(v, keyOfLength3) {
		t.Error("expected NoneOf to be falsy when all test.k are not length 3")
	}
}

func TestForEach(t *testing.T) {
	s := []string{"some", "exciting", "values"}
	builder := strings.Builder{}

	f := func(s string) {
		builder.Write([]byte(s))
	}

	ForEach(s, f)

	if builder.String() != "someexcitingvalues" {
		t.Errorf("expected 'someexcitingvalues' got %v instead", builder.String())
	}
}

func TestCount(t *testing.T) {
	v := []int32{1, 2, 2, 5, 2, 19, 2}

	c := Count(v, 2)
	if c != 4 {
		t.Errorf("expected 4, got %v instead", c)
	}

	s := []string{"h", "e", "h", "l"}
	sc := Count(s, "h")

	if sc != 2 {
		t.Errorf("expected count of 2, got %v instead", sc)
	}
}

func TestCountIf(t *testing.T) {
	evens := []int64{-4, 2, 3, 15, 12, 144}
	isEven := func(x int64) bool {
		return x%2 == 0
	}

	countEvens := CountIf(evens, isEven)
	if countEvens != 4 {
		t.Errorf("expected 4 evens, got %v instead", countEvens)
	}
	type test struct {
		k string
		v int64
	}
	tests := []test{
		{
			k: "user_id_1",
			v: 1234,
		},
		{
			k: "user_id_1",
			v: 9999,
		},
		{
			k: "user_id_2",
			v: 12312312,
		},
	}

	countUserID1 := CountIf(tests, func(x test) bool {
		return x.k == "user_id_1"
	})
	if countUserID1 != 2 {
		t.Errorf("expected count user ID 1 to be 2 got %v instead", countUserID1)
	}
}

func TestGroupBy(t *testing.T) {
	stringsValues := []string{
		"hello",
		"another",
		"hello",
		"string",
		"hello",
	}

	sm := GroupBy(stringsValues)
	if len(sm) != 3 || sm["hello"] != 3 || sm["another"] != 1 || sm["string"] != 1 {
		t.Errorf("failed to group by with string map: %v", sm)
	}
}

func TestUniqueCopy(t *testing.T) {
	v := []int64{1, 1, 1, 2, 2, 3, 4, 5}
	uniqueV := UniqueCopy(v)
	if len(uniqueV) != 5 {
		t.Errorf("expected 5 unique values, got %v instead", uniqueV)
	}
}

func TestLinearSearch(t *testing.T) {
	if LinearSearch([]string{"a", "b", "c"}, "c") != 2 {
		t.Error("linear search failed to return correct index")
	}

	if LinearSearch([]int32{1, 2, 4, 5}, 3) != -1 {
		t.Error("linear search failed to return -1 when value did not exist")
	}
}

func TestBinarySearch(t *testing.T) {
	if x := BinarySearch([]int32{
		1, 3, 5, 7, 9, 12, 15, 19, 25,
	}, 3); x != 1 {
		t.Errorf("expected 1 got %v", x)
	}

	if x := BinarySearch([]int64{
		2, 4, 5, 9, 12, 23, 50, 1023,
	}, 29); x != -1 {
		t.Errorf("expected -1, got %v instead", x)
	}

	if x := BinarySearch([]float64{}, 10); x != -1 {
		t.Errorf("expected -1, got %v instead", x)
	}
}

func TestMergeSort(t *testing.T) {
	data := [...]int{-15, 2, 13, -20, -32, 23, 53, 102, -13}
	x := sort.IntSlice(data[0:])
	x = MergeSort(x)
	if !sort.IsSorted(x) {
		t.Error("failed to sort data")
	}
	floats := [...]float64{-12351.2, 123.30, 9923.2, -2300.5, math.Inf(-1)}
	y := sort.Float64Slice(floats[0:])
	y = MergeSort(y)
	if !sort.IsSorted(y) {
		t.Error("failed to sort floats")
	}

}
