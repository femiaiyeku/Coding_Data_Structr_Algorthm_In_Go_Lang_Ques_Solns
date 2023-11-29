/*



Write a function that takes in two strings and returns the minimum number of edit operations that need to be performed on the first string to obtain the second string.
There are three edit operations: insertion of a character, deletion of a character, and substitution of a character for another.

Sample Input: str1 = "abc", str2 = "yabd"

Sample Output: 2 // insert "y"; substitute "c" for "d"







Hints
Hint 1
Try building a two-dimensional array of the minimum numbers of edits for pairs of substrings of the input strings. Let the rows of the array represent substrings of the second input string str2. Let the first row represent the empty string. Let each row i thereafter represent the substrings of str2 from 0 to i, with i excluded. Let the columns similarly represent the first input string str1.
Hint 2
Build up the array mentioned in Hint #1 one row at a time. In other words, find the minimum numbers of edits between all the substrings of str1 represented by the columns and the empty string represented by the first row, then between all the substrings of str1 represented by the columns and the first letter of str2 represented by the second row, etc., until you compare both full strings. Find a formula that relates the minimum number of edits at any given point to previous numbers.
Hint 3
At any position (i, j) in the two-dimensional array, if str2[i] is equal to str1[j], then the edit distance at position (i, j) is equal to the one at position (i - 1, j - 1), since adding str2[i] and str1[j] to the substrings represented at position (i - 1, j - 1) does not require any additional edit operation. If str2[i] is not equal to str1[j] however, then the edit distance at position (i, j) is equal to 1 + the minimum of the edit distances at positions (i - 1, j), (i, j - 1), and (i - 1, j - 1). Why is that the case?
Hint 4
Do you really need to store the entire two-dimensional array mentioned in Hint #1? Identify what stored values you actually use throughout the process of building the array and come up with a way of storing only what you need and nothing more.


*/

package main

import (
	"fmt"
)

func editDistance(str1, str2 string) int {
	// Create a DP table to store the edit distances for all possible substrings
	dp := make([][]int, len(str1)+1)
	for i := range dp {
		dp[i] = make([]int, len(str2)+1)
	}

	// Fill the DP table using top-down approach
	for i := 0; i <= len(str1); i++ {
		for j := 0; j <= len(str2); j++ {
			if i == 0 || j == 0 {
				// If either string is empty, the edit distance is equal to the length of the other string
				dp[i][j] = max(i, j)
			} else if str1[i-1] == str2[j-1] {
				// If the characters are equal, no edit operation is needed
				dp[i][j] = dp[i-1][j-1]
			} else {
				// Perform the minimum of insertion, deletion, or substitution
				dp[i][j] = min(1+dp[i-1][j], 1+dp[i][j-1], 1+dp[i-1][j-1])
			}
		}
	}

	// Return the edit distance for the entire strings
	return dp[len(str1)][len(str2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b, c int) int {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

func main() {
	// Set the input values.
	str1 := "abc"
	str2 := "yabd"

	// Calculate the minimum edit distance between the strings.
	minEditDistance := editDistance(str1, str2)

	// Print the minimum edit distance.
	fmt.Println(minEditDistance)
}
