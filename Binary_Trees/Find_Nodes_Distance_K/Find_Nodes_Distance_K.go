/*


You're given the root node of a Binary Tree, a target value of a node that's contained in the tree, and a positive integer k .
 Write a function that returns the values of all the nodes that are exactly distance k from the node with target value.
The distance between two nodes is defined as the number of edges that must be traversed to go from one node to the other.
 For example, the distance between a node and its immediate left or right child is 1 . The same holds in reverse: the distance between a node and its parent is 1 .
 In a tree of three nodes where the root node has a left and right child, the left and right children are distance 2 from each other.
Each BinaryTree node has an integer value , a left child node, and a right child node. Children nodes can either be Bi naryTree nodes themselves or None / nul 1 .
Note that all BinaryTree node values will be unique, and your function can return the output values in any order.


Sample Input
tree =   1
        /   \
        2     3
        /   \     \
        4     5    6
                   /   \
                  7     8
target = 3
k = 2


Sample Output
[2, 7, 8]
// There are three nodes with value 2, 7, and 8 at distance 2 from node with value 3
// Note that there are two other nodes that are distance 2 from node 3:
//   - The node with value 1 is at distance 2 since it takes two "turns" to reach it from node 3.
//   - The node with value 6 is at distance 2 since it takes two "turns" to reach it from node 3.


Hints
Hint 1

Would it be easier to solve this problem if you had information about every node's parent node?

Hint 2

One approach to this problem is to find the parent nodes of all nodes in the tree. With this information you can perform a breadth-first search starting at the target node and traverse through each neighbor (left, right, and parent node) of every node, keeping track of your distance from the target node at each iteration. Once you reach a node that is distance k from the target node, you can add it to your output array. You'll have to also keep track of which nodes you've visited so as to avoid visiting the same nodes over and over again.

Hint 3

Another approach is to use a recursive depth-first-search algorithm as follows:

Case #1: when currentNode == target, search the subtree rooted at currentNode for all nodes that are k distance from currentNode.
Case #2: when target is in the left subtree of currentNode at distance L + 1, look for nodes that are distance k - L - 1 in the right subtree of currentNode.
Case #3: when target is in the right subtree of currentNode at distance L + 1, do the same thing as in case #2 but in the opposite subtree.
Case #4: when target is neither in the left nor in right subtree of currentNode, stop recursing.



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

func findNodesAtDistanceK(tree *BinaryTreeNode, target int, k int) []int {
	// Create a queue to store the nodes that need to be visited.
	queue := []*BinaryTreeNode{}

	// Add the target node to the queue.
	queue = append(queue, tree)

	// Create a set to store the nodes that have already been visited.
	visited := map[*BinaryTreeNode]bool{}

	// Perform a breadth-first search of the tree.
	for len(queue) > 0 {
		// Get the next node from the queue.
		node := queue[0]
		queue = queue[1:]

		// If the node has already been visited, skip it.
		if visited[node] {
			continue
		}

		// Mark the node as visited.
		visited[node] = true

		// If the distance from the target node to the current node is equal to k, add the current node to the result list.
		if k == 0 {
			result := append(result, node.Value)
		}

		// Add the current node's left and right children to the queue if they are not null.
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		// Decrement the distance k.
		k--
	}

	// Return the result list.
	return result
}

func main() {
	// Create the binary tree.
	tree := &BinaryTreeNode{Value: 1}
	tree.Left = &BinaryTreeNode{Value: 2}
	tree.Right = &BinaryTreeNode{Value: 3}
	tree.Left.Left = &BinaryTreeNode{Value: 4}
	tree.Left.Right = &BinaryTreeNode{Value: 5}
	tree.Right.Right = &BinaryTreeNode{Value: 6}
	tree.Right.Right.Left = &BinaryTreeNode{Value: 7}
	tree.Right.Right.Right = &BinaryTreeNode{Value: 8}

	// Find all the nodes that are exactly distance 2 from the node with target value 3.
	result := findNodesAtDistanceK(tree, 3, 2)

	// Print the result.
	fmt.Println(result)
}
