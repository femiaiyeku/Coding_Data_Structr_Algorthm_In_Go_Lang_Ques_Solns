/*


Write a function that takes in an array of strings and groups anagrams together.
Anagrams are strings made up of exactly the same letters, where order doesn't matter. For example, "cinema" and "iceman" are anagrams; similarly, "foo" and "ofo" are anagrams.
Your function should return a list of anagram groups in no particular order.


Sample Input
words = ["yo", "act", "flop", "tac", "foo", "cat", "oy", "olfp"]

Sample Output
[["yo", "oy"], ["flop", "olfp"], ["act", "tac", "cat"], ["foo"]]


Hints
Hint 1

Try rearranging every input string such that each string's letters are ordered in alphabetical order. What can you do with the resulting strings?

Hint 2

For any two of the resulting strings mentioned in Hint #1 that are equal to each other, their original strings (with their letters in normal order) must be anagrams. Realizing this, you could bucket all of these resulting strings together, all the while keeping track of their original strings, to find the groups of anagrams.

Hint 3

Can you simply store the resulting strings mentioned in Hint #1 in a hash table and find the groups of anagrams using this hash table?



This code first creates a map to store anagrams together. It iterates through the words and sorts each word using the sortString function. The sorted word is used as the key in the map, and the corresponding value is a list of words that share the same sorted word. Finally, it iterates through the map and appends each list of anagrams to the anagramGroups array.


*/

package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(words []string) [][]string {
	anagramGroups := [][]string{}

	anagramMap := make(map[string][]string)
	for _, word := range words {
		sortedWord := sortString(word)
		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}

	for _, words := range anagramMap {
		anagramGroups = append(anagramGroups, words)
	}

	return anagramGroups
}

func sortString(word string) string {
	chars := []byte(word)
	sort.Sort(sort.Bytes(chars))

	return string(chars)
}

func main() {
	words := []string{"yo", "act", "flop", "tac", "foo", "cat", "oy", "olfp"}
	anagramGroups := groupAnagrams(words)
	fmt.Println("Anagram groups:", anagramGroups)
}
