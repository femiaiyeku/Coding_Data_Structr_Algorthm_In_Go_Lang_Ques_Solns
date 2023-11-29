/*

Write a function that takes in a non-empty string and that returns a boolean representing whether the string is a
palindrome.
A palindrome is defined as a string that's written the same forward and backward. Note that single-character
strings are palindromes.
Sample Input

string = "abcdcba"

Sample Output

true // it's written the same forward and backward






This code first reverses the input string using a nested loop and stores the reversed string in a separate variable. Then, it compares the original string to the reversed string using the strings.EqualFold function, which ignores case sensitivity. If the strings are equal, the original string is a palindrome, and the function returns true; otherwise, it returns false.






*/

package main

import (
	"fmt"
	"strings"
)

func isPalindrome(str string) bool {
	reversedStr := reverseString(str)
	return strings.EqualFold(str, reversedStr)
}

func reverseString(str string) string {
	reversedStr := ""

	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}

	return reversedStr
}

func main() {
	str := "abcdcba"
	isPal := isPalindrome(str)
	fmt.Println("Is palindrome:", isPal)
}
