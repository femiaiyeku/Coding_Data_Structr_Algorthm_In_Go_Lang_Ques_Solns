/*



Write a function that takes in a square-shaped n x n two-dimensional array of only 1 s and Os and returns a boolean representing
whether the input matrix contains a square whose borders are made up of only Os.
Note that a 1 x 1 square doesn't count as a valid square for the purpose of this question.
 In other words, a singular e in the input matrix doesn't constitute a square whose borders are made up of only Os; a square of zeroes has to be at least 2 x 2.


Sample Input:
matrix = [
    [1, 1, 1, 0, 1, 0],
    [0, 0, 0, 0, 0, 1],
    [0, 1, 0, 1, 1, 1],
    [0, 0, 0, 1, 0, 1],
    [1, 1, 1, 0, 1, 0],
    [0, 1, 0, 0, 0, 0]
]

Sample Output:

true



Hints
Hint 1

For the purpose of this question, a square is defined by its topmost and bottommost rows and by its leftmost and rightmost columns. Given a pair of rows and a pair of columns that form a valid square, you can easily determine if the relevant square is a square of zeroes with two for loops.

Hint 2

You can apply the logic described in Hint #1 on every valid square in the input matrix in order to solve this problem. To find every valid square, you can either traverse the matrix iteratively with three nested loops, or you can start out at the outtermost square and recursively go inwards in the matrix, checking the squares obtained by moving each corner of a square inwards. If you go with this recursive approach, you'll need to use a cache to avoid doing many duplicate computations.

Hint 3

The operation described in Hint #1 is a computationally expensive one to have to repeat for every single square in the matrix. Can you precompute certain values to make this operation a constant-time operation?

Hint 4

You can make the operation described in Hint #1 a constant-time operation by precomputing some values in the matrix. Specifically, you can precompute two values for every element in the matrix: the number of 0s to the right of each element (including the element itself) and the number of 0s below each element (including the element itself). You can compute these values by iterating through the matrix starting at the bottom right corner and moving your way up by traversing each row from right to left; applying some simple dynamic programming techniques will allow you to build up these values trivially. Once you have these values precomputed, you can perform the operation described in Hint #1 in constant time just by looking at the number of 0s below any square's two top corners and the number of 0s to the right of the same square's two left corners.





*/

package main

import (
	"fmt"
)

func hasSquareWithZeroBorders(matrix [][]int) bool {
	// Check if the matrix is a square
	if len(matrix) != len(matrix[0]) {
		return false
	}

	// Check if the matrix is at least 2 x 2
	if len(matrix) < 2 {
		return false
	}

	// Check for a square with zero borders
	for i := 1; i < len(matrix)-1; i++ {
		for j := 1; j < len(matrix[0])-1; j++ {
			if matrix[i][j] == 1 {
				continue
			}

			// Check the top border
			topBorderAllZeros := true
			for k := 0; k < j; k++ {
				if matrix[i][k] == 1 {
					topBorderAllZeros = false
					break
				}
			}

			// Check the left border
			leftBorderAllZeros := true
			for k := 0; k < i; k++ {
				if matrix[k][j] == 1 {
					leftBorderAllZeros = false
					break
				}
			}

			// Check the right border
			rightBorderAllZeros := true
			for k := j + 1; k < len(matrix[0]); k++ {
				if matrix[i][k] == 1 {
					rightBorderAllZeros = false
					break
				}
			}

			// Check the bottom border
			bottomBorderAllZeros := true
			for k := i + 1; k < len(matrix); k++ {
				if matrix[k][j] == 1 {
					bottomBorderAllZeros = false
					break
				}
			}

			// If all borders are zeros, a square with zero borders exists
			if topBorderAllZeros && leftBorderAllZeros && rightBorderAllZeros && bottomBorderAllZeros {
				return true
			}
		}
	}

	// If no square with zero borders was found, return false
	return false
}

func main() {
	// Set the input matrix.
	matrix := [][]int{
		{1, 1, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 1},
		{0, 1, 0, 1, 1, 1},
		{0, 0, 0, 1, 0, 1},
		{1, 1, 1, 0, 1, 0},
		{0, 1, 0, 0, 0, 0},
	}

	// Check if the matrix contains a square with zero borders.
	hasSquare := hasSquareWithZeroBorders(matrix)

	// Print the result.
	fmt.Println(hasSquare)
}
