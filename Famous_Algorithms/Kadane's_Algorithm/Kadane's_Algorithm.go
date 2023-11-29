/*


Write a function that takes in a non-empty array of integers and returns the maximum sum that can be obtained by
summing up all of the integers in a non-empty subarray of the input array. A subarray must only contain adjacent
numbers (numbers next to each other in the input array).



Sample Input

array = [3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4]



Sample Output

19 // sum of the subarray: [1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1]







*/

package main

import (
	"fmt"
)

func maxSubarraySum(array []int) int {
	// Initialize current and maximum sums
	currentSum := 0
	maxSum := array[0]

	// Iterate through the array
	for _, num := range array {
		// Update current sum
		currentSum = max(currentSum+num, num)

		// Update maximum sum
		maxSum = max(maxSum, currentSum)
	}

	// Return the maximum sum
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	// Set the input array
	array := []int{3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4}

	// Find the maximum sum of a subarray
	maxSum := maxSubarraySum(array)

	// Print the maximum sum
	fmt.Println(maxSum)
}
