/*


Write a function that takes in an array of integers and returns an array of length 2 representing the largest range of integers contained 1n that array.
The first number 1n the output array should be the first number 1n the range, while the second number should be the last number 1n the range.
A range of numbers 1s defined as a set of numbers that come right after each other 1n the set of real integers. For instance, the output array [2, 6] represents the range {2, 3, 4, 5, 6} , which is a range of length S. Note that numbers don't need to be sorted or adJacent 1n the input array 1n order to form a range.
You can assume that there will only be one largest range.
Sample Input
array = [1, 11, 3, 0, 15, 5, 2, 4, 10, 7, 12, 6]
Sample Output
[0, 7]






*/

package main

import (
	"fmt"
	"sort"
)

func largestRange(nums []int) []int {
	// Sort the array in ascending order.
	sort.Ints(nums)

	// Initialize the largest range and its length.
	largestRange := []int{nums[0], nums[0]}
	largestRangeLength := 1

	// Iterate over the array and keep track of the current range and its length.
	currentRange := []int{nums[0], nums[0]}
	currentRangeLength := 1

	for i := 1; i < len(nums); i++ {
		// If the current number is within the current range, increment the current range length.
		if nums[i] <= currentRange[1]+1 {
			currentRangeLength++
			currentRange[1] = nums[i]
		} else {
			// If the current range length is greater than the largest range length, update the largest range.
			if currentRangeLength > largestRangeLength {
				largestRange = currentRange
				largestRangeLength = currentRangeLength
			}

			// Start a new range with the current number.
			currentRange = []int{nums[i], nums[i]}
			currentRangeLength = 1
		}
	}

	// If the current range length is greater than the largest range length, update the largest range.
	if currentRangeLength > largestRangeLength {
		largestRange = currentRange
	}

	return largestRange
}

func main() {
	array := []int{1, 11, 3, 0, 15, 5, 2, 4, 10, 7, 12, 6}

	largestRange := largestRange(array)

	fmt.Println(largestRange)
}
