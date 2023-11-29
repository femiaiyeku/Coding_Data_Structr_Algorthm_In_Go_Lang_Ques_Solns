/*



Write a function that takes in a potentially invalid Binary Search Tree (BSTI and returns a boolean representing whether the BST is valid.
Each BST node has an integer value , a left child node, and a r-ight child node. A node is said to be a valid BST node 1f and only if 1t satisfies the BST property: its value is strictly greater than the values of every node to its left; its value is less than or equal to the values of every node to its right; and its children nodes are either valid BST nodes themselves or None / null.
A BST is valid 1f and only if all of its nodes are valid BST nodes.


Sample Input:

tree = 10
        / \
         5  15
        / \  / \
         2  5 13 22
        /       \
         1        14

Sample Output: True




Hints
Hint 1
Every node in the BST has a maximum possible value and a minimum possible value. In other words, the value of any given node in the BST must be strictly smaller than some value (the value of its closest right parent) and must be greater than or equal to some other value (the value of its closest left parent).
Hint 2
Validate the BST by recursively calling the validateBst function on every node, passing in the correct maximum and minimum possible values to each. Initialize those values to be -Infinity and +Infinity.



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

func isValidBST(tree *BST) bool {
	// Validate the BST recursively.
	return validateBST(tree.Root, nil, nil)
}

func validateBST(node *BSTNode, minValue, maxValue *int) bool {
	// If the node is nil, then it is a valid BST.
	if node == nil {
		return true
	}

	// Check if the node's value is greater than the minimum value and less than or equal to the maximum value.
	if minValue != nil && node.Value <= *minValue || maxValue != nil && node.Value > *maxValue {
		return false
	}

	// Recursively validate the left and right subtrees.
	return validateBST(node.Left, minValue, &node.Value) && validateBST(node.Right, &node.Value, maxValue)
}

func main() {
	// Create a new BST.
	tree := &BST{}

	// Insert some values into the BST.
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(13)
	tree.Insert(22)
	tree.Insert(1)
	tree.Insert(14)

	// Validate the BST.
	isValidBST := isValidBST(tree)

	// Print the result.
	fmt.Println(isValidBST)
}
