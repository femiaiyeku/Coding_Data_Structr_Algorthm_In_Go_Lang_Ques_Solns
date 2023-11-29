/*


Write a function that takes in an array of integers and returns a sorted version
that array. Use the Selection Sort algorithm to sort the array.
If you're unfamiliar with Selection Sort, we recommend watching the Conceptual
Overview section of this question's video explanation before starting to code.

Sample Input:
array = [8, 5, 2, 9, 5, 6, 3]

Sample Output:
[2, 3, 5, 5, 6, 8, 9]



This code defines a function selectionSort that takes an array of integers as input and sorts the array in ascending order using the Selection Sort algorithm. The Selection Sort algorithm works by repeatedly finding the minimum element in the unsorted portion of the array and swapping it with the first element in that portion.

The main function creates an array of integers and calls the selectionSort function to sort the array. It then prints the sorted array.







*/

package main

import (
	"fmt"
)

func selectionSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		minIndex := i

		for j := i + 1; j < len(array); j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}

		temp := array[i]
		array[i] = array[minIndex]
		array[minIndex] = temp
	}
}

func main() {
	array := []int{8, 5, 2, 9, 5, 6, 3}

	selectionSort(array)

	fmt.Println("Sorted array:", array)
}
