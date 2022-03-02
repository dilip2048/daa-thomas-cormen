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

func main() {
	x := 8 // 8th fibonacci number is 21

	// define an memo array for memoization
	memo := make([]int, x+1)
	fibo := BottomDownfibonacci(x, memo)
	fmt.Printf("The nth fibonacci number is %d\n", fibo)
}
