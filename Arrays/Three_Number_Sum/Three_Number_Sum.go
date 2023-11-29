/*




Write a function that takes in a non-empty array of distinct integers and an integer representing a target sum. The
function should find all triplets in the array that sum up to the target sum and return a two-dimensional array of all
these triplets. The numbers in each triplet should be ordered in ascending order, and the triplets themselves
should be ordered in ascending order with respect to the numbers they hold.
If no three numbers sum up to the target sum, the function should return an empty array.


Sample Input:
array = [12, 3, 1, 2, -6, 5, -8, 6]
targetSum = 0


Sample Output:
[[-8, 2, 6], [-8, 3, 5], [-6, 1, 5]]



This function works by first sorting the array in ascending order. This is necessary to ensure that the function can efficiently find all triplets that sum up to the target sum. The function then initializes the result array.

The function then iterates over the array. For each element in the array, the function initializes the left and right pointers. The function then iterates over the array again, starting from the left pointer and ending at the right pointer. For each iteration of the loop, the function calculates the total sum of the three elements. If the total sum is equal to the target sum, the function adds the triplet to the result array. Otherwise, the function increments or decrements the left or right pointer, depending on whether the total sum is less than or greater than the target sum.

Once the function has finished iterating over the array, it returns the result array.


*/

package main

import (
	"fmt"
	"sort"
)

func findTriplets(array []int, targetSum int) [][]int {
	// Sort the array in ascending order.
	sort.Ints(array)

	// Initialize the result array.
	result := [][]int{}

	// Iterate over the array.
	for i := 0; i < len(array)-2; i++ {
		// Initialize the left and right pointers.
		leftPointer := i + 1
		rightPointer := len(array) - 1

		// Iterate over the array again, starting from the left pointer and ending at the right pointer.
		for leftPointer < rightPointer {
			// Calculate the total sum of the three elements.
			totalSum := array[i] + array[leftPointer] + array[rightPointer]

			// If the total sum is equal to the target sum, add the triplet to the result array.
			if totalSum == targetSum {
				result = append(result, []int{array[i], array[leftPointer], array[rightPointer]})

				// Increment the left pointer.
				leftPointer++

				// Decrement the right pointer.
				rightPointer--
			} else if totalSum < targetSum {
				// Increment the left pointer.
				leftPointer++
			} else {
				// Decrement the right pointer.
				rightPointer--
			}
		}
	}

	// Return the result array.
	return result
}

func main() {
	array := []int{12, 3, 1, 2, -6, 5, -8, 6}
	targetSum := 0

	result := findTriplets(array, targetSum)

	fmt.Println(result)
}
