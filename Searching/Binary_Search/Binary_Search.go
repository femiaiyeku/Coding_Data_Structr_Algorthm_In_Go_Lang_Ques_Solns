/*

Write a function that takes in a sorted array of integers as well as a target integer. The function should use the Binary Search
algorithm to determine if the target integer is contained in the array and should return its index if it is, otherwise -1.

Sample Input:

array = [0, 1, 21, 33, 45, 45, 61, 71, 72, 73]

target = 33

Sample Output:
3



Hints
Hint 1
The Binary Search algorithm works by finding the number in the middle of the input array and comparing it to the target number. Given that the array is sorted, if this middle number is smaller than the target number, then the entire left part of the array is no longer worth exploring since the target number can no longer be in it; similarly, if the middle number is greater than the target number, then the entire right part of the array is no longer worth exploring. Applying this logic recursively eliminates half of the array until the number is found or until the array runs out of numbers.
Hint 2
Write a helper function that takes in two additional arguments: a left pointer and a right pointer representing the indices at the extremities of the array (or subarray) that you are applying Binary Search on. The first time this helper function is called, the left pointer should be zero and the right pointer should be the final index of the input array. To find the index of the middle number mentioned in Hint #1, simply round down the number obtained from: (left + right) / 2. Apply this logic recursively until you find the target number or until the left pointer becomes greater than the right pointer.
Hint 3
Can you implement this algorithm iteratively? Are there any advantages to doing so?




This code defines a function binarySearch that takes a sorted array of integers and a target integer as input and returns the index of the target integer if it is contained in the array, or -1 if it is not. The function uses the binary search algorithm to efficiently search for the target integer.

The main function creates an array of integers and a target integer, and then calls the binarySearch function to search for the target integer in the array. It then prints the result.

*/

package main

import (
	"fmt"
)

func binarySearch(array []int, target int) int {
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := (low + high) / 2

		if array[mid] == target {
			return mid
		} else if array[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	array := []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}
	target := 33

	index := binarySearch(array, target)

	if index != -1 {
		fmt.Printf("Target %d found at index %d\n", target, index)
	} else {
		fmt.Println("Target not found")
	}
}
