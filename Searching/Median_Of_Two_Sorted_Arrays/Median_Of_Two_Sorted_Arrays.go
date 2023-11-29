/*



You're given two sorted arrays of integers arrayOne and arrayTwo. Write a
function that returns the median of these arrays.
You can assume both arrays have at least one value. However, they could be of
different lengths. If the median is a decimal value, it should not be rounded
(beyond the inevitable rounding of floating point math).

Sample Input:
arrayOne = [-1, 5, 10, 20, 28, 3]
arrayTwo = [26, 134, 135, 15, 17]

Sample Output:
16



This code defines a function findMedian that takes two sorted arrays of integers as input and returns the median of the combined arrays. The function first checks if the total length of the two arrays is even or odd. If it is even, it returns the average of the two middle elements of the combined arrays. If it is odd, it returns the middle element of the combined arrays.

The main function creates two arrays of integers and calls the findMedian function to find the median of the two arrays. It then prints the result.





*/

package main

import (
	"fmt"
)

func findMedian(arrayOne []int, arrayTwo []int) float64 {
	totalLength := len(arrayOne) + len(arrayTwo)
	isEven := totalLength%2 == 0

	if isEven {
		mid1 := (totalLength - 1) / 2
		mid2 := mid1 + 1

		return (float64(arrayOne[mid1]) + float64(arrayTwo[mid2])) / 2.0
	} else {
		mid := (totalLength - 1) / 2

		if len(arrayOne) > len(arrayTwo) {
			return float64(arrayOne[mid])
		} else {
			return float64(arrayTwo[mid])
		}
	}
}

func main() {
	arrayOne := []int{-1, 5, 10, 20, 28, 3}
	arrayTwo := []int{26, 134, 135, 15, 17}

	median := findMedian(arrayOne, arrayTwo)
	fmt.Println("Median of the two arrays:", median)
}
