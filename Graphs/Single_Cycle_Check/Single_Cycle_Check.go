/*


You're given an array of integers where each integer represents a Jump of its value in the array.
For instance, the integer 2 represents a jump of two indices forward in the array; the integer -3 represents a jump of three indices backward in the array.
If a Jump spills past the array's bounds, it wraps over to the other side. For instance, a jump of -1 at index 0 brings us to the last index in the array.
Similarly, a jump of 1 at the last index in the array brings us to index e .
Write a function that returns a boolean representing whether the jumps in the array form a single cycle.
A single cycle occurs if, starting at any index in the array and following the jumps,
every element in the array is visited exactly once before landing back on the starting index.

Sample Input:
array = [2, 3, 1, -4, -4, 2]

Sample Output:

True




Hints
Hint 1
In order to check if the input array has a single cycle, you have to jump through all of the elements in the array. Could you keep a counter, jump through elements in the array, and stop once you've jumped through as many elements as are contained in the array?
Hint 2
Assume the input array has length n. If you start at index 0 and jump through n elements, what are the simplest conditions that you can check to see if the array doesn't have a single cycle?
Hint 3
Given Hint #2, there are 2 conditions that need to be met for the input array to have a single cycle. First, the starting element (in this case, the element at index 0) cannot be jumped through more than once, at the very beginning, so long as you haven't jumped through all of the other elements in the array. Second, the (n + 1)th element you jump through, where n is the length of the array, must be the first element you visited: the element at index 0 in this case.





*/

package main

import "fmt"

func hasSingleCycle(array []int) bool {
	// Track visited indices to avoid infinite loops
	visited := make(map[int]bool)

	// Start at any index in the array
	startIndex := 0

	// Keep track of the current index and the next index
	currentIndex := startIndex
	nextIndex := (currentIndex + array[currentIndex]) % len(array)

	// Continue traversing the array until we reach a cycle or encounter an infinite loop
	for currentIndex != nextIndex || !visited[currentIndex] {
		// Mark the current index as visited
		visited[currentIndex] = true

		// Update the current and next indices
		currentIndex = nextIndex
		nextIndex = (currentIndex + array[currentIndex]) % len(array)
	}

	// Check if we reached a cycle or an infinite loop
	if currentIndex == startIndex {
		// Cycle detected
		return true
	} else {
		// Infinite loop detected
		return false
	}
}

func main() {
	array := []int{2, 3, 1, -4, -4, 2}

	hasCycle := hasSingleCycle(array)
	fmt.Println(hasCycle)
}
