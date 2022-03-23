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
func tailRecursivequicksort(arr []int, left int, right int) {
	for left < right {
		pi := tailPartition(arr, left, right)
		tailRecursivequicksort(arr, left, right-1)
		left = pi + 1
	}
}

// this method will partition the array around pivot and return pivot's index
func tailPartition(arr []int, left int, right int) int {
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

//func main() {
//	array := []int{3, 5, 6, 2, 1, 9, 8}
//	tailRecursivequicksort(array, 0, 6)
//	fmt.Printf("%v", array)
//}
