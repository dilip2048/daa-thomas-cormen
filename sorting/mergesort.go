/*

This is a divide and conquer algorithms.

* the merge part takes O(n) time
* the combine part take O(nlogn) time
* cons is it take extra space to heapsort the subarrays.

*/

package main

import "fmt"

// This method sorts and merges the sorted array
func merge(a []int, left int, middle int, right int) {
	// the size of left subarray
	n1 := middle - left + 1
	// the size of right subarray
	n2 := right - middle

	// defining temporary left and right slice/array
	var L = make([]int, n1)
	var R = make([]int, n2)

	// assigning elements to left and right array
	for i := 0; i < n1; i++ {
		L[i] = a[left+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = a[middle+1+j]
	}

	i := 0
	j := 0
	k := left

	// compare the left and right array and heapsort it
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			a[k] = L[i]
			i++
		} else {
			a[j] = R[j]
			j++
		}
		k++
	}

	// copy the remaining elements to the array
	for i < n1 {
		a[k] = L[i]
		i++
		k++
	}
	for j < n2 {
		a[k] = R[j]
		j++
		k++
	}
}

// this method will print the sorted array using merge heapsort algo
func printMSortedArray(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
}

// this method will divide the array and then merge them.
func doMergeSort(a []int, left int, right int) {
	if left < right {
		middle := (left + right) / 2
		doMergeSort(a, left, middle)
		doMergeSort(a, middle+1, right)
		merge(a, left, middle, right)
	}
}

func main() {
	a := []int{3, 2, 5, 1, 7, 2}
	doMergeSort(a, 0, len(a)-1)
	printMSortedArray(a)
}
