/*

Write a function that takes in an array of integers and returns a boolean representing whether the array is monotonic.
An array is said to be monotonic if its elements, from left to right, are entirely non-increasing or entirely non-decreasing.
Non-increasing elements aren't necessarily exclusively decreasing; they simply don't increase. Similarly, non-decreasing elements aren't necessarily exclusively increasing; they simply don't decrease.
Note that empty arrays and arrays of one element are monotonic.

Sample Input
array = [-1, -5, -10, -1100, -1100, -1101, -1102, -9001]

Sample Output
true

Hint 1
You can solve this question by iterating through the input array from left to right once.
Hint 2
Try iterating through the input array from left to right, in search of two adjacent integers that can indicate whether the array is trending upward or downward. Once you've found the tentative trend of the array, at each element thereafter, compare the element to the previous one; if this comparison breaks the previously identified trend, the array isn't monotonic.
Hint 3
Alternatively, you can start by assuming that the array is both entirely non-decreasing and entirely non-increasing. As you iterate through each element, perform a check to see if you can invalidate one or both of your assumptions.


Hints
Hint 1

You can solve this question by iterating through the input array from left to right once.
Hint 2

Try iterating through the input array from left to right, in search of two adjacent integers that can indicate whether the array is trending upward or downward. Once you've found the tentative trend of the array, at each element thereafter, compare the element to the previous one; if this comparison breaks the previously identified trend, the array isn't monotonic.
Hint 3

Alternatively, you can start by assuming that the array is both entirely non-decreasing and entirely non-increasing. As you iterate through each element, perform a check to see if you can invalidate one or both of your assumptions.





*/

package main

import (
	"fmt"
)

func isMonotonic(array []int) bool {
	// Check if the array is empty or has one element.
	if len(array) == 0 || len(array) == 1 {
		return true
	}

	// Initialize the direction of the array.
	direction := 0

	// Iterate over the array and check if the direction of the array changes.
	for i := 1; i < len(array); i++ {
		if array[i] > array[i-1] {
			if direction == -1 {
				return false
			}

			direction = 1
		} else if array[i] < array[i-1] {
			if direction == 1 {
				return false
			}

			direction = -1
		}
	}

	// If the direction of the array never changes, then the array is monotonic.
	return true
}

func main() {
	array := []int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001}

	isMonotonic := isMonotonic(array)

	fmt.Println(isMonotonic)
}
