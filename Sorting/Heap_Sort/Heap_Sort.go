/*


Write a function that takes in an array of integers and returns a sorted version of that array.
Use the Heap Sort algorithm to sort the array.
If you're unfamiliar with Heap Sort,
we recommend watching the Conceptual Overview section of this question's video explanation before starting to code.

Sample Input
array = [8, 5, 2, 9, 5, 6, 3]

Sample Output
[2, 3, 5, 5, 6, 8, 9]



Hints
Hint 1

Divide the input array into two subarrays in place. The second subarray should be sorted at all times and should start with a length of 0, while the first subarray should be transformed into a max (or min) heap and should satisfy the heap property at all times.

Hint 2

Note that the largest (or smallest) value of the heap should be at the very beginning of the newly-built heap. Start by swapping this value with the last value in the heap; the largest (or smallest) value in the array should now be in its correct position in the sorted subarray, which should now have a length of 1; the heap should now be one element smaller, with its first element out of place. Apply the "sift down" method of the heap to re-position this out-of-place value.

Hint 3

Repeat the step mentioned in Hint #2 until the heap is left with only one value, at which point the entire array should be sorted.





This code defines a custom type IntHeap that implements the heap.Interface interface. This allows us to use the heap.Init, heap.Push, and heap.Pop functions with our custom heap type.

The heapSort function first creates an IntHeap from the input array and then initializes it using the heap.Init function. Next, it repeatedly pops elements from the heap and adds them to the end of the array, effectively sorting the array in ascending order.

The main function creates an array of integers and calls the heapSort function to sort the array. It then prints the sorted array.








*/

package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func heapSort(array []int) {
	h := &IntHeap(array)
	heap.Init(h)

	for len(*h) > 0 {
		array[len(array)-1] = heap.Pop(h).(int)
	}
}

func main() {
	array := []int{8, 5, 2, 9, 5, 6, 3}

	heapSort(array)

	fmt.Println("Sorted array:", array)
}
