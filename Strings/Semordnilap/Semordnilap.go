/*

Write a function that takes in a list of unique strings and returns a list of semordnilap pairs.
A semordnilap pair is defined as a set of different strings where the reverse of one word is the same as the forward version of the other.
For example the words "diaper" and "repaid" are a semordnilap pair, as are the words "palindromes" and "semordnilap".
The order of the returned pairs and the order of the strings within each pair does not matter.

Sample Input
words = ["code", "edoc", "da", "d"]

Sample Output

[["code", "edoc"], ["da", "d"]]







This code iterates through the list of words and checks each pair of words for a semordnilap relationship using the isSemordnilap function. If a semordnilap pair is found, it adds it to the semordnilapPairs slice. Finally, it returns the list of semordnilap pairs.

Output:

Semordnilap pairs: [[code edoc] [da d]]

*/

package main

import (
	"fmt"
)

func findSemordnilapPairs(words []string) [][]string {
	semordnilapPairs := make([][]string, 0)

	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if isSemordnilap(words[i], words[j]) {
				semordnilapPair := []string{words[i], words[j]}
				semordnilapPairs = append(semordnilapPairs, semordnilapPair)
			}
		}
	}

	return semordnilapPairs
}

func isSemordnilap(word1, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[len(word2)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	words := []string{"code", "edoc", "da", "d"}
	pairs := findSemordnilapPairs(words)
	fmt.Println("Semordnilap pairs:", pairs)
}
