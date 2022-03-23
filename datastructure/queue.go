/*
it follows first in first out paradigm

https://www.tutorialspoint.com/data_structures_algorithms/dsa_queue.htm#:~:text=Queue%20is%20an%20abstract%20data,first%20will%20be%20accessed%20first.
*/

package main

import (
	"errors"
	"fmt"
)

// ArrayQueue Generic type structure to create queue using simple array
type ArrayQueue struct {
	front    int
	rear     int
	capacity int
	arr      []interface{}
}

// This will create a new queue
func newArrayQueue(capacity int) *ArrayQueue {
	var queue ArrayQueue
	queue.front = 0
	queue.rear = -1
	queue.capacity = capacity
	return &queue
}

// This method will check if the queue is empty
func isQueueEmtpty(queue *ArrayQueue) bool {
	return queue.rear == -1
}

// This method will check if the queue is full
func isQueueFull(queue *ArrayQueue) bool {
	return queue.rear == queue.capacity
}

// This method will return to top element in the queue
func peek(queue *ArrayQueue) interface{} {
	return queue.arr[queue.front]
}

// This method will push an element into queue
func enqueue(queue *ArrayQueue, a interface{}) error {
	if isQueueFull(queue) {
		return errors.New("queue is full")
	} else {
		queue.arr = append(queue.arr, a)
		queue.rear++
	}
	return nil
}

// This method will dequeue the top element from the queue
func dequeue(queue *ArrayQueue) (interface{}, error) {
	var element interface{}
	if isQueueEmtpty(queue) {
		return nil, errors.New("queue is empty")
	} else {
		element = queue.arr[queue.front]
		for i := 0; i < queue.rear; i++ {
			queue[i] = queue[i+1]
		}
		queue.rear--1
	}
	return element, nil
}

func main() {
	// created a queue object
	queue := newArrayQueue(5)
	// push elements into the queue
	_ = enqueue(queue, 10)
	_ = enqueue(queue, 11)
	_ = enqueue(queue, 12)
	_ = enqueue(queue, 13)

	// dequeue from queue
	x, err := dequeue(queue)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\nThe popped element is :%d\n", x)

	// check if queue is empty
	b := isQueueEmtpty(queue)
	if b {
		fmt.Println("The queue is empty")
	} else {
		fmt.Println("The queue is not empty")
	}

	// print the front element of the queue
	top := peek(queue)
	fmt.Printf("the front element of the queue is %d", top)
}
