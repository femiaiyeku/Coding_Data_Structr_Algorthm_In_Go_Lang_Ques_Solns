/*



You're given a list of string tokens representing a mathematical expression using Reverse Polish Notation. Reverse Polish Notation is a notation where operators come after operands, instead of between them. For example 2 4 + would evaluate to 6.
Parenthesis are always implicit in P.everse Polish Notation, meaning an expression is evaluated from left to right_ All of the
operators for this problem take two operands, which will always be the two values immediately preceding the operator. For example, 18 4 - 7 / would evaluate to ( (18 - 4) / 7) or 2 .

Write a function that takes this list of tokens and returns the result. Your function should support four operators: + , - and / for addition, subtraction, multiplication, and division respectively.



Division should always be treated as integer division, rounding towards zero. For example, 3 / 2 evaluates to 1 and -3 / 2 evaluates to -1 . You can assume the input will always be valid Reverse Polish Notation, and it will always result in a valid number. Your code should not edit this input list.

Sample Input

tokens = ["SO", "3", "17', •+•, "2", - , "/")

Sample Output
2



This code uses a stack to keep track of the operands. It iterates through the tokens and checks if each token is an operator or an operand. If it's an operator, it pops the two top elements from the stack, performs the operation, and pushes the result back onto the stack. If it's an operand, it pushes the operand onto the stack. Finally, the last element on the stack is the result of the RPN expression.




*/

package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, token := range tokens {
		if isOperator(token) {
			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			operand1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result := 0
			switch token {
			case "+":
				result = operand1 + operand2
			case "-":
				result = operand1 - operand2
			case "*":
				result = operand1 * operand2
			case "/":
				result = operand1 / operand2
			}

			stack = append(stack, result)
		} else {
			operand, _ := strconv.Atoi(token)
			stack = append(stack, operand)
		}
	}

	return stack[0]
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

func main() {
	tokens := []string{"3", "17", "+", "2", "-"}
	result := evalRPN(tokens)
	fmt.Println("Result:", result)
}
