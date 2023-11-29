/*


You're given a list of intergers nums.Write a function that returns a boolean representing wehether there exists a zero-sum subarray of nums.

A zero-sum subarray is any subarray whose elements sum up to 0..A subarray is a contiguous part of an array.For the purpose of this question,
a subarray can be as small as one element and as long as the original array.


Sample Input
nums = [-5, -5. 2, 3, -2]

Sample Output
True // The subarray [-5, 2, 3] sums up to 0



This function works by first creating a map to store the prefix sums of the array. The prefix sum of an array is the sum of all the elements of the array up to and including a particular index.

The function then initializes the prefix sum to 0. Then, the function iterates over the array and updates the prefix sum. If the prefix sum is equal to 0, then there exists a zero-sum subarray, and the function returns true.

If the prefix sum is not equal to 0, the function checks if the prefix sum is already present in the map. If it is, then there exists a zero-sum subarray, and the function returns true. Otherwise, the function adds the prefix sum to the map.

After the function has finished iterating over the array, it checks if the map contains the value 0. If it does, then there exists a zero-sum subarray, and the function returns true. Otherwise, the function returns false.


*/

package main

import (
	"fmt"
)

func hasZeroSumSubarray(nums []int) bool {
	// Create a map to store the prefix sums of the array.
	prefixSums := make(map[int]bool)

	// Initialize the prefix sum.
	prefixSum := 0

	// Iterate over the array.
	for _, num := range nums {
		// Update the prefix sum.
		prefixSum += num

		// If the prefix sum is equal to 0, then there exists a zero-sum subarray.
		if prefixSum == 0 {
			return true
		}

		// If the prefix sum is already present in the map, then there exists a zero-sum subarray.
		if prefixSums[prefixSum] {
			return true
		}

		// Add the prefix sum to the map.
		prefixSums[prefixSum] = true
	}

	// If there is no zero-sum subarray, return false.
	return false
}

func main() {
	nums := []int{-5, -5, 2, 3, -2}

	hasZeroSumSubarray := hasZeroSumSubarray(nums)

	fmt.Println(hasZeroSumSubarray)
}
