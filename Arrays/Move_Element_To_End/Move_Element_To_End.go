/*

You re given an array of integers and an integer .
Write a function that moves all instances of that integer 1n the array to the end of the array and returns the array The function should perform this in place (1 e,
It should mutate the input array) and doesnt need to maintain the order of the other integers

Sample Input:
array = [2, 1, 2, 2, 2, 3, 4, 2]

Sample Output:
[1, 3, 4, 2, 2, 2, 2, 2] // the numbers 1, 3, and 4 could be ordered differently


This function works by initializing two pointers, one at the beginning of the array and the other at the end of the array. The function then iterates over the array until the two pointers meet. For each element in the array, the function checks if the element is equal to the element to move. If so, the function swaps the elements at the two pointers. This effectively moves the element to move to the end of the array.

The function continues to iterate over the array until the two pointers meet. At this point, all of the elements to move have been moved to the end of the array. The function then returns the modified array.


*/

package main

import (
	"fmt"
)

func moveElementToEnd(array []int, element int) []int {
	// Initialize two pointers.
	leftPointer := 0
	rightPointer := len(array) - 1

	// Iterate over the array until the two pointers meet.
	for leftPointer < rightPointer {
		// If the element at the left pointer is equal to the element to move, swap the elements at the left pointer and the right pointer.
		if array[leftPointer] == element {
			array[leftPointer], array[rightPointer] = array[rightPointer], array[leftPointer]
			rightPointer--
		} else {
			leftPointer++
		}
	}

	// Return the modified array.
	return array
}

func main() {
	array := []int{2, 1, 2, 2, 2, 3, 4, 2}
	element := 2

	array = moveElementToEnd(array, element)

	fmt.Println(array)
}
