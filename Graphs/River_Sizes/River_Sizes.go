/*


You're given a two-dimensional array (a matrix) of potentially unequal height and width containing only 0 s and 1 s. Each 0 represents land, and each 1 represents part of a river .. 0.. river consists of any number of 1 s that are either horizontally or vertically adjacent (but not diagonally adjacent). The number of adjacent 1 s forming a river determine its size.
Note that a river can twist. In other words, it doesn't have to be a straight vertical line or a straight horizontal line; it can be L-shaped, for example.
Write a function that returns an array of the sizes of all rivers represented in the input matrix. The sizes don't need to be in any particular order.

Sample Input:

matrix = [
    [1, 0, 0, 1, 0],
    [1, 0, 1, 0, 0],
    [0, 0, 1, 0, 1],
    [1, 0, 1, 0, 1],
    [1, 0, 1, 1, 0]
]

Sample Output:

    [1, 2, 2, 2, 5] # The numbers could be ordered differently.




Hints
Hint 1
Since you must return the sizes of rivers, which consist of horizontally and vertically adjacent 1s in the input matrix, you must somehow keep track of groups of neighboring 1s as you traverse the matrix. Try treating the matrix as a graph, where each element in the matrix is a node in the graph with up to 4 neighboring nodes (above, below, to the left, and to the right), and traverse it using a popular graph-traversal algorithm like Depth-first Search or Breadth-first Search.
Hint 2
By traversing the matrix using DFS or BFS as mentioned in Hint #1, any time that you encounter a 1 you can traverse the entire river that this 1 is a part of (and keep track of its size) by simply iterating through the given node's neighboring nodes and their own neighboring nodes so long as the nodes are 1s.
Hint 3
Naturally, many nodes in the graph mentioned in Hint #1 will have overlapping neighboring nodes, and as you traverse the matrix, you will undoubtedly encounter nodes that you have previously visited. In order to prevent mistakenly calculating the same river's size multiple times and to avoid doing needless computational work, try keeping track of every node that you visit in an auxiliary data structure and only performing important computations on unvisited nodes. What data structure would be ideal here?



*/

package main

import (
	"fmt"
	"sort"
)

func findRiverSizes(matrix [][]int) []int {
	riverSizes := []int{}

	// Mark visited cells to avoid counting river multiple times
	visited := make(map[[2]int]bool)

	// Explore and count rivers from each 1 in the matrix
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 1 && !visited[[2]int{i, j}] {
				currentRiverSize := exploreAndCountRiver(matrix, i, j, visited)
				riverSizes = append(riverSizes, currentRiverSize)
			}
		}
	}

	// Sort the river sizes in ascending order
	sort.Sort(riverSizes)

	return riverSizes
}

func exploreAndCountRiver(matrix [][]int, i, j int, visited map[[2]int]bool) int {
	riverSize := 1

	// Check if the current cell is out of bounds or already visited
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) || visited[[2]int{i, j}] {
		return 0
	}

	// Mark the current cell as visited
	visited[[2]int{i, j}] = true

	// Explore and count adjacent river cells
	for _, direction := range directions {
		newI := i + direction[0]
		newJ := j + direction[1]

		// Explore adjacent cells that are part of the same river
		if matrix[newI][newJ] == 1 {
			riverSize += exploreAndCountRiver(matrix, newI, newJ, visited)
		}
	}

	return riverSize
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

	riverSizes := findRiverSizes(matrix)
	fmt.Println(riverSizes)
}
