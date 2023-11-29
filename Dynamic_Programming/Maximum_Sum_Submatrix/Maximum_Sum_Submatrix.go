/*
You're given a two-dimensional array (a matrix) of potentially unequal height and width that's filled with integers.
You're also given a positive integer size . Write a function that returns the maximum sum that can be generated from a sub matrix with dimensions size * size ,
For example, consider the following matrix:


[
    [2, 4],
    [5, 6],
    [-3, 2],
]

If size = 2 , then the 2x2 sub matrices to consider are:

[
    [2, 4],
    [5, 6],
]

[
    [5, 6],
    [-3, 2],
]

The sum of the elements in the first sub matrix is 17 , and the sum of the elements in the second sub matrix is 10 .
In this example, your function should return 17 .

Note: You can assume that size is less than or equal to the height and width of the input matrix.

Sample Input

matrix = [
    [5, 3, -1, 5],
    [-7, 3, 7, 4],
    [12, 8, 0, 0],
    [1, -8, -8, 2],
]

size = 2


Sample Output

18


Hints
Hint 1

The brute-force approach to solve this problem involves simply considering all possible submatrices of size size * size, determining their sums, and finally returning the maximum sum. This approach is acceptable, but it isn't optimal. Why isn't it optimal?

Hint 2

The approach stated in Hint #1 isn't optimal because it repeats some additions. When considering submatrices of any size larger than 1, it's almost always the case that some these matrices will have overlapping elements, meaning that we'll repeatedly add up the same numbers. If we were to use the brute-force approach, we would get a time complexity of O(width * height * size). To achieve a more optimal time complexity, we need to avoid readding elements that have already been added. Can you think of a way to solve this problem in O(width * height) time?

Hint 3

To avoid doing repeated addition, we have to use auxiliary space. Ideally, this extra space will allow us to determine the sum of a submatrix of any size in constant time. Start by creating a matrix with the same dimensions as the input matrix (we call this matrix sums). The element at position i, j (where i is the row and j is the column) in this new matrix should be the sum of all the elements in the submatrix whose top left corner is at 0, 0 and whose bottom right corner is at i, j. How can you quickly fill up this new matrix, and how can you then use it to determine the sum of a submatrix of any size in constant time?

Hint 4

The sum of a matrix whose bottom right corner is at i, j (where size <= i <= j) is simply sums[i][j] - sums[i - size][j] - sums[i][j - size] + sums[i - size][j - size]. See the Conceptual Overview section of this question's video explanation for a more in-depth explanation.






*/

package main

import (
	"fmt"
	"math"
)

func maxSubmatrixSum(matrix [][]int, size int) int {
	// Handle the case when the input matrix is empty or `size` is invalid
	if len(matrix) == 0 || size <= 0 {
		return 0
	}

	// Create a prefix sum matrix to store the cumulative sum of elements in each submatrix
	prefixSumMatrix := make([][]int, len(matrix))
	for i := range matrix {
		prefixSumMatrix[i] = make([]int, len(matrix[0]))
	}

	// Initialize the prefix sum matrix
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i == 0 && j == 0 {
				prefixSumMatrix[i][j] = matrix[i][j]
			} else if i == 0 {
				prefixSumMatrix[i][j] = matrix[i][j] + prefixSumMatrix[i][j-1]
			} else if j == 0 {
				prefixSumMatrix[i][j] = matrix[i][j] + prefixSumMatrix[i-1][j]
			} else {
				prefixSumMatrix[i][j] = matrix[i][j] + prefixSumMatrix[i][j-1] + prefixSumMatrix[i-1][j] - prefixSumMatrix[i-1][j-1]
			}
		}
	}

	// Initialize the maximum sum and the starting coordinates of the maximum submatrix
	maxSum := math.MinInt32
	startIndex := []int{0, 0}

	// Iterate over all possible top-left corners of the submatrix
	for i := 0; i <= len(matrix)-size; i++ {
		for j := 0; j <= len(matrix[0])-size; j++ {
			// Calculate the sum of the current submatrix using the prefix sum matrix
			submatrixSum := prefixSumMatrix[i+size-1][j+size-1]
			if i > 0 {
				submatrixSum -= prefixSumMatrix[i-1][j+size-1]
			}
			if j > 0 {
				submatrixSum -= prefixSumMatrix[i+size-1][j-1]
			}
			if i > 0 && j > 0 {
				submatrixSum += prefixSumMatrix[i-1][j-1]
			}

			// Update the maximum sum and the starting coordinates if necessary
			if submatrixSum > maxSum {
				maxSum = submatrixSum
				startIndex = []int{i, j}
			}
		}
	}

	// Return the maximum sum and the starting coordinates of the maximum submatrix
	return maxSum, startIndex
}

func main() {
	// Set the input matrix and size.
	matrix := [][]int{
		{5, 3, -1, 5},
		{-7, 3, 7, 4},
		{12, 8, 0, 0},
		{1, -8, -8, 2},
	}
	size := 2

	// Find the maximum sum of a `size` x `size` submatrix in the given matrix.
	maxSum, startIndex := maxSubmatrixSum(matrix, size)

	// Print the maximum sum and the starting coordinates of the maximum submatrix.
	fmt.Println(maxSum)
	fmt.Println(startIndex)
}
