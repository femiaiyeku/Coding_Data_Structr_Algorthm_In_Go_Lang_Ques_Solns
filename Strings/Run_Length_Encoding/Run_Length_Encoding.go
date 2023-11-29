/*


Write a function that takes in a non-empty string and returns its run-length encoding.
From Wikipedia, "run-length encoding is a form of lossless data compression in which runs of data are stored as a
single data value and count, rather than as the original run." For this problem, a run of data is any sequence of
consecutive, identical characters. So the run "AAA" would be run-length-encoded as "3A"
To make things more complicated, however, the input string can contain all sorts of special characters, including
numbers. And since encoded data must be decodable, this means that we can't naively run-length-encode long
runs. For example, the run "AAAAAAAAAAAA" (12 A s), can't naively be encoded as "12A", since this string can
be decoded as either "AAAAAAAAAAAA" or "1AA". Thus, long runs (runs of 10 or more characters) should be
encoded in a split fashion; the aforementioned run should be encoded as "9A3A"


Sample Input:
string = "AAAAAAAAAAAAABBCCCCDD"

Sample Output:
"9A4A2B4C2D"




Hints
Hint 1
Traverse the input string and count the length of each run. As you traverse the string, what should you do when you reach a run of length 9 or the end of a run?
Hint 2
When you reach a run of length 9 or the end of a run, store the computed count for the run as well as its character (you'll likely need a list for these computed counts and characters), and reset the count to 1 before continuing to traverse the string.
Hint 3
Make sure that your solution correctly handles the last run in the string.





This code iterates through the string and keeps track of the current character and the current count of consecutive occurrences of that character. If the current character changes, it appends the current count and the previous character to the encoded string and resets the current character and count. If the current count reaches 10, it splits the long run into smaller chunks of 9 and 1. Finally, it appends the remaining count and character to the encoded string and returns the encoded string.


*/

package main

import (
	"fmt"
	"strconv"
)

func runLengthEncoding(str string) string {
	encodedString := ""
	currentCharacter := str[0]
	currentCount := 1

	for i := 1; i < len(str); i++ {
		if str[i] != currentCharacter {
			encodedString += strconv.Itoa(currentCount) + string(currentCharacter)
			currentCharacter = str[i]
			currentCount = 1
		} else {
			currentCount++

			// Split long runs (10 or more characters)
			if currentCount == 10 {
				encodedString += strconv.Itoa(9) + string(currentCharacter)
				currentCount = 1
			}
		}
	}

	encodedString += strconv.Itoa(currentCount) + string(currentCharacter)

	return encodedString
}

func main() {
	str := "AAAAAAAAAAAAABBCCCCDD"
	encodedString := runLengthEncoding(str)
	fmt.Println("Run-length encoding:", encodedString)
}
