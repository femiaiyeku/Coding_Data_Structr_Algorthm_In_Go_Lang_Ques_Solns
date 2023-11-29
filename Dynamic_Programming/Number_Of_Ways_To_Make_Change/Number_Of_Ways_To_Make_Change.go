/*

Given an array of distinct positive integers representing coin denominations and a single non-negative integer n representing a target amount of money,
write a function that returns the number of ways to make change for that target amount using the given coin denominations.
Note that an unlimited amount of coins is at your disposal.


Sample Input: n = 6, denoms = [1, 5]

Sample Output: 2 // 1x1 + 1x5 and 6x1


Hints
Hint 1
Try building an array of the number of ways to make change for all amounts between 0 and n inclusive. Note that there is only one way to make change for 0: that is to not use any coins.
Hint 2
Build up the array mentioned in Hint #1 one coin denomination at a time. In other words, find the number of ways to make change for all amounts between 0 and n with only one denomination, then with two, etc., until you use all denominations.
Optimal Space & Time Complexity
O(nd) time | O(n) space - where n is the target amount and d is the number of coin denominations






*/

package main

import (
	"fmt"
)

func countWaysToMakeChange(targetAmount int, denominations []int) int {
	// Initialize a DP table to store the number of ways to make change for each amount up to the target amount
	dp := make([]int, targetAmount+1)

	// Initialize the DP table with the default values of 0
	for i := 0; i <= targetAmount; i++ {
		dp[i] = 0
	}

	// Base case: 1 way to make change for 0 amount
	dp[0] = 1

	// Fill the DP table using bottom-up approach
	for i := 1; i <= targetAmount; i++ {
		for _, denomination := range denominations {
			if denomination <= i {
				// Add the number of ways to make change for the remaining amount using the current denomination
				dp[i] += dp[i-denomination]
			}
		}
	}

	// Return the number of ways to make change for the target amount
	return dp[targetAmount]
}

func main() {
	// Set the input values.
	targetAmount := 6
	denominations := []int{1, 5}

	// Find the number of ways to make change for the given target amount using the given coin denominations.
	numWays := countWaysToMakeChange(targetAmount, denominations)

	// Print the number of ways to make change for the target amount.
	fmt.Println(numWays)
}
