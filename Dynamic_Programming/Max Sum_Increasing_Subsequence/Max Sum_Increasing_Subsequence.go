/*


Write a function that takes in a non-empty array of integers and returns the greatest sum that can be generated from a strictly-increas1ng subsequence
In the array as well as an array of the numbers 1n that subsequence.
A subsequence of an array 1s a set of numbers that aren't necessarily adJacent 1n the array but that are 1n the same order as they appear in the array.
For instance, the numbers [l, 3, 4] form a subsequence of the array [l, 2, 3, 4] , and so do the numbers [2, 4] .
Note that a single number 1n an array and the array itself are both valid subsequences of the array.
You can assume that there will only be one increas1ng subsequence with the greatest sum.


Sample Input:
array = [10, 70, 20, 30, 50, 11, 30]


Sample Output:
[110, [10, 20, 30, 50]]     // The subsequence [10, 20, 30, 50] is strictly 1ncreasing and yields the greatest sum: 110.


Hints
Hint 1

Try building an array of the same length as the input array. At each index in this new array, store the maximum sum that can be generated from an increasing subsequence ending with the number found at that index in the input array.

Hint 2

Can you efficiently keep track of potential sequences in another array? Instead of storing entire sequences, try storing the indices of previous numbers. For example, at index 3 in this other array, store the index of the before-last number in the max-sum increasing subsequence ending with the number at index 3.





*/

package main

import (
	"fmt"
)

type Subsequence struct {
	sum         int
	numbers     []int
	predecessor *Subsequence
}

func greatestSumIncreasingSubsequence(array []int) *Subsequence {
	// Create a DP table to store the maximum sums and their corresponding predecessors for all possible suffixes of the array
	dp := make([]*Subsequence, len(array))

	// Initialize the DP table with the default values of 0 and nil
	for i := range dp {
		dp[i] = &Subsequence{sum: 0, numbers: nil, predecessor: nil}
	}

	// Fill the DP table using bottom-up approach
	for i := 1; i < len(array); i++ {
		for j := 0; j < i; j++ {
			if array[i] > array[j] && dp[i].sum < dp[j].sum+array[i] {
				dp[i] = &Subsequence{sum: dp[j].sum + array[i], numbers: append([]int{array[i]}, dp[j].numbers...), predecessor: dp[j]}
			}
		}
	}

	// Find the index of the subsequence with the greatest sum
	maxIndex := 0
	for i := 1; i < len(dp); i++ {
		if dp[i].sum > dp[maxIndex].sum {
			maxIndex = i
		}
	}

	// Construct the greatest sum increasing subsequence using backtracking
	greatestSumSubsequence := dp[maxIndex]
	subsequenceNumbers := greatestSumSubsequence.numbers
	for greatestSumSubsequence.predecessor != nil {
		subsequenceNumbers = append([]int{greatestSumSubsequence.predecessor.numbers[len(greatestSumSubsequence.predecessor.numbers)-1]}, subsequenceNumbers...)
		greatestSumSubsequence = greatestSumSubsequence.predecessor
	}

	// Reverse the constructed subsequence
	for i, j := 0, len(subsequenceNumbers)-1; i < j; i, j = i+1, j-1 {
		subsequenceNumbers[i], subsequenceNumbers[j] = subsequenceNumbers[j], subsequenceNumbers[i]
	}

	// Return the greatest sum increasing subsequence
	return &Subsequence{sum: greatestSumSubsequence.sum, numbers: subsequenceNumbers, predecessor: nil}
}

func main() {
	// Set the input values.
	array := []int{10, 70, 20, 30, 50, 11, 30}

	// Find the greatest sum of a strictly increasing subsequence in the array and the subsequence itself.
	greatestSumSubsequence := greatestSumIncreasingSubsequence(array)

	// Print the greatest sum of a strictly increasing subsequence and the subsequence itself.
	fmt.Println(greatestSumSubsequence.sum, greatestSumSubsequence.numbers)
}
