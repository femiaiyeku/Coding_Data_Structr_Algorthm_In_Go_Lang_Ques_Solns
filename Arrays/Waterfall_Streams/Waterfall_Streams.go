/*


You're given a two-dimensional array that represents the structure of an indoor waterfall and a positive integer that represents the column that the waterfall's water source will start at. More specifically, the water source will start directly above the structure and will flow downwards.
Each row in the array contains 0 s and 1 s, where a 0 represents a free space and a 1 represents a block that water can't pass through. You can imagine that the last row of the array contains buckets that the water will eventually flow into; thus, the last row of the array will always contain only 0 s. You can also imagine that there are walls on both sides of the structure, meaning that water will never leave the structure; it will either be trapped against a wall or flow into one of the buckets in the last row.
As water flows downwards, if it hits a block, it splits evenly to the left and right-hand side of that block. In other words, 50% of the water flows left and 50% of it flows right. If a water stream is unable to flow to the left or to the right (because of a block or a wall), the water stream in question becomes trapped and can no longer continue to flow in that direction; it effectively gets stuck in the structure and can no longer flow downwards, meaning that 50% of the previous water stream is forever lost.
Lastly, the input array will always contain at least two rows and one column, and the space directly below the water source (in the first row of the array) will always be empty, allowing the water to start flowing downwards.

Explanation
Time complexity: O(wA2*h)
Space Complexity: O(w)
The only part of this solution that consumes space is the recursive call stack which in the worst case
will scale with the width of the array where the 1 s form a ladder from the top right to the bottom left corner.
This is a breadth first sea ch algorithm. The time complexity is the exact same as that of the A Igo Expert solution.
The traversal method is divided up into downward and sideways traversal for a clean readable implementation.

Sample Input
array = [
    [0, 0, 0, 0, 0, 0, 0],
    [1, 0, 0, 1, 1, 1, 0],
    [0, 0, 0, 0, 1, 1, 0],
    [0, 0, 0, 0, 0, 0, 0],
    [1, 1, 0, 1, 1, 0, 1],
    [0, 1, 0, 0, 0, 0, 1],
    [0, 1, 0, 0, 0, 0, 0],
]
source = 3

Sample Output
[0, 0, 0, 0, 0, 0, 0]
[1, 1, 1, 0, 0, 0, 0]
[0, 0, 0, 0, 0, 0, 0]
[0, 0, 1, 1, 1, 0, 0]
[0, 0, 0, 0, 1, 1, 0]
[0, 0, 0, 0, 0, 1, 0]
[0, 0, 0, 0, 0, 1, 0]



Hints
Hint 1

Try not to overthink the solution to this problem. If you were to manually go through an example of water flowing downwards through the waterfall structure, what steps would you follow exactly? Can you simply transcribe these steps into code?

Hint 2

To start simple, consider how you would solve this problem if there were only two rows. How would you make water flow from the first row to the second row with your code? Can you make a slight modification to this approach in order to solve this problem for any number of rows?

Hint 3

You'll want to traverse through the input array, all the while keeping track of where and how much water flows. To do this, you'll need to represent water with some value (-1, for example, to distinguish it from the other values in the array). Iterate through the input array, row by row, column by column, specifically looking at each current row and the row above it. When you see water in the row above, you'll have to reiterate through both the row above and the current row to see where the water will flow to next (i.e., whether there are open spaces allowing the water to flow sideways and / or downwards), mutating these rows along the way whenever water does flow. You'll have the make sure to keep track of the percentage of water that's flowing whenever water gets split in half.



This function works by first initializing the result array. The function then copies the original array to the result array. Next, the function creates a queue to store the water streams.

The function then starts a while loop that runs until the queue is empty. In each iteration of the loop, the function gets the next water stream from the queue. Then, the function gets the row and column of the current water stream.

If the current water stream is in the last row, then it has reached the bottom and can no longer flow downwards. The function skips to the next iteration of the loop.

If there is a block below the current water stream, then the water stream splits evenly to the left and right-hand side of the block. The function adds the two new water streams to the queue.

If there is no block below the current water stream, then the water stream flows downwards. The function increments the water level at the position of the current water stream in the result array. The function then



*/

package main

import (
	"fmt"
)

func simulateWaterfall(array [][]int, source int) [][]int {
	// Initialize the result array.
	result := make([][]int, len(array))

	// Copy the original array to the result array.
	for i, row := range array {
		result[i] = make([]int, len(row))
		copy(result[i], row)
	}

	// Create a queue to store the water streams.
	queue := [][]int{{0, source}}

	// While the queue is not empty, keep flowing the water.
	for len(queue) > 0 {
		// Get the next water stream from the queue.
		currentWaterStream := queue[0]
		queue = queue[1:]

		// Get the row and column of the current water stream.
		currentRow := currentWaterStream[0]
		currentColumn := currentWaterStream[1]

		// If the current water stream is in the last row, then it has reached the bottom and can no longer flow downwards.
		if currentRow == len(array)-1 {
			continue
		}

		// If there is a block below the current water stream, then the water stream splits evenly to the left and right-hand side of the block.
		if array[currentRow+1][currentColumn] == 1 {
			queue = append(queue, [][]int{{currentRow + 1, currentColumn - 1}, {currentRow + 1, currentColumn + 1}}...)
			continue
		}

		// If there is no block below the current water stream, then the water stream flows downwards.
		result[currentRow+1][currentColumn]++
		queue = append(queue, [][]int{{currentRow + 1, currentColumn}}...)
	}

	// Return the result array.
	return result
}

func main() {
	array := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 1, 1, 0},
		{0, 0, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 1, 1, 0, 1},
		{0, 1, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 0},
	}
	source := 3

	result := simulateWaterfall(array, source)

	fmt.Println(result)
}
