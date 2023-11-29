/*


ou're given an array of arrays where each subarray holds two integer values and represents an item; the first integer is
the item's value, and the second integer is the item's weight. You're also given an integer representing the maximum
capacity of a knapsack that you have.
Your goal is to fit items in your knapsack without having the sum of their weights exceed the knapsack's capacity, all the
while maximizing their combined value. Note that you only have one of each item at your disposal.
Write a function that returns the maximized combined value of the items that you should pick as well as an array of the
indices of each item picked.
If there are multiple combinations of items that maximize the total value in the knapsack, your function can return any
of them.


Sample Input:
items = [[1, 2], [4, 3], [5, 6], [6, 7]]
capacity = 10

Sample Output:
[10, [1, 3]] // items [4, 3] and [6, 7]

// The values in the array are ordered according to the indices in the input array







*/

package main

import (
	"fmt"
)

type Item struct {
	value  int
	weight int
}

func knapsack(items []Item, capacity int) (int, []int) {
	// Create a 2D DP table to store the maximum value for each possible weight
	dp := make([][]int, len(items)+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// Fill the DP table using bottom-up approach
	for i := 1; i <= len(items); i++ {
		for w := 1; w <= capacity; w++ {
			if items[i-1].weight > w {
				dp[i][w] = dp[i-1][w]
			} else {
				dp[i][w] = max(dp[i-1][w], items[i-1].value+dp[i-1][w-items[i-1].weight])
			}
		}
	}

	// Backtrack to find the items included in the optimal solution
	selectedItems := []int{}
	currentWeight := capacity

	for i := len(items); i > 0; i-- {
		if dp[i][currentWeight] != dp[i-1][currentWeight] {
			selectedItems = append(selectedItems, i-1)
			currentWeight -= items[i-1].weight
		}
	}

	// Return the maximum value and the indices of the selected items
	return dp[len(items)][capacity], selectedItems
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Set the input values.
	items := []Item{{1, 2}, {4, 3}, {5, 6}, {6, 7}}
	capacity := 10

	// Find the maximized combined value and the indices of the selected items.
	maxValue, selectedItems := knapsack(items, capacity)

	// Print the maximized combined value and the indices of the selected items.
	fmt.Println(maxValue, selectedItems)
}
