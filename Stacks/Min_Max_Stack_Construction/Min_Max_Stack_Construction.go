/*


Write a MinMaxStack class for a Min Max Stack. The class should support:

• 	Pushing and popping values on and off the stack.
• 	Peeking at the value at the top of the stack.
• 	Getting both the minimum and the maximum values in the stack at any given point in time.

All class methods, when considered independently, should run in constant time and with constant space.


Sample Usuage

M1nMaxStack(): -push(S): -getM1 n(): 5 getMax(): 5 peek(): 5













*/

package main

import (
	"fmt"
)

type MinMaxNode struct {
	value int
	min   int
	max   int
}

type MinMaxStack struct {
	stack []MinMaxNode
}

func NewMinMaxStack() *MinMaxStack {
	return &MinMaxStack{stack: []MinMaxNode{}}
}

func (stack *MinMaxStack) Push(value int) {
	var minValue int
	if len(stack.stack) > 0 {
		minValue = stack.stack[len(stack.stack)-1].min
	} else {
		minValue = value
	}

	var maxValue int
	if len(stack.stack) > 0 {
		maxValue = stack.stack[len(stack.stack)-1].max
	} else {
		maxValue = value
	}

	newNode := MinMaxNode{value: value, min: minValue, max: maxValue}
	stack.stack = append(stack.stack, newNode)
}

func (stack *MinMaxStack) Pop() int {
	if len(stack.stack) == 0 {
		panic("Stack is empty")
	}

	poppedNode := stack.stack[len(stack.stack)-1]
	stack.stack = stack.stack[:len(stack.stack)-1]

	return poppedNode.value
}

func (stack *MinMaxStack) Peek() int {
	if len(stack.stack) == 0 {
		panic("Stack is empty")
	}

	return stack.stack[len(stack.stack)-1].value
}

func (stack *MinMaxStack) GetMin() int {
	if len(stack.stack) == 0 {
		panic("Stack is empty")
	}

	return stack.stack[len(stack.stack)-1].min
}

func (stack *MinMaxStack) GetMax() int {
	if len(stack.stack) == 0 {
		panic("Stack is empty")
	}

	return stack.stack[len(stack.stack)-1].max
}

func main() {
	stack := NewMinMaxStack()
	stack.Push(5)
	fmt.Println("Peek:", stack.Peek())
	fmt.Println("Min:", stack.GetMin())
	fmt.Println("Max:", stack.GetMax())
}
