package main

/*
this method will take O(n) times which better than traditional algorithm, which takes 2^n.

here we use memoization to store the value of the subproblem, and hence it is called as dynamic problem.
*/

import "fmt"

// TopDownfibonacci method will recursively add and return nth fibonacci number
func TopDownfibonacci(n int, memo []int) int {
	//base case to handle 0 and 1
	if n <= 1 {
		memo[n] = n
	} else {
		memo[n] = TopDownfibonacci(n-1, memo) + TopDownfibonacci(n-2, memo)
	}
	return memo[n]
}

func main() {
	x := 8 // 8th fibonacci number is 21

	// define an memo array for memoization
	memo := make([]int, x+1)

	fibo := TopDownfibonacci(x, memo)
	fmt.Printf("The nth fibonacci number is %d\n", fibo)
}
