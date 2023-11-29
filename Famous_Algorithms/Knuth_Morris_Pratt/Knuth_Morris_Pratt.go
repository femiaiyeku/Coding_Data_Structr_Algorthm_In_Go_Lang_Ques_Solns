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



Hints
Hint 1

The Knuth—Morris—Pratt algorithm works by identifying patterns in the potential substring and exploiting them to avoid doing needless character comparisons when searching for the substring in the main string. For instance, take the string "ababac" and the substring "abac"; comparing these strings will fail at the fourth character, where "b" is not equal to "c". Instead of having to restart our comparisons at the second character of the main string, however, we notice that the substring "ab", which is at the beginning of our potential substring, just appeared near our point of failure in the main string. How can we use this to our advantage?

Hint 2

Start by traversing the potential substring and building out a pattern table. This 1-dimensional array should store, for every position in the substring, the last index at which a matching pattern has been seen; more specifically, this index should be the ending index of a prefix in the substring that is also a suffix at the given position. For example, the string "abcababcd" should yield the following pattern table: [-1, -1, -1, 0, 1, 0, 1, 2, -1].

Hint 3

After the pattern table mentioned in Hint #2 has been built, traverse the main string and the potential substring with two separate pointers. When characters match, move the pointers forward. When characters don't match, check if the pointer in the substring is at the very beginning of the substring; if it is, then there is no match and you can move the pointer of the main string forward until there is a match; if it isn't, then move it to the position that comes right after the last seen pattern stored at the previous index in the pattern table.








*/

package main

import (
	"fmt"
)

func kmpSearch(text, pattern string) bool {
	// Calculate the prefix function for the pattern
	prefixTable := computePrefixTable(pattern)

	// Initialize variables for searching
	i := 0
	j := 0

	// Iterate through the text
	for i < len(text) {
		// Compare the current characters of the text and pattern
		if text[i] == pattern[j] {
			i++
			j++

			// Check if the pattern has been matched completely
			if j == len(pattern) {
				return true
			}
		} else {
			// Shift j using the prefix table
			if j > 0 {
				j = prefixTable[j-1]
			}
		}
	}

	// The pattern was not found in the text
	return false
}

func computePrefixTable(pattern string) []int {
	prefixTable := make([]int, len(pattern))

	// Initialize the first element of the prefix table
	prefixTable[0] = 0

	// Calculate the prefix table for the remaining characters
	for i := 1; i < len(pattern); i++ {
		j := prefixTable[i-1]

		for j > 0 && pattern[i] != pattern[j] {
			j = prefixTable[j-1]
		}

		if pattern[i] == pattern[j] {
			j++
		}

		prefixTable[i] = j
	}

	return prefixTable
}

func main() {
	// Set the input strings
	text := "aefoaefcdaefcdaed"
	pattern := "aefcdaed"

	// Check if the first string contains the second one
	contains := kmpSearch(text, pattern)

	// Print the result
	fmt.Println(contains)
}
