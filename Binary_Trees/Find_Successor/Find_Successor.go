/*

Write a function that takes in a Bincry Tree (where nodes have an additional pointer to their parent node) as well as a node contained in that tree and returns the given node's successor.
A node's successor is the next node to be visited (immediately after the given node) when traversing its tree using the in-order tree-traversal technique. A node has no successor if it's the last node to be visited in the in-order traversal.
If a node has no successor, your function should return None / nul 1 .
Each BinaryTree node has an integer value, a parent node, a left child node, and a right child node. Children nodes can either be Bi naryTree nodes themselves or None / null.



Sample Input
tree =
         1
        /   \
        2     3
       / \
      4   5
    /
    6
node = 5

Sample Output
1


Hints
Hint 1

Start by performing an in-order traversal of the tree and storing the nodes in an array as you go. Then, traverse the nodes that you've stored; once you find the input node, return the node immediately after it in the array.

Hint 2

Can you think of a more time-efficient way to solve this problem without performing the entire in-order traversal?

Hint 3

Use the fact that each node has a pointer to its parent to solve this problem in O(h) time, where h is the height of the tree.

Hint 4

If the given node has a right subtree, then the next node in the in-order traversal is simply the leftmost node in that right subtree. If it doesn't have a right subtree, then we need to traverse up the tree looking for an ancestor of this node that contains the node in question in its left subtree. The first node that we find that contains the input node in its left subtree is the one that will be visited next in the in-order traversal. If we reach the root node, and the input node isn't in the root node's left subtree, then the input node has no successor, because it must be the rightmost node of entire tree.


*/

package main

import (
	"fmt"
)

type BinaryTreeNode struct {
	Value  int
	Parent *BinaryTreeNode
	Left   *BinaryTreeNode
	Right  *BinaryTreeNode
}

func getSuccessor(node *BinaryTreeNode) *BinaryTreeNode {
	// If the node has a right child, then its successor is the leftmost node in its right subtree.
	if node.Right != nil {
		return getLeftmostNode(node.Right)
	}

	// If the node does not have a right child, then its successor is its parent, but only if its parent is not to its left.
	if node.Parent != nil && node.Parent.Left != node {
		return node.Parent
	}

	// Otherwise, the node has no successor.
	return nil
}

func getLeftmostNode(node *BinaryTreeNode) *BinaryTreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func main() {
	// Create the binary tree.
	tree := &BinaryTreeNode{Value: 1}
	tree.Left = &BinaryTreeNode{Value: 2, Parent: tree}
	tree.Right = &BinaryTreeNode{Value: 3, Parent: tree}
	tree.Left.Left = &BinaryTreeNode{Value: 4, Parent: tree.Left}
	tree.Left.Right = &BinaryTreeNode{Value: 5, Parent: tree.Left}
	tree.Left.Left.Left = &BinaryTreeNode{Value: 6, Parent: tree.Left.Left}

	// Find the successor of node 5.
	successor := getSuccessor(tree.Left.Right)

	// If the successor is not nil, print its value. Otherwise, print "None".
	if successor != nil {
		fmt.Println(successor.Value)
	} else {
		fmt.Println("None")
	}
}
