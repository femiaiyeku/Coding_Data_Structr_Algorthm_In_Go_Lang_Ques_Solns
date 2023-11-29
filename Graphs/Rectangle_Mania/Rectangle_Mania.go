/*

Write a function that takes In a list of Cartesian coordinates {i.e., {x, y) coordinates)
and returns the number of rectangles formed by these coordinates.
A rectangle must have its four corners amongst the coordinates in order to be counted,
and we only care about rectangles with sides parallel to the x and y axes
{i.e., with horizontal and vertical sides--no diagonal sides).
You can also assume that no coordinate will be farther than 100 units from the origin


Sample Input
coordinates = [
    [0, 0],
    [0, 1],
    [1, 1],
    [1, 0],
    [2, 1],
    [2, 0],
    [3, 1],
    [3, 0]
]

Sample Output

6 // 6 rectangles can be formed by these coordinates:

// 1: [[0, 0], [0, 1], [1, 1], [1, 0]]








*/

package main

import "fmt"

func countRectangles(coordinates [][]int) int {
	count := 0

	// Create a map to store the coordinates
	coordinateMap := make(map[[2]int]bool)
	for _, coord := range coordinates {
		coordinateMap[[2]int{coord[0], coord[1]}] = true
	}

	// Iterate through the coordinates and check for valid rectangles
	for _, coord1 := range coordinates {
		for _, coord2 := range coordinates {
			if coord1[0] == coord2[0] {
				continue // Skip points with the same x-coordinate
			}

			// Check if the other two points exist to form a rectangle
			if coordinateMap[[2]int{coord1[0], coord2[1]}] && coordinateMap[[2]int{coord2[0], coord1[1]}] {
				count++
			}
		}
	}

	return count
}

func main() {
	coordinates := [][]int{
		{0, 0},
		{0, 1},
		{1, 1},
		{1, 0},
		{2, 1},
		{2, 0},
		{3, 1},
		{3, 0},
	}

	rectangleCount := countRectangles(coordinates)
	fmt.Println(rectangleCount)
}
