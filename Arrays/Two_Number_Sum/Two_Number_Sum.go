/*

Write a function that takes in a non-ernpty array of distinct integers and an integer
representing a target surn. If any two nurnbers in the input array surn up to the target surn,
the function should return thern in an array, in any order. If no two nurnbers surn up to the
target surn, the function should return an ernpty array.
Note that the target sum has to be obtained by summing two different integers in the
array; you can't add a single integer to itself in order to obtain the target sum.
You can assume that there will be at most one pair of nurnbers summing up to the target
sum

Sample Input
array = [5, 1, 22, 25, 6, -1, 8, 10)
sequence = [l, 6, -1, 10)

Sample Output

true


This function works by first sorting the array in ascending order. This is necessary to ensure that the function can efficiently find the pair of numbers that sum up to the target sum. The function then initializes a map to store the elements of the array.

The function then iterates over the array. For each element, the function checks if the target sum minus the current element is present in the map. If it is, then there is a pair of numbers in the array that sum up to the target sum, and the function returns true. Otherwise, the function adds the current element to the map.

Once the function has finished iterating over the array, it checks if the map contains the target sum. If it does, then there is a pair of numbers in the array that sum up to the target sum, and the function returns true. Otherwise, the function returns false.


*/

package main

import (
	"fmt"
	"sort"
)

func isPairPresent(array []int, targetSum int) bool {
	// Sort the array in ascending order.
	sort.Ints(array)

	// Initialize a map to store the elements of the array.
	elementMap := make(map[int]bool)

	// Iterate over the array.
	for i := 0; i < len(array); i++ {
		// Get the current element.
		currentElement := array[i]

		// If the target sum minus the current element is present in the map, then there is a pair of numbers in the array that sum up to the target sum.
		if elementMap[targetSum-currentElement] {
			return true
		}

		// Add the current element to the map.
		elementMap[currentElement] = true
	}

	// If there is no pair of numbers in the array that sum up to the target sum, return false.
	return false
}

func main() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	targetSum := 11

	isPresent := isPairPresent(array, targetSum)

	fmt.Println(isPresent)
}
