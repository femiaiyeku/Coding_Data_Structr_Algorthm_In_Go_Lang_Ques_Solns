/*

Write a function that takes in an array of integers and returns the length of the longest peak in the array.
A peak is defined as adjacent integers in the array that are strictly increasing until they reach a tip (the highest value in the peak),
at which point they become strictly decreasing. At least three integers are required to form a peak.
For example, the integers 1, 4, 10, 2 form a peak, but the integers 4, 0, 10 don't and neither do the integers
1, 2, 2, 0 . Similarly, the integers 1, 2, 3 don't form a peak because there aren't any strictly decreasing integers after the 3.

Sample Input
array = [1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3]

Sample Output
6 // 0, 10, 6, 5, -1, -3



Hints
Hint 1

You can solve this question by iterating through the array from left to right once.
Hint 2

Iterate through the array from left to right, and treat every integer as the potential tip of a peak. To be the tip of a peak, an integer has to be strictly greater than its adjacent integers. What can you do when you find an actual tip?
Hint 3

As you iterate through the array from left to right, whenever you find a tip of a peak, expand outwards from the tip until you no longer have a peak. Given what peaks look like and how many peaks can therefore fit in an array, realize that this process results in a linear-time algorithm. Make sure to keep track of the longest peak you find as you iterate through the array.




This function works by first sorting the input list. This is necessary to ensure that the function can efficiently find the two missing numbers. The function then iterates over the sorted list and compares each number to its expected value. If the two numbers are not equal, then the expected number is a missing number. The function adds the missing number to the list of missing numbers.

Once the function has finished iterating over the sorted list, it returns the list of missing numbers.

*/

package main

import (
	"fmt"
	"sort"
)

type Peak struct {
	startIndex int
	endIndex   int
	length     int
}

func findLongestPeak(array []int) int {
	// Sort the array in ascending order.
	sort.Ints(array)

	// Create a slice to store the peaks.
	peaks := []Peak{}

	// Iterate over the array and find all of the peaks.
	for i := 1; i < len(array)-1; i++ {
		if array[i] > array[i-1] && array[i] > array[i+1] {
			// Find the start and end indices of the peak.
			startIndex := i - 1
			endIndex := i + 1
			for startIndex >= 0 && array[startIndex] < array[startIndex+1] {
				startIndex--
			}
			for endIndex < len(array) && array[endIndex] > array[endIndex-1] {
				endIndex++
			}

			// Create a peak object and add it to the slice of peaks.
			peak := Peak{startIndex, endIndex, endIndex - startIndex}
			peaks = append(peaks, peak)
		}
	}

	// Find the longest peak.
	longestPeak := Peak{}
	for _, peak := range peaks {
		if peak.length > longestPeak.length {
			longestPeak = peak
		}
	}

	// Return the length of the longest peak.
	return longestPeak.length
}

func main() {
	array := []int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3}

	longestPeakLength := findLongestPeak(array)

	fmt.Println(longestPeakLength)
}
