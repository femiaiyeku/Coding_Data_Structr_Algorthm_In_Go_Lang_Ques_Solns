/*



Write a function that takes in a string of words separated by one or more whites paces and returns a string that has these words in reverse order. For example, given the string "tim is great" , your function should return "great is tim"
For this problem, a word can contain special characters, punctuation, and numbers. The words 1n the string will be separated by
one or more wh1tespaces, and the reversed string must contain the same wh1tespaces as the original string. For example, given the string "whi tespaces 4" you would be expected to return "4 whitespaces11
Note that you're nor allowed to to use any built-in split or reverse methods/functions. However, you are allowed to use a built-in join method/function.
Also note that the input string 1sn·t guaranteed to always contain words

Sample Input:
string = "AlgoExpert is the best!"

Sample Output:
"best! the is AlgoExpert"



This code iterates through the input string and keeps track of the start index of a word. When it encounters a whitespace, it checks if the previous word exists and adds it to the words slice if it does. Then, it updates the startIndex to the index after the whitespace. It continues this process until it reaches the end of the string. Finally, it reverses the words slice and joins them using a space as a delimiter, resulting in the reversed string.




*/

package main

import (
	"fmt"
)

func reverseWords(str string) string {
	words := make([]string, 0)
	startIndex := 0

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			if startIndex < i {
				words = append(words, str[startIndex:i])
			}
			startIndex = i + 1
		}
	}

	if startIndex < len(str) {
		words = append(words, str[startIndex:])
	}

	reversedString := ""

	for i := len(words) - 1; i >= 0; i-- {
		if i != 0 {
			reversedString += " "
		}
		reversedString += words[i]
	}

	return reversedString
}

func main() {
	str := "AlgoExpert is the best!"
	reversedString := reverseWords(str)
	fmt.Println("Reversed string:", reversedString)
}
