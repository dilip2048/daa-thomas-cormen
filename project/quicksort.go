package main

import "fmt"

// this method will put the pivot at the sorted position and return its index
// parameters:
// 		a := the array where pivot is the sorted
// 		left := the beginning of the list
// 		right := the end of the list
func partition1(a []int, left int, right int) int {
	pivot := a[left]

	for left < right {
		// while you don't find bigger element than pivot, increase left index
		for a[pivot] <= a[left] {
			left++
		}
		// while you don't find bigger element than pivot, decrease right index
		for a[pivot] >= a[right] {
			right--
		}
		if left < right {
			temp := a[left]
			a[left] = a[right]
			a[right] = temp
		}
	}
	//swap pivot to the correct position, j
	temp := a[right]
	a[right] = a[left]
	a[left] = temp
	return right
}

// this method will sort the array and place pivot at than correct position.
// we will then run another quicksort on the partitioned array.
func quicksort1(arr []int, left int, right int) {
	if left < right {
		pivot := randomizedPartition(arr, left, right)
		quicksort1(arr, left, pivot)
		quicksort1(arr, pivot+1, right)
	}
}

func main() {
	array := []int{13, 43, 44, 3, 57}
	quicksort1(array, 0, 4)
	fmt.Printf("%v", array)
}
