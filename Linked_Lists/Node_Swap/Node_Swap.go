/*


Write a function that takes in the head of a Singly Linked List, swaps every pair of adjacent nodes in place (i.e.,
doesn't create a brand new list), and returns its new head.
If the input Linked List has an odd number of nodes, its final node should remain the same.
Each LinkedList node has an integer value as well as a next node pointing to the next node in the list or
to None/null if it's the tail of the list.
You can assume that the input Linked List will always have at least one node; in other words, the head will never
be None / null.

Sample Input

head = 0 -> 1 -> 2 -> 3 -> 4 -> 5 // the head node with value 0


Sample Output

1 -> 0 -> 3 -> 2 -> 5 -> 4 // the head node with value 1







*/

package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func swapAdjacentNodes(head *Node) *Node {
	prev := &Node{next: head}
	current := head
	next := head.next

	for next != nil {
		prev.next = next
		current.next = next.next

		next.next = current
		prev = current
		current = next

		next = current.next
	}

	return prev.next
}

func main() {
	head := &Node{value: 0, next: &Node{value: 1, next: &Node{value: 2, next: &Node{value: 3, next: &Node{value: 4, next: &Node{value: 5}}}}}}

	newHead := swapAdjacentNodes(head)

	for current := newHead; current != nil; current = current.next {
		fmt.Print(current.value, " -> ")
	}
}
