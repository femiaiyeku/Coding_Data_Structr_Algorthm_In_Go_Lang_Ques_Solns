/*



You're given the head of a Singly Linked List of arbitrary length k .
 Write a function that zips the Linked List in place (i.e., doesn't create a brand new list) and returns its head.
A Linked List is zipped if its nodes are in the following order, where k 1s the length of the Linked List:

1st node -> kth node -> 2nd node -> (k - 1)th node -> 3rd node -> (k - 2)th node -> ...


Each Linked List node has an integer value as well as a next node pointing to the next node in the 11st or to None I null 1f it's the tail of the list.
You can assume that the input Linked List will always have at least one node; in other words, the head will never be None I null

Sample Input:
head = 0 -> 1 -> 2 -> 3 -> 4 -> 5 // the head node with value 0

Sample Output:
0 -> 5 -> 1 -> 4 -> 2 -> 3 // the head node with value 0














*/

package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func zipLinkedList(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	slow := head
	fast := head

	var middle *Node

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if middle == nil {
			middle = slow
		}
	}

	if slow == nil {
		return head
	}

	middleNext := middle.next
	middle.next = nil

	tail := slow
	current := middle.next

	for current != nil {
		next := current.next

		current.next = tail
		tail = current

		current = next
	}

	tail.next = middleNext

	return head
}

func main() {
	head := &Node{value: 0, next: &Node{value: 1, next: &Node{value: 2, next: &Node{value: 3, next: &Node{value: 4, next: &Node{value: 5}}}}}}

	newHead := zipLinkedList(head)

	for current := newHead; current != nil; current = current.next {
		fmt.Print(current.value, " -> ")
	}
}
