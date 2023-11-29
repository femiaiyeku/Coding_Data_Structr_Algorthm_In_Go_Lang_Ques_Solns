/*



You're given two Linked Lists of potentially unequal length. These Linked Lists potentially merge at a shared
intersection node. Write a function that returns the intersection node or returns None/null if there is no
intersection.
Each LinkedList node has an integer value as well as a next node pointing to the next node in the list or
to None/null if it's the tail of the list.
Note: Your function should return an existing node. It should not modify either Linked List, and it should not
create any new Linked Lists.



Sample Input:
headOne = 0 -> 1 -> 2 -> 3 -> 4 -> 5 -> 6 // the head node with value 2
headTwo = 10 -> 11 -> 12 -> 13 // the head node with value 12


Sample Output:
12 -> 13 // the node with value 12
// Note that the intersection node can be anywhere between the two lists and the










*/


package main

type Node struct {
	value int
	next  *Node
}

func findIntersection(headOne *Node, headTwo *Node) *Node {
	// Create two sets to track the nodes in each list
	seenNodesOne := make(map[*Node]bool)
	seenNodesTwo := make(map[*Node]bool)

	// Traverse both lists simultaneously
	currentOne := headOne
	currentTwo := headTwo

	for currentOne != nil && currentTwo != nil {
		// If we've seen a node before, it's the intersection node
		if seenNodesOne[currentOne] || seenNodesTwo[currentTwo] {
			return currentOne
		}

		// Add the current nodes to the respective sets
		seenNodesOne[currentOne] = true
		seenNodesTwo[currentTwo] = true

		// Move to the next nodes in each list
		currentOne = currentOne.next
		currentTwo = currentTwo.next
	}

	// If we reach the end of either list without finding an intersection, there is no intersection
	return nil
}

func main() {
	headOne := &Node{value: 0, next: &Node{value: 1, next: &Node{value: 2, next: &Node{value: 3, next: &Node{value: 4, next: &Node{value: 5, next: &Node{value: 6}}}}}}}
	headTwo := &Node{value: 10, next: &Node{value: 11, next: &Node{value: 12, next: &Node{value: 13}}}}}}

	intersectionNode := findIntersection(headOne, headTwo)
	if intersectionNode != nil {
		fmt.Println("Intersection node:", intersectionNode.value)
	} else {
		fmt.Println("No intersection found")
	}
}

