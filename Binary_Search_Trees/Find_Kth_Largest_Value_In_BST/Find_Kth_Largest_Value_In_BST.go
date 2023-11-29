/*



Write a function that takes in a Binary Search Tree (BST) and a positive integer k and returns the kth largest integer contained in the BST.
You can assume that there will only' be integer values in the BST and that k is less than or equal to the number of nodes in the tree.
Also, for the purpose of this question, duplicate integers will be treated as distinct values. In other words,
the second largest value in a BST containing values {5, 7, 7} will be 7 -not 5 .
Each BST node has an integer value, a left child node, and a right child node.
A node is said to be a valid BST node if and only if it satisfies the BST property:
 its value is strictly greater than the values of every node to its left;
 its value is less than or equal to the values of every node to its right; and its children nodes are either valid BST nodes themselves or None /
null.

Sample Input

tree =  15
        / \
        5 20
        / \ / \
        2 5 17 22
        / \
        1 3


        k = 3

Sample Output

17




*/

package main

import (
	"fmt"
)

type BSTNode struct {
	Value int
	Left  *BSTNode
	Right *BSTNode
}

type BST struct {
	Root *BSTNode
}

func kthLargest(bst *BST, k int) int {
	// Create a stack to store the nodes of the BST in reverse in-order traversal order.
	stack := []BSTNode{}

	// Perform a reverse in-order traversal of the BST.
	node := bst.Root
	for node != nil || len(stack) > 0 {
		// If the current node is not nil, push it onto the stack and move to its left child.
		if node != nil {
			stack = append(stack, *node)
			node = node.Left
		} else {
			// Pop the top node from the stack and assign it to the current node.
			node = &stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Check if the current node is the kth largest node.
			if k == 1 {
				return node.Value
			}

			// Move to the current node's right child.
			node = node.Right
		}
	}

	// If we reach here, then the BST does not have at least k nodes.
	return -1
}

func main() {
	// Create a new BST.
	bst := &BST{}

	// Insert some values into the BST.
	bst.Insert(15)
	bst.Insert(5)
	bst.Insert(20)
	bst.Insert(2)
	bst.Insert(5)
	bst.Insert(17)
	bst.Insert(22)
	bst.Insert(1)
	bst.Insert(3)

	// Find the kth largest element in the BST.
	k := 3
	kthLargestElement := kthLargest(bst, k)

	// Print the kth largest element.
	fmt.Println(kthLargestElement)
}
