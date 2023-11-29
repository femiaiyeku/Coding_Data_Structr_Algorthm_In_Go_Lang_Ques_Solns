/*



Given two binary trees, merge them and return the resulting tree.
If two nodes overlap during the merger then sum the values, otherwise use the existing node.
Note that your solution can either mutate the existing trees or return a new tree.

tree1 =  1
        / \
       3   2
       /   \
       7    4

tree2 =  2
        / \
        5   9
        /   / \
        2   7  6

output =    2
            / \
            8   11
            / \  / \
            9  4 7  6



*/

package main

import "fmt"

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func mergeBinaryTrees(tree1 *BinaryTreeNode, tree2 *BinaryTreeNode) *BinaryTreeNode {
	// If either tree is empty, return the other tree.
	if tree1 == nil {
		return tree2
	}
	if tree2 == nil {
		return tree1
	}

	// Merge the values of the two nodes.
	tree1.Value += tree2.Value

	// Recursively merge the left and right subtrees.
	tree1.Left = mergeBinaryTrees(tree1.Left, tree2.Left)
	tree1.Right = mergeBinaryTrees(tree1.Right, tree2.Right)

	// Return the merged tree.
	return tree1
}

func main() {
	// Create the two binary trees.
	tree1 := &BinaryTreeNode{Value: 1}
	tree1.Left = &BinaryTreeNode{Value: 3}
	tree1.Right = &BinaryTreeNode{Value: 2}
	tree1.Left.Left = &BinaryTreeNode{Value: 7}
	tree1.Left.Right = &BinaryTreeNode{Value: 4}

	tree2 := &BinaryTreeNode{Value: 2}
	tree2.Left = &BinaryTreeNode{Value: 5}
	tree2.Right = &BinaryTreeNode{Value: 9}
	tree2.Left.Left = &BinaryTreeNode{Value: 2}
	tree2.Right.Left = &BinaryTreeNode{Value: 7}
	tree2.Right.Right = &BinaryTreeNode{Value: 6}

	// Merge the two binary trees.
	mergedTree := mergeBinaryTrees(tree1, tree2)

	// Print the merged binary tree.
	fmt.Println(mergedTree)
}
