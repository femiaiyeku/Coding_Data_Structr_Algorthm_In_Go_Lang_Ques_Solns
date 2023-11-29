/*

Write a function that takes in a non-empty array of non-negative integers and a non-negative integer representing a target sum.The function should find the
longest subarray where the values collectively sum up to equal the target  sum. Return an array containing the starting index and ending index of this subarray both inclusive.
If there is no subarray that sums up top the target sunm, the function should return an empty array. You can assume that the given inputs will only ever have one answer.

Sample Input:
array = [1,2,3,4,3,3,1,2,1,2]

targetSum = 10


Sample Output:

[4, 8] // The longest subarray that sums to 10 starts at index 4 and ends at index 8.





*/

package main

import (
	"fmt"
)

func findLongestSubarrayWithTargetSum(array []int, targetSum int) []int {
	// Create a map to store the current sum and the starting index of the subarray at that sum.
	sumIndexMap := make(map[int]int)

	// Initialize the current sum and the starting index of the longest subarray.
	currentSum := 0
	longestStartIndex := 0
	longestEndIndex := 0
	longestSubarrayLength := 0

	// Iterate over the array and keep track of the current sum.
	for i := 0; i < len(array); i++ {
		currentSum += array[i]

		// Check if the current sum is equal to the target sum.
		if currentSum == targetSum {
			// Update the longest subarray if necessary.
			if i+1 > longestSubarrayLength {
				longestSubarrayLength = i + 1
				longestStartIndex = 0
				longestEndIndex = i
			}
		}

		// Check if the current sum minus the target sum is in the map.
		if currentSum-targetSum >= 0 && sumIndexMap[currentSum-targetSum] > 0 {
			// Update the longest subarray if necessary.
			if i-sumIndexMap[currentSum-targetSum] > longestSubarrayLength {
				longestSubarrayLength = i - sumIndexMap[currentSum-targetSum]
				longestStartIndex = sumIndexMap[currentSum-targetSum] + 1
				longestEndIndex = i
			}
		}

		// Add the current sum to the map.
		sumIndexMap[currentSum] = i
	}

	// If there is no subarray that sums to the target sum, return an empty array.
	if longestSubarrayLength == 0 {
		return []int{}
	}

	// Return an array containing the starting index and ending index of the longest subarray.
	return []int{longestStartIndex, longestEndIndex}
}

func main() {
	array := []int{1, 2, 3, 4, 3, 3, 1, 2, 1, 2}

	targetSum := 10

	longestSubarray := findLongestSubarrayWithTargetSum(array, targetSum)

	fmt.Println(longestSubarray)
}
