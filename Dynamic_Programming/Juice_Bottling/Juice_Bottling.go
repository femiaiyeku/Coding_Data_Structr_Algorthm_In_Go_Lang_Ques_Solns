/*

You're given an array of integers prices of length n with the retail prices of various quantities of juice. Each index in this array corresponds to the price of that amount of juice.
For example, prices [2] would be the retail price of 2 units of juice.
You have n - 1 total units of juice. For example, if the length of prices is 5, then you would have 4 total units of juice.
Write a function to determine the optimal way to bottle the Juice such that it maximizes revenue.
This function should return a list of all of the Juice quantities required in ascending order.
Note that the first value in the prices array will always be 0, because there is no value in no juice. All other values will be positive integers.
Additionally, a larger quantity of juice will not always be more expensive than a smaller quantity. For simplicity, all of the test cases only have one possible solution.





Sample Input:

prices = [0, 1, 3, 2]

Sample Output:

[1, 2]

Explanation:

The optimal solution is to bottle 1 unit of juice, then 2 units of juice.









*/

package main

import (
	"fmt"
	"sort"
)

func optimizeJuiceBottling(prices []int) []int {
	// Create a memoization table to store the optimal revenue for each quantity of juice
	memo := make(map[int]int)

	// Recursively calculate the optimal revenue for each quantity of juice
	optimalQuantities := optimizeJuiceBottlingHelper(prices, memo, 1)

	// Sort the optimal quantities in ascending order
	sort.Ints(optimalQuantities)

	// Return the list of optimal quantities
	return optimalQuantities
}

func optimizeJuiceBottlingHelper(prices []int, memo map[int]int, remainingJuice int) []int {
	// If there is no remaining juice, there are no optimal quantities
	if remainingJuice == 0 {
		return []int{}
	}

	// If the optimal revenue for the remaining juice has already been calculated, return it
	if optimalRevenue, ok := memo[remainingJuice]; ok {
		return memo[remainingJuice]
	}

	// Initialize the optimal quantities and revenue
	optimalQuantities := []int{}
	optimalRevenue := 0

	// Try bottling each possible quantity of juice
	for quantity := 1; quantity <= remainingJuice; quantity++ {
		// Calculate the revenue from bottling the current quantity
		currentRevenue := prices[quantity]

		// Recursively calculate the optimal revenue for the remaining juice
		remainingQuantities := optimizeJuiceBottlingHelper(prices, memo, remainingJuice-quantity)

		// Calculate the total revenue from bottling the current quantity and the remaining juice
		totalRevenue := currentRevenue + memo[remainingJuice-quantity]

		// If the total revenue is greater than the optimal revenue so far, update the optimal quantities and revenue
		if totalRevenue > optimalRevenue {
			optimalQuantities = append(optimalQuantities, quantity)
			optimalRevenue = totalRevenue
		}
	}

	// Store the optimal revenue for the remaining juice in the memoization table
	memo[remainingJuice] = optimalRevenue

	// Return the optimal quantities
	return optimalQuantities
}

func main() {
	// Set the input values.
	prices := []int{0, 1, 3, 2}

	// Determine the optimal way to bottle the juice.
	optimalQuantities := optimizeJuiceBottling(prices)

	// Print the optimal quantities.
	fmt.Println(optimalQuantities)
}
