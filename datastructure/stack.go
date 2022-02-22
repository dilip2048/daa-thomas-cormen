/*
It follows Last in Last Out(LIFO)

Push: O(1)
Pop: O(1)
*/

package main

import (
	"errors"
	"fmt"
)

// ArrayStack Generic type structure to create stack using simple array
type ArrayStack struct {
	top      int
	capacity int
	arr      []interface{}
}

// This will create a new stack
func newArrayStack(capacity int) *ArrayStack {
	var stack ArrayStack
	stack.top = -1
	stack.capacity = capacity
	return &stack
}

// This method will check if the stack is empty
func isEmtpty(stack *ArrayStack) bool {
	return stack.top == -1
}

// This method will check if the stack is full
func isFull(stack *ArrayStack) bool {
	return stack.top == stack.capacity-1
}

// This method will return to top element in the stack
func top(stack *ArrayStack) interface{} {
	return stack.arr[stack.top]
}

// This method will push an element into stack
func push(stack *ArrayStack, a interface{}) error {
	if isFull(stack) {
		return errors.New("stack is full")
	} else {
		stack.arr = append(stack.arr, a)
		stack.top++
	}
	return nil
}

// This method will pop the top element from the stack
func pop(stack *ArrayStack) (interface{}, error) {
	var element interface{}
	if isEmtpty(stack) {
		return nil, errors.New("stack is empty")
	} else {
		element = stack.arr[stack.top]
		stack.top--
	}
	return element, nil
}

func main() {
	// created a stack object
	stack := newArrayStack(5)
	// push elements into the stack
	_ = push(stack, 10)
	_ = push(stack, 11)
	_ = push(stack, 12)
	_ = push(stack, 13)

	// pop from stack
	x, err := pop(stack)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\nThe popped element is :%d\n", x)

	// check if stack is empty
	b := isEmtpty(stack)
	if b {
		fmt.Println("The stack is empty")
	} else {
		fmt.Println("This is not empty")
	}

	// print the top element of the stack
	top := top(stack)
	fmt.Printf("the top element of the stack is %d", top)
}
