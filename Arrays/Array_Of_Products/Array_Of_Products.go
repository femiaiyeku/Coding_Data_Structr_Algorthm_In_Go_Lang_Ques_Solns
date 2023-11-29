/*

Write a function that takes in a non-empty array of integers and returns an array of the same length,
where each element in the output array is equal to the product of every other number in the input array.
In other words, the value at output [ i] is equal to the product of every number in the input array other than i nput[i]
Note that you're expected to solve this problem without using division.

Sample Input
array = [5, 1, 4, 2]

Sample Output
[8, 40, 10, 20]
// 8 is equal to 1 x 4 x 2
// 40 is equal to 5 x 4 x 2
// 10 is equal to 5 x 1 x 2
// 20 is equal to 5 x 1 x 4


To use the function, simply pass in the array of integers as an argument. The function will return an array of the same length, where each element in the output array is equal to the product of every other number in the input array.

For example, using the sample input provided in the question, we would get the following output:




*/

package main

func productExceptSelf(array []int) []int {

	// Create two arrays to store the prefix product and the suffix product of the input array.
	prefixProduct := make([]int, len(array))
	suffixProduct := make([]int, len(array))

	// Calculate the prefix product.
	prefixProduct[0] = 1
	for i := 1; i < len(array); i++ {
		prefixProduct[i] = prefixProduct[i-1] * array[i-1]
	}

	// Calculate the suffix product.
	suffixProduct[len(array)-1] = 1
	for i := len(array) - 2; i >= 0; i-- {
		suffixProduct[i] = suffixProduct[i+1] * array[i+1]
	}

	// Create the output array.
	outputArray := make([]int, len(array))
	for i := 0; i < len(array); i++ {
		outputArray[i] = prefixProduct[i] * suffixProduct[i]
	}

	// Return the output array.
	return outputArray
}
