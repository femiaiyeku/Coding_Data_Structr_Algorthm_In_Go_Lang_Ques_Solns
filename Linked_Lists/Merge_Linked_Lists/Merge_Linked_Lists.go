/*

Write a function that takes in the heads of two Singly Linked Lists that are in sorted order, respectively. The
function should merge the lists in place (.e., It shouldn't create a brand new list) and return the head of the
merged list; the merged list should be in sorted order.
Each Linked List node has an integer value as well as a next node pointing to the next node in the list or
to None null if it's the tail of the list.
You can assume that the input linked lists will always have at least one node; in other words, the heads will never
be None / null.

Sample Input:
headOne = 2 -> 6 -> 7 -> 8  // the head node with value 2

headTwo = 1 -> 3 -> 4 -> 5 -> 9 -> 10  // the head node with value 1

Sample Output:

1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 10  // the new head node with value 1









*/

package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func mergeSortedLists(headOne *Node, headTwo *Node) *Node {
	if headOne == nil {
		return headTwo
	}

	if headTwo == nil {
		return headOne
	}

	var mergedHead *Node
	var current *Node

	if headOne.value <= headTwo.value {
		mergedHead = headOne
		current = headOne
		headOne = headOne.next
	} else {
		mergedHead = headTwo
		current = headTwo
		headTwo = headTwo.next
	}

	for headOne != nil && headTwo != nil {
		if headOne.value <= headTwo.value {
			current.next = headOne
			headOne = headOne.next
		} else {
			current.next = headTwo
			headTwo = headTwo.next
		}

		current = current.next
	}

	if headOne != nil {
		current.next = headOne
	} else {
		current.next = headTwo
	}

	return mergedHead
}

func main() {
	headOne := &Node{value: 2, next: &Node{value: 6, next: &Node{value: 7, next: &Node{value: 8}}}}
	headTwo := &Node{value: 1, next: &Node{value: 3, next: &Node{value: 4, next: &Node{value: 5, next: &Node{value: 9, next: &Node{value: 10}}}}}}

	mergedHead := mergeSortedLists(headOne, headTwo)
	for current := mergedHead; current != nil; current = current.next {
		fmt.Println(current.value)
	}
}
