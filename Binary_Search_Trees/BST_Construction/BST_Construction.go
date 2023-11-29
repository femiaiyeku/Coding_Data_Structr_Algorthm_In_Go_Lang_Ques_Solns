/*

Write a BST class for a Binary Search Tree. The class should support
• 	Inserting values wrth the insert method.
• 	Removing values wrth the remove method; thrs method should only remove the first instance of a grven value.
• 	Searching for values with the contains method.

Note that you can't remove values from a srngle-node tree. In other words. callrng the remove method on a srngle-node tree should simply not do anything.
Each BST node has an rnteger value , a left child node. and a right ch rid node.
 A node rs sard to be a valrd BST node rf and only if it satrsfres the BST property:
   rts value rs strictly greater than the values of every node to rts left;
   rts value is less than or equal to the values of every node to its rrght;
 and its children nodes are erther valid BST nodes themselves or None / null

Sample Input:

tree = 10
        / \
            5  15
            / \  / \
                2  5 13 22
                /       \
                    1        14

Sample Output (after inserting 12):

tree = 10
        / \
            5  15
            / \  / \
                2  5 13 22
                /       \
                    1        14
                    /    \
                null  12

Sample Output (after removing 10):
tree = 12
        / \
            5  15
            / \  / \
                2  5 13 22
                /       \
                    1        14
                    /    \
                null  null

Sample Output (after searching for 15): true



Hints
Hint 1
As you try to insert, find, or a remove a value into, in, or from a BST, you will have to traverse the tree's nodes. The BST property allows you to eliminate half of the remaining tree at each node that you traverse: if the target value is strictly smaller than a node's value, then it must be (or can only be) located to the left of the node, otherwise it must be (or can only be) to the right of that node.
Hint 2
Traverse the BST all the while applying the logic described in Hint #1. For insertion, add the target value to the BST once you reach a leaf (None / null) node. For searching, if you reach a leaf node without having found the target value that means the value isn't in the BST. For removal, consider the various cases that you might encounter: the node you need to remove might have two children nodes, one, or none; it might also be the root node; make sure to account for all of these cases.
Hint 3
What are the advantages and disadvantages of implementing these methods iteratively as opposed to recursively?



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

func (bst *BST) Insert(value int) {
	// Create a new node.
	newNode := &BSTNode{Value: value}

	// If the tree is empty, set the root node to the new node.
	if bst.Root == nil {
		bst.Root = newNode
		return
	}

	// Find the appropriate place to insert the new node.
	currentNode := bst.Root
	for currentNode != nil {
		if value < currentNode.Value {
			// Insert the new node to the left.
			if currentNode.Left == nil {
				currentNode.Left = newNode
				return
			}

			// Move to the left child node.
			currentNode = currentNode.Left
		} else {
			// Insert the new node to the right.
			if currentNode.Right == nil {
				currentNode.Right = newNode
				return
			}

			// Move to the right child node.
			currentNode = currentNode.Right
		}
	}
}

func (bst *BST) Remove(value int) {
	// Find the node to remove.
	nodeToRemove := bst.FindNode(value)
	if nodeToRemove == nil {
		return
	}

	// If the node to remove has no children, simply remove it.
	if nodeToRemove.Left == nil && nodeToRemove.Right == nil {
		// Remove the node from its parent.
		if nodeToRemove.Parent == nil {
			bst.Root = nil
		} else if nodeToRemove.Parent.Left == nodeToRemove {
			nodeToRemove.Parent.Left = nil
		} else {
			nodeToRemove.Parent.Right = nil
		}

		return
	}

	// If the node to remove has one child, replace it with its child.
	if nodeToRemove.Left == nil {
		// Replace the node to remove with its right child.
		nodeToRemove.Value = nodeToRemove.Right.Value
		nodeToRemove.Right = nil
		return
	} else if nodeToRemove.Right == nil {
		// Replace the node to remove with its left child.
		nodeToRemove.Value = nodeToRemove.Left.Value
		nodeToRemove.Left = nil
		return
	}

	// If the node to remove has two children, replace it with its successor.
	successor := bst.FindSuccessor(nodeToRemove)
	nodeToRemove.Value = successor.Value
	bst.Remove(successor.Value)
}

func (bst *BST) Contains(value int) bool {
	// Find the node with the given value.
	node := bst.FindNode(value)
	return node != nil
}

func (bst *BST) FindNode(value int) *BSTNode {
	// Start at the root node.
	currentNode := bst.Root

	// Search for the node with the given value.
	for currentNode != nil {
		if value < currentNode.Value {
			currentNode = currentNode.Left
		} else if value > currentNode.Value {
			currentNode = currentNode.Right
		} else {
			return currentNode
		}
	}

	// The node with the given value was not found.
	return nil
}

func (bst *BST) FindSuccessor(nodeToRemove *BSTNode) *BSTNode {
	// Start at the right child of the node to remove.
	successor := nodeToRemove.Right

	// Find the leftmost child of the successor.
	for successor.Left != nil {
		successor = successor.Left
	}

	return successor
}

func (bst *BST) Print() {
	// Print the tree in a pre-order traversal.
	bst.PreOrderTraversal(bst.Root)
	fmt.Println()
}

func (bst *BST) PreOrderTraversal(node *BSTNode) {
	if node == nil {
		return
	}

	fmt.Printf("%d ", node.Value)
	bst.PreOrderTraversal(node.Left)
	bst.PreOrderTraversal(node.Right)
}

func main() {
	// Create a new BST.
	bst := &BST{}

	// Insert some values into the BST.
	bst
}
