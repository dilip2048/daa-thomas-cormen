/*
The power is calculated using divide and conquer approach
*/
package main

import "fmt"

// This method will find the power of x of base n
func findPower(x, n int) int {
	if n == 0 {
		return 1
	} else if n%2 == 0 {
		return findPower(x, n/2) * findPower(x, n/2)
	} else {
		return findPower(x, (n-1)/2) * findPower(x, (n-1)/2) * x
	}
}

func main() {
	x := 2
	n := 4
	p := findPower(x, n)
	fmt.Printf("x = %d", p)
}
