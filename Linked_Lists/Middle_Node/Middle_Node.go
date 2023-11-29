/*


You're given a Linked List with at least one node. Write a function that returns the middle node of the Linked List.
If there are two middle nodes (i.e. an even length list), your function should return the second of these nodes.
Each LinkedList node has an integer value as well as a next node pointing to the next node in the list or
to None / null if it's the tail of the list.

Sample Input:

head = 0 -> 1 -> 2 -> 3 -> 4 -> 5 // the head node with value 0

Sample Output:

3 -> 4 -> 5 // the head node with value 3



*/

package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func findMiddleNode(head *Node) *Node {
	slow := head
	fast := head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}

func main() {
	head := &Node{value: 0, next: &Node{value: 1, next: &Node{value: 2, next: &Node{value: 3, next: &Node{value: 4, next: &Node{value: 5}}}}}}

	middleNode := findMiddleNode(head)
	fmt.Println("Middle node:", middleNode.value)
}
