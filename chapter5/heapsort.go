package main

import "fmt"

// Max Heapify maintain max heap invariate
func maxHeapify(a []int, i int) {
	var largest int
	// get left child of a[i]
	l := i*2 + 1
	// get right chaild of a[i]
	r := i*2 + 2
	if l <= len(a) && a[l] > a[i] {
		largest = l
	} else {
		largest = i
	}
	if r < len(a) && a[r] > a[largest] {
		largest = r
	}
	// heap invariate is not satisfied if below's condition is not met
	if largest != i {
		// exchange a[i] and largest element by swapping
		temp := a[i]
		a[i] = a[largest]
		a[largest] = temp
		maxHeapify(a, largest)
	}
}

// prints heap
func printHeap(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()
}

func main() {
	A := []int{1, 5, 3, 2}
	printHeap(A)
	maxHeapify(A, 0)
	printHeap(A)
}
