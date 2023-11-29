/*




You're given a two-dimensional ar·ay (a matrix) of potentially unequal height and width containing only 0 sand 1 s.
The matrix represents a two-toned image, where each 1 represents black and each 0 represents white.
 An island is defined as any number of 1 s that are horizonta ly or vertically adjacent (but not diagonally adjacent) and that don't touch the border of the image. In other words, a group of 1orizontally or vertically adjacent 1 s isn't an island if any of those 1 s are in the first row, last row, first column, or last column of the input matrix.
Note that an island can twist. In other words, it doesn't have to be a straight vertical line or a straight horizontal line; it can be L­shaped, for example.
You can think of islands as patches of black that don't touch the border of the two-toned image.
Write a function that returns a modified version of the input matrix, where all of the islands are removed. You remove an island by replacing it with 0 s.
Naturally, you're allowed to mutate the input matrix.


Sample Input:

matrix = [
    [1, 0, 0, 0, 0, 0],
    [0, 1, 0, 1, 1, 1],
    [0, 0, 1, 0, 1, 0],
    [1, 1, 0, 0, 1, 0],
    [1, 0, 1, 1, 0, 0],
    [1, 0, 0, 0, 0, 1],
    ]


Sample Output:

[
    [1, 0, 0, 0, 0, 0],
    [0, 0, 0, 1, 1, 1],
    [0, 0, 0, 0, 1, 0],
    [1, 1, 0, 0, 1, 0],
    [1, 0, 0, 0, 0, 0],
    [1, 0, 0, 0, 0, 1],
]



// The islands that were removed can be clearly seen here:

[
    [ ,  ,  ,  ,  ,  ],
    [ , 1,  ,  ,  ,  ],
    [ ,  , 1,  ,  ,  ],
    [ ,  ,  ,  ,  ,  ],
    [ ,  , 1, 1,  ,  ],
    [ ,  ,  ,  ,  ,  ],
]


// There are other valid solutions, too.






*/

package main

import "fmt"

func removeIslands(matrix [][]int) {
	// Mark visited cells to avoid removing the same island multiple times
	visited := make(map[[2]int]bool)

	// Explore and remove islands from each 1 in the matrix
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 1 && !visited[[2]int{i, j}] {
				exploreAndRemoveIsland(matrix, i, j, visited)
			}
		}
	}
}

func exploreAndRemoveIsland(matrix [][]int, i, j int, visited map[[2]int]bool) {
	// Check if the current cell is out of bounds or already visited
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) || visited[[2]int{i, j}] {
		return
	}

	// Mark the current cell as visited
	visited[[2]int{i, j}] = true

	// Check if the current cell is an island border cell
	if i == 0 || i == len(matrix)-1 || j == 0 || j == len(matrix[0])-1 {
		matrix[i][j] = 0 // Remove the island border cell
		return
	}

	// Explore adjacent cells
	for _, direction := range directions {
		newI := i + direction[0]
		newJ := j + direction[1]

		// Explore adjacent cells that are part of the same island
		if matrix[newI][newJ] == 1 {
			exploreAndRemoveIsland(matrix, newI, newJ, visited)
		}
	}
}

var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func main() {
	matrix := [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 1},
		{0, 0, 1, 0, 1, 0},
		{1, 1, 0, 0, 1, 0},
		{1, 0, 1, 1, 0, 0},
		{1, 0, 0, 0, 0, 1},
	}

	removeIslands(matrix)

	// Print the modified matrix
	fmt.Println(matrix)
}
