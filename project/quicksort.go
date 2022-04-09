package main

// this method will put the pivot at the sorted position and return its index
// parameters:
// 		a := the array to be sort
// 		left := the beginning of the list
// 		right := the end of the list
func partition1(a []interface{}, left int, right int) int {
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

// this method will sort the array. this is generic method that will take array of data type integers, and floats
// we will then run another quicksort on the partitioned array.
func quicksort1(arr []interface{}, left int, right int) {
	if left < right {
		pivot := partition1(arr, left, right)
		quicksort1(arr, left, pivot)
		quicksort1(arr, pivot+1, right)
	}
}
