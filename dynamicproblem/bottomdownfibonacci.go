package main

/*
this method will take O(n) times which better than traditional algorithm, which takes 2^n.

here we use memoization to store the value of the subproblem, and hence it is called as dynamic problem. Bottom down is a
iterative method
*/

import "fmt"

// BottomDownfibonacci method will recursively add and return nth fibonacci number
func BottomDownfibonacci(n int, memo []int) int {
	//base cases to handle 0th and 1st position
	memo[0] = 0
	memo[1] = 1
	for k := 2; k <= n; k++ {
		memo[k] = memo[k-1] + memo[k-2]
	}
	return memo[n]
}

func OptimizedBottomUp(n int) int {
	a := 0
	b := 1
	if n <= 1 {
		return n
	}
	for i := 2; i < n; i++ {
		c := a + b
		a = b
		b = c
	}
	return b // b will be the nth Fibonacii number, c is the n + 1th fibo number
}

func main() {
	x := 8 // 8th fibonacci number is 21

	// define an memo array for memoization
	memo := make([]int, x+1)
	fibo := BottomDownfibonacci(x, memo)
	fmt.Printf("The nth fibonacci number is %d\n", fibo)

	n := OptimizedBottomUp(x)
	fmt.Printf("(Optimized)The nth fibonacci number is %d\n", n)
}
