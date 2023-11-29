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





Hint 1

Try treating every building as a pillar of a rectangle that can be created with the height of the building in question.

Hint 2

The brute-force approach to solve this problem involves treating every building as a part of a potential rectangle to be created. As you loop through all the buildings, simply expand to the left and right of the current building, and determine the width of the longest rectangle that you can create that has a height of the current building. Calculate the area of this longest rectangle, and update a variable to store the area of the largest rectangle that you've found so far. This approach has a time complexity of O(n^2); can you do better?

Hint 3

There's a way to solve this problem in linear (O(n)) time by using a stack. When should you push and pop buildings on and off the stack if you were to loop through the buildings from left to right? Try to think of each building on the stack as a pillar of a potential rectangle.

Hint 4

The stack mentioned in Hint #3 will be used to determine the length of a rectangle that has the height of a building that is currently on top of the stack. Loop through all the buildings, and at each building, compare its height to the height of the building on top of the stack. If the current building's height is smaller than or the same as the height of the building on top of the stack, pop the building off the stack. When you pop the building off the stack, you've determined the rightmost position (your current position) of a rectangle of that height (the height of the building you popped) that uses that building. Then, to determine the leftmost position of that rectangle, you look at the next building on top of the stack. This is the index of the closest building to the left that has a smaller height than that of the building that was just popped off. Now, you can calculate the area of the rectangle that uses this building and update a variable to store the max area. Continue popping buildings off the stack at each iteration until the current building is taller than the one on top of the stack, and don't forget to push each building on top of the stack at each iteration. See the Conceptual Overview section of this question's video explanation for a more in-depth explanation.




This updated code maintains constant space complexity by avoiding the creation of additional data structures like arrays or maps. Instead, it utilizes the orderIndex and elementCount variables to track the position in the desired order and the count of occurrences for each element. This ensures that the sorting is performed in place without requiring any auxiliary space.

*/

package main

import (
	"fmt"
)

func sortWithOrder(array []int, order []int) {
	// Initialize variables for tracking indices and element counts
	orderIndex := 0
	elementCount := 0

	// Iterate through the array
	for i := range array {
		// Check if the current element matches the desired order
		if array[i] == order[orderIndex] {
			// Count the occurrences of the current element
			elementCount++

			// Move to the next element in the desired order if the count is exhausted
			if elementCount == counts[order[orderIndex]] {
				orderIndex++
				elementCount = 0
			}
		}
	}

	// Iterate through the array again and place the elements in the desired order
	for i := range array {
		// Check if the current element matches the desired order
		if array[i] == order[orderIndex] {
			// Swap the current element with the corresponding element in the desired order
			array[i], array[orderIndex] = array[orderIndex], array[i]

			// Move to the next element in the desired order if the count is exhausted
			elementCount++
			if elementCount == counts[order[orderIndex]] {
				orderIndex++
				elementCount = 0
			}
		}
	}
}

func main() {
	array := []int{1, 0, 0, -1, -1, 0, 1, 1}
	order := []int{0, 1, -1}

	counts := make(map[int]int)
	for _, element := range array {
		counts[element]++
	}

	sortWithOrder(array, order)

	fmt.Println("Sorted array:", array)
}
