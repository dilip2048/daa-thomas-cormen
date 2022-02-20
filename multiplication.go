/*
The naive method would take O(n^3) to multiply the 2 matrix of size n * p and p * m.

We will also use block matrix multiplication which reduces the solution to O(n^2)
*/

package main

import (
	"fmt"
)

// This method with multiple two matrices and return a final matrix
func multiply(matrix1, matrix2 [][]int) {
	row1 := len(matrix1)
	row2 := len(matrix2)
	column2 := len(matrix2[0])

	// this will be the result matrix of multiplied two matrices
	// the resulted matrix will have the size of row1 * column2
	result := make([][]int, row1)
	for i := 0; i < row1; i++ {
		result[i] = make([]int, column2)
	}

	for i := 0; i < row1; i++ {
		for j := 0; j < column2; j++ {
			for k := 0; k < row2; k++ {
				// k = column of first matrix = row of second matrix
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
			fmt.Printf("%d\t", result[i][j])
		}
		fmt.Println()
	}
}

func main() {
	// defined two matrices
	matrix1 := [][]int{{1, 2, 3}, {4, 5, 6}}
	matrix2 := [][]int{{1}, {4}, {2}}
	multiply(matrix1, matrix2)
}
