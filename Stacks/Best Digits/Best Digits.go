/*

Write a function that takes a positive integer represented as a string number and an integer numDigits-
Remove numDigits from the string so that the number represented by the string is as large as possible
afterwards.
Note that the order of the remaining digits cannot be changed. You can assume numDigits will always be
than the length of number and greater than or equal to 0.

Sample Input

number = "1432219"

numDigits = 3

Sample Output
"2219"






This code first initializes an array remainingDigits to store the remaining digits after removing the specified number. Then, it iterates through the digits of the original number. For each digit, it checks if there is a larger digit among the remaining digits. If not, it adds the current digit to the remainingDigits array. Finally, it constructs the remaining number from the remainingDigits array and returns it.




*/

package main

import (
	"fmt"
	"strconv"
)

func removeDigits(number string, numDigits int) string {
	if numDigits == 0 {
		return number
	}

	// Initialize an array to store the remaining digits
	remainingDigits := make([]byte, len(number)-numDigits)
	index := 0

	// Iterate through the digits of the number
	for i := 0; i < len(number); i++ {
		currentDigit, _ := strconv.Atoi(string(number[i]))

		// Check if the current digit is the largest among the remaining digits
		largestFound := false
		for j := index; j < len(remainingDigits); j++ {
			remainingDigit := int(remainingDigits[j])

			if remainingDigit > currentDigit {
				largestFound = true
				break
			}
		}

		// If the current digit is the largest, add it to the remaining digits
		if !largestFound {
			remainingDigits[index] = byte(currentDigit)
			index++
		}
	}

	return string(remainingDigits)
}

func main() {
	number := "1432219"
	numDigits := 3

	result := removeDigits(number, numDigits)
	fmt.Println("Result:", result)
}
