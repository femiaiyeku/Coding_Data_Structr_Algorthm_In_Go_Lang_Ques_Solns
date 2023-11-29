/*

Write a function that takes in a non-empty sorted array of distinct integers, constructs a BST from the integers, and returns the root of the BST.
The function should minimize the height of the BST.  You've been provided with a BST class that you'll have to use to construct the BST.
Each BST node has an integer value , a left child node, and a right child node. A node is said to be a valid BST node
if and only if it satisfies the BST property: its value is strictly greater than the values of every node to its left;
its value is less than or equal to the values of every node to its right; and its children nodes are either valid BST nodes themselves or None /
null. A BST is valid if and only if all of its nodes are valid BST nodes.
Note that the BST class already has an insert method which you can use if you want.

Sample Input
array = [1, 2, 5, 7, 10, 13, 14, 15, 22]

Sample Output
            10
            /    \
              2      14
            /   \   /   \
            1     5  13   15
                    \       \
                    7       22

            // This is one example of a BST with min height
            // that you could create from the input array.
            // You could create other BSTs with min height
            // from the same array; for example:
            //            10
            //          /     \
            //        5        15
            //      /   \     /   \
            //     2     7   13   22
            //   /           \
            //  1            14




			Hints
Hint 1
In order for the BST to have the smallest height possible, it needs to be balanced; in other words, it needs to have roughly the same number of nodes in its left subtree as in its right subtree.
Hint 2
How can you use the sorted nature of the input array to construct a balanced BST?
Hint 3
Grab the middle element of the array, and make that element be the root node of the BST. Then, grab the middle element between the beginning of the array and the first middle element, and make that element be the root of the BST's left subtree; similarly, make the middle element between the end of the array and the first middle element be the root of the BST's right subtree. Continue this approach until you run out of elements in the array.



*/

package main

import "fmt"

type BSTNode struct {
	Value int
	Left  *BSTNode
	Right *BSTNode
}

type BST struct {
	Root *BSTNode
}

func buildMinHeightBST(array []int) *BSTNode {
	// If the array is empty, return nil.
	if len(array) == 0 {
		return nil
	}

	// Find the middle index of the array.
	middleIndex := len(array) / 2

	// Create a new BST node with the middle element of the array.
	newNode := &BSTNode{Value: array[middleIndex]}

	// Recursively build the left and right subtrees of the new node.
	newNode.Left = buildMinHeightBST(array[:middleIndex])
	newNode.Right = buildMinHeightBST(array[middleIndex+1:])

	// Return the new node.
	return newNode
}

func main() {
	// Create a new array of distinct integers.
	array := []int{1, 2, 5, 7, 10, 13, 14, 15, 22}

	// Build a BST with minimal height from the array.
	root := buildMinHeightBST(array)

	// Print the BST in pre-order traversal.
	printPreOrder(root)
}

func printPreOrder(node *BSTNode) {
	if node == nil {
		return
	}

	fmt.Println(node.Value)
	printPreOrder(node.Left)
	printPreOrder(node.Right)
}
