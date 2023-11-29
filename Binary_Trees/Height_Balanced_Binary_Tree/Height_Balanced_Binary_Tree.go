/*

You're given the root node of a Binary Tree. Write a function that returns true if this Binary Tree is height balanced and false if it isn't.
A Binary Tree is height balanced if for each node in the tree, the difference between the height of its left subtree and the height of its right subtree is at most 1 .
Each BinaryTree node has an integer value , a left child node, and a right child node. Children nodes can either be Bi naryTree nodes themselves or None / null .


Sample Input

tree =    1
        /   \
       2     3
     /   \     \
        4     5     6
             / \
            7   8

Sample Output

true







*/

package main

import (
	"fmt"
	"math"
)

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func isHeightBalanced(tree *BinaryTreeNode) bool {
	// If the tree is empty, return true.
	if tree == nil {
		return true
	}

	// Calculate the heights of the left and right subtrees.
	leftSubtreeHeight := calculateHeight(tree.Left)
	rightSubtreeHeight := calculateHeight(tree.Right)

	// If the difference between the heights of the left and right subtrees is greater than 1, then the tree is not height-balanced.
	if math.Abs(float64(leftSubtreeHeight-rightSubtreeHeight)) > 1 {
		return false
	}

	// Recursively check if the left and right subtrees are height-balanced.
	return isHeightBalanced(tree.Left) && isHeightBalanced(tree.Right)
}

func calculateHeight(tree *BinaryTreeNode) int {
	// If the tree is empty, return 0.
	if tree == nil {
		return 0
	}

	// Calculate the heights of the left and right subtrees.
	leftSubtreeHeight := calculateHeight(tree.Left)
	rightSubtreeHeight := calculateHeight(tree.Right)

	// Return the height of the taller subtree plus 1.
	return int(math.Max(float64(leftSubtreeHeight), float64(rightSubtreeHeight)) + 1)
}

func main() {
	// Create the binary tree.
	tree := &BinaryTreeNode{Value: 1}
	tree.Left = &BinaryTreeNode{Value: 2}
	tree.Right = &BinaryTreeNode{Value: 3}
	tree.Left.Left = &BinaryTreeNode{Value: 4}
	tree.Left.Right = &BinaryTreeNode{Value: 5}
	tree.Right.Right = &BinaryTreeNode{Value: 6}
	tree.Right.Right.Left = &BinaryTreeNode{Value: 7}
	tree.Right.Right.Right = &BinaryTreeNode{Value: 8}

	// Check if the binary tree is height-balanced.
	isHeightBalanced := isHeightBalanced(tree)

	// Print whether the binary tree is height-balanced.
	fmt.Println(isHeightBalanced)
}
