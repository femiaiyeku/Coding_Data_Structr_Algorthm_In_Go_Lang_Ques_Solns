/*


Given an array of positive integers representing coin denominations and a single non-negative integer n representing
a target amount of money, write a function that returns the smallest number of coins needed to make change for (to
sum up to) that target amount using the given coin denominations.
Note that you have access to an unlimited amount of coins. In other words, if the denominations are [1, 5, 10],
you have access to an unlimited amount of 1 s, 5 s, and 10 s.
If it's impossible to make change for the target amount, return

Sample Input: n = 7, denoms = [1, 5, 10]

Sample Output: 3 // 2x1 + 1x5


Hints
Hint 1
Try building an array of the minimum number of coins needed to make change for all amounts between 0 and n inclusive. Note that no coins are needed to make change for 0: in order to make change for 0, you do not need to use any coins.

Hint 2
Build up the array mentioned in Hint #1 one coin denomination at a time. In other words, find the minimum number of coins needed to make change for all amounts between 0 and n with only one denomination, then with two, etc., until you use all denominations.




*/

package main

import (
	"fmt"
	"math"
)

func minCoins(targetAmount int, denominations []int) int {
	// Initialize a DP table to store the minimum number of coins for each amount up to the target amount
	dp := make([]int, targetAmount+1)

	// Initialize the DP table with the default values of infinity
	for i := 0; i <= targetAmount; i++ {
		dp[i] = math.MaxInt32
	}

	// Base case: 0 coins needed to make change for 0 amount
	dp[0] = 0

	// Fill the DP table using bottom-up approach
	for i := 1; i <= targetAmount; i++ {
		// For each denomination, check if it's possible to make change for the remaining amount
		for _, denomination := range denominations {
			if denomination <= i {
				// Calculate the number of coins needed using the current denomination and the minimum number of coins for the remaining amount
				coinsNeeded := dp[i-denomination] + 1

				// Update the DP table with the minimum number of coins found so far
				if coinsNeeded < dp[i] {
					dp[i] = coinsNeeded
				}
			}
		}
	}

	// Check if it's impossible to make change for the target amount
	if dp[targetAmount] == math.MaxInt32 {
		return -1
	}

	// Return the minimum number of coins needed to make change for the target amount
	return dp[targetAmount]
}

func main() {
	// Set the input values.
	targetAmount := 7
	denominations := []int{1, 5, 10}

	// Find the smallest number of coins needed to make change for the given target amount using the given coin denominations.
	minCoinsNeeded := minCoins(targetAmount, denominations)

	// Print the minimum number of coins needed.
	if minCoinsNeeded == -1 {
		fmt.Println("Impossible to make change")
	} else {
		fmt.Println(minCoinsNeeded)
	}
}
