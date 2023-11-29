/*



A city wants to build a new public park, and you've been tasked with finding the largest park they can build without
disturbing existing infrastructure.
Write a function that takes in a two-dimensional array (a matrix) Land representing the total land of the city from
a top-down view. Each value in land is a boolean; false values are pleces of land not currently in use, while true
values are pleces of land currently in use by other infrastructure. Return the area of the largest possible park.
The largest possible park will be placed exclusively on unused land (false values). Moreover, the city wants the
park to be a perfect rectangle. If there is no available land, your function should return 0.
Note that you're always given at least one piece of land to use (there will always be at least one true value in land).
Also note that you can't place a park on the edges of the city, since that would require extending the park beyond
the boundaries of the city.
Sample Input
land = [
[false, true, true, false],
[false, false, true, false],
[true, true, false, true],
[false, false, true, false],
]
Sample Output
6

// The park should be made up of the land at row indices 1-2 and column indices 2-3, as follows:
// [
// [false, true, true, false],
// [false, false, true, false],
// [true, true, false, true],
// [false, false, true, false],
// ]
// The area of this park is 2 * 3 = 6.







This code first initializes a matrix potentialAreaMatrix to store the maximum potential area of a park starting from each cell. Then, it iterates through the land, calculating the potential area of a park starting from each unused land cell and updating the potentialAreaMatrix. Finally, it iterates through the potentialAreaMatrix to find the largest potential area, which is the largest possible rectangular park area.







*/

package main

import (
	"fmt"
)

func largestRectangleParkArea(land [][]bool) int {
	// Initialize variables for tracking the largest park area and current park area
	largestParkArea := 0
	currentParkArea := 0

	// Initialize a matrix to store the maximum potential area of a park starting from each cell
	potentialAreaMatrix := make([][]int, len(land))
	for i := range potentialAreaMatrix {
		potentialAreaMatrix[i] = make([]int, len(land[0]))
	}

	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[0]); j++ {
			// Check if the current cell is unused land
			if !land[i][j] {
				// Calculate the width of the potential park from the current cell
				width := 1

				// Check if there are unused land cells to the right
				for k := j + 1; k < len(land[0]) && !land[i][k]; k++ {
					width++
				}

				// Calculate the height of the potential park from the current cell
				height := 1

				// Check if there are unused land cells below
				for k := i + 1; k < len(land) && !land[k][j]; k++ {
					height++
				}

				// Calculate the potential area of the park
				currentParkArea = width * height

				// Update the potential area matrix
				potentialAreaMatrix[i][j] = currentParkArea

				// Check if the current park area is larger than the largest
				if currentParkArea > largestParkArea {
					largestParkArea = currentParkArea
				}
			}
		}
	}

	return largestParkArea
}

func main() {
	land := [][]bool{
		{false, true, true, false},
		{false, false, true, false},
		{true, true, false, true},
		{false, false, true, false},
	}

	largestParkArea := largestRectangleParkArea(land)
	fmt.Println("Largest rectangular park area:", largestParkArea)
}
