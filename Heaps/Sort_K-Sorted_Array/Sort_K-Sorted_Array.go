/*


Write a function that takes in a non-negative integer k and a k-sorted array of integers and returns the sorted version of the array.
 Your function can either sort the array in place or create an entirely new array.
A k-sorted array is a partially sorted array in which all elements are at most k positions away from their sorted position.
 For example, the array (3, 1, 2, 2] is k-sorted with k = 3 , because each element in the array is at most 3 positions away from its sorted position.
Note that you're expected to come up with an algorithm that can sort the k-sorted array faster than in O{nlog{n)) time.


Sample Input:
array = [3, 2, 1, 5, 4, 7, 6, 5]
k = 3

Sample Output:
[1, 2, 3, 4, 5, 5, 6, 7]


Hints
Hint 1

What does the k parameter tell you? How can you use it to come up with an algorithm that runs in O(nlog(k))?

Hint 2

Since the input array is k-sorted, try repeatedly sorting k elements at a time and inserting the minimum element of all those k elements into its final sorted position in the array.

Hint 3

What auxiliary data structure would be helpful to quickly determine the minimum element of k elements?

Hint 4

As you iterate through the array, use a min-heap to keep track of the most recent k elements. At each iteration, remove the minimum value from the heap, insert it into its final sorted position in the array, and add the current element in the array to the heap. Continue this process until the heap is empty.




*/

package main

import (
	"container/heap"
	"fmt"
)

type minHeap struct {
	array []int
	size  int
	k     int // k-sorted
}

func newMinHeap(array []int, k int) *minHeap {
	h := &minHeap{
		array: array,
		size:  len(array),
		k:     k,
	}

	heap.Init(h)
	return h
}

func (h *minHeap) Len() int {
	return h.size
}

func (h *minHeap) Less(i, j int) bool {
	return h.array[i] < h.array[j]
}

func (h *minHeap) Swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func (h *minHeap) Pop() interface{} {
	old := h.array[:h.size]
	n := len(old)
	h.array = old[:n-1]
	h.size = n - 1

	for i := 0; i <= h.size-h.k-1; i++ {
		heap.SiftDown(h, i)
	}

	return old[n-1]
}

func (h *minHeap) Push(x interface{}) {
	if h.size == len(h.array) {
		panic("Heap is full")
	}

	h.array[h.size] = x.(int)
	h.size++

	heap.SiftUp(h, h.size-1)
}

func kSort(array []int, k int) []int {
	h := newMinHeap(array, k)

	for len(h.array) > 0 {
		smallest := heap.Pop(h).(int)
		array[len(array)-1] = smallest
		array = array[:len(array)-1]
	}

	return array
}

func main() {
	array := []int{3, 2, 1, 5, 4, 7, 6, 5}
	k := 3

	sortedArray := kSort(array, k)
	fmt.Println(sortedArray) // [1, 2, 3, 4, 5, 5, 6, 7]
}
