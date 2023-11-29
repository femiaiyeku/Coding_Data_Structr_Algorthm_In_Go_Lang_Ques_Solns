/*



You're given an array of positive integers representing the prices of a single stock on various days (each index in the
array represents a different day). You're also given an integer k, which represents the number of transactions you're
allowed to make. One transaction consists of buying the stock on a given day and selling it on another, later day.
Write a function that returns the maximum profit that you can make by buying and selling the stock, given k
transactions.
Note that you can only hold one share of the stock at a time; in other words, you can't buy more than one share of the
stock on any given day, and you can't buy a share of the stock if you're still holding another share. Also, you don't need
to use all k transactions that you're allowed.


Sample Input:

prices = [5, 11, 3, 50, 60, 90]

k = 2

Sample Output:

93 // Buy: 5, Sell: 11; Buy: 3, Sell: 90



Hints
Hint 1

Try building a two-dimensional array of the maximum profits you can make on each day with zero, one, two, etc., k transactions. Let columns represent days and rows represent the number of transactions.

Hint 2

Build up the array mentioned in Hint #1 one row at a time. In other words, find the maximum profits that you can make on each day with zero transactions first, then with one transaction, etc., until you reach k transactions. Find a formula that relates the maximum profit at any given point to previous profits. Once you find that formula, identify certain values that you repeatedly need and that you can temporarily store to optimize your algorithm.

Hint 3

Do you really need to store the entire two-dimensional array mentioned in Hint #1? Identify what stored values you actually use throughout the process of building the array and come up with a way of storing only what you need and nothing more.


*/

package main

import (
	"fmt"
)

type Transaction struct {
	buyPrice  int
	sellPrice int
}

func maxProfitWithKTransactions(prices []int, k int) int {
	// Create two DP tables to store the maximum profits for holding and not holding a stock on each day
	holdingProfit := make([]int, len(prices))
	notHoldingProfit := make([]int, len(prices))

	// Initialize the DP tables with the default values of 0
	for i := 0; i < len(prices); i++ {
		holdingProfit[i] = 0
		notHoldingProfit[i] = 0
	}

	// Fill the DP tables using bottom-up approach
	for t := 1; t <= k; t++ {
		for i := 0; i < len(prices); i++ {
			if i > 0 {
				holdingProfit[i] = max(holdingProfit[i-1], notHoldingProfit[i-1]-prices[i])
			}
			notHoldingProfit[i] = max(notHoldingProfit[i-1], holdingProfit[i-1]+prices[i])
		}
	}

	// Return the maximum profit from the notHoldingProfit table on the last day
	return notHoldingProfit[len(prices)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Set the input values.
	prices := []int{5, 11, 3, 50, 60, 90}
	k := 2

	// Find the maximum profit that can be made by buying and selling the stock, given k transactions.
	maxProfit := maxProfitWithKTransactions(prices, k)

	// Print the maximum profit.
	fmt.Println(maxProfit)
}
