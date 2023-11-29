/*


Write a function that, given a string, returns its longest palindromic substring.
 A palindrome 1s defined as a string thats written the same forward and backward.
 Note that single-character strings are palindromes.
 You can assume that there will only be one longest palindromic substring

Sample Input
string = "abaxyzzyxf"

Sample Output

"xyzzyx"




Hints
Hint 1

Try generating all possible substrings of the input string and checking for their palindromicity. What is the runtime of the isPalindrome check? What is the total runtime of this approach?

Hint 2

Recognize that a palindrome is a string that is symmetrical with respect to its center, which can either be a character (in the case of odd-length palindromes) or an empty string (in the case of even-length palindromes). Thus, you can check the palindromicity of a string by simply expanding from its center and making sure that characters on both sides are indeed mirrored.

Hint 3

Traverse the input string, and at each index, apply the logic mentioned in Hint #2. What does this accomplish? Is the runtime of this approach better?


This code uses dynamic programming to efficiently find the longest palindromic substring. It creates a 2D array dp to store whether each substring is a palindrome. The dp array is filled in a bottom-up manner, starting with single-character substrings and moving on to longer substrings. Finally, it iterates through the dp array to find the longest palindromic substring and returns it.



*/

package main

import (
	"fmt"
)

func longestPalindrome(str string) string {
	n := len(str)
	dp := make([][]bool, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	longestPalLength := 1
	longestPalStartIndex := 0

	for length := 2; length <= n; length++ {
		for startIndex := 0; startIndex <= n-length; startIndex++ {
			endIndex := startIndex + length - 1

			if str[startIndex] == str[endIndex] {
				if length == 2 {
					dp[startIndex][endIndex] = true
				} else {
					dp[startIndex][endIndex] = dp[startIndex+1][endIndex-1]
				}

				if dp[startIndex][endIndex] && length > longestPalLength {
					longestPalLength = length
					longestPalStartIndex = startIndex
				}
			}
		}
	}

	return str[longestPalStartIndex:(longestPalStartIndex + longestPalLength)]
}

func main() {
	str := "abaxyzzyxf"
	longestPalindrome := longestPalindrome(str)
	fmt.Println("Longest palindromic substring:", longestPalindrome)
}
