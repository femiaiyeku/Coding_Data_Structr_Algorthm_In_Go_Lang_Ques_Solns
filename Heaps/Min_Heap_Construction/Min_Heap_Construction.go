/*

Implement a MinHeap class that supports:


• 	Building a Min Heap from ar input array of integers.
• 	Inserting integers in the heaJ.
• 	Removing the heap's minimum/ root value.
• 	Peeking at the heap's minimum/ root value.
• 	Sifting integers up and dowr, the heap, which is to be used when inserting and removing values.

Note that the heap should be repr􀀌sented in the form of an array.
If you're unfamiliar with Min Heaps, we recommend watching the Conceptual Overview section of this question's video explanation before starting to code.

Sample Usage

// All operations below are performed sequentially.

MinHeap(array)
// - buildHeap(array):
//   - takes in an array of integers and builds a min heap from them.
//   - returns the input array modified in place.
buildHeap([9, 4, 7, 1, -2, 6, 5])
//   -> [-2, 1, 5, 9, 4, 6, 7]
// - insert(8):
//   - inserts the value 8 into the heap.
insert(8)
// - remove():
//   - removes the minimum value (that is, the root value) from the heap.
remove()
//   -> -2
// - peek():
//   - returns the minimum value from the heap without removing it.
peek()
//   -> 1
// - insert(-3):
//   - inserts the value -3 into the heap.
insert(-3)
//   -> [-3, 1, 5, 9, 4, 6, 7, 8]

Hints & Tips
• 	The heapify procedure is a common procedure used in heaps.
• 	You can use the siftDown procedure to implement the remove procedure.
• 	You can use the siftUp procedure to implement the insert procedure.
• 	You can use the siftDown procedure to implement the buildHeap procedure.
• 	You can use the siftUp procedure to implement the insert procedure.




Hints
Hint 1
For the buildHeap(), remove(), and insert() methods of the Heap, you will need to use the siftDown() and siftUp() methods. These two methods should essentially allow you to take any node in the heap and move it either down or up in the heap until it's in its final, appropriate position. This can be done by comparing the node in question to its child nodes in the case of siftDown() or to its parent node in the case of siftUp().
Hint 2
In an array-based Heap, you can easily access a node's children nodes and parent node by using the nodes' indices. If a node is located at index i, then its children nodes are located at indices 2 * i + 1 and 2 * i + 2, and its parent node is located at index Math.floor((i - 1) / 2).
Hint 3
To implement the buildHeap() method, you can either sift every node in the input array down to its final, correct position, or you can sift every node in the input array up to its final, correct position. What are the runtime implications of both approaches? Which methods (siftDown() or siftUp()) will insert() and remove() utilize? What about peek()?






*/

package main

import "fmt"

type MinHeap []int

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	*h = old[:n-1]
	return old[n-1]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
	heap.Init(h)
}

func buildHeap(array []int) {
	h := &MinHeap{array}
	for i := len(array) / 2; i >= 0; i-- {
		heap.SiftDown(h, i)
	}
}

func insert(heap *MinHeap, value int) {
	heap.Push(value)
	heap.SiftUp(len(heap) - 1)
}

func remove(heap *MinHeap) int {
	if heap.Len() == 0 {
		panic("Heap is empty")
	}
	value := heap.Pop()
	heap.SiftDown(0)
	return value
}

func peek(heap *MinHeap) int {
	if heap.Len() == 0 {
		panic("Heap is empty")
	}
	return (*heap)[0]
}

func siftDown(heap *MinHeap, i int) {
	for {
		leftChild := 2*i + 1
		rightChild := 2*i + 2

		if leftChild >= heap.Len() {
			break
		}

		minChild := leftChild

		if rightChild < heap.Len() && (*heap)[rightChild] < (*heap)[leftChild] {
			minChild = rightChild
		}

		if (*heap)[i] <= (*heap)[minChild] {
			break
		}

		heap.Swap(i, minChild)
		i = minChild
	}
}

func siftUp(heap *MinHeap, i int) {
	for i > 0 {
		parent := (i - 1) / 2

		if (*heap)[i] >= (*heap)[parent] {
			break
		}

		heap.Swap(i, parent)
		i = parent
	}
}

func main() {
	array := []int{9, 4, 7, 1, -2, 6, 5}
	buildHeap(array)
	fmt.Println(array) // [-2, 1, 5, 9, 4, 6, 7]

	insert(heap, 8)
	fmt.Println(heap) // [-2, 1, 5, 9, 4, 6, 7, 8]

	value := remove(heap)
	fmt.Println(value) // -2

	value = peek(heap)
	fmt.Println(value) // 1

	insert(heap, -3)
	fmt.Println(heap) // [-3, 1, 5, 9, 4, 6, 7, 8]
}
