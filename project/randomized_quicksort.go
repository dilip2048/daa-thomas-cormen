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

func randomizedPartition(a []interface{}, left, right int) int {
	// take random element as pivot
	rand := rand.Intn(right-left) + left
	temp := a[right]
	a[right] = a[rand]
	a[rand] = temp
	return partition(a, left, right)
}

// this method will sort the array and place pivot at than correct position.
// we will then run another randomizedquicksort on the partitioned array.
func randomizedquicksort(arr []interface{}, left int, right int) {
	if left < right {
		pivot := randomizedPartition(arr, left, right)
		randomizedquicksort(arr, left, pivot-1)
		randomizedquicksort(arr, pivot+1, right)
	}
}

func partition(a []interface{}, left int, right int) int {
	pivot := a[right]
	i := left - 1
	for j := left; j <= right-1; j++ {
		switch piv := pivot.(type) {
		case float64:
			if a[j].(float64) < piv {
				i++
				//swap
				a[i], a[j] = a[j], a[i]
			}
		case int:
			if a[j].(int) < piv {
				i++
				//swap
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	//swap pivot with pth index
	a[right], a[i+1] = a[i+1], a[right]
	return i + 1
}
