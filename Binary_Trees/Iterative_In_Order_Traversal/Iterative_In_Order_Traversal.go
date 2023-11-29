/*


Write a function that takes in a Binary Tree (where nodes have an additional pointer to their parent node) and traverses
it iteratively using the in-order tree-traversal technique; the traversal should specifically not use recursion. As the tree is
being traversed, a callback function passed in as an argument to the main function should be called on each node (i.e.,
callback(currentNode)).
Each BinaryTree node has an integer value, a parent node, a left child node, and a right child node.
Children nodes can either be BinaryTree nodes themselves or None / null.

Sample Input:

tree =    1
        /   \
        2     3
        /   \  / \
        4     6  7
          \
            9
callback = someCallbackFunction

Sample Output:

[4, 9, 2, 6, 1, 3, 7] // the order in which the callback function is called on the nodes

// The tree from the sample input is:
//        1
//      /   \    <-- parent-child relationships
//     2     3
//    / \   / \
//   4   6 7   null
//    \
//     9

// The callback function is called in the following order on the nodes of the tree:
// someCallbackFunction(4)
// someCallbackFunction(9)
// someCallbackFunction(2)
// someCallbackFunction(6)
// someCallbackFunction(1)
// someCallbackFunction(3)
// someCallbackFunction(7)


Hints
Hint 1

Start by realizing that in-order traversal always traverses left child nodes before parent nodes before right child nodes. In other words, you will somehow have to traverse the entire left side of the input Binary Tree before calling the input callback on the root node and before traversing the entire right side.

Hint 2

While each node in the input Binary Tree does have a "parent" property, allowing you to traverse your way back up the tree if need be, the difficulty arises when you must choose which node to actually call the input callback on. For instance, on your way back up the left side of the input tree, how do you know whether to traverse the right side of a node or to keep going up? Is there something that you can keep track of in order to know which node to call the input callback back on next at any time during the life of your algorithm?

Hint 3

Try keeping track of three nodes at all times: a current node (the node currently being traversed), a previous node (the node traversed just before the current one), and a next node (the next node to be traversed). Determine which node to traverse next and when to call the input callback on the current node by analyzing the previous node. For instance, if the previous node is actually the current node's left child node, then you know that you must call the callback on the current node and that you must then explore the right side of the current node before going back up. Figure out all of the possible scenarios, and develop an algorithm to handle all of these scenarios.







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

func iterativeInOrderTraversal(tree *BinaryTreeNode, callback func(*BinaryTreeNode)) {
	// Create a stack to store the nodes that need to be visited.
	stack := []*BinaryTreeNode{}

	// Initialize the current node to the root node.
	currentNode := tree

	// Push the current node onto the stack.
	stack = append(stack, currentNode)

	// While the stack is not empty, do the following:
	for len(stack) > 0 {
		// Pop the top node from the stack.
		currentNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Call the callback function on the current node.
		callback(currentNode)

		// If the current node has a right child, push it onto the stack.
		if currentNode.Right != nil {
			stack = append(stack, currentNode.Right)
		}

		// If the current node has a left child, push it onto the stack.
		if currentNode.Left != nil {
			stack = append(stack, currentNode.Left)
		}
	}
}

func main() {
	// Create the binary tree.
	tree := &BinaryTreeNode{Value: 1}
	tree.Left = &BinaryTreeNode{Value: 2}
	tree.Right = &BinaryTreeNode{Value: 3}
	tree.Left.Left = &BinaryTreeNode{Value: 4}
	tree.Left.Right = &BinaryTreeNode{Value: 6}
	tree.Right.Left = &BinaryTreeNode{Value: 7}
	tree.Left.Right.Right = &BinaryTreeNode{Value: 9}

	// Define a callback function.
	callback := func(node *BinaryTreeNode) {
		fmt.Println(node.Value)
	}

	// Traverse the binary tree iteratively using the in-order tree-traversal technique and call the callback function on each node.
	iterativeInOrderTraversal(tree, callback)
}
