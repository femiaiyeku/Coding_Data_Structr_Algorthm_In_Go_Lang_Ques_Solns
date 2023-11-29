/*



Write three functions that take in a Binary Search Tree (BST) and an empty array, traverse the BST,
add its nodes' values to the input array, and return that array. The three functions should traverse the BST using the in-order, pre-order,
and post-order tree­traversal techniques, respectively. If you're unfamiliar with tree-traversal techniques, we recommend watching the Conceptual
Overview section of this question's video explanation  before starting to code.
Each BST node has an integer value, a left child node, and a right child node. A node is said to be a valid BST node if and only if it satisfies the BST property:
its value is strictly greater than the values of every node to its left;
its value is less than or equal to the values of ever/ node to its right; and its children nodes are either valid BST nodes themselves or None / null.


Sample Input:

tree = 10
        / \
         5  15
        / \  / \
         2  5 13 22
        /       \
         1        14

Sample Output:

inOrderTraverse: [1, 2, 5, 5, 10, 13, 14, 15, 22]
preOrderTraverse: [10, 5, 2, 1, 5, 15, 13, 14, 22]
postOrderTraverse: [1, 2, 5, 5, 14, 13, 22, 15, 10]



Hints
Hint 1
Try traversing the BST node by node, all the while keeping track of the node with the value closest to the target value. Calculating the absolute value of the difference between a node's value and the target value should allow you to check if that node is closer than the current closest one.
Hint 2
Make use of the BST property to determine what side of any given node has values close to the target value and is therefore worth exploring.
Hint 3
What are the advantages and disadvantages of solving this problem iteratively as opposed to recursively?




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

func inOrderTraverse(bst *BST, array []int) []int {
	if bst.Root == nil {
		return array
	}

	// Traverse the left subtree.
	array = inOrderTraverse(bst.Root.Left, array)

	// Add the current node's value to the array.
	array = append(array, bst.Root.Value)

	// Traverse the right subtree.
	array = inOrderTraverse(bst.Root.Right, array)

	return array
}

func preOrderTraverse(bst *BST, array []int) []int {
	if bst.Root == nil {
		return array
	}

	// Add the current node's value to the array.
	array = append(array, bst.Root.Value)

	// Traverse the left subtree.
	array = preOrderTraverse(bst.Root.Left, array)

	// Traverse the right subtree.
	array = preOrderTraverse(bst.Root.Right, array)

	return array
}

func postOrderTraverse(bst *BST, array []int) []int {
	if bst.Root == nil {
		return array
	}

	// Traverse the left subtree.
	array = postOrderTraverse(bst.Root.Left, array)

	// Traverse the right subtree.
	array = postOrderTraverse(bst.Root.Right, array)

	// Add the current node's value to the array.
	array = append(array, bst.Root.Value)

	return array
}

func main() {
	// Create a new BST.
	bst := &BST{}

	// Insert some values into the BST.
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(2)
	bst.Insert(5)
	bst.Insert(13)
	bst.Insert(22)
	bst.Insert(1)
	bst.Insert(14)

	// Traverse the BST in-order and print the results.
	array := []int{}
	array = inOrderTraverse(bst, array)
	fmt.Println("inOrderTraverse:", array)

	// Traverse the BST pre-order and print the results.
	array = []int{}
	array = preOrderTraverse(bst, array)
	fmt.Println("preOrderTraverse:", array)

	// Traverse the BST post-order and print the results.
	array = []int{}
	array = postOrderTraverse(bst, array)
	fmt.Println("postOrderTraverse:", array)
}
