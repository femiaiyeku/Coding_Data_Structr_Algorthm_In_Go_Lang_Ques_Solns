/*

Write a function that takes in a sorted array of integers as well as a target integer.
The function should use a variation of the Binary Search algorithm to find a range of indices in between
which the target number 1s contained 1n the array and should return this range in the form of an array
The first number 1n the output array should represent the first index at which the target number 1s located,
while the second number should represent the last index at which the target number 1s located The function
should return [-1, -1] if the integer isn't contained in the array
If you're unfamiliar with Binary Search, we recommend watching the Conceptual Overview section of the Binary
Search question's video explanation before starting to code.

Sample Input:
array = [0, 1, 21, 33, 45, 45, 45, 45, 45, 45, 61, 71, 73]
target = 45


Sample Output:
[4, 9]






Hints
Hint 1

The Binary Search algorithm involves a left pointer and a right pointer and using those pointers to find the middle number in an array in an effort to find a target number. Unlike with normal Binary Search however, here you cannot simply find the middle number of the array, compare it to the target, and stop once you find it because you are looking for a range rather than a single number. Instead, realize that whenever you find the middle number in the array, the following two scenarios are possible: either the middle number is not equal to the target number, in which case you must proceed with normal Binary Search, or the middle number is equal to the target number, in which case you must figure out if this middle number is an extremity of the range or not.

Hint 2

Try applying an altered version of Binary Search twice: once to find the left extremity of the range and once to find the right extremity of the range. How can you accomplish this? What are the time complexity implications of this approach?

Hint 3

Can you implement this algorithm iteratively? Are there any advantages to doing so?


This code defines two functions: findRange, which takes a sorted array and a target integer as input and returns a slice containing the indices of the first and last occurrences of the target integer in the array, and findLeftmostIndex, which takes a sorted array and a target integer as input and returns the index of the leftmost occurrence of the target integer in the array.

The findRange function first calls findLeftmostIndex to find the index of the leftmost occurrence of the target integer in the array. If the leftmost index is -1, it means that the target integer is not present in the array, and it returns [-1, -1]. Otherwise, it calls findRightmostIndex to find the index of the rightmost occurrence of the target integer in the array, starting from the leftmost index. Finally, it returns a slice containing the leftmost and rightmost indices.

The findLeftmostIndex function uses a modified binary search algorithm to find the leftmost occurrence of the target integer. The modified binary search algorithm works like the standard binary search algorithm, but it stops when it finds the first occurrence of the target integer, rather than continuing to search for the middle element.

The findRightmostIndex function uses another modified binary search algorithm to find the rightmost occurrence of the target integer. The modified binary search algorithm works like the standard binary search algorithm, but it stops when it finds the last occurrence of the target integer, rather than continuing to search for the middle element.



*/

package main

import (
	"fmt"
)

func findRange(array []int, target int) []int {
	leftIndex := findLeftmostIndex(array, target)
	if leftIndex == -1 {
		return []int{-1, -1}
	}

	rightIndex := findRightmostIndex(array, target, leftIndex)

	return []int{leftIndex, rightIndex}
}

func findLeftmostIndex(array []int, target int) int {
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := (low + high) / 2

		if array[mid] == target && (mid == 0 || array[mid-1] != target) {
			return mid
		} else if array[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func findRightmostIndex(array []int, target int, startIndex int) int {
	low := startIndex
	high := len(array) - 1

	for low <= high {
		mid := (low + high) / 2

		if array[mid] == target && (mid == len(array)-1 || array[mid+1] != target) {
			return mid
		} else if array[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return high
}

func main() {
	array := []int{0, 1, 21, 33, 45, 45, 45, 45, 45, 45, 61, 71, 73}
	target := 45

	rangeIndices := findRange(array, target)
	fmt.Println("Range of indices:", rangeIndices)
}
