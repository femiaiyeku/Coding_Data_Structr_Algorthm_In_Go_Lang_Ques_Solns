/*

You're given the position of two knight pieces on an infinite chess board. Write a function that returns the minimum number of turns
required before one of the knights is able to capture the other knight, assuming the knights are working together to achieve this goal.
The position of each knight is given as a list of 2 values, the x and y coordinates. A knight can make 1 of 8 possible moves on any given turn.
Each of these moves involves moving in an "L" shape. This means they can either move 2 squares horizontally and 1 square vertically, or
they can move 1 square horizontally and 2 squares vertically. For example, if a knight is currently at position [0, 0], then it can move to any
of these 8 locations on its next move:
[
[-2, 1], [-1, 2], [1, 2], [2, 1], [2, -1], [1, -2], [-1, -2], [-2, -1]
]
A knight is able to capture the other knight when it is able to move onto the square currently occupled by the other knight.
Each turn allows each knight to move up to one time. For example, if both knights moved towards each other once, and then knightA
captures knightB on its next move, two turns would have been used (even though knightB never made its second move).


Sample Input:
knightA = [0, 0]
knightB = [4, 2]

Sample Output:
1 // knightA can move to [2, 1], and then knightB captures KnightA on its next turn by moving to [2, 1]


This code first checks if either knight can capture the other knight on their first move. If so, it returns one.

Then, it creates a queue of knights to process. The queue initially contains the positions of both knights. It also creates a visited map to keep track of the positions that have already been visited.

While the queue is not empty, it removes the first knight from the queue and checks all of its possible moves. If a move would allow the knight to capture the other knight, it returns the length of the queue, which represents the number of turns required to reach that position.

If a move does not allow the knight to capture the other knight, it checks if the position has already been visited. If it has not, it adds the position to the visited map and the queue.

Finally, if the queue is empty and no knight has been able to capture the other knight, it returns -1, indicating that no possible moves exist.


*/

package main

import (
	"fmt"
)

func minMovesToCapture(knightA, knightB []int) int {
	if abs(knightA[0]-knightB[0]) == 1 && abs(knightA[1]-knightB[1]) == 2 ||
		abs(knightA[0]-knightB[0]) == 2 && abs(knightA[1]-knightB[1]) == 1 {
		return 1
	}

	queue := [][]int{{knightA[0], knightA[1]}, {knightB[0], knightB[1]}}
	visited := make(map[[2]int]bool)

	for len(queue) > 0 {
		currentKnight := queue[0]
		queue = queue[1:]

		for _, move := range possibleMoves {
			newX := currentKnight[0] + move[0]
			newY := currentKnight[1] + move[1]

			if abs(newX-knightB[0]) == 1 && abs(newY-knightB[1]) == 2 ||
				abs(newX-knightB[0]) == 2 && abs(newY-knightB[1]) == 1 {
				return len(queue)
			}

			if visited[[2]int{newX, newY}] {
				continue
			}

			visited[[2]int{newX, newY}] = true
			queue = append(queue, []int{newX, newY})
		}
	}

	return -1 // No possible moves
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

var possibleMoves = [][]int{
	{-2, 1},
	{-1, 2},
	{1, 2},
	{2, 1},
	{2, -1},
	{1, -2},
	{-1, -2},
	{-2, -1},
}

func main() {
	knightA := []int{0, 0}
	knightB := []int{4, 2}

	minMoves := minMovesToCapture(knightA, knightB)

	fmt.Println(minMoves)
}
