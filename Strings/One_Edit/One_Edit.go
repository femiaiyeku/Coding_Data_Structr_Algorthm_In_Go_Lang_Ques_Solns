/*

You're given two strings stringone and stringTwo. Write a function that determines if these two strings can
be made equal using only one edit.
There are 3 possible edits:
• Replace: One character in one string is swapped for a different character.
• Add: One character is added at any index in one string.
• Remove: One character is removed at any index in one string.
Note that both strings will contain at least one character. If the strings are the same, your function should return
true.

Sample Input
stringOne = "abc"

stringTwo = "ab"

Sample Output
true

Explanation
We can make stringOne equal to stringTwo by removing the "c" from stringOne.
Alternatively, we can make stringOne equal to stringTwo by deleting the "b" from stringTwo.
Both of these edits are one edit away from stringOne and stringTwo.




This code handles three cases:

If the strings are the same length, it checks if they have exactly one different character.

If the strings are different lengths, it makes sure the difference in length is not more than one. Then, it iterates through the strings and checks if they have the same characters at the same indices. If they differ at one index, it checks if the rest of the strings match. If they do, it means that one character can be added or removed to make the strings equal.

If none of the above conditions are met, it returns false





*/

package main

import (
	"fmt"
	"strings"
)

func canOneEdit(stringOne, stringTwo string) bool {
	if len(stringOne) == len(stringTwo) {
		if strings.Compare(stringOne, stringTwo) == 0 {
			return true
		}

		differences := 0

		for i := 0; i < len(stringOne); i++ {
			if stringOne[i] != stringTwo[i] {
				differences++
			}
		}

		return differences == 1
	} else {
		if len(stringOne) > len(stringTwo) {
			stringOne, stringTwo = stringTwo, stringOne
		}

		if len(stringTwo)-len(stringOne) > 1 {
			return false
		}

		for i := 0; i < len(stringOne); i++ {
			if stringOne[i] != stringTwo[i] {
				return stringOne[i:] == stringTwo[i+1:]
			}
		}

		return true
	}
}

func main() {
	stringOne := "abc"
	stringTwo := "ab"

	canEdit := canOneEdit(stringOne, stringTwo)
	fmt.Println("Can make strings equal with one edit:", canEdit)
}
