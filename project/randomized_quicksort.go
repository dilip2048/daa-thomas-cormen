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
	pivot := a[left]
	i := left
	j := right
	for i < j {
		switch piv := pivot.(type) {
		case float64:
			// while you don't find bigger element than pivot, increase left index
			for a[i].(float64) <= piv {
				i++
			}
			// while you don't find bigger element than pivot, decrease right index
			for a[j].(float64) > piv {
				j--
			}
			if i < j {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		case int:
			for i < len(a) && a[i].(int) <= piv {
				i++
			}
			for a[j].(int) > piv {
				j--
			}
			if i < j {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}

	}
	//swap pivot to the correct position, j
	temp := a[left]
	a[left] = a[j]
	a[j] = temp
	return j
}
