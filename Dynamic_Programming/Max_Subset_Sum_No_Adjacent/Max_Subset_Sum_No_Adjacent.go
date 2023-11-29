/*

Write a function that takes in an array of positive integers and returns the maximum sum of non-adjacent elements in the array. If the input array is empty,
the function should return e .

Sample Input: array = [75, 105, 120, 75, 90, 135]

Sample Output: 330 // 75 + 120 + 135




Hints
Hint 1
Try building an array of the same length as the input array. At each index in this new array, store the maximum sum that can be generated using no adjacent numbers located between index 0 and the current index.

Hint 2
Can you come up with the formula for calculating the maximum sum that can be generated at any given index in the array? Once you've come up with the formula, try storing the maximum sum that can be generated at each index in the input array.

Hint 3
Try storing not just the maximum sum that can be generated at each index, but also the numbers that generate the maximum sum. You can do this by storing arrays of numbers at each index in the input array, or by storing objects with both the sum and the numbers that generate that sum.

Optimal Space & Time Complexity
O(n) time | O(n) space - where n is the length of the input array



*/

package main

import (
	"fmt"
)

func maxNonAdjacentSum(array []int) int {
	// Handle the empty array case
	if len(array) == 0 {
		return 0
	}

	// Create a DP table to store the maximum sums for including and excluding the current element at each index
	dp := make([]int, len(array))

	// Initialize the DP table with the default values of 0
	for i := 0; i < len(array); i++ {
		dp[i] = 0
	}

	// Fill the DP table using bottom-up approach
	dp[0] = array[0]
	for i := 1; i < len(array); i++ {
		// Include the current element if it's greater than the maximum sum excluding the previous element
		if array[i] > dp[i-1] {
			dp[i] = array[i]
		} else {
			// Exclude the current element and consider the maximum sum excluding the previous element
			dp[i] = dp[i-1]
		}

		// Check if including the current element and the maximum sum of non-adjacent elements excluding the current element
		// and the previous element is greater than the current maximum sum
		if dp[i] < array[i]+dp[i-2] {
			dp[i] = array[i] + dp[i-2]
		}
	}

	// Return the maximum sum from the DP table
	return dp[len(array)-1]
}

func main() {
	// Set the input values.
	array := []int{75, 105, 120, 75, 90, 135}

	// Find the maximum sum of non-adjacent elements in the array.
	maxSum := maxNonAdjacentSum(array)

	// Print the maximum sum of non-adjacent elements.
	fmt.Println(maxSum)
}
