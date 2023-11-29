/*


Write a function that takes in a string made up of brackets ((. [..). ], and )) and other optional
characters. The function should return a boolean representing whether the string is balanced with regards to
brackets.
A string is said to be balanced If It has as many opening brackets of a certain type as it has closing brackets of that
type and if no bracket is unmatched. Note that an opening bracket can't match a corresponding closing bracket
that comes before it, and similarly, a closing bracket can't match a corresponding opening bracket that comes afte
it. Also, brackets can't overlap each other as in [(]).

Sample Input:
string = "([])(){}(())()()"

Sample Output:
true // it's balanced






This code uses a stack to keep track of the opening brackets. It iterates through the string and checks for opening and closing brackets. If it encounters an opening bracket, it pushes it onto the stack. If it encounters a closing bracket, it checks if the top element on the stack is the matching opening bracket. If not, the string is not balanced. If it is, it pops the opening bracket from the stack. Finally, if the stack is empty after processing the entire string, the string is balanced.



*/

package main

import (
	"fmt"
)

func isBalanced(str string) bool {
	stack := []string{}

	for _, char := range str {
		bracket := string(char)

		if isOpeningBracket(bracket) {
			stack = append(stack, bracket)
		} else if isClosingBracket(bracket) {
			if len(stack) == 0 {
				return false
			}

			topBracket := stack[len(stack)-1]
			if !isMatchingBracket(topBracket, bracket) {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func isOpeningBracket(bracket string) bool {
	return bracket == "(" || bracket == "[" || bracket == "{"
}

func isClosingBracket(bracket string) bool {
	return bracket == ")" || bracket == "]" || bracket == "}"
}

func isMatchingBracket(openingBracket string, closingBracket string) bool {
	if openingBracket == "(" && closingBracket == ")" {
		return true
	} else if openingBracket == "[" && closingBracket == "]" {
		return true
	} else if openingBracket == "{" && closingBracket == "}" {
		return true
	}

	return false
}

func main() {
	str := "([])(){}(())()()"
	balanced := isBalanced(str)
	fmt.Println("String is balanced:", balanced)
}
