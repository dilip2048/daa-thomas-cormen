package main

/*
this method will take 2^n times.
*/

import "fmt"

// this method will recursively add and return nth fibonacci number
func fibonacci(n int) int {
	//base case to handle 0 and 1
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	x := 8 // 8th fibonacci number is 21 considering first position as 0
	fibo := fibonacci(x)
	fmt.Printf("The nth fibonacci number is %d\n", fibo)
}
