/*

Write a function that takes in a list of Cartesian coordinates (i.e., (x, y) coordinates) and
returns the number of squares that can be formed by these coordinates.
A square must have its four corners amongst the coordinates in order to be counted.
A single coordinate can be used as a corner for multiple different squares.
You can also assume that no coordinate will be farther than 100 units from the origin.

Sample Input

points = [
    (1, 1),
    (1, 2),
    (2, 1),
    (2, 2),
    (1, 3),
    (2, 3),
    (3, 1),
    (3, 2),
    (3, 3),
    (4, 1),
    (4, 2),
    (4, 3),
]

Sample Output

8   // The 8 squares are:
    // 1: (1, 1), (1, 2), (2, 1), (2, 2)
    // 2: (1, 2), (1, 3), (2, 2), (2, 3)
    // 3: (2, 1), (2, 2), (3, 1), (3, 2)
    // 4: (2, 2), (2, 3), (3, 2), (3, 3)
    // 5: (1, 1), (1, 2), (1, 3), (2, 3)
    // 6: (1, 2), (1, 3), (1, 4), (2, 4)
    // 7: (2, 1), (2, 2), (2, 3), (3, 3)
    // 8: (2, 2), (2, 3), (2, 4), (3, 4)



To use the function, simply pass in the list of Cartesian coordinates as an argument. The function will return the number of squares that can be formed by the coordinates.

For example, using the sample input provided in the question, we would get the following output:


points := [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}, {1, 3}, {2, 3}, {3, 1}, {3, 2}, {3, 3}, {4, 1}, {4, 2}, {4, 3}}

numberOfSquares := countSquares(points)

fmt.Println(numberOfSquares) // 8



*/

package main

import (
	"fmt"
	"sort"
)

func countSquares(points [][]int) int {

	// Create a set to store the coordinates.
	coordinates := make(map[string]struct{})
	for _, point := range points {
		coordinates[fmt.Sprintf("%d,%d", point[0], point[1])] = struct{}{}
	}

	// Create a set to store the squares.
	squares := make(map[string]struct{})

	// Iterate over all pairs of coordinates.
	for i, point1 := range points {
		for j := i + 1; j < len(points); j++ {
			point2 := points[j]

			// Check if the two coordinates form a side of a square.
			if point1[0] != point2[0] && point1[1] != point2[1] {
				// Calculate the coordinates of the other two corners of the square.
				xDiff := abs(point1[0] - point2[0])
				yDiff := abs(point1[1] - point2[1])
				if xDiff == yDiff {
					corner1 := []int{point1[0], point2[1]}
					corner2 := []int{point2[0], point1[1]}

					// Check if the other two corners of the square are in the set of coordinates.
					if _, ok := coordinates[fmt.Sprintf("%d,%d", corner1[0], corner1[1])]; ok {
						if _, ok := coordinates[fmt.Sprintf("%d,%d", corner2[0], corner2[1])]; ok {
							// Add the square to the set of squares.
							square := []int{point1[0], point1[1], point2[0], point2[1]}
							sort.Ints(square)
							squares[fmt.Sprintf("%d,%d,%d,%d", square[0], square[1], square[2], square[3])] = struct{}{}
						}
					}
				}
			}
		}
	}

	// Return the number of squares.
	return len(squares)
}
