/*

Write a function that takes in the head of a Singly Linked List and an integer k , shifts the list in place (i.e., doesn't create a brand new list) by k positions, and returns its new head.
Shifting a Linked List means moving its nodes forward or backward and wrapping them around the list where appropriate. For example, shifting a Linked List forward by one position would make its tail become the new head of the linked list.
Whether nodes are moved forward or backward is determined by whether k is positive or negative.
Each L inkedl i st node has an integer value as well as a next node pointing to the next node in the list or to None I null if it's the tail of the list.
You can assume that the input Linked List will always have at least one node; in other words, the head will never be None / null.

Sample Input:

head = 0 -> 1 -> 2 -> 3 -> 4 -> 5 // the head node with value 0
k = 2

Sample Output:

4 -> 5 -> 0 -> 1 -> 2 -> 3 // the new head node with value 4




*/

package main

func shiftLinkedList(head *Node, k int) *Node {
	listLength := 1
	listTail := head

	for listTail.next != nil {
		listTail = listTail.next
		listLength++
	}

	offset := abs(k) % listLength
	if offset == 0 {
		return head
	}

	newTailPosition := listLength - offset
	if k > 0 {
		newTailPosition = offset
	}

	newTail := head
	for i := 1; i < newTailPosition; i++ {
		newTail = newTail.next
	}

	newHead := newTail.next
	newTail.next = nil
	listTail.next = head

	return newHead
}

func abs(k int) int {
	if k < 0 {
		return -k
	}
	return k
}
