/*
You·re given two non-empty strings. The first one is a pattern consisting of only "x" sand/ or "y" s; the other one is a normal string of alphanumeric characters. Write a function that checks whether the normal string matches the pattern.
A string S0 1s said to match a pattern 1f replacing all "x" s 1n the pattern with some non-empty substring Sl of S0 and replacing all "y" s 1n the pattern with some non-empty substring S2 of S0 yields the same string S0
If the input stnng doesn't match the input pattern. the function should return an empty array; otherwise. 1t should return an array holding the strings Sl and S2 that represent "x" and "y" 1n the normal string, 1n that order. If the pattern doesn't contain any "x" s or "y" s, the respective letter should be represented by an empty string in the final array that you return.
You can assume that there will never be more than one pair of strings S1 and S2 that appropriately represent "x" and
"y" 1n the normal string.


Sample Input

pattern = "xxyxxy"
string = "gogopowerrangergogopowerranger"

Sample Output

["go", "powerranger"]






Hints
Hint 1

Start by checking if the pattern starts with an "x". If it doesn't, consider generating a new pattern that swaps all "x"s for "y"s and vice versa; this might greatly simplify the rest of your algorithm. Make sure to keep track of whether or not you do this swap, as your final answer will be affected by it.

Hint 2

Use a hash table to store the number of "x"s and "y"s that appear in the pattern, and keep track of the position of the first "y". Knowing how many "x"s and "y"s appear in the pattern, paired with the length of the main string which you have access to, will allow you to quickly test out various possible lengths for "x" and "y". Knowing where the first "y" appears in the pattern will allow you to actually generate potential substrings.

Hint 3

Traverse the main string and try different combinations of substrings that could represent "x" and "y". For each potential combination, map the new pattern mentioned in Hint #1 and see if it matches the main string.




This code iterates through the pattern and the normal string simultaneously. For each 'x' in the pattern, it checks if the corresponding character in the normal string matches the current xSubstr value. If it doesn't, it sets the matched flag to false. Similarly, for each 'y' in the pattern, it checks if the corresponding character in the normal string matches the current ySubstr value. If it doesn't, it sets the matched flag to false.

If the matched flag remains true, it means that all the occurrences of 'x' and 'y' in the pattern have been matched with corresponding substrings in the normal string. It then constructs the final matched pattern by replacing all 'x's and 'y's with their corresponding substrings and compares it to the original normal string. If they match, it returns the matched substrings as an array; otherwise, it returns an empty array.



*/

package main

import (
	"fmt"
	"strings"
)

func checkPattern(pattern, str string) []string {
	xSubstr := ""
	ySubstr := ""
	matched := true

	for i, char := range pattern {
		if char == 'x' {
			if xSubstr == "" {
				xSubstr = str[i]
			} else if xSubstr != str[i] {
				matched = false
				break
			}
		} else if char == 'y' {
			if ySubstr == "" {
				ySubstr = str[i]
			} else if ySubstr != str[i] {
				matched = false
				break
			}
		}
	}

	if !matched {
		return []string{}
	}

	pattern = strings.ReplaceAll(pattern, "x", xSubstr)
	pattern = strings.ReplaceAll(pattern, "y", ySubstr)

	if pattern == str {
		return []string{xSubstr, ySubstr}
	}

	return []string{}
}

func main() {
	pattern := "xxyxxy"
	str := "gogopowerrangergogopowerranger"

	matchedSubstrs := checkPattern(pattern, str)
	fmt.Println("Matched substrings:", matchedSubstrs)
}
