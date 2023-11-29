/*


You're given a Binary Search Tree (BST) that has at least 2 nodes and that only hasnodes with unique values (no duplicate values).
 Exactly two nodes in the BST havehad their values swapped, therefore breaking the BST.
 Write a function thatreturns a repaired version of the tree with all values on the correct nodes.Your function can mutate the original tree;
 you do not need to create a new one.Moreover, the shape of the returned tree should be exactly the same as that of theoriginal input tree.
 Each BST node has an integer value, a left child node, and a right childnode.
 A node is said to be a valid BST node if and only if it satisfies the BSTproperty:
 its value is strictly greater than the values of every node to its left;
itsvalue is less than or equal to the values of every node to its right;
and itschildren nodes are either valid BST nodes themselves or None / null.

Sample Input

tree =  10
         /  \
         7   20
        / \  / \
        3 12 8  22
        /    \
      2      14


Sample Output

          10
         /  \
         7   20
        / \  / \
        3 12 8  22
        /    \
      2      14





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

func repairBST(tree *BST) {
	// Find the two nodes that have been swapped.
	firstSwappedNode, secondSwappedNode := findSwappedNodes(tree.Root)

	// Swap the values of the two nodes.
	firstSwappedNode.Value, secondSwappedNode.Value = secondSwappedNode.Value, firstSwappedNode.Value
}

func findSwappedNodes(node *BSTNode) (*BSTNode, *BSTNode) {
	// If the node is nil, return nil.
	if node == nil {
		return nil, nil
	}

	// Recursively find the swapped nodes in the left and right subtrees.
	leftSwappedNode, rightSwappedNode := findSwappedNodes(node.Left), findSwappedNodes(node.Right)

	// If the swapped nodes have been found, return them.
	if leftSwappedNode != nil && rightSwappedNode != nil {
		return leftSwappedNode, rightSwappedNode
	}

	// If the swapped nodes have not been found, check the current node.
	// If the current node's value is less than the value of its left child or greater than the value of its right child, then the current node is one of the swapped nodes.
	if node.Value < node.Left.Value || node.Value > node.Right.Value {
		if leftSwappedNode == nil {
			return node, nil
		} else {
			return node, rightSwappedNode
		}
	}

	// If the swapped nodes have not been found, return nil.
	return nil, nil
}

func main() {
	// Create a new BST.
	tree := &BST{}

	// Insert some values into the BST.
	tree.Insert(10)
	tree.Insert(7)
	tree.Insert(20)
	tree.Insert(3)
	tree.Insert(12)
	tree.Insert(8)
	tree.Insert(22)
	tree.Insert(2)
	tree.Insert(14)

	// Swap the values of two nodes in the BST.
	tree.Root.Left.Value, tree.Root.Right.Right.Value = tree.Root.Right.Right.Value, tree.Root.Left.Value

	// Repair the BST.
	repairBST(tree)

	// Print the BST in pre-order traversal.
	printPreOrder(tree.Root)
}

func printPreOrder(node *BSTNode) {
	if node == nil {
		return
	}

	fmt.Println(node.Value)
	printPreOrder(node.Left)
	printPreOrder(node.Right)
}
