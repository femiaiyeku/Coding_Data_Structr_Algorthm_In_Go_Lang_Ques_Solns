/*

Write a function that takes in a big string and an array of small strings, all of which are smaller in length than the
big string. The function should return an array of booleans, where each boolean represents whether the small
string at that index in the array of small strings is contained in the big string.
Note that you can't use language-built-in string-matching methods.


Sample Input
bigString = "this is a big string"
smallStrings = ["this", "yo", "is", "a", "bigger", "string", "kappa"]

Sample Output
[true, false, true, true, false, true, false]

// Note: the booleans in the array don't have to be in any particular order.







Hints
Hint 1

A simple way to solve this problem is to iterate through all of the small strings, checking if each of them is contained in the big string by iterating through the big string's characters and comparing them to the given small string's characters with a couple of loops. Is this approach efficient from a time-complexity point of view?

Hint 2

Try building a suffix-trie-like data structure containing all of the big string's suffixes. Then, iterate through all of the small strings and check if each of them is contained in the data structure you've created. What are the time-complexity ramifications of this approach?

Hint 3

Try building a trie containing all of the small strings. Then, iterate through the big string's characters and check if any part of the big string is a string contained in the trie you've created. Is this approach better than the one described in Hint #2 from a time-complexity point of view?




This code iterates through the small strings and searches for each small string in the big string using a simple pattern matching algorithm. It checks for the first character of the small string in the big string and then compares the remaining characters one by one. If a match is found, it sets the corresponding boolean in the foundStrings array to true. Finally, it returns the array of booleans representing which small strings were found in the big string.



*/

package main

import (
	"fmt"
)

func findSmallStrings(bigString string, smallStrings []string) []bool {
	foundStrings := make([]bool, len(smallStrings))

	for i, smallString := range smallStrings {
		found := false

		for j := 0; j < len(bigString); j++ {
			if bigString[j] == smallString[0] {
				if len(smallString) == 1 {
					found = true
					break
				}

				matchFound := true
				for k := 1; k < len(smallString); k++ {
					if bigString[j+k] != smallString[k] {
						matchFound = false
						break
					}
				}

				if matchFound {
					found = true
					break
				}
			}
		}

		foundStrings[i] = found
	}

	return foundStrings
}

func main() {
	bigString := "this is a big string"
	smallStrings := []string{"this", "yo", "is", "a", "bigger", "string", "kappa"}

	foundStrings := findSmallStrings(bigString, smallStrings)
	fmt.Println("Found strings:", foundStrings)
}
