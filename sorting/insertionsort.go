/*
Tuesday, 01/18

Invariant: All left elements are sorted in each iteration in insertion heapsort.
Example: Sorting playing card

x*/

package main

import "fmt"

// this method will take array as input and return sorted array
func performInsertionSort(a []int) []int {
	for j := 1; j < len(a); j++ {
		key := a[j]
		i := j - 1
		for i >= 0 && a[i] > key {
			a[i+1] = a[i]
			i = i - 1
		}
		// place key at this correct position
		a[i+1] = key
	}
	return a
}

// this method will print the sorted array
func iPrintArray(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
}

func main() {
	a := []int{4, 3, 5, 2, 1, 10}
	s := performInsertionSort(a)
	iPrintArray(s)
}
