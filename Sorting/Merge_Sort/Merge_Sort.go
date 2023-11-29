/*


Write a function that takes in an array of integers and returns a sorted version of that array.
Use the Merge Sort algorithm to sort the array.
If you're unfamiliar with Merge Sort,
 we recommend watching the Conceptual Overview section of this question's video explanation before starting to code.

Sample Input
array = [8, 5, 2, 9, 5, 6, 3]

Sample Output

[2, 3, 5, 5, 6, 8, 9]


Hints
Hint 1

Merge Sort works by cutting an array in two halves, respectively sorting those two halves by performing some special logic, and then merging the two newly-sorted halves into one sorted array. The respective sorting of the two halves is done by reapplying the Merge Sort algorithm / logic on each half until single-element halves are obtained; these single-element arrays are sorted by nature and can very easily be merged back together.

Hint 2

Divide the input array in two halves by finding the middle-most index in the array and slicing the two halves around that index. Then, recursively apply Merge Sort to each half, and finally merge them into one single, sorted array by iterating through their values and progressively adding them to the new array in ascending order.

Hint 3

Your implementation of Merge Sort almost certainly uses a lot of auxiliary space and likely does not sort the input array in place. What is the space complexity of your algorithm? Can you implement a version of the algorithm using only one additional array of the same length as the input array, and can this version sort the input array in place?




This code defines a function mergeSort that takes an array of integers as input and sorts the array in ascending order using the Merge Sort algorithm. The Merge Sort algorithm works by recursively dividing the array into smaller subarrays, sorting the subarrays, and then merging the sorted subarrays back together.

The main function creates an array of integers and calls the mergeSort function to sort the array. It then prints the sorted array.



*/

package main

import (
	"fmt"
)

func mergeSort(array []int) {
	if len(array) <= 1 {
		return
	}

	mid := len(array) / 2
	left := make([]int, mid)
	right := make([]int, len(array)-mid)

	copy(left, array[:mid])
	copy(right, array[mid:])

	mergeSort(left)
	mergeSort(right)

	merge(array, left, right)
}

func merge(array []int, left []int, right []int) {
	i := 0
	j := 0
	k := 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			array[k] = left[i]
			i++
		} else {
			array[k] = right[j]
			j++
		}

		k++
	}

	for i < len(left) {
		array[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		array[k] = right[j]
		j++
		k++
	}
}

func main() {
	array := []int{8, 5, 2, 9, 5, 6, 3}

	mergeSort(array)

	fmt.Println("Sorted array:", array)
}
