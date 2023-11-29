/*


Write a function that takes in a non-empty, unordered array of positive Integers and returns the array's
majority element without sorting the array and without using more than constant space.
An array's majority element is an element of the array that appears in over half of its indices. Note that the
most common element of an array (the element that appears the most times in the array) isn't necessarily the
array's majority element; for example, the arrays [3, 2, 2, 1] and [3, 4, 2, 2, 13 both have 2 as
their most common element, yet neither of these arrays have a majority element, because neither 2 nor any
other element appears in over half of the respective arrays' indices.
You can assume that the input array will always have a majority element.


Sample Input
array= [1, 2, 3, 2, 2, 1, 2]

Sample Output

2   // 2 appears more than 3 times in the input array, and no other number appears more than 3 times in the input array.




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
