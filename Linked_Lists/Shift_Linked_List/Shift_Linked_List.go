/*

Write a function that takes in the head of a Singly Linked List and an integer k , shifts the list in place (i.e., doesn't create a brand new list) by k positions, and returns its new head.
Shifting a Linked List means moving its nodes forward or backward and wrapping them around the list where appropriate. For example, shifting a Linked List forward by one position would make its tail become the new head of the linked list.
Whether nodes are moved forward or backward is determined by whether k is positive or negative.
Each Linked list node has an integer value as well as a next node pointing to the next node in the list or to None I null if it's the tail of the list.
You can assume that the input Linked List will always have at least one node; in other words, the head will never be None / null.

Sample Input:

head = 0 -> 1 -> 2 -> 3 -> 4 -> 5 // the head node with value 0
k = 2

Sample Output:

4 -> 5 -> 0 -> 1 -> 2 -> 3 // the new head node with value 4









*/

package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func shiftLinkedList(head *Node, k int) *Node {
	if k == 0 {
		return head
	}

	listLength := getListLength(head)
	k = k % listLength

	if k < 0 {
		k += listLength
	}

	current := head
	prev := &Node{}

	for i := 0; i < k; i++ {
		prev = current
		current = current.next
	}

	prev.next = nil

	newHead := current
	current = current.next

	for current != nil {
		newHead.next = current
		current = current.next
	}

	return newHead
}

func getListLength(head *Node) int {
	length := 0

	for current := head; current != nil; current = current.next {
		length++
	}

	return length
}

func main() {
	head := &Node{value: 0, next: &Node{value: 1, next: &Node{value: 2, next: &Node{value: 3, next: &Node{value: 4, next: &Node{value: 5}}}}}}
	k := 2

	newHead := shiftLinkedList(head, k)

	for current := newHead; current != nil; current = current.next {
		fmt.Print(current.value, " -> ")
	}
}
