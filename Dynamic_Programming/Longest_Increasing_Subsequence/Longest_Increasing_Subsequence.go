/*

Given a non-empty array of integers, write a function that returns the longest strictly-increasing subsequence in the
array.
A subsequence of an array is a set of numbers that aren't necessarily adjacent in the array but that are in the same
order as they appear in the array. For instance, the numbers [1, 3, 4] form a subsequence of the array
[1, 2, 3, 4], and so do the numbers [2, 4]. Note that a single number in an array and the array itself are both
valid subsequences of the array.
You can assume that there will only be one longest increasing subsequence.

Sample Input:

array = [5, 7, -24, 12, 10, 2, 3, 12, 5, 6, 35]

Sample Output:

[-24, 2, 3, 5, 6, 35]



Hints
Hint 1

Try building an array of the same length as the input array. At each index in this new array, store the length of the longest increasing subsequence ending with the number found at that index in the input array.

Hint 2

Can you efficiently keep track of potential sequences in another array? Instead of storing entire sequences, try storing the indices of previous numbers. For example, at index 3 in this other array, store the index of the before-last number in the longest increasing subsequence ending with the number at index 3.

Hint 3

You can optimize your algorithm by taking a slightly different approach. Instead of building the array mentioned in Hint #1, try building an array whose indices represent lengths of subsequences and whose values represent the smallest numbers in the input array that can end a subsequence of a given length. Traverse the input array, and for each number determine what the length L of the longest increasing subsequence ending with that number is; store the position of that number at index L in the new array that you're building. Find a way to use binary search to build this array.






*/

package main

import (
	"fmt"
)

func longestIncreasingSubsequence(array []int) []int {
	// Create a DP table to store the lengths of the longest increasing subsequences for all possible suffixes of the array
	dp := make([]int, len(array))

	// Initialize the DP table with the default value of 1, indicating that each element is a subsequence of itself
	for i := range dp {
		dp[i] = 1
	}

	// Fill the DP table using bottom-up approach
	for i := 1; i < len(array); i++ {
		for j := 0; j < i; j++ {
			if array[i] > array[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}

	// Find the index of the element with the longest increasing subsequence
	maxIndex := 0
	for i := 1; i < len(dp); i++ {
		if dp[i] > dp[maxIndex] {
			maxIndex = i
		}
	}

	// Construct the longest increasing subsequence using backtracking
	lis := []int{}
	maxLength := dp[maxIndex]

	for i := maxIndex; i >= 0 && maxLength > 0; i-- {
		if dp[i] == maxLength && (maxLength == 1 || array[i] > array[lis[len(lis)-1]]) {
			lis = append(lis, array[i])
			maxLength--
		}
	}

	// Reverse the constructed subsequence
	return reverse(lis)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	// Set the input values.
	array := []int{5, 7, -24, 12, 10, 2, 3, 12, 5, 6, 35}

	// Find the longest strictly increasing subsequence in the array.
	lis := longestIncreasingSubsequence(array)

	// Print the longest strictly increasing subsequence.
	fmt.Println(lis)
}
