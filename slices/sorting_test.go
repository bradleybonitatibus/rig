package slices

import (
	"math"
	"sort"
	"testing"
)

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
