/*

Write a function that takes in a Binary Tree, transforms it into a Right Sibling Tree, and returns its root.
A Right Sibling Tree is obtained by making every node in a Binary Tree have its right property point to its right sibling instead of its right child.
 A node's right sibling is the node immediately to its right on the same level or None / null if there is no node immediately to its right.
Note that once the transformation is complete, some nodes might no longer have a node pointing to them. For example, in the sample output below,
the node with value 10 no longer has any inbound pointers and is effectively unreachable.
The transformation should be done in place, meaning that the original data structure should be mutated (no new structure should be created).
Each BinaryTree node has an integer value , a left child node, and a right child node. Children nodes can either be Bi naryTree nodes themselves or None / null .



Sample Input

tree =    1
        /   \
        2     3
        / \   / \
        4   5 6   7
        / \   \    / \
        8   9   10 11 12
                /
                13


Sample Output

tree =    1
        /   \
        2     3
        / \   / \
        4   5 6   7
        / \     / \
        8   9   10 11
                /
                12
                /
                13


Hints
Hint 1

Try to identify a pattern or formula that determines how to reach a given node's right sibling.

Hint 2

There are two patterns: if a node is the left child of another node, its right sibling is its parent's right child; if a node is the right child of another node, its right sibling is its parent's right sibling's left child.

Hint 3

You'll need to a find a way to quickly access a node's parent's right child and a node's parent's right sibling; this won't be trivial because the second one implies that the parent node's original right pointer has been overwritten.

Hint 4

Recursively traverse the binary tree and sequence the transformation operations as follows: at any given node, recursively transform its left subtree into a right sibling tree, then edit the given node's right pointer to point to its right sibling, and then finally recursively transform its right subtree into a right sibling tree. This sequencing of operations will allow left child nodes to always access their parent's right child (before their parent's right pointer gets overwritten to point to the parent's right sibling) and will allow right child nodes to always access their parent's right sibling (after their parent's right pointer has gotten overwritten to point to the parent's right sibling).








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

func transformToRightSiblingTree(tree *BinaryTreeNode) {
	// If the tree is empty, return.
	if tree == nil {
		return
	}

	// Create a queue to store the nodes that need to be visited.
	queue := []*BinaryTreeNode{tree}

	// While the queue is not empty, do the following:
	for len(queue) > 0 {
		// Get the next node from the queue.
		currentNode := queue[0]
		queue = queue[1:]

		// Set the current node's right property to point to its right sibling.
		currentNode.Right = nil
		for currentNode.Right == nil && currentNode.Parent != nil {
			currentNode = currentNode.Parent
		}
		if currentNode.Parent != nil {
			currentNode.Right = currentNode.Parent.Right
		}

		// Add the current node's left and right children to the queue.
		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}
		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		}
	}
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
	tree.Right.Right.Right = &BinaryTreeNode{Value: 10}
	tree.Right.Right.Right.Left = &BinaryTreeNode{Value: 11}
	tree.Right.Right.Right.Left.Left = &BinaryTreeNode{Value: 12}
	tree.Right.Right.Right.Left.Left.Left = &BinaryTreeNode{Value: 13}

	// Transform the binary tree into a right sibling tree.
	transformToRightSiblingTree(tree)

	// Print the right sibling tree.
	fmt.Println(tree)
}
