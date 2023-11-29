/*



Write a function that takes in an array of integers and returns a sorted version of
that array. Use the Insertion Sort algorithm to sort the array.
If you're unfamiliar with Insertion Sort, we recommend watching the Conceptual
Overview section of this question's video explanation before starting to code

Sample Input:
array = [8, 5, 2, 9, 5, 6, 3]

Sample Output:
[2, 3, 5, 5, 6, 8, 9]




This code defines a function insertionSort that takes an array of integers as input and sorts the array in ascending order using the Insertion Sort algorithm. The Insertion Sort algorithm works by repeatedly inserting elements into their correct positions in a partially sorted array.

The main function creates an array of integers and calls the insertionSort function to sort the array. It then prints the sorted array.





*/

package main

import (
	"fmt"
)

func insertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		current := array[i]
		j := i - 1

		for j >= 0 && array[j] > current {
			array[j+1] = array[j]
			j--
		}

		array[j+1] = current
	}
}

func main() {
	array := []int{8, 5, 2, 9, 5, 6, 3}

	insertionSort(array)

	fmt.Println("Sorted array:", array)
}
