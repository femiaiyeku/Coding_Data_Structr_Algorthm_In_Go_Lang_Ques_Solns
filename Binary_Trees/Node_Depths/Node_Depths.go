/*

The distance between a node in a Binary Tree and the tree's root is called the node's depth.
rite a function that takes in a Binar1 Tree and returns the sum of its nodes· depths.
ach BinaryTree node has an integer value , a left child node, and a right child node. Children nodes can either be Bi naryTree nodes themselves or None I nulL


Sample Input
tree =  1
         /  \
        2    3
         / \  / \
        4  5 6  7
         / \
        8   9

Sample Output
16

// The depth of the node with value 2 is 1.
// The depth of the node with value 3 is 1.
// The depth of the node with value 4 is 2.
// The depth of the node with value 5 is 2.
// Etc..

// Summing all of these depths yields 16.






*/

package main

import "fmt"

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func sumNodeDepths(tree *BinaryTreeNode) int {
	// If the tree is empty, return 0.
	if tree == nil {
		return 0
	}

	// Recursively calculate the sum of the depths of the left and right subtrees.
	leftSubtreeDepthSum := sumNodeDepths(tree.Left)
	rightSubtreeDepthSum := sumNodeDepths(tree.Right)

	// Calculate the depth of the current node.
	currentNodeDepth := 1

	// Return the sum of the depths of the current node, its left subtree, and its right subtree.
	return currentNodeDepth + leftSubtreeDepthSum + rightSubtreeDepthSum
}

func main() {
	// Create the binary tree.
	tree := &BinaryTreeNode{Value: 1}
	tree.Left = &BinaryTreeNode{Value: 2}
	tree.Right = &BinaryTreeNode{Value: 3}
	tree.Left.Left = &BinaryTreeNode{Value: 4}
	tree.Left.Right = &BinaryTreeNode{Value: 5}
	tree.Right.Left = &BinaryTreeNode{Value: 6}
	tree.Right.Right = &BinaryTreeNode{Value: 7}
	tree.Left.Left.Left = &BinaryTreeNode{Value: 8}
	tree.Left.Left.Right = &BinaryTreeNode{Value: 9}

	// Calculate the sum of the depths of all nodes in the binary tree.
	sumOfNodeDepths := sumNodeDepths(tree)

	// Print the sum of the depths of all nodes in the binary tree.
	fmt.Println(sumOfNodeDepths)
}
