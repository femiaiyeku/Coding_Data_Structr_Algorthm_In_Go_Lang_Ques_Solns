/*


Given a non-empty string of lowercase letters and a non-negative integer representing a key, write a function that
returns a new string obtained by shifting every letter in the input string by k positions in the alphabet, where k is
the key.
Note that letters should "wrap" around the alphabet; in other words, the letter shifted by one returns the letter

Sample Input:
string = "xyz"
key = 2

Sample Output:
"zab"


This code iterates through the characters of the input string and shifts each character by the given key. It checks if the new character goes beyond the lowercase alphabet ('z') and wraps it around if necessary. Finally, it appends the shifted characters to the shiftedString variable and returns the result.





*/

package main

import (
	"fmt"
)

func shiftCipher(str string, key int) string {
	shiftedString := ""

	for _, char := range str {
		newChar := char + byte(key)

		if newChar > 'z' {
			newChar -= 'z' + 1
		}

		shiftedString += string(newChar)
	}

	return shiftedString
}

func main() {
	inputString := "xyz"
	key := 2

	shiftedString := shiftCipher(inputString, key)
	fmt.Println("Shifted string:", shiftedString)
}
