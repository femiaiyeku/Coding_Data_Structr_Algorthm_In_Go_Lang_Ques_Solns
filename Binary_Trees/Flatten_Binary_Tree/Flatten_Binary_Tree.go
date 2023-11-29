/*


Write a function that takes 1n a Binary Tree, flattens 1t, and returns its leftmost node.
A flattened Binary Tree 1s a structure that's nearly 1dent1cal to a Doubly Linked List (except that nodes have left and right pointers instead of prev and next pointers),
where nodes follow the original tree·s left-to-right order.
Note that ,f the ,nput Binary Tree happens to be a valid Binary Search Tree, the nodes ,n the flattened tree will be sorted.
The flattening should be done 1n place, meaning that the original data structure should be mutated (no new structure should be created).
Each BinaryTree node has an integer value, a left child node, and a right child node. Children nodes can either be BinaryTree nodes themselves or None / null.


Sample Input
tree =     5
         /     \
         2       7
         /   \    /  \
        1     4  6    8
     /       / \
     0      3       9


     // This tree has been flattened:
    nodeOne= 5 // the leftmost node 1n the flattened tree
    nodeTwo= 2 // the second leftmost node 1n the flattened tree
    nodeThree= 3 // the third leftmost node 1n the flattened tree
Sample Output

true  // the flattened tree's node's values are 0, 1, 2, 3, 4, 5, 6, 7, 8, 9



Hints
Hint 1

You can solve this problem pretty easily by traversing the tree using the in-order tree-traversal technique, gathering all of the nodes in an array, and then iterating through them from left to right and connecting them accordingly. Can you solve this problem without storing an entire array of the tree's nodes?

Hint 2

Try to figure out what the relation between two adjacent nodes in the in-order-traversal order is, as far as positioning in the tree is concerned.

Hint 3

At any given node in the in-order-traversal order, the node immediately to its left is the rightmost node of its left subtree, and the node immediately the its right is the leftmost node of its right subtree.

Hint 4

Write a function that recursively gets the leftmost and rightmost nodes of a given node's left subtree and right subtree and that connects the left subtree's rightmost node to the given node and the right subtree's leftmost node to the given node.


*/

package main

import "fmt"

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func flattenBinaryTree(tree *BinaryTreeNode) *BinaryTreeNode {
	// If the tree is empty, return nil.
	if tree == nil {
		return nil
	}

	// Flatten the left subtree and return the leftmost node.
	leftmostLeftNode := flattenBinaryTree(tree.Left)

	// Flatten the right subtree and assign it to the leftmost node's right pointer.
	leftmostLeftNode.Right = flattenBinaryTree(tree.Right)

	// Set the current node's left and right pointers to nil.
	tree.Left = nil
	tree.Right = nil

	// Return the current node.
	return tree
}

func main() {
	// Create the binary tree.
	tree := &BinaryTreeNode{Value: 5}
	tree.Left = &BinaryTreeNode{Value: 2}
	tree.Right = &BinaryTreeNode{Value: 7}
	tree.Left.Left = &BinaryTreeNode{Value: 1}
	tree.Left.Right = &BinaryTreeNode{Value: 4}
	tree.Right.Left = &BinaryTreeNode{Value: 6}
	tree.Right.Right = &BinaryTreeNode{Value: 8}
	tree.Left.Left.Left = &BinaryTreeNode{Value: 0}
	tree.Left.Right.Left = &BinaryTreeNode{Value: 3}
	tree.Right.Left.Right = &BinaryTreeNode{Value: 9}

	// Flatten the binary tree.
	flattenedTree := flattenBinaryTree(tree)

	// Check if the flattened tree's node values are in the correct order.
	nodeValues := []int{}
	currentNode := flattenedTree
	for currentNode != nil {
		nodeValues = append(nodeValues, currentNode.Value)
		currentNode = currentNode.Right
	}

	// Print the flattened tree's node values.
	fmt.Println(nodeValues)

	// Check if the flattened tree is valid.
	isValid := true
	currentNode = flattenedTree
	for currentNode != nil && currentNode.Right != nil {
		if currentNode.Value > currentNode.Right.Value {
			isValid = false
			break
		}
		currentNode = currentNode.Right
	}

	// Print whether the flattened tree is valid.
	fmt.Println(isValid)
}
