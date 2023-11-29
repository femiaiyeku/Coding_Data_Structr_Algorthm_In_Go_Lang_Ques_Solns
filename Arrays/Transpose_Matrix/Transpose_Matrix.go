/*


You're given a 2D array of integers matrix. Write a function that returns the transpose of the matrix.
The transpose of a matrix is a flipped version of the original matrix across its main diagonal (which runs from top-
left to bottom-right); it switches the row and column indices of the original matrix.
You can assume the input matrix always has at least 1 value; however its width and height are not necessarily the
same.


Sample Input
matrix = [
[1, 2],
]

Sample Output
[
[1],
[2],
]




This function works by first initializing the transpose matrix. The function then iterates over the rows of the original matrix. For each row, the function creates a new row in the transpose matrix. The function then iterates over the columns of the original matrix and sets the element at (i, j) in the transpose matrix to the element at (j, i) in the original matrix. Finally, the function adds the new row to the transpose matrix.

Once the function has finished iterating over the rows of the original matrix, it returns the transpose matrix.

*/

package main

import (
	"fmt"
)

func transpose(matrix [][]int) [][]int {
	// Initialize the transpose matrix.
	transposeMatrix := make([][]int, len(matrix[0]))

	// Iterate over the rows of the original matrix.
	for i := range matrix {
		// Create a new row in the transpose matrix.
		transposeRow := make([]int, len(matrix))

		// Iterate over the columns of the original matrix.
		for j := range matrix[0] {
			// Set the element at (i, j) in the transpose matrix to the element at (j, i) in the original matrix.
			transposeRow[i] = matrix[j][i]
		}

		// Add the new row to the transpose matrix.
		transposeMatrix[i] = transposeRow
	}

	// Return the transpose matrix.
	return transposeMatrix
}

func main() {
	matrix := [][]int{
		{1, 2},
	}

	transposeMatrix := transpose(matrix)

	fmt.Println(transposeMatrix)
}
