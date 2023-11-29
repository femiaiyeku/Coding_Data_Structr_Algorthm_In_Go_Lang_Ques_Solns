/*

Write a function that takes in a non-empty sorted array of distinct integers, constructs a BST from the integers, and returns the root of the BST.
The function should minimize the height of the BST.  You've been provided with a BST class that you'll have to use to construct the BST.
Each BST node has an integer value , a left child node, and a right child node. A node is said to be a valid BST node
if and only if it satisfies the BST property: its value is strictly greater than the values of every node to its left;
its value is less than or equal to the values of every node to its right; and its children nodes are either valid BST nodes themselves or None /
null. A BST is valid if and only if all of its nodes are valid BST nodes.
Note that the BST class already has an insert method which you can use if you want.

Sample Input
array = [1, 2, 5, 7, 10, 13, 14, 15, 22]

Sample Output
            10
            /    \
              2      14
            /   \   /   \
            1     5  13   15
                    \       \
                    7       22

            // This is one example of a BST with min height
            // that you could create from the input array.
            // You could create other BSTs with min height
            // from the same array; for example:
            //            10
            //          /     \
            //        5        15
            //      /   \     /   \
            //     2     7   13   22
            //   /           \
            //  1            14



Hints
Hint 1

Think about the properties of a BST. Looking at the pre-order-traversal nodes (values), how can you determine the right child of a particular node?

Hint 2

The right child of any BST node is simply the first node in the pre-order traversal whose value is larger than or equal to the particular node's value. From this, we know that the nodes in the pre-order traversal that come before the right child of a node must be in the left subtree of that node.

Hint 3

Once you determine the right child of any given node, you're able to generate the entire left and right subtrees of that node. You can do so by recursively creating the left and right child nodes of each subsequent node using the fact stated in Hint #2. A node that has no left and right children is naturally a leaf node.

Hint 4

To solve this problem with an optimal time complexity, you need to realize that it's unnecessary to locate the right child of every node. You can simply keep track of the pre-order-traversal position of the current node that needs to be created and try to insert that node as the left or right child of the relevant previously visited node. Since this tree is a BST, every node must satisfy the BST property; by somehow keeping track of lower and upper bounds for node values, you should be able to determine if a node can be inserted as the left or right child of another node. With this approach, you can solve this problem in linear time. See this question's video explanation for a more detailed explanation of this approach.


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

func buildMinHeightBST(array []int) *BSTNode {
	// If the array is empty, return nil.
	if len(array) == 0 {
		return nil
	}

	// Find the middle index of the array.
	middleIndex := len(array) / 2

	// Create a new BST node with the middle element of the array.
	newNode := &BSTNode{Value: array[middleIndex]}

	// Recursively build the left and right subtrees of the new node.
	newNode.Left = buildMinHeightBST(array[:middleIndex])
	newNode.Right = buildMinHeightBST(array[middleIndex+1:])

	// Return the new node.
	return newNode
}

func main() {
	// Create a new array of distinct integers.
	array := []int{1, 2, 5, 7, 10, 13, 14, 15, 22}

	// Build a BST with minimal height from the array.
	root := buildMinHeightBST(array)

	// Print the BST in pre-order traversal.
	printPreOrder(root)
}

func printPreOrder(node *BSTNode) {
	if node == nil {
		return
	}

	fmt.Println(node.Value)
	printPreOrder(node.Left)
	printPreOrder(node.Right)
}
