/*
Given a non-empty string, write a function that returns the minimum number of cuts needed to perform on the string
such that each remaining substring is a palindrome.
A palindrome is defined as a string that's written the same forward as backward. Note that single-character strings are
palindromes

Sample Input:

string = "noonabbad"

Sample Output:

2 // noon | abba | d


Hints
Hint 1

Try building a two-dimensional array of the palindromicities of all substrings of the input string. Let the value stored at row i and at column j represent the palindromicity of the substring starting at index i and ending at index j.

Hint 2

Checking for palindromicity is typically an O(n) time operation. Can you eliminate this step and build the same two-dimensional array mentioned in Hint #1 a different way? Realize that the substring whose starting and ending indices are (i, j) is only a palindrome if string[i] is equal to string[j] and if the substring denoted by (i + 1, j - 1) is also a palindrome.

Hint 3

Build a one-dimensional array of the same length as the input string. At each index i in this array compute and store the minimum number of cuts needed for the substring whose starting and ending indices are (0, i). Use previously calculated values as well as the two-dimensional array mentioned in Hint #1 to find each value in this array.



*/

package main

import (
	"fmt"
)

// Function to check if a given string is a palindrome
func isPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

// Function to find the minimum number of cuts needed to partition a string into palindromes
func minCuts(str string) int {
	// Create a DP table to store the minimum number of cuts for each substring
	dp := make([]int, len(str)+1)
	for i := 0; i <= len(str); i++ {
		dp[i] = len(str) - i
	}

	// Fill the DP table using bottom-up approach
	for i := 0; i < len(str); i++ {
		if isPalindrome(str[:i+1]) {
			dp[i+1] = 0
		} else {
			for j := 1; j <= i+1; j++ {
				if isPalindrome(str[j : i+1]) {
					dp[i+1] = min(dp[i+1], dp[j]+1)
				}
			}
		}
	}

	// Return the minimum number of cuts for the entire string
	return dp[len(str)]
}

// Function to find the minimum value among two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Set the input string.
	str := "noonabbad"

	// Find the minimum number of cuts needed to partition the string into palindromes.
	minCutsNeeded := minCuts(str)

	// Print the minimum number of cuts needed.
	fmt.Println(minCutsNeeded)
}
