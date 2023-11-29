/*

The distance between a node in a Binary Tree and the tree's root is called the node's depth.
Write a function that takes in a Binary Tree and returns the sum of all of its subtrees· nodes· depths.
Each BinaryTree node has an integer value , a left child node, and a right child node. Children nodes can either be BinaryTree nodes themselves or None / null


Sample Input:
tree = 1
        / \
        2   3
        / \ / \
        4  5 6  7
         / \
        8  9


Sample Output:
26


// The depth of the node with value 2 is 1.
// The depth of the node with value 3 is 1.
// The depth of the node with value 4 is 2.
// The depth of the node with value 5 is 2.
// The depth of the node with value 6 is 2.
// The depth of the node with value 7 is 2.
// The depth of the node with value 8 is 3.
// The depth of the node with value 9 is 3.
// Therefore, the sum of all of the nodes' depths is 1




Hints
Hint 1

You can calculate the sum of a tree's node depths with a simple recursive function. Iterate through every node in the tree, call the simple recursive function on each node to caculate the sum of the node depths of the tree rooted at the node in question, and add up all of the sums to obtain the final sum.

Hint 2

You can solve this question in linear time by coming up with a relation between a tree's sum of node depths and the sums of node depths of the trees rooted at its left and right child nodes.

Hint 3

The depth of a node relative to a node X is 1 value smaller than its depth relative to node X's parent node Y. It follows that, if a subtree rooted at node X has a sum of node depths S, you can get the sum of those node depths relative to node Y by calculating: S + number-of-nodes-in-subtree-rooted-at-X, since this effectively increments all of the node depths relative to node X by 1.

Hint 4

From Hint #3, we can deduce the formula: nodeDepths(node) = nodeDepths(node.left) + numberOfNodesInLeftSubtree + nodeDepths(node.right) + numberOfNodesInRightSubtree. We can easily count the number of nodes in each subtree with a single pass in the input tree, and then we can apply this formula to calculate all of the node depths in linear time and finally sum them up.





*/

package main

import "fmt"

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func sumOfAllNodesDepths(tree *BinaryTreeNode) int {
	// Create a stack to store the nodes of the tree in reverse in-order traversal order.
	stack := []*BinaryTreeNode{}

	// Perform a reverse in-order traversal of the tree.
	node := tree
	for node != nil || len(stack) > 0 {
		// If the current node is not nil, push it onto the stack and move to its left child.
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			// Pop the top node from the stack and assign it to the current node.
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Calculate the depth of the current node.
			depth := 0
			for node.Parent != nil {
				depth++
				node = node.Parent
			}

			// Add the depth of the current node to the sum of all nodes' depths.
			sumOfAllNodesDepths += depth

			// Move to the current node's right child.
			node = node.Right
		}
	}

	// Return the sum of all nodes' depths.
	return sumOfAllNodesDepths
}

func main() {
	// Create a new binary tree.
	tree := &BinaryTreeNode{Value: 1}
	tree.Left = &BinaryTreeNode{Value: 2}
	tree.Right = &BinaryTreeNode{Value: 3}
	tree.Left.Left = &BinaryTreeNode{Value: 4}
	tree.Left.Right = &BinaryTreeNode{Value: 5}
	tree.Right.Left = &BinaryTreeNode{Value: 6}
	tree.Right.Right = &BinaryTreeNode{Value: 7}
	tree.Left.Left.Left = &BinaryTreeNode{Value: 8}
	tree.Left.Left.Right = &BinaryTreeNode{Value: 9}

	// Calculate the sum of all nodes' depths.
	sumOfAllNodesDepths := sumOfAllNodesDepths(tree)

	// Print the sum of all nodes' depths.
	fmt.Println(sumOfAllNodesDepths)
}
