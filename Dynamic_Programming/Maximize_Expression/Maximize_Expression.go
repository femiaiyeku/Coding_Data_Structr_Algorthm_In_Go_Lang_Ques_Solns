/*


Write a function that takes 1n an array of integers and returns the largest possible value for the expression
array[a] - array[b] + array[c] - array[d] ,where a, b, c ,and d areindices ofthe array and a< b < c < d

if  the input array has fewer than 4 elements, your function should return 0

Sample Input

array = [3, 6, 1, -3, 2, 7]


Sample Output

4 // by choosing a = 1, b = 3, c = 4, and d = 5


Hints

Hint 1

The brute-force approach to solving this problem is to simply iterate through every valid choice of a, b, c, and d and to evaluate the expression at each iteration. While doing this, you can keep track of the maximum value that you find and return it after considering all possibilities. This solution runs in O(n^4) time; can you think of a way to solve this faster?

Hint 2

You can solve this problem using dynamic programming with a time complexity of O(n); however, you'll need to use external space.

Hint 3

If you know what the maximum possible value of a is at each index in the array, you can find the maximum possible value of a - b at each individual index in the array in O(1) time (or in O(n) time for all indices). The same thing holds for finding the maximum possible value of a - b + c if you know the maximum possible value of a - b at each index. How does this fact help you solve the entire problem in O(n) time?

Hint 4

Start by finding the maximum possible value of a at each index in the array, meaning the maximum value of a that you can obtain at each index i if a is chosen from an index between 0 and i, inclusive. Store all of these values in an array, and use them to help you determine the maximum possible value of a - b at each index. Do the same for a - b + c (using the results from a - b) and a - b + c - d (using the results from a - b + c). Once you make it to a - b + c - d, you'll be able to determine the maximum value of the expression.

*/

package main

import (
	"fmt"
)

func largestPossibleValue(array []int) int {
	// Handle the case when the input array has fewer than 4 elements
	if len(array) < 4 {
		return 0
	}

	// Initialize the minimum and maximum values
	minA := array[0]
	maxD := array[len(array)-1]

	// Create a DP table to store the maximum values of the expression for each index
	dp := make([]int, len(array))

	// Initialize the DP table with the default values of 0
	for i := 0; i < len(array); i++ {
		dp[i] = 0
	}

	// Fill the DP table using bottom-up approach
	for i := 1; i < len(array); i++ {
		// Update the minimum value of 'a'
		if array[i] < minA {
			minA = array[i]
		}

		// Calculate the maximum value of the expression for the current index 'i'
		dp[i] = max(dp[i-1], maxD-minA)

		// Update the maximum value of 'd'
		if array[len(array)-1-i] > maxD {
			maxD = array[len(array)-1-i]
		}
	}

	// Return the maximum value from the DP table
	return dp[len(array)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Set the input values.
	array := []int{3, 6, 1, -3, 2, 7}

	// Find the largest possible value for the expression `array[a] - array[b] + array[c] - array[d]`.
	largestPossibleValue := largestPossibleValue(array)

	// Print the largest possible value.
	fmt.Println(largestPossibleValue)
}
