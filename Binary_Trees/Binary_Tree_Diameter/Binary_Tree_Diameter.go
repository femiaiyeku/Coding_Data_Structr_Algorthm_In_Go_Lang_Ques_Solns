/*



Write a function that takes in a Binary Tree and returns its diameter. The diameter of a binary tree is defined as the length of its longest path, even if that path doesn't pass through the root of the tree.
A path is a collection of connected 1odes in a tree, where no node is connected to more than two other nodes. The length of a path is the number of edges between the path's first node and its last node.
Each BinaryTree node has an integer value , a left child node, and a right child node. Children nodes can either be Bi naryTree nodes themselves or None / null


Sample Input

tree =  1
        /   \
        3     2
        /   \
        7     4
        /       \
        8         5
        /           \
        9             6

Sample Output
6 // 9 -> 8 -> 7 -> 3 -> 4 -> 5 -> 6



Hints
Hint 1
How can you use the height of a binary tree and the heights of its subtrees to calculate its diameter?
Hint 2

The length of the longest path that goes through the root of a binary tree is the sum of the heights of its left and right subtrees (left subtree height + right subtree height). The diameter of a binary tree can be calculated by taking the maximum of: 1) the maximum subtree diameter (max(left subtree diameter, right subtree diameter)); and 2) the length of the longest path that goes through the root (left subtree height + right subtree height).

Hint 3

Implement a variation of depth-first search that recursively keeps track of both the diameter and the height of a each subtree in the input binary tree. Follow Hint #2 to continuously compute these diameters.





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

func diameterOfBinaryTree(tree *BinaryTreeNode) int {
	// Calculate the heights of the left and right subtrees.
	leftSubtreeHeight := calculateHeight(tree.Left)
	rightSubtreeHeight := calculateHeight(tree.Right)

	// The diameter of the tree is the maximum of the following:
	// 1. The diameter of the left subtree.
	// 2. The diameter of the right subtree.
	// 3. The sum of the heights of the left and right subtrees.
	return int(math.Max(float64(leftSubtreeHeight+rightSubtreeHeight+1), math.Max(float64(diameterOfBinaryTree(tree.Left)), float64(diameterOfBinaryTree(tree.Right)))))
}

func calculateHeight(tree *BinaryTreeNode) int {
	if tree == nil {
		return 0
	}

	return int(math.Max(float64(calculateHeight(tree.Left)+1), float64(calculateHeight(tree.Right)+1)))
}

func main() {
	// Create a new binary tree.
	tree := &BinaryTreeNode{Value: 1}
	tree.Left = &BinaryTreeNode{Value: 3}
	tree.Right = &BinaryTreeNode{Value: 2}
	tree.Left.Left = &BinaryTreeNode{Value: 7}
	tree.Left.Right = &BinaryTreeNode{Value: 4}
	tree.Right.Left = &BinaryTreeNode{Value: 8}
	tree.Right.Right = &BinaryTreeNode{Value: 5}
	tree.Left.Left.Left = &BinaryTreeNode{Value: 9}
	tree.Left.Left.Right = &BinaryTreeNode{Value: 6}

	// Calculate the diameter of the tree.
	diameter := diameterOfBinaryTree(tree)

	// Print the diameter of the tree.
	fmt.Println(diameter)
}
