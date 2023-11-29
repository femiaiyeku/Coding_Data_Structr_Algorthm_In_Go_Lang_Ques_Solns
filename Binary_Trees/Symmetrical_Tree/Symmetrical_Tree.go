/*

Write a function that takes in a Binary Tree and returns if that tree is symmetrical. A tree is symmetrical if the left and right subtrees are mirror images of each other.
Each BinaryTree node has an irteger value , a left child node, and a right child node. Children nodes can either be Bi naryTree nodes themselves or None / null .

Sample Input:
tree =   1
        / \
       2   2
        / \ / \
       3  4 4  3
      / \     / \
    5    6   6   5

Sample Output:
True






*/

package main

import "fmt"

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func isSymmetricalBinaryTree(tree *BinaryTreeNode) bool {
	// If the tree is empty, return true.
	if tree == nil {
		return true
	}

	// Recursively check if the left and right subtrees are symmetrical.
	return isSymmetricalBinaryTreeHelper(tree.Left, tree.Right)
}

func isSymmetricalBinaryTreeHelper(left *BinaryTreeNode, right *BinaryTreeNode) bool {
	// If either the left or right subtree is empty, return false.
	if left == nil || right == nil {
		return false
	}

	// If the values of the left and right nodes are not equal, return false.
	if left.Value != right.Value {
		return false
	}

	// Recursively check if the left and right subtrees of the left and right nodes are symmetrical.
	return isSymmetricalBinaryTreeHelper(left.Left, right.Right) && isSymmetricalBinaryTreeHelper(left.Right, right.Left)
}

func main() {
	// Create the binary tree.
	tree := &BinaryTreeNode{Value: 1}
	tree.Left = &BinaryTreeNode{Value: 2}
	tree.Right = &BinaryTreeNode{Value: 2}
	tree.Left.Left = &BinaryTreeNode{Value: 3}
	tree.Left.Right = &BinaryTreeNode{Value: 4}
	tree.Right.Left = &BinaryTreeNode{Value: 4}
	tree.Right.Right = &BinaryTreeNode{Value: 3}
	tree.Left.Left.Left = &BinaryTreeNode{Value: 5}
	tree.Left.Left.Right = &BinaryTreeNode{Value: 6}
	tree.Right.Right.Left = &BinaryTreeNode{Value: 6}
	tree.Right.Right.Right = &BinaryTreeNode{Value: 5}

	// Check if the binary tree is symmetrical.
	isSymmetrical := isSymmetricalBinaryTree(tree)

	// Print whether or not the binary tree is symmetrical.
	fmt.Println(isSymmetrical)
}
