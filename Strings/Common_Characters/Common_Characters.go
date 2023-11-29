/*


Write a function that takes in a non-empty list of non-empty strings and returns a list of characters that are common to all strings in the list, ignoring multiplicity.
Note that the strings are not guaranteed to only contain alphanumeric charaaers. The list you return can be in any order.

Sample Input

strings = ["abc", "bcd", "cbaccd"]

Sample Output

    ["b", "c"]


This code iterates through the list of strings and maintains a map of common characters. If a character is present in all strings and hasn't been added to the commonChars map yet, it is added to the map. Finally, it creates a list of the common characters from the commonChars map and returns it.







*/

package main

import (
	"fmt"
)

func findCommonCharacters(strList []string) []string {
	commonChars := make(map[byte]bool)
	result := []string{}

	for _, str := range strList {
		seenChars := make(map[byte]bool)

		for _, char := range str {
			if !seenChars[char] {
				commonChars[char] = true
				seenChars[char] = true
			}
		}
	}

	for char := range commonChars {
		result = append(result, string(char))
	}

	return result
}

func main() {
	strList := []string{"abc", "bcd", "cbaccd"}
	commonChars := findCommonCharacters(strList)
	fmt.Println("Common characters:", commonChars)
}
