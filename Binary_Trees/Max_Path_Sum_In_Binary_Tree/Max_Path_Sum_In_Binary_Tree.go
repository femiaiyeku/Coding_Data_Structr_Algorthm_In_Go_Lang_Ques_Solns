/*
Write a function that takes in a Binary Tree and returns its max path sum.
A path is a collection of connected nodes in a tree, where no node is connected to more than two other nodes; a path
sum is the sum of the values of the nodes in a particular path.
Each BinaryTree node has an integer value, a left child node, and a right child node. Children nodes can
either be BinaryTree nodes themselves or None / null


Sample Input:
tree =    1
        /   \
        2     3
     /   \     /  \
    4     5    6  7


Sample Output:
18
// 5 + 2 + 1 + 3 + 7





Hints
Hint 1

If you were to imagine each node in a Binary Tree as the root of the Binary Tree, temporarily eliminating all of the nodes that come above it, how would you find the max path sum for each of these newly imagined Binary Trees? In simpler terms, how can you find the max path sum for each subtree in the Binary Tree?

Hint 2

For every node in a Binary Tree, there are four options for the max path sum that includes its value: the node's value alone, the node's value plus the max path sum of its left subtree, the node's value plus the max path sum of its right subtree, or the node's value plus the max path sum of both its subtrees.

Hint 3

A recursive algorithm that computes each node's max path sum and uses it to compute its parents' nodes' max path sums seems appropriate, but realize that you cannot have a path going through a node and both its subtrees as well as that node's parent node. In other words, the fourth option mentioned in Hint #2 poses a challenge to implementing a recursive algorithm that solves this problem. How can you get around it?






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

func maxPathSum(tree *BinaryTreeNode) int {
	// If the tree is empty, return 0.
	if tree == nil {
		return 0
	}

	// Recursively calculate the maximum path sums of the left and right subtrees.
	leftSubtreeMaxPathSum := maxPathSum(tree.Left)
	rightSubtreeMaxPathSum := maxPathSum(tree.Right)

	// Calculate the maximum path sum of the current node, including its left and right subtrees.
	currentMaxPathSum := int(math.Max(float64(leftSubtreeMaxPathSum), float64(rightSubtreeMaxPathSum))) + tree.Value

	// Calculate the maximum path sum of the current node, including only itself.
	maxPathSumFromCurrentNode := tree.Value

	// Return the maximum of the current path sum, the maximum path sum from the current node, and the maximum path sums of the left and right subtrees.
	return int(math.Max(float64(currentMaxPathSum), float64(maxPathSumFromCurrentNode)))
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

	// Find the maximum path sum in the binary tree.
	maxPathSum := maxPathSum(tree)

	// Print the maximum path sum.
	fmt.Println(maxPathSum)
}
