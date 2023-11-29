/*

Write a function that takes in a pos tive integer numberOfTags and returns a list of all the valid strings that you can generate with that number of matched <div></di v> tags.
A string is valid and contains matched <div></div> tags if for every opening tag <div> , there's a closing tag </div> that comes after the opening tag and that isn't used as a closing tag for another opening tag. Each output string should contain exactly numberOfTags opening tags and numberOfTags closing tags.
For example, given numberOfTags = 2 , the valid strings to return would be:

["<div></div><div></div>", "<div><div></div></div>"]

Note that the output strings don't need to be in any particular order.


Sample Input

numberOfTags = 3

Sample Output

    [
    "<div><div><div></div></div></div>",
    "<div><div></div><div></div></div>",
    "<div><div></div></div><div></div>",
    "<div></div><div><div></div></div>",
    "<div></div><div></div><div></div>"
    ]

    // The strings could be ordered differently.

    // For example, the following order would also be valid:



Hints
Hint 1

The brute-force approach to solve this problem is to generate every single possible string that contains numberOfTags tags and to then check all of those strings to see if they're valid. Can you think of a better way to do this?

Hint 2

To solve this problem optimally, you'll have to incrementally build valid strings by adding

and
tags to already valid partial strings. While doing this, you can avoid creating strings that will never lead to a valid final string by following two rules:
If a string has fewer opening tags than numberOfTags, it's valid to add an opening tag to the end of it.
If a string has fewer closing tags than opening tags, it's valid to add a closing tag to the end of it.
Hint 3

Using the rules defined in Hint #2, write a recursive algorithm that generates all possible valid strings. You'll need to keep track of how many opening and closing tags each partial string has available (at each recursive call), and you'll simply follow the rules outlined in Hint #2. Once a string has no more opening and closing tags available, you can add it to your final list of strings. Your first call to the function will start with an empty string as the partial string and with numberOfTags as the number of opening and closing tags available. For example, after you add an opening tag to a partial string, you'll recursively call the function like this: recursiveFunction(partialStringWithExtraOpeningTag, openingTags - 1, closingTags)


This code first defines two arrays, openingTags and closingTags, each containing the opening and closing tags <div> and </div>, respectively. Then, it calls the generateValidStringsRecursive function to generate all the valid strings.

The generateValidStringsRecursive function takes four arguments:

openingTags: An array of opening tags
closingTags: An array of closing tags
currentTag: The index of the current tag to add to the string
currentString: The current string being built
validStrings: A pointer to a slice of strings that will store the valid strings
The function first checks if the current tag is the last tag in the array. If so, it adds the current string to the validStrings slice and returns. Otherwise, it checks if the current string starts with the closing tag corresponding to the current opening tag. If so, it returns because this string is not valid.

Otherwise, it calls itself recursively twice: once with the opening tag added to the current string, and once with the closing tag added to the current string. This process continues until all valid strings have been generated.







*/

package main

import (
	"fmt"
	"strings"
)

func generateValidStrings(numberOfTags int) []string {
	if numberOfTags <= 0 {
		return []string{}
	}

	openingTags := make([]string, numberOfTags)
	closingTags := make([]string, numberOfTags)

	for i := 0; i < numberOfTags; i++ {
		openingTags[i] = "<div>"
		closingTags[i] = "</div>"
	}

	validStrings := []string{}
	generateValidStringsRecursive(openingTags, closingTags, 0, "", &validStrings)

	return validStrings
}

func generateValidStringsRecursive(openingTags []string, closingTags []string, currentTag int, currentString string, validStrings *[]string) {
	if currentTag == len(openingTags) {
		*validStrings = append(*validStrings, currentString)
		return
	}

	if strings.HasPrefix(currentString, closingTags[currentTag]) {
		return
	}

	generateValidStringsRecursive(openingTags, closingTags, currentTag+1, currentString+openingTags[currentTag], validStrings)
	generateValidStringsRecursive(openingTags, closingTags, currentTag+1, currentString+closingTags[currentTag], validStrings)
}

func main() {
	numberOfTags := 3
	validStrings := generateValidStrings(numberOfTags)
	fmt.Println(validStrings)
}
