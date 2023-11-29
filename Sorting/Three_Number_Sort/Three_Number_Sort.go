/*


You're given an array of Integers and another array of three distinct Integers. The first array is guaranteed
only contain Integers that are in the second array, and the second array represents a desired order for the
integers in the first array. For example, a second array of [x, y, z1 represents a desired order of
[X, X, X, Y YY, 2, 2, 2] In the first array.
Write a function that sorts the first array according to the desired order in the second array.
The function should perform this in place (i.e., It should mutate the input array), and it shouldn't use any
auxiliary space (1.e., it should run with constant space: 0(1) space).
Note that the desired order won't necessarily be ascending or descending and that the first array won't
necessarily contain all three integers found in the second array-It might only contain one or two.

Sample Input
array = [1, 0, 0, -1, -1, 0, 1, 1]
order = [0, 1, -1]

Sample Output
[0, 0, 0, 1, 1, 1, -1, -1]





This code first creates a map counts to store the count of occurrences for each element in the array. Then, it iterates through the order array and replaces the elements in the array with the corresponding elements from the order array, decrementing the count for each element as it's placed. This ensures that the elements are placed in the desired order while maintaining the original counts.





*/

package main

import (
	"fmt"
)

func sortWithOrder(array []int, order []int) {
	counts := make(map[int]int)
	for _, element := range array {
		counts[element]++
	}

	index := 0
	for _, element := range order {
		for counts[element] > 0 {
			array[index] = element
			index++
			counts[element]--
		}
	}
}

func main() {
	array := []int{1, 0, 0, -1, -1, 0, 1, 1}
	order := []int{0, 1, -1}

	sortWithOrder(array, order)

	fmt.Println("Sorted array:", array)
}
