/*
Invariant: Pivot is at the correct position

Best Case: After first pass, if the pivot element is at the middle position in the array. T(n) = O(nlogn)

Worst Case: When the array is either already sorted either in ascending order or descending order.
It will take T(n)=T(n-1)+ n = O(n^2)

Reference: Thomas Coreman book. Page:179
*/

package main

import (
	"math/rand"
)

func randomizedPartition(a []int, left, right int) int {
	// take random element as pivot
	rand := rand.Intn(right-left) + left
	temp := a[right]
	a[right] = a[rand]
	a[rand] = temp
	return partition(a, left, right)
}

// this method will sort the array and place pivot at than correct position.
// we will then run another randomizedquicksort on the partitioned array.
func randomizedquicksort(arr []int, left int, right int) {
	if left < right {
		pivot := randomizedPartition(arr, left, right)
		randomizedquicksort(arr, left, pivot-1)
		randomizedquicksort(arr, pivot+1, right)
	}
}

func partition(arr []int, left int, right int) int {
	pivot := arr[right]
	p := left - 1
	for i := left; i < right; i++ {
		// if element is found lower than pivot swap it with pth element
		if arr[i] <= pivot {
			//swap
			p++
			temp := arr[p]
			arr[p] = arr[i]
			arr[i] = temp
		}
	}
	//swap pivot with pth index
	temp := arr[right]
	arr[right] = arr[p+1]
	arr[p+1] = temp
	return p + 1
}
