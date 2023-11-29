/*


Write a function that takes in a Binary Tree and returns a list of its branch sums ordered from leftmost branch sum to rightmost branch sum.
A branch sum is the sum of all values in a Binary Tree branch.
A Binary Tree branch is a path of nodes in a tree that starts at the root node and ends at any leaf node.
Each BinaryTree node has an integer value, a left child node, and a right child node.
Children nodes can either be Bi naryTree nodes themselves or None / nul 1 .

Sample Input
tree = 1
        / \
         2   3
        / \ / \
         4  5 6  7
         / \
            8 9
Sample Output
[15, 16, 18, 10, 11]
// 15 == 1 + 2 + 4 + 8
// 16 == 1 + 2 + 4 + 9
// 18 == 1 + 2 + 5 + 10
// 10 == 1 + 3 + 6
// 11 == 1 + 3 + 7


Hints
Hint 1
Try traversing the Binary Tree in a depth-first-search-like fashion.
Hint 2
Recursively traverse the Binary Tree in a depth-first-search-like fashion, and pass a running sum of the values of every previously-visited node to each node that you're traversing.
Hint 3
As you recursively traverse the tree, if you reach a leaf node (a node with no "left" or "right" Binary Tree nodes), add the relevant running sum that you've calculated to a list of sums (which you'll also have to pass to the recursive function). If you reach a node that isn't a leaf node, keep recursively traversing its children nodes, passing the correctly updated running sum to them.




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

func branchSums(tree *BinaryTreeNode) []int {
	// Create a stack to store the nodes of the tree in pre-order traversal order.
	stack := []*BinaryTreeNode{}

	// Create a list to store the branch sums.
	branchSums := []int{}

	// Perform a pre-order traversal of the tree.
	currentBranchSum := 0
	node := tree
	for node != nil || len(stack) > 0 {
		// If the current node is not nil, add its value to the current branch sum and push it onto the stack.
		if node != nil {
			currentBranchSum += node.Value
			stack = append(stack, node)
			node = node.Left
		} else {
			// Pop the top node from the stack and assign it to the current node.
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// If the current node is a leaf node, add the current branch sum to the list of branch sums.
			if node.Left == nil && node.Right == nil {
				branchSums = append(branchSums, currentBranchSum)
			}

			// Move to the current node's right child.
			node = node.Right
		}
	}

	// Return the list of branch sums.
	return branchSums
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

	// Calculate the branch sums.
	branchSums := branchSums(tree)

	// Print the branch sums.
	fmt.Println(branchSums)
}
