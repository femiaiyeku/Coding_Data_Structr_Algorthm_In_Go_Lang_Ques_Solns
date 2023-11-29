/*

Write a function that takes in an array of words and returns the smallest array of characters needed to form all of
the words. The characters don't need to be in any particular order.
For example, the characters ["y", "r", "o", "u"] are needed to form the words
["your", "you", "or", "yo"].
Note: the input words won't contain any spaces; however, they might contain punctuation and/or special
characters.

Sample Input
words = ["this", "that", "did", "deed", "them!", "a"]

Sample Output
["t", "t", "h", "i", "s", "a", "d", "d", "e", "e", "m", "!"]






This code iterates through the array of words and maintains a map to keep track of all the unique characters used in the words. Then, it creates a new array to store the characters and appends each unique character from the map to the array. Finally, it returns the array of characters.



*/

package main

import "fmt"

func findSmallestCharacters(words []string) []string {
	characters := make(map[byte]bool)

	for _, word := range words {
		for _, char := range word {
			characters[byte(char)] = true
		}
	}

	result := make([]string, 0, len(characters))
	for char, _ := range characters {
		result = append(result, string(char))
	}

	return result
}

func main() {
	words := []string{"this", "that", "did", "deed", "them!", "a"}
	characters := findSmallestCharacters(words)
	fmt.Println("Smallest characters:", characters)
}
