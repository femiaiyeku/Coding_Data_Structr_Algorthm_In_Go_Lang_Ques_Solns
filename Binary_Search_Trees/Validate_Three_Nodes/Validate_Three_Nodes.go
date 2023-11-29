/*


You're given three nodes that are contained in the same Binary Search Tree: nodeOne , node Two , and node Three .
Write a function that returns a boolean representing whether one of nodeOne or node Three is an ancestor of node Two and the other node is a descendant of nodeTwo . For example,
if your function determines that nodeOne is an ancestor of nodeTwo , then it needs to see if node Three is a descendant of node Two . If your function determines that node Three is an ancestor,
then it needs to see if nodeOne is a descendant.
A descendant of a node N is defined as a node contained in the tree rooted at N .
A node N is an ancestor of another node M if M is a descendant of N .
It isn't guaranteed that nodeOne or node Three will be ancestors or descendants of node Two ,
 but it is guaranteed that all three nodes will be unique and will never be None / null .
In other words, you'll be given valid input nodes. Each BST node has an integer value , a left child node, and a right child node.
A node is said to be a valid BST node if and only if it satisfies the BST property: its value is strictly greater than the values of every node to its left;
its value is less than or equal to the values of every node to its right; and its children nodes are either valid BST nodes themselves or None /
null.


Sample Input
tree =    5
        /   \
         2     7
        / \   / \
         1   4 6   8
         /   /
         0   3

nodeOne = 5
nodeTwo = 2
nodeThree = 3

Sample Output

true




*/

package main

import "fmt"

type BSTNode struct {
	Value int
	Left  *BSTNode
	Right *BSTNode
}

type BST struct {
	Root *BSTNode
}

func isAncestorAndDescendant(tree *BST, nodeOne *BSTNode, nodeTwo *BSTNode, nodeThree *BSTNode) bool {
	// Check if nodeOne is an ancestor of nodeTwo.
	if isAncestor(tree.Root, nodeOne, nodeTwo) {
		// Check if nodeThree is a descendant of nodeTwo.
		if isDescendant(tree.Root, nodeTwo, nodeThree) {
			return true
		}
	}

	// Check if nodeThree is an ancestor of nodeTwo.
	if isAncestor(tree.Root, nodeThree, nodeTwo) {
		// Check if nodeOne is a descendant of nodeTwo.
		if isDescendant(tree.Root, nodeTwo, nodeOne) {
			return true
		}
	}

	// Otherwise, nodeOne and nodeThree are not an ancestor and descendant of nodeTwo.
	return false
}

func isAncestor(node *BSTNode, ancestor *BSTNode, descendant *BSTNode) bool {
	// If the node is nil, then it is not an ancestor of the descendant.
	if node == nil {
		return false
	}

	// If the node is the ancestor, then return true.
	if node == ancestor {
		return true
	}

	// Recursively check if the ancestor is in the left or right subtree.
	return isAncestor(node.Left, ancestor, descendant) || isAncestor(node.Right, ancestor, descendant)
}

func isDescendant(node *BSTNode, descendant *BSTNode, ancestor *BSTNode) bool {
	// If the node is nil, then it is not a descendant of the ancestor.
	if node == nil {
		return false
	}

	// If the node is the descendant, then return true.
	if node == descendant {
		return true
	}

	// Recursively check if the descendant is in the left or right subtree.
	return isDescendant(node.Left, descendant, ancestor) || isDescendant(node.Right, descendant, ancestor)
}

func main() {
	// Create a new BST.
	tree := &BST{}

	// Insert some values into the BST.
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(8)
	tree.Insert(0)
	tree.Insert(3)

	// Create the three nodes.
	nodeOne := tree.Root
	nodeTwo := tree.Root.Left
	nodeThree := tree.Root.Left.Right

	// Check if one of nodeOne or nodeThree is an ancestor of nodeTwo and the other node is a descendant of nodeTwo.
	isAncestorAndDescendant := isAncestorAndDescendant(tree, nodeOne, nodeTwo, nodeThree)

	// Print the result.
	fmt.Println(isAncestorAndDescendant)
}
