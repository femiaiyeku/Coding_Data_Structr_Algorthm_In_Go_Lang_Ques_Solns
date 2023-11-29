/*



Write a function that takes in two strings and returns their longest common subsequence.
A subsequence of a string is a set of characters that aren't necessarily adjacent in the string but that are in the same
order as they appear in the string. For instance, the characters ["a", "c", "d"] form a subsequence of the string
"abcd", and so do the characters ["b", "d"]. Note that a single character in a string and the string itself are both
valid subsequences of the string.
You can assume that there will only be one longest common subsequence.


Sample Input:


str1 = "ZXVVYZW"
str2 = "XKYKZPW"


Sample Output:


["X", "Y", "Z", "W"]




Hints
Hint 1

Try building a two-dimensional array of the longest common subsequences of substring pairs of the input strings. Let the rows of the array represent substrings of the second input string str2. Let the first row represent the empty string. Let each row i thereafter represent the substrings of str2 from 0 to i, with i excluded. Let the columns similarly represent the first input string str1.

Hint 2

Build up the array mentioned in Hint #1 one row at a time. In other words, find the longest common subsequences for all the substrings of str1 represented by the columns and the empty string represented by the first row, then for all the substrings of str1 represented by the columns and the first letter of str2 represented by the second row, etc., until you compare both full strings. Find a formula that relates the longest common subsequence at any given point to previous subsequences.

Hint 3

Do you really need to build and store subsequences at each point in the two-dimensional array mentioned in Hint #1? Try storing booleans to determine whether or not a letter at a given point in the two-dimensional array is part of the longest common subsequence as well as pointers to determine what should come before this letter in the final subsequence. Use these pointers to backtrack your way through the array and to build up the longest common subsequence at the end of your algorithm.



*/

package main

import (
	"fmt"
)

func longestCommonSubsequence(str1, str2 string) string {
	// Create a DP table to store the lengths of the longest common subsequences for all possible substrings
	dp := make([][]int, len(str1)+1)
	for i := range dp {
		dp[i] = make([]int, len(str2)+1)
	}

	// Fill the DP table using bottom-up approach
	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// Construct the longest common subsequence using backtracking
	lcs := ""
	i := len(str1)
	j := len(str2)

	for i > 0 && j > 0 {
		if str1[i-1] == str2[j-1] {
			lcs = string(str1[i-1]) + lcs
			i--
			j--
		} else if dp[i][j] == dp[i-1][j] {
			i--
		} else {
			j--
		}
	}

	// Reverse the constructed subsequence
	return reverse(lcs)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	// Set the input values.
	str1 := "ZXVVYZW"
	str2 := "XKYKZPW"

	// Find the longest common subsequence of the strings.
	lcs := longestCommonSubsequence(str1, str2)

	// Print the longest common subsequence.
	fmt.Println(lcs)
}
