/*


Write a function that takes in an integer matrix of potentially unequal height and width and returns the minimum number of passes required to convert all negative integers
 in the matrix to positive integers.
A negative integer in the matrix can only be converted to a positive integer if one or more of its adjacent elements is positive.
An adjacent element is an element thct is to the left, to the right, above, or below the current element in the matrix.
Converting a negative to a positive simply involves multiplying it by -1
Note that the 0 value is neither p3sitive nor negative, meaning that a 0 can't convert an adjacent negative to a positive.
A single pass through the matrix in􀀟olves converting all the negative integers that can be converted at a particular point in time.
For example, consider the following input matrix:

[

[0, -2, -1],

    [-5, 2, 0],

    [-6, -2, 0]

    ]

After a first pass, only 3 values can be converted to positives:

[

[0, 2, -1],

[5, 2, 0],

    [6, 2, 0]

    ]

After a second pass, the remaining negative values can all be converted to positives:

    [

    [0, 2, 1],

    [5, 2, 0],

        [6, 2, 0]

        ]

Note that the input matrix will always contain at least one element. If the negative integers in the input matrix can't all be converted to positives, regardless of how many passes are run, your function should return -1.

Sample Input:

matrix = [

[0, -1, -3, 2, 0],

    [1, -2, -5, -1, -3],

    [3, 0, 0, -4, -1]

    ]

Sample Output:

3






*/

package main

import "fmt"

func findMinPasses(matrix [][]int) int {
	passes := 0

	for {
		madeChanges := false

		// Copy the original matrix to avoid modifying the original matrix
		tempMatrix := make([][]int, len(matrix))
		for i := range tempMatrix {
			tempMatrix[i] = make([]int, len(matrix[0]))
			for j := range tempMatrix[i] {
				tempMatrix[i][j] = matrix[i][j]
			}
		}

		// Check each negative integer and convert it to a positive if possible
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[0]); j++ {
				if matrix[i][j] < 0 {
					if hasPositiveAdjacent(matrix, i, j) {
						tempMatrix[i][j] = -matrix[i][j]
						madeChanges = true
					}
				}
			}
		}

		// Update the matrix if changes were made
		if madeChanges {
			matrix = tempMatrix
			passes++
		} else {
			break
		}
	}

	// Check if all negative integers were converted
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] < 0 {
				return -1 // Not all negative integers can be converted
			}
		}
	}

	return passes
}

func hasPositiveAdjacent(matrix [][]int, i, j int) bool {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, direction := range directions {
		newI := i + direction[0]
		newJ := j + direction[1]

		if newI >= 0 && newI < len(matrix) && newJ >= 0 && newJ < len(matrix[0]) && matrix[newI][newJ] > 0 {
			return true
		}
	}

	return false
}

func main() {
	matrix := [][]int{
		{0, -1, -3, 2, 0},
		{1, -2, -5, -1, -3},
		{3, 0, 0, -4, -1},
	}

	minPasses := findMinPasses(matrix)
	fmt.Println(minPasses)
}
