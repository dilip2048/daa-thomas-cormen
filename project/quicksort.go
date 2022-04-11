package main

// this method will put the pivot at the sorted position and return its index
// parameters:
// 		a := the array to be sort
// 		left := the beginning of the list
// 		right := the end of the list
func partition1(a []interface{}, left int, right int) int {
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

// this method will sort the array. this is generic method that will take array of data type integers, and floats
// we will then run another quicksort on the partitioned array.
func quicksort1(arr []interface{}, left int, right int) {
	if left < right {
		pivot := partition1(arr, left, right)
		quicksort1(arr, left, pivot-1)
		quicksort1(arr, pivot+1, right)
	}
}
