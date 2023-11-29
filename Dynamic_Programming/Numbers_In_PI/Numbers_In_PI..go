/*

Given a string representation of the first n d1g1ts of Pi and a 11st of positive integers (all 1n string format),
write a function that returns the smallest number of spaces that can be added to the n d1g1ts of P1 such that all resulting numbers are found 1n the list of integers.
Note that a single number can appear multiple times 1n the resulting numbers. For example, 1f P1 1s "3141" and the numbers are [" l", "3", "4"] ,
the number "l" 1s allowed to appear twice 1n the 11st of resulting numbers after three spaces are added: "3 I 1 I 4 I l"
If no number of spaces to be added exists such that all resulting numbers are found 1n the 11st of integers, the function should return -1


Sample Input
pi = "3141592653589793238462643383279"
numbers = [
    "314159265358979323846",
    "26433",
    "8",
    "3279",
    "314159265",
    "35897932384626433832",
    "79",
]

Sample Output
2 // "314159265 | 35897932384626433832 | 79"




Hints
Hint 1

You'll need to look numbers up quickly; is the input array the best data structure for this?

Hint 2

Dump every favorite number in a hash table for fast look-up. Iterate through the digits of Pi, checking if every prefix of the n digits is a favorite number. What should you do if you find that a prefix of the n digits of Pi is a favorite number?

Hint 3

Going off of Hint #2, if you find a prefix of the n digits of Pi that is a favorite number, try adding 1 space after it and then recursively calculating the smallest number of spaces in the suffix that comes after it. Do this for every prefix, and you'll find the answer. Can this method be optimized with a cache?


*/

package main

import (
	"fmt"
	"strings"
)

func minSpaces(pi, numbers string) int {
	// Create a map to store the frequency of each number in the list of integers
	numberFrequency := make(map[string]int)
	for _, number := range strings.Split(numbers, ",") {
		numberFrequency[number]++
	}

	// Initialize the minimum number of spaces to -1 (impossible case)
	minSpaces := -1

	// Try adding spaces at each position in Pi
	for i := 0; i < len(pi); i++ {
		// Split Pi into two parts based on the current position
		leftPart := pi[:i]
		rightPart := pi[i:]

		// Create a set to store the unique numbers formed by adding spaces
		uniqueNumbers := make(map[string]bool)

		// Add all possible combinations of numbers formed by adding spaces
		for j := 0; j <= len(rightPart); j++ {
			currentNumber := leftPart + " " + rightPart[:j]
			uniqueNumbers[currentNumber] = true
		}

		// Check if all numbers in the uniqueNumbers set are present in the numberFrequency map
		allNumbersFound := true
		for number := range uniqueNumbers {
			if numberFrequency[number] == 0 {
				allNumbersFound = false
				break
			}
		}

		// If all numbers are found, update the minimum number of spaces
		if allNumbersFound {
			numberOfSpaces := strings.Count(pi[:i]+strings.Join(strings.Split(pi[i:], " "), " "), " ")
			if minSpaces == -1 || numberOfSpaces < minSpaces {
				minSpaces = numberOfSpaces
			}
		}
	}

	// Return the minimum number of spaces or -1 if impossible
	return minSpaces
}

func main() {
	// Set the input values.
	pi := "3141592653589793238462643383279"
	numbers := "314159265358979323846,26433,8,3279,314159265,35897932384626433832,79"

	// Find the smallest number of spaces to be added to Pi.
	minSpacesNeeded := minSpaces(pi, numbers)

	// Print the minimum number of spaces or -1 if impossible.
	if minSpacesNeeded == -1 {
		fmt.Println("Impossible to add spaces")
	} else {
		fmt.Println(minSpacesNeeded)
	}
}
