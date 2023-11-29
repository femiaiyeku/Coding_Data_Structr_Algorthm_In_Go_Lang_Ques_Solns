/*

You're given a string of available characters and a string representing a document that you need to generate.
Write a function that determines if you can generate the document using the available characters. If you can
generate the document, your function should return true; otherwise, it should return false.
You're only able to generate the document if the frequency of unique characters in the characters string is greater
than or equal to the frequency of unique characters in the document string. For example, if you're given
characters= "abcabc" and document "aabbccc" you cannot generate the document because you're
missing one c.
The document that you need to create may contain any characters, including special characters, capital letters,
numbers, and spaces.

Note: you can always generate the empty string ("").


Sample Input

characters = "Bste!hetsi ogEAxpelrt x "

document = "AlgoExpert is the Best!"

Sample Output

true



This code iterates through the document string and counts the occurrences of each unique character. Then, it iterates through the characters string and counts the occurrences of each unique character. Finally, it compares the counts for each unique character in the document string to the corresponding counts in the characters string. If the count in the document string is greater than or equal to the count in the characters string, then the document can be generated; otherwise, it cannot.


*/

package main

import (
	"fmt"
)

func canGenerateDocument(characters, document string) bool {
	documentChars := make(map[byte]int)

	for _, char := range document {
		documentChars[byte(char)]++
	}

	availableChars := make(map[byte]int)

	for _, char := range characters {
		availableChars[byte(char)]++
	}

	for char, count := range documentChars {
		if availableChars[char] < count {
			return false
		}
	}

	return true
}

func main() {
	characters := "Bste!hetsi ogEAxpelrt x "
	document := "AlgoExpert is the Best!"

	canGenerate := canGenerateDocument(characters, document)
	fmt.Println("Can generate document:", canGenerate)
}
