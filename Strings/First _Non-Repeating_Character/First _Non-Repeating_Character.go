/*



Write a function that takes in a string of lowercase English-alphabet letters and returns the index of the string's first non-repeating rhr1rr1rtPr
The first non-repeating character i􀀖 the first character in a string that occurs only once.
If the input string doesn't have any non-repeating characters, your function should return -1 .


Sample Input


string = "abcdcaf"


Sample Output


1 // The first non-repeating character is "b" and is found at index 1.



This code iterates through the string and maintains a map of character counts. If a character count is 1, it means that character appears only once and is the first non-repeating character. It returns the index of that character. If no character with a count of 1 is found, it returns -1.




*/

package main

import (
	"fmt"
)

func findFirstNonRepeatingChar(str string) int {
	charCounts := make(map[byte]int)

	for _, char := range str {
		charCounts[byte(char)]++
	}

	for i, char := range str {
		if charCounts[byte(char)] == 1 {
			return i
		}
	}

	return -1
}

func main() {
	str := "abcdcaf"
	index := findFirstNonRepeatingChar(str)
	fmt.Println("First non-repeating character index:", index)
}
