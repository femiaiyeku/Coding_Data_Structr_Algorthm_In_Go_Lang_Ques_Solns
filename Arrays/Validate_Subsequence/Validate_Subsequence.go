/*



Given two non•empty arrays of integers, write a function that determines whether the second array is a subsequence of the first one.
A subsequence of an array is a set of numbers that aren't necessarily adjacent in the array but that are in the same order as they appear in the array. For instance, the numbers
[l, 3, 4] form a subsequence of the array [l, 2, 3, 4] , and so do the numbers (2, 4] Note that a single number in an array and the array itself are both valid
subsequences of the array

Sample Input:
array = [5, 1, 22, 25, 6, -1, 8, 10]
sequence = [1, 6, -1, 10]

Sample Output:
true



This function works by initializing the subsequence index. The function then iterates over the first array. For each element in the first array, the function checks if it matches the current element in the subsequence. If it does, the function increments the subsequence index.

If the subsequence index is equal to the length of the subsequence, then the subsequence is present in the first array, and the function returns true. Otherwise, the function continues iterating over the first array.

Once the function has finished iterating over the first array, it checks if the subsequence index is equal to the length of the subsequence. If it is, then the subsequence is present in the first array, and the function returns true. Otherwise, the function returns false.

*/

package main

import (
	"fmt"
)

func isSubsequence(array []int, sequence []int) bool {
	// Initialize the subsequence index.
	subsequenceIndex := 0

	// Iterate over the first array.
	for i := range array {
		// If the current element in the first array matches the current element in the subsequence, increment the subsequence index.
		if array[i] == sequence[subsequenceIndex] {
			subsequenceIndex++
		}

		// If the subsequence index is equal to the length of the subsequence, then the subsequence is present in the first array.
		if subsequenceIndex == len(sequence) {
			return true
		}
	}

	// If the subsequence index is not equal to the length of the subsequence, then the subsequence is not present in the first array.
	return false
}

func main() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}

	isSubsequence := isSubsequence(array, sequence)

	fmt.Println(isSubsequence)
}
