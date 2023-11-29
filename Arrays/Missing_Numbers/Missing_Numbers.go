/*

You're given an unordered list of unique integers nums in the range [1, n] , where n represents the length of nums + 2 . This means that two numbers in this range are missing from the list.
Write a function that takes in this list and returns a new list with the two missing numbers, sorted numerically.

Example:
Input:
nums = [3, 1, 4]
Output:
[2, 5]




*/

package main

import (
	"fmt"
	"sort"
)

func findMissingNumbers(nums []int) []int {
	// Sort the input list.
	sort.Ints(nums)

	// Initialize the two missing numbers.
	missingNumbers := []int{}

	// Iterate over the sorted list and find the two missing numbers.
	for i := 1; i <= len(nums)+2; i++ {
		if i != nums[i-1] {
			missingNumbers = append(missingNumbers, i)
		}
	}

	// Return the two missing numbers.
	return missingNumbers
}

func main() {
	nums := []int{3, 1, 4}

	missingNumbers := findMissingNumbers(nums)

	fmt.Println(missingNumbers)
}
