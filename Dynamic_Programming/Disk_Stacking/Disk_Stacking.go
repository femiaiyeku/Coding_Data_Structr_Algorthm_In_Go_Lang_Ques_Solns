/*

You're given a non-empty array of arrays where each subarray holds three integers and represents a disk. These
integers denote each disk's width, depth, and height, respectively. Your goal is to stack up the disks and to maximize the
total height of the stack. A disk must have a strictly smaller width, depth, and height than any other disk below it.
Write a function that returns an array of the disks in the final stack, starting with the top disk and ending with the
bottom disk. Note that you can't rotate disks; in other words, the integers in each subarray must represent
[width, depth, height] at all times.
You can assume that there will only be one stack with the greatest total height.


Sample Input:
disks = [[2, 1, 2], [3, 2, 3], [2, 2, 8], [2, 3, 4], [1, 3, 1], [4, 4, 5]]

Sample Output:
[[2, 1, 2], [3, 2, 3], [4, 4, 5]]






*/

package main

import (
	"fmt"
	"sort"
)

type Disk struct {
	width  int
	depth  int
	height int
}

func stackDisks(disks [][3]int) []Disk {
	// Sort the disks by height in descending order
	sort.Slice(disks, func(i, j int) bool {
		return disks[i][2] > disks[j][2]
	})

	// Create a stack to store the disks
	stack := []Disk{}

	// Keep track of the maximum height of the stack
	maxHeight := 0

	// Iterate through the disks
	for _, disk := range disks {
		// If the disk's width, depth, and height are all strictly smaller than the maximum height of the stack,
		// then add the disk to the stack and update the maximum height of the stack.
		if disk[0] < maxHeight && disk[1] < maxHeight && disk[2] < maxHeight {
			stack = append(stack, Disk{width: disk[0], depth: disk[1], height: disk[2]})
			maxHeight = disk[2]
		}
	}

	// Return the stack of disks in the final order
	return stack
}

func main() {
	// Set the input values.
	disks := [][]int{{2, 1, 2}, {3, 2, 3}, {2, 2, 8}, {2, 3, 4}, {1, 3, 1}, {4, 4, 5}}

	// Stack the disks and maximize the total height of the stack.
	stackedDisks := stackDisks(disks)

	// Print the stacked disks.
	fmt.Println(stackedDisks)
}
