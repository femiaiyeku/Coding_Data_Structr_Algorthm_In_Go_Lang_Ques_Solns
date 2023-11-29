/*


Write a function that takes in the root nodes of two Binary Trees and returns a boolean representing whether their leaf traversals are the same.
The leaf traversal of a Binary Tree traverses only its leaf nodes from left to right. A leaf node is any node that has no left or right children.
For example, the leaf traversal of the following Binary Tree is 1, 3, 2 .


Sample Input:
tree1 =

        1
         / \
        2   3
         /   / \
        4   6   7
         / \   \
        9   10  8

tree2 =   1
            / \
          2   3
                /   / \
                 4   6   7
                        / \   \
                       9   10  8

Sample Output:
true




Hints
Hint 1

To traverse the leaf nodes of a tree from left to right, you'll need to use a pre-order traversal.

Hint 2

The simplest approach to solving this problem is to perform a pre-order traversal on both trees, to store their leaf nodes in arrays in the order in which they're visited, and to then compare the two resulting arrays. This solutions works, but it's not optimal from a space-complexity perspective. Can you think of a way to solve this problem using less extra space? It's possible to solve this with O(h1 + h2) space or better, where h1 is the height of tree1 and h2 is the height of tree2.

Hint 3

To solve this problem with a more optimal space complexity, you can perform pre-order traversals on both trees at the same time. As you traverse the trees, you need to look for the next leaf node in each tree and pause the traversal as soon as you find it. Once you've found the next leaf node in both trees, you can compare their values and see if they match; if they do, continue the traversal , and repeat the process. If they don't match, the leaf traversals aren't the same, and you can return false.

Hint 4

Another unique way to solve this problem is to connect all of the leaf nodes in each individual tree so as to form two linked lists. Since the leaf nodes don't have any children, you can use their right pointers to store the next leaf nodes in the leaf traversals. And since you're reusing the input trees to store the leaf traversals, the only extra space you'll be using comes from the recursion used in the traversal of the trees. Once both trees have their leaf nodes connected, you can iterate through the linked lists and check if they're the same. To compare the linked lists, you'll need to keep track of their heads (the first leaf node in each tree).




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

func areLeafTraversalsSame(tree1 *BinaryTreeNode, tree2 *BinaryTreeNode) bool {
	// Create two stacks to store the leaf nodes of the two binary trees in order.
	stack1 := []*BinaryTreeNode{}
	stack2 := []*BinaryTreeNode{}

	// Perform a pre-order traversal of the first binary tree and push all of its leaf nodes onto the first stack.
	node := tree1
	for node != nil {
		if node.Left == nil && node.Right == nil {
			stack1 = append(stack1, node)
		}

		if node.Left != nil {
			node = node.Left
		} else if node.Right != nil {
			node = node.Right
		} else {
			node = nil
		}
	}

	// Perform a pre-order traversal of the second binary tree and push all of its leaf nodes onto the second stack.
	node = tree2
	for node != nil {
		if node.Left == nil && node.Right == nil {
			stack2 = append(stack2, node)
		}

		if node.Left != nil {
			node = node.Left
		} else if node.Right != nil {
			node = node.Right
		} else {
			node = nil
		}
	}

	// Compare the leaf nodes of the two binary trees.
	for i := 0; i < len(stack1) && i < len(stack2); i++ {
		if stack1[i] != stack2[i] {
			return false
		}
	}

	// If the two stacks have the same length, then the two binary trees have the same leaf traversals.
	return len(stack1) == len(stack2)
}

func main() {
	// Create the two binary trees.
	tree1 := &BinaryTreeNode{Value: 1}
	tree1.Left = &BinaryTreeNode{Value: 2}
	tree1.Right = &BinaryTreeNode{Value: 3}
	tree1.Left.Left = &BinaryTreeNode{Value: 4}
	tree1.Left.Right = &BinaryTreeNode{Value: 6}
	tree1.Right.Left = &BinaryTreeNode{Value: 7}
	tree1.Left.Left.Left = &BinaryTreeNode{Value: 9}
	tree1.Left.Left.Right = &BinaryTreeNode{Value: 10}
	tree1.Right.Left.Right = &BinaryTreeNode{Value: 8}

	tree2 := &BinaryTreeNode{Value: 1}
	tree2.Left = &BinaryTreeNode{Value: 2}
	tree2.Right = &BinaryTreeNode{Value: 3}
	tree2.Right.Left = &BinaryTreeNode{Value: 4}
	tree2.Right.Right = &BinaryTreeNode{Value: 6}
	tree2.Right.Left.Left = &BinaryTreeNode{Value: 7}
	tree2.Right.Left.Right = &BinaryTreeNode{Value: 9}
	tree2.Right.Left.Right.Right = &BinaryTreeNode{Value: 10}
	tree2.Right.Right.Right = &BinaryTreeNode{Value: 8}

	// Check if the two binary trees have the same leaf traversals.
	areLeafTraversalsSame := areLeafTraversalsSame(tree1, tree2)

	// Print the result.
	fmt.Println(areLeafTraversalsSame)
}
