/*

Binary searching is also a divide and conquer problem
It takes t(n) = t(n/2) + O(n) time which falls under case 2 - O(log n)

*/

package main

import "fmt"

// This method will search for a number in a sorted array
func binarySearch(a []int, left, right, x int) int {
	middle := (left + right) / 2
	if x == a[middle] {
		return middle
	} else if x > a[middle] {
		return binarySearch(a, middle+1, right, x)
	} else if x < a[middle] {
		return binarySearch(a, left, middle, x)
	}
	return -1
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	index := binarySearch(a, 0, len(a)-1, 3)
	fmt.Printf("index = %d", index)
}
