/*


Write a function that takes in a string made up of parentheses ( (and) ). The function should return an integer
representing the length of the longest balanced substring with regards to parentheses.
A string is said to be balanced if it has as many opening parentheses as it has closing parentheses and if no
parenthesis is unmatched. Note that an opening parenthesis can't match a closing parenthesis that comes before
it, and similarly, a closing parenthesis can't match an opening parenthesis that comes after it.

Sample Input
string = "(()))("

Sample Output
4 // The longest balanced substring is "(())".





This code iterates through the string and maintains a stack to keep track of open parentheses. For each opening parenthesis, it increments the currentLength counter and pushes a true value onto the stack. For each closing parenthesis, it checks if the stack is empty or if the top element is not a true value. If it is, it pops the element from the stack and increments the currentLength counter by 2. If the stack is empty or the top element is not a true value, it resets the currentLength counter to 0. Finally, it returns the maximum length of any balanced substring found during the iteration.


*/

package main

import (
	"fmt"
	"stack"
)

func findLongestBalancedSubstring(str string) int {
	maxLength := 0
	stack := stack.New()
	currentLength := 0

	for _, char := range str {
		if char == '(' {
			stack.Push(true)
			currentLength++
		} else {
			if stack.Len() > 0 && stack.Peek().(bool) {
				stack.Pop()
				currentLength += 2

				if currentLength > maxLength {
					maxLength = currentLength
				}
			} else {
				currentLength = 0
			}
		}
	}

	return maxLength
}

func main() {
	str := "(()))("
	maxLength := findLongestBalancedSubstring(str)
	fmt.Println("Length of longest balanced substring:", maxLength)
}
