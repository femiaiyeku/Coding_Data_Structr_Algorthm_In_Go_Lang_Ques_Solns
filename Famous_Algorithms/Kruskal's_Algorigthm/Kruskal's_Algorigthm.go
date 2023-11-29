/*

Write a function that takes in two strings and checks if the first string contains the second one using the Knuth-Morris
-Pratt algorithm. The function should return a boolean.
If you're unfamiliar with the Knuth-Morris-Pratt Algorithm, we recommend watching the Conceptual Overview section
of this question's video explanation before starting to code


Sample Input:

string = "aefoaefcdaefcdaed"
substring = "aefcdaed"

Sample Output:

true





*/

package main

import (
	"fmt"
)

// kmpSearch checks if the first string (text) contains the second string (pattern) using the Knuth-Morris-Pratt algorithm.
func kmpSearch(text, pattern string) bool {
	pi := computePrefixFunction(pattern)
	i, j := 0, 0
	for i < len(text) {
		if text[i] == pattern[j] {
			i++
			j++
			if j == len(pattern) {
				return true
			}
		} else {
			if j == 0 {
				i++
			} else {
				j = pi[j-1]
			}
		}
	}
	return false
}

// computePrefixFunction computes the prefix function for the Knuth-Morris-Pratt algorithm.
func computePrefixFunction(pattern string) []int {
	pi := make([]int, len(pattern))
	i := 0
	j := 1
	for j < len(pattern) {
		if pattern[i] == pattern[j] {
			pi[j] = i + 1
			i++
			j++
		} else {
			if i == 0 {
				pi[j] = 0
				j++
			} else {
				i = pi[i-1]
			}
		}
	}
	return pi
}

func main() {
	text := "aefoaefcdaefcdaed"
	pattern := "aefcdaed"

	if kmpSearch(text, pattern) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
