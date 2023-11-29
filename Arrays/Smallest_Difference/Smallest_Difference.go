/*

Write a function that takes in two non-empty arrays of integers, finds the pair of numbers (one from each array) whose absolute difference is closest to zero, and returns an array containing these two numbers, with the number from the first array in the first position.
Note that the absolute difference of two integers is the distance between them on the real number line. For example, the absolute difference of -5 and 5 is 10, and the absolute difference of -5 and -4 is 1.
You can assume that there will only be one pair of numbers with the smallest difference.

Sample Input
arrayOne = [-1, 5, 10, 20, 28, 3]
arrayTwo = [26, 134, 135, 15, 17]

Sample Output
[28, 26]



Hints
Hint 1
Instead of generating all possible pairs of numbers, try somehow only looking at pairs that you know could actually have the smallest difference. How can you accomplish this?
Hint 2
Would it help if the two arrays were sorted? If the arrays were sorted and you were looking at a given pair of numbers, could you efficiently find the next pair of numbers to look at? What are the runtime implications of sorting the arrays?
Hint 3
Start by sorting both arrays, as per Hint #2. Put a pointer at the beginning of both arrays and evaluate the absolute difference of the pointer-numbers. If the difference is equal to zero, then you've found the closest pair; otherwise, increment the pointer of the smaller of the two numbers to find a potentially better pair. Continue until you get a pair with a difference of zero or until one of the pointers gets out of range of its array.


This function works by first sorting the two arrays. This is necessary to ensure that the function can efficiently find the pair of numbers with the smallest difference. The function then initializes the minimum absolute difference and the closest pair.

The function then iterates over the two arrays. For each pair of elements in the two arrays, the function calculates the absolute difference between the two elements. If the absolute difference is less than the minimum absolute difference, the function updates the minimum absolute difference and the closest pair.

The function continues to iterate over the two arrays until it reaches the end of either array. The function then returns the closest pair.



*/

package main

import (
	"fmt"
	"math"
	"sort"
)

func findClosestPair(arrayOne []int, arrayTwo []int) []int {
	// Sort the two arrays.
	sort.Ints(arrayOne)
	sort.Ints(arrayTwo)

	// Initialize the minimum absolute difference and the closest pair.
	minAbsoluteDifference := math.Inf(1)
	closestPair := []int{0, 0}

	// Iterate over the two arrays.
	i := 0
	j := 0
	for i < len(arrayOne) && j < len(arrayTwo) {
		// Calculate the absolute difference between the current elements in the two arrays.
		absoluteDifference := int(math.Abs(float64(arrayOne[i]) - float64(arrayTwo[j])))

		// If the absolute difference is less than the minimum absolute difference, update the minimum absolute difference and the closest pair.
		if absoluteDifference < minAbsoluteDifference {
			minAbsoluteDifference = absoluteDifference
			closestPair[0] = arrayOne[i]
			closestPair[1] = arrayTwo[j]
		}

		// If the current element in the first array is less than the current element in the second array, increment the index of the current element in the first array.
		if arrayOne[i] < arrayTwo[j] {
			i++
		} else {
			// Otherwise, increment the index of the current element in the second array.
			j++
		}
	}

	// Return the closest pair.
	return closestPair
}

func main() {
	arrayOne := []int{-1, 5, 10, 20, 28, 3}
	arrayTwo := []int{26, 134, 135, 15, 17}

	closestPair := findClosestPair(arrayOne, arrayTwo)

	fmt.Println(closestPair)
}
