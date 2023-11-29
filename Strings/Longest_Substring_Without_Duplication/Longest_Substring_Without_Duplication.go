/*

Write a function that takes in a string and returns its longest substring without duplicate characters.
You can assume that there will only be one longest substring without duplication.

Sample Input

string = "clementisacap"

Sample Output

"mentisac"


Hints
Hint 1

Try traversing the input string and storing the last position at which you see each character in a hash table. How can this help you solve the given problem?

Hint 2

As you traverse the input string, keep track of a starting index variable. This variable, as its name suggests, should represent the most recent index from which you could start a substring with no duplicate characters, ending at your current index. Use the hash table mentioned in Hint #1 to update this variable correctly, and update the longest substring as you go.





This code uses a sliding window technique to find the longest substring without duplicate characters. It maintains a currentUniqueSubstring variable to store the current substring being considered and a seenCharacters map to keep track of the characters that have already been seen. As it iterates through the string, it checks if the current character has already been seen. If it has, it resets the currentUniqueSubstring and the seenCharacters map, except for the characters before the first occurrence of the repeated character. This ensures that the current unique substring doesn't contain any duplicate characters. Finally, it compares the length of the currentUniqueSubstring to the length of the longestUniqueSubstring and updates the longestUniqueSubstring if necessary.

*/

package main

import (
	"fmt"
)

func longestUniqueSubstring(str string) string {
	longestUniqueSubstring := ""
	currentUniqueSubstring := ""
	seenCharacters := make(map[byte]bool)

	for i := range str {
		currentChar := str[i]

		if seenCharacters[currentChar] {
			// If the current character has already been seen, we need to reset the current unique substring
			if len(currentUniqueSubstring) > len(longestUniqueSubstring) {
				longestUniqueSubstring = currentUniqueSubstring
			}

			currentUniqueSubstring = ""

			// We don't need to check the characters before the first occurrence of the repeated character
			// because they will be included in the next iteration of the loop
			for j := i - 1; j >= 0 && str[j] != currentChar; j-- {
				seenCharacters[str[j]] = false
			}
		}

		currentUniqueSubstring += string(currentChar)
		seenCharacters[currentChar] = true
	}

	if len(currentUniqueSubstring) > len(longestUniqueSubstring) {
		longestUniqueSubstring = currentUniqueSubstring
	}

	return longestUniqueSubstring
}

func main() {
	str := "clementisacap"
	longestUniqueSubstring := longestUniqueSubstring(str)
	fmt.Println("Longest substring without duplicate characters:", longestUniqueSubstring)
}
