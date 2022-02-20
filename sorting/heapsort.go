/*
 Heap is a nearly complete binary tree i.e left is filled from left to right

 Invariant:
	minHeap: Parent is lower than the children
	maxHeap: Parent is greater than than children
*/

package main

import "fmt"

func heapsort(a []int) {
	buildMaxHeap(a)
	for i := len(a) - 1; i >= 1; i-- {
		// swapping root with the end element
		temp := a[0]
		a[0] = a[i]
		a[i] = temp
		// run heapify on the root node every time
		maxHeapify(a[:i-1], 0)
	}

}

func buildMaxHeap(a []int) {
	for i := (len(a) - 1) / 2; i >= 0; i-- {
		maxHeapify(a, i)
	}
}

// Max Heapify maintain max heap invariate
func maxHeapify(a []int, i int) {
	largest := i
	// get leftChild child of a[i]
	leftChild := i*2 + 1
	// get rightChild chaild of a[i]
	rightChild := i*2 + 2
	if leftChild < len(a) && a[leftChild] > a[i] {
		largest = leftChild
	}
	if rightChild < len(a) && a[rightChild] > a[largest] {
		largest = rightChild
	}
	// heap invariate is not satisfied if below's condition is not met
	if largest != i {
		// exchange a[i] and largest element by swapping
		// swapping in golang
		a[i], a[largest] = a[largest], a[i]
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
	A := []int{3, 2, 4, 5}
	fmt.Print("The arry before heapsort: ")
	printHeap(A)
	heapsort(A)
	fmt.Print("The arry after heapsort: ")
	printHeap(A)
}
