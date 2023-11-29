/*


Phone Number Mnemonics

Almost every digit is associated with some leners in the alphabet; this allows certain phone numbers to spell out actual words. For example, the phone number 8464747328 can be written as timisgreat; similarly, the phone number 2686463 can be wrinen as antoi ne or as ant6463 .
It's important to note that a phone number doesn't represent a single sequence of letters, but rather multiple combinations of letters. For instance, the digit 2 can represent three different letters (a, b, and c).
A mnemonic is defined as a pattern of letters, ideas, or associations that assist in remembering something. Companies oftentimes use a mnemonic for their phone number to make it easier to remember.
Given a stringified phone number of any non-zero length, write a function that returns all mnemonics for this phone number, in any order.
For this problem, a valid mnemonic may only contain letters and the digits e and 1 . In other words, if a digit is able to be represented by a lener, then it must be. Digits e and 1 are the only two digits that don't have letter representations on the keypad.
Note that you should rely on the keypad illustrated above for digit-letter associations.


Sample Input    : "1905"

Sample Output   : [
                    "1w0j",
                    "1w0k",
                    "1w0l",
                    "1x0j",
                    "1x0k",
                    "1x0l",
                    "1y0j",
                    "1y0k",
                    "1y0l",
                    "1z0j",
                    "1z0k",
                    "1z0l",
                ]





Hints
Hint 1

The first thing you'll need to do is create a mapping from digits to letters. You can do this by creating a hash table mapping all string digits to lists of their respective characters.

Hint 2

This problem can be solved fairly easily using recursion. Try generating all characters for the first digit in the phone number one at a time, and for each character, recursively performing the same action on the the next digit, and then on the digit after that, and so on and so forth until you've done so for all digits in the phone number.

Hint 3

You can recursively generate characters one digit at a time and store the intermediate results in a array. Once you've reached the last digit in the phone number, you can add the currently generated mnemonic (stored in the previously mentioned array) to a final array that stores all the results.




This code first defines a function findMnemonics that takes a phone number as a string and returns a slice of strings containing all the mnemonics for the phone number. The function first initializes an empty slice of strings to store the mnemonics. Then, it splits the phone number into a slice of digits and calls the findMnemonicsRecursive function to recursively find all the mnemonics.

The findMnemonicsRecursive function takes a slice of digits, an index into the slice, a current mnemonic, and a pointer to a slice of strings to store the mnemonics. The function first checks if the index has reached the end of the slice, in which case it appends the current mnemonic to the slice of mnemonics and returns.

Otherwise, it gets the letters for the current digit and iterates over the letters. For each letter, it creates a new mnemonic by appending the letter to the current mnemonic and then recursively calls itself to find the mnemonics for the remaining digits.

The getLetters function takes a digit as a string and returns a string of all the letters associated with that digit. The function uses a switch statement to map each digit to its corresponding letters.

The main function calls the findMnemonics function with the given phone number and prints the result.

*/

package main

import (
	"fmt"
	"strings"
)

func findMnemonics(phoneNumber string) []string {
	mnemonics := make([]string, 0)
	phoneDigits := strings.Split(phoneNumber, "")

	findMnemonicsRecursive(phoneDigits, 0, "", &mnemonics)

	return mnemonics
}

func findMnemonicsRecursive(phoneDigits []string, index int, currentMnemonic string, mnemonics *[]string) {
	if index == len(phoneDigits) {
		*mnemonics = append(*mnemonics, currentMnemonic)
		return
	}

	digit := phoneDigits[index]
	letters := getLetters(digit)

	for _, letter := range letters {
		newMnemonic := currentMnemonic + string(letter)
		findMnemonicsRecursive(phoneDigits, index+1, newMnemonic, mnemonics)
	}
}

func getLetters(digit string) string {
	switch digit {
	case "0":
		return "0"
	case "1":
		return "1"
	case "2":
		return "abc"
	case "3":
		return "def"
	case "4":
		return "ghi"
	case "5":
		return "jkl"
	case "6":
		return "mno"
	case "7":
		return "pqrs"
	case "8":
		return "tuv"
	case "9":
		return "wxyz"
	default:
		return ""
	}
}

func main() {
	phoneNumber := "1905"

	mnemonics := findMnemonics(phoneNumber)
	fmt.Println("Mnemonics for", phoneNumber, ":", mnemonics)
}
