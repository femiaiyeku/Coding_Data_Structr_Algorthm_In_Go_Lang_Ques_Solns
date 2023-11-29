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

type Node struct {
	value int
	next  *Node
}

func mergeLinkedLists(head1 *Node, head2 *Node) *Node {
	dummyHead := &Node{}
	curr := dummyHead

	for head1 != nil && head2 != nil {
		if head1.value < head2.value {
			curr.next = head1
			head1 = head1.next
		} else {
			curr.next = head2
			head2 = head2.next
		}

		curr = curr.next
	}

	// Add the remaining nodes from the first linked list, if any.
	if head1 != nil {
		curr.next = head1
	}

	// Add the remaining nodes from the second linked list, if any.
	if head2 != nil {
		curr.next = head2
	}

	return dummyHead.next
}
