/*

You're given a two-dimensional array (a matrix) of potentially unequal height and width containing only es and 1 s.
Each 1 represents water, and each represents part of a land mass. A land mass consists of any number of es
that are either horizontally or vertically adjacent (but not diagonally adjacent). The number of adjacentes forming a
land mass determine its size.
Note that a land mass can twist. In other words, it doesn't have to be a straight vertical line or a straight horizontal line;
it can be L-shaped, for example.
Write a function that returns the largest possible land mass size after changing exactly one 1 to a
given matrix will always contain at least one 1, and you may mutate the matrix.

Sample Input
matrix = [
    [1, 0, 0, 1, 0],
    [1, 0, 1, 0, 0],
    [0, 0, 1, 0, 1],
    [1, 0, 1, 0, 1],
    [1, 0, 1, 1, 0],
]

Sample Output
6
// The 1 at row index 0 and column index 2 can be changed to a 1 to connect to the other 1s and form a land mass of size 6:

// [
    '1', 0, '0', 1, 0,
    '1', 0, '1', 0, 0,
    '0', 0, '1', 0, 1,
    '1', 0, '1', 0, 1,
    '1', 0, '1', '1', 0,
]








*/

package main

import "fmt"

func findLargestLandmass(matrix [][]int) int {
	originalMatrix := matrix
	maxLandmassSize := 0

	// Try changing each 0 to a 1 and find the resulting landmass size
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				// Change the current 0 to a 1
				matrix[i][j] = 1

				// Find the landmass size after changing the 0 to a 1
				landmassSize := findLandmassSize(matrix)

				// Update the maximum landmass size if necessary
				if landmassSize > maxLandmassSize {
					maxLandmassSize = landmassSize
				}

				// Restore the original matrix
				matrix = originalMatrix
			}
		}
	}

	return maxLandmassSize
}

func findLandmassSize(matrix [][]int) int {
	landmassSize := 0

	// Mark visited cells to avoid counting landmass multiple times
	visited := make(map[[2]int]bool)

	// Explore the landmass from each 1 in the matrix
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 1 && !visited[[2]int{i, j}] {
				landmassSize += exploreLandmass(matrix, i, j, visited)
			}
		}
	}

	return landmassSize
}

func exploreLandmass(matrix [][]int, i, j int, visited map[[2]int]bool) int {
	// Check if the current cell is out of bounds or already visited
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) || visited[[2]int{i, j}] {
		return 0
	}

	// Mark the current cell as visited
	visited[[2]int{i, j}] = true

	// Count the current cell as part of the landmass
	landmassSize := 1

	// Explore adjacent cells
	for _, direction := range directions {
		newI := i + direction[0]
		newJ := j + direction[1]

		// Explore adjacent landmass cells
		if matrix[newI][newJ] == 1 {
			landmassSize += exploreLandmass(matrix, newI, newJ, visited)
		}
	}

	return landmassSize
}

var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func main() {
	matrix := [][]int{
		{1, 0, 0, 1, 0},
		{1, 0, 1, 0, 0},
		{0, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 1, 0},
	}

	largestLandmassSize := findLargestLandmass(matrix)
	fmt.Println(largestLandmassSize)
}
