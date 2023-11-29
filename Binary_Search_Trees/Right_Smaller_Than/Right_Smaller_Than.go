/*


Write a function that takes in an array of integers and returns an array of the same length,
where each element in the output array corresponds to the number of integers in the input array
that are to the right of the relevant index and that are strictly smaller than the integer at that index.

In other words, the value at output [ i] represents the number of integers that are to the right of smallerthan input[i] .

More formally, the value at output[i] is equivalent to the number of elements that are to the right of i and that are strictly smaller than input[i] .

Sample Input

array = [8, 5, 11, -1, 3, 4, 2]

Sample Output

[5, 4, 4, 0, 1, 1, 0]

// There are 5 integers smaller than 8 to the right of it.
// There are 4 integers smaller than 5 to the right of it.
// There are 4 integers smaller than 11 to the right of it.
// ...



Hints
Hint 1

The naive solution to this problem involves a simple pair of nested for loops and runs in O(n^2) time, where n is the length of the input array. This problem doesn't seem like it can be solved in linear time, so what time complexity could we realistically achieve if we somehow optimized our algorithm?

Hint 2

The only better time complexity than the O(n^2) one of our naive solution that we could realistically achieve would be an O(nlog(n)) complexity. What data structure has log(n) operations, and how could it help for this problem?

Hint 3

A Binary Search Tree supports log(n) insertions and has the relevant property of every left-subtree-node having a smaller value than a given node's value; can we construct a BST from the input array in such a way that it leads us to the result array that we're looking for?

Hint 4

Construct a BST by inserting the input array's integers one by one, in reverse order (from right to left). At each insertion, once a new BST node is positioned in the BST, the number of nodes in its parent node's left subtree (plus the parent node itself, if its value is smaller than the inserted node's value) is the number of "right-smaller-than" elements for the element being inserted.

Hint 5

Going off of Hint #4, you can construct a special type of BST that stores the size of every node's left subtree. This value can then be used to obtain the right-smaller-than numbers for every element in the array.




*/

package main

import (
	"fmt"
	"sort"
)

func countSmallerToTheRight(array []int) []int {
	// Create a copy of the input array.
	sortedArray := make([]int, len(array))
	copy(sortedArray, array)

	// Sort the copied array in ascending order.
	sort.Ints(sortedArray)

	// Create an array to store the number of smaller integers to the right of each element in the input array.
	countSmallerToTheRight := make([]int, len(array))

	// Iterate over the input array from right to left.
	for i := len(array) - 1; i >= 0; i-- {
		// Find the index of the current element in the sorted array.
		indexInSortedArray := sort.SearchInts(sortedArray, array[i])

		// The number of smaller integers to the right of the current element is the number of elements to the right of the current element in the sorted array.
		countSmallerToTheRight[i] = len(sortedArray) - indexInSortedArray - 1
	}

	return countSmallerToTheRight
}

func main() {
	// Create a new array of integers.
	array := []int{8, 5, 11, -1, 3, 4, 2}

	// Count the number of smaller integers to the right of each element in the array.
	countSmallerToTheRight := countSmallerToTheRight(array)

	// Print the output array.
	fmt.Println(countSmallerToTheRight)
}
