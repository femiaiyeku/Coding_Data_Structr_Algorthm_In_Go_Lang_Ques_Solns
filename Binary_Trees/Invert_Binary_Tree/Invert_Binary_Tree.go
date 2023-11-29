/*



Write a function that takes in a Binary Tree and inverts it. In other words, the function should swap every left node in the tree for its corresponding right node.
Each BinaryTree node has an integer value, a left child node, and a right child node. Children nodes can either be Bi naryTree nodes themselves or None / null

Sample Input

tree =    1
         /     \
         2       3
        /  \    /  \
        4    5  6    7
        / \
       8   9



Sample Output

tree =    1
            /     \
            3       2
            /  \    /  \
            7    6  5    4
                        / \
                       9   8





*/

package main

import "fmt"

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func invertBinaryTree(tree *BinaryTreeNode) {
	// If the tree is empty, return.
	if tree == nil {
		return
	}

	// Swap the left and right subtrees.
	tree.Left, tree.Right = tree.Right, tree.Left

	// Recursively invert the left and right subtrees.
	invertBinaryTree(tree.Left)
	invertBinaryTree(tree.Right)
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

	// Invert the binary tree.
	invertBinaryTree(tree)

	// Print the inverted binary tree.
	fmt.Println(tree)
}
