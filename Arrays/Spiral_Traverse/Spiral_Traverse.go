/*

Write a function that takes in an n x m two-dimensional array (that can be square-shaped when n == m) and
returns a one­-dimensional array of all the array's elements in spiral order.
Spiral order starts at the top left corner of the two-dimensional array, goes to the right,
and proceeds in a spiral pattern all the way until every element has been visited.

Sample Input:
array = [
    [1, 2, 3, 4],
    [12, 13, 14, 5],
    [11, 16, 15, 6],
    [10, 9, 8, 7],
]

Sample Output:
[1, 2, 3, 4, 5, 6, 7, 8, 9, 10 ,11, 12, 13, 14, 15, 16]



Hints
Hint 1

You can think of the spiral that you have to traverse as a set of rectangle perimeters that progressively get smaller (i.e., that progressively move inwards in the two-dimensional array).
Hint 2

Going off of Hint #1, declare four variables: a starting row, a starting column, an ending row, and an ending column. These four variables represent the bounds of the first rectangle perimeter in the spiral that you have to traverse. Traverse that perimeter using those bounds, and then move the bounds inwards. End your algorithm once the starting row passes the ending row or the starting column passes the ending column.
Hint 3

You can solve this problem both iteratively and recursively following very similar logic.



This function works by initializing a result array. The function then initializes the current position and the directions in which the function will move. The function also initializes the current direction.

The function then loops until all of the elements in the array have been visited. For each iteration of the loop, the function adds the current element to the result array and marks the current element as visited. The function then tries to move in the current direction. If the next position is out of bounds or has already been visited, the function changes direction.

The function then moves to the next position and continues the loop. Once the loop has finished, the function returns the result array.


*/

package main

import (
	"fmt"
)

func spiralOrder(array [][]int) []int {
	// Initialize the result array.
	result := make([]int, 0)

	// Initialize the current position.
	currentRow := 0
	currentColumn := 0

	// Initialize the directions in which the function will move.
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	// Initialize the current direction.
	currentDirection := 0

	// Loop until all of the elements in the array have been visited.
	for len(result) < len(array)*len(array[0]) {
		// Add the current element to the result array.
		result = append(result, array[currentRow][currentColumn])

		// Mark the current element as visited.
		array[currentRow][currentColumn] = -1

		// Try to move in the current direction.
		nextRow := currentRow + directions[currentDirection][0]
		nextColumn := currentColumn + directions[currentDirection][1]

		// If the next position is out of bounds or has already been visited, change direction.
		if nextRow < 0 || nextRow >= len(array) || nextColumn < 0 || nextColumn >= len(array[0]) || array[nextRow][nextColumn] == -1 {
			currentDirection = (currentDirection + 1) % 4
		}

		// Move to the next position.
		currentRow = nextRow
		currentColumn = nextColumn
	}

	// Return the result array.
	return result
}

func main() {
	array := [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7},
	}

	result := spiralOrder(array)

	fmt.Println(result)
}
