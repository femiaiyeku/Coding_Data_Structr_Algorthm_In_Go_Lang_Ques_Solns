/*

Write a function that takes in a Binary Tree with at least one node and checks if that Binary Tree can be split into two Binary Trees of equal sum by removing a single edge.
If this split is possible, return the new sum of each Binary Tree, otherwise return 0. Note that you do not need to return the edge that was removed.
The sum of a Binary Tree is the sum of all values in that Binary Tree.
Each BinaryTree node has an integer value , a left child node,
and a right child node. Children nodes can either be Bi naryTree nodes themselves or None / nul 1 .




tree = 1
      / \
      3  -2
    / \   / \
  -6   -5 5  2
  /
  2

    Sample Output:
    6
    // 6 == 1 - 3 + -6 + 2




*/

package main

import (
	"fmt"
)

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func canSplitBinaryTreeIntoEqualSumTrees(tree *BinaryTreeNode) int {
	// If the tree is empty, return 0.
	if tree == nil {
		return 0
	}

	// Calculate the sum of the tree.
	treeSum := calculateTreeSum(tree)

	// If the tree sum is odd, then it is not possible to split the tree into two binary trees of equal sum.
	if treeSum%2 != 0 {
		return 0
	}

	// Recursively check if the left and right subtrees of the tree can be split into two binary trees of equal sum.
	leftSubtreeCanSplit := canSplitBinaryTreeIntoEqualSumTrees(tree.Left)
	rightSubtreeCanSplit := canSplitBinaryTreeIntoEqualSumTrees(tree.Right)

	// If either the left or right subtree can be split into two binary trees of equal sum, then the current tree can also be split.
	if leftSubtreeCanSplit > 0 || rightSubtreeCanSplit > 0 {
		return treeSum / 2
	}

	// Otherwise, it is not possible to split the current tree into two binary trees of equal sum.
	return 0
}

func calculateTreeSum(tree *BinaryTreeNode) int {
	// If the tree is empty, return 0.
	if tree == nil {
		return 0
	}

	// Calculate the sum of the left and right subtrees.
	leftSubtreeSum := calculateTreeSum(tree.Left)
	rightSubtreeSum := calculateTreeSum(tree.Right)

	// Return the sum of the current node, its left subtree, and its right subtree.
	return tree.Value + leftSubtreeSum + rightSubtreeSum
}

func main() {
	// Create the binary tree.
	tree := &BinaryTreeNode{Value: 1}
	tree.Left = &BinaryTreeNode{Value: 3}
	tree.Right = &BinaryTreeNode{Value: -2}
	tree.Left.Left = &BinaryTreeNode{Value: -6}
	tree.Left.Right = &BinaryTreeNode{Value: -5}
	tree.Right.Left = &BinaryTreeNode{Value: 5}
	tree.Right.Right = &BinaryTreeNode{Value: 2}
	tree.Left.Left.Left = &BinaryTreeNode{Value: 2}

	// Check if the binary tree can be split into two binary trees of equal sum.
	canSplit := canSplitBinaryTreeIntoEqualSumTrees(tree)

	// If the binary tree can be split, print the new sum of each binary tree.
	if canSplit > 0 {
		fmt.Println(canSplit)
	} else {
		fmt.Println(0)
	}
}
