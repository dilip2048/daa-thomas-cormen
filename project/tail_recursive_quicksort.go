/*
Reference: algorithm from page 188 of Thomas Coreman

Invariant: Pivot is at the correct position

Best Case: After first pass, if the pivot element is at the middle position in the array. T(n) = O(nlogn)

Worst Case: When the array is either already sorted either in ascending order or descending order.
It will take T(n)=T(n-1)+ n = O(n^2)

*/

package main

// this method will sort the array and place pivot at than correct position.
// we will then run another quicksort on the partitioned array.
// Last thing we do is recursion in tail recursion
func tailRecursivequicksort(a []interface{}, left int, right int) {
	for left < right {
		pi := tailPartition(a, left, right)
		tailRecursivequicksort(a, left, right-1)
		left = pi + 1
	}
}

// this method will partition the array around pivot and return pivot's index
func tailPartition(a []interface{}, left int, right int) int {
	pivot := a[right]
	p := left - 1
	for i := left; i < right; i++ {
		switch piv := pivot.(type) {
		case float64:
			// if element is found lower than pivot swap it with pth element
			if a[i].(float64) <= piv {
				//swap
				p++
				a[p], a[i] = a[i], a[p]
			}
		case int:
			// if element is found lower than pivot swap it with pth element
			if a[i].(int) <= piv {
				//swap
				p++
				a[p], a[i] = a[i], a[p]
			}
		}
	}
	//swap pivot with pth index
	a[right], a[p+1] = a[p+1], a[right]
	return p + 1
}
