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



*/

package main

func contains_all(big_string string, small_strings []string) []bool {

	result := make([]bool, len(small_strings))

	for i, small_string := range small_strings {
		found := false
		for j := 0; j <= len(big_string)-len(small_string); j++ {
			if big_string[j:j+len(small_string)] == small_string {
				found = true
				break
			}
		}

		result[i] = found
	}

	return result
}
