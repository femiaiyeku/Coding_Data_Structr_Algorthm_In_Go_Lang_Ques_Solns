/*


Write a function that takes in an array of integers and returns a sorted version of that array.
Use the Bubble Sort algorithm to sort the array.
If you're unfamiliar with Bubble Sort,
we recommend watching the Conceptual Overview section of this question's video explanation before starting to code.


Sample Input:

array = [8, 5, 2, 9, 5, 6, 3]

Sample Output:

[2, 3, 5, 5, 6, 8, 9]



This code defines a function bubbleSort that takes an array of integers as input and sorts the array in ascending order using the Bubble Sort algorithm. The Bubble Sort algorithm works by repeatedly comparing adjacent elements in the array and swapping them if they are in the wrong order. This process is repeated until the array is sorted.

The main function creates an array of integers and calls the bubbleSort function to sort the array. It then prints the sorted array.




*/

package main

import (
	"fmt"
)

func bubbleSort(array []int) {
	isSorted := false

	for !isSorted {
		isSorted = true

		for i := 0; i < len(array)-1; i++ {
			if array[i] > array[i+1] {
				temp := array[i]
				array[i] = array[i+1]
				array[i+1] = temp

				isSorted = false
			}
		}
	}
}

func main() {
	array := []int{8, 5, 2, 9, 5, 6, 3}

	bubbleSort(array)

	fmt.Println("Sorted array:", array)
}
