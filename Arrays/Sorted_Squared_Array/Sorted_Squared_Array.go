/*

Write a function that takes 1n a non empty array of integers that are sorted 1n ascending order and
returns a new array of the same length with the squares of the original integers also sorted 1n ascending order

Sample Input
array = [1, 2, 3, 5, 6, 8, 9]

Sample Output

[1, 4, 9, 25, 36, 64, 81]




This function works by first initializing a new array. The function then iterates over the original array and squares each element. The function then sorts the new array in ascending order. Finally, the function returns the new array.


*/

package main

import (
	"fmt"
	"sort"
)

func squaresOfSortedArray(array []int) []int {
	// Initialize the new array.
	result := make([]int, len(array))

	// Iterate over the original array and square each element.
	for i, element := range array {
		result[i] = element * element
	}

	// Sort the new array in ascending order.
	sort.Ints(result)

	// Return the new array.
	return result
}

func main() {
	array := []int{1, 2, 3, 5, 6, 8, 9}

	result := squaresOfSortedArray(array)

	fmt.Println(result)
}
