/*

You're given a binary expression tree. Write a function to evaluate this tree mathematically and return a single resulting integer.
All leaf nodes in the tree represent operands, which will always be positive integers. All of the other nodes represent operators. There are 4 operators supported, each of which is represented by a negative integer:



-	1 : Addition operator, adding the left and right subtrees.
-	2 : Subtraction operator, subtracting the right subtree from the left subtree.
-	3 : Division operator, dividing the left subtree by the right subtree. If the result is a decimal, it should be rounded towards zero.
-	4 : Multiplication operator, multiplying the left and right subtrees.



You can assume the tree will always be a valid expression tree. Each operator also works as a grouping symbol, meaning the bonom of the tree is always evaluated first, regardless of the operator.



Sample Input

tree =  -1
         /  \
        -2  -3
         / \  / \
        -4  2 8 3
       / \
      2  3

Sample Output

6








*/

package main

import (
	"fmt"
	"math"
)

type ExpressionTreeNode struct {
	Operator int
	Left     *ExpressionTreeNode
	Right    *ExpressionTreeNode
}

func evaluateExpressionTree(tree *ExpressionTreeNode) int {
	// If the tree is a leaf node, then it represents an operand, so return its value.
	if tree.Left == nil && tree.Right == nil {
		return tree.Operator
	}

	// Evaluate the left and right subtrees.
	leftSubtreeValue := evaluateExpressionTree(tree.Left)
	rightSubtreeValue := evaluateExpressionTree(tree.Right)

	// Perform the operation on the left and right subtrees.
	switch tree.Operator {
	case 1: // Addition
		return leftSubtreeValue + rightSubtreeValue
	case 2: // Subtraction
		return leftSubtreeValue - rightSubtreeValue
	case 3: // Division
		return int(math.Round(float64(leftSubtreeValue) / float64(rightSubtreeValue)))
	case 4: // Multiplication
		return leftSubtreeValue * rightSubtreeValue
	default:
		panic("Invalid operator")
	}
}

func main() {
	// Create the expression tree.
	tree := &ExpressionTreeNode{Operator: 1}
	tree.Left = &ExpressionTreeNode{Operator: 2}
	tree.Right = &ExpressionTreeNode{Operator: 3}
	tree.Left.Left = &ExpressionTreeNode{Operator: 4}
	tree.Left.Right = &ExpressionTreeNode{Operator: 2}
	tree.Right.Left = &ExpressionTreeNode{Operator: 8}
	tree.Right.Right = &ExpressionTreeNode{Operator: 3}
	tree.Left.Left.Left = &ExpressionTreeNode{Operator: 2}
	tree.Left.Left.Right = &ExpressionTreeNode{Operator: 3}

	// Evaluate the expression tree.
	result := evaluateExpressionTree(tree)

	// Print the result.
	fmt.Println(result)
}
