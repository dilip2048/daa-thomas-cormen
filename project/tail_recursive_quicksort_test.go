package main

import (
	"reflect"
	"testing"
)

func TestRecursiveQuickSortBaseCase(t *testing.T) {
	array := []int{0}
	quicksort(array, 0, 0)
	if !reflect.DeepEqual(array, []int{0}) {
		t.Fail()
	}
}

func TestRecursiveQuickSort1(t *testing.T) {
	array := []int{3, 5, 6, 2, 1, 9, 8, 4, 7, 0}
	quicksort(array, 0, 9)
	if !reflect.DeepEqual(array, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Fail()
	}
}

func TestRecursiveQuickSort2(t *testing.T) {
	array := []int{3, 5, 6, 2, 1, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 8, 4, 7, 0}
	quicksort(array, 0, 19)
	if !reflect.DeepEqual(array, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}) {
		t.Fail()
	}
}
