package slices

import "testing"

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
