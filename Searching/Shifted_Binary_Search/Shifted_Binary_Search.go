/*

Write a function that takes 1n a sorted array of distinct integers as well as a target integer.
 The caveat is that the integers in the array have been shifted by some amount; 1n other words,
 they've been moved to the left or to the right by one or more pos1t1ons. For example, [l, 2, 3, 4] might have turned into [3, 4, 1, 2]
The function should use a variation of the Binary Search algorithm to determine 1f the target integer 1s contained 1n the array and should return its index if 1t is,
otherwise -1 . If you·re unfamiliar with Binary Search,
we recommend watching the Conceptual Overview section of the Binary Search question·s video explanation before starting to code.

Sample Input
array = [45, 61, 71, 72, 73, 0, 1, 21, 33, 37]
target = 33

Sample Output
8

Example Explanation
The target integer 33 1s located at 1ndex 8 1n the array.
This approach finds the correct indices of shifted sorted array and then applies naive bunary search. First ,
find the position of smallest element in shifted array which will be the number of steps the array has shifted.
And then using modulas operator map "mid" to it's position in given shifted array.

Time complexity: O(log(n))
Space complexity: 0(1)




Hints
Hint 1

The Binary Search algorithm involves a left pointer and a right pointer and using those pointers to find the middle number in an array. Unlike with a normal sorted array however, you cannot simply find the middle number of the array and compare it to the target here, because the shift could lead you in the wrong direction. Instead, realize that whenever you find the middle number in the array, the following two scenarios are possible (assuming the middle number is not equal to the target number, in which case you're done): either the left-pointer number is smaller than or equal to the middle number, or it is bigger. Figure out a way to eliminate half of the array depending on the scenario.

Hint 2

In the scenario where the left-pointer number is smaller than or equal to the middle number, two other scenarios can arise: either the target number is smaller than the middle number and greater than or equal to the left-pointer number, or it's not. In the first scenario, the right half of the array can be eliminated; in the second scenario, the left half can be eliminated. Figure out the scenarios that can arise if the left-pointer number is greater than the middle number and apply whatever logic you come up with recursively until you find the target number or until you run out of numbers in the array.

Hint 3

Can you implement this algorithm iteratively? Are there any advantages to doing so?


This code first defines a function findMin that finds the index of the smallest element in the array. This element represents the number of positions that the array has been shifted. The function uses a modified binary search algorithm to find the smallest element.

Next, it defines a function findTarget that takes the shifted array and the target integer as input and returns the index of the target integer if it is contained in the array, or -1 if it is not. The function first calculates the adjusted index of the middle element of the array, taking into account the shift. It then compares the adjusted element to the target integer and updates the search range accordingly.

Finally, it calls the findMin function to find the shifted array and then calls the findTarget function to find the index of the target integer. It then prints the result.



*/

package main

import (
	"fmt"
)

func findMin(array []int) int {
	low := 0
	high := len(array) - 1

	for low < high {
		mid := (low + high) / 2

		if array[mid] < array[high] {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

func findTarget(array []int, target int) int {
	minIndex := findMin(array)

	low := 0
	high := len(array) - 1

	for low <= high {
		mid := (low + high) / 2

		adjustedMid := (mid + minIndex) % len(array)
		if array[adjustedMid] == target {
			return adjustedMid
		} else if array[adjustedMid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	array := []int{45, 61, 71, 72, 73, 0, 1, 21, 33, 37}
	target := 33

	targetIndex := findTarget(array, target)

	if targetIndex != -1 {
		fmt.Printf("Target %d found at index %d\n", target, targetIndex)
	} else {
		fmt.Println("Target not found")
	}
}
