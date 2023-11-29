/*

Write a function that takes in the head of a Singly Linked List, reverses the list in place (i.e., doesn't create a brand
new list), and returns its new head.
Each LinkedList node has an integer value as well as a next] node pointing to the next node in the list or
to None / null if it's the tail of the list.
You can assume that the input Linked List will always have at least one node; in other words, the head will never
be None / null.

Sample Input:

head = 0 -> 1 -> 2 -> 3 -> 4 -> 5 // the head node with value 0

Sample Output:

5 -> 4 -> 3 -> 2 -> 1 -> 0 // the new head node with value 5






*/

package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func reverseLinkedList(head *Node) *Node {
	var prev *Node
	current := head

	for current != nil {
		next := current.next

		current.next = prev
		prev = current

		current = next
	}

	return prev
}

func main() {
	head := &Node{value: 0, next: &Node{value: 1, next: &Node{value: 2, next: &Node{value: 3, next: &Node{value: 4, next: &Node{value: 5}}}}}}

	newHead := reverseLinkedList(head)

	for current := newHead; current != nil; current = current.next {
		fmt.Print(current.value, " -> ")
	}
}
