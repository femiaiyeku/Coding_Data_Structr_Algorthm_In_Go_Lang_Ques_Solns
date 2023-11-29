/*

You're given a two-dimensional array that represents a 9x9 partially filled Sudoku board.
Write a function that returns the solved Sudoku board.
Sudoku is a famous number-placement puzzle in which you need to fill a 9x9 grid with integers in the range of 1-9 .
Each 9x9 Sudoku board is split into 9 3x3 subgrids, as seen in the illustration below, and starts out partially filled.

The objective is to fill the grid such that each row, column, and 3x3 subgrid contains the numbers 1-9 exactly once.
In other words, no row may contain the same digit more than once, no column may contain the same digit more than once, and none of the 9 3x3 subgrids may contain the same digit more than once.
Your input for this problem will always be a partially filled 9x9 two-dimensional array that represents a solvable Sudoku puzzle.
 Every element in the array will be an integer in the range of 0-9 , where a 0 represents an empty square that must be filled by your algorithm.
Note that you may modify the input array and that there will always be exactly one solution to each input Sudoku board. '

Sample Input

board = [
[7, 8, 0, 4, 0, 0, 1, 2, 0],
[6, 0, 0, 0, 7, 5, 0, 0, 9],
[0, 0, 0, 6, 0, 1, 0, 7, 8],
[0, 0, 7, 0, 4, 0, 2, 6, 0],
[0, 0, 1, 0, 5, 0, 9, 3, 0],
[9, 0, 4, 0, 6, 0, 0, 0, 5],
[0, 7, 0, 3, 0, 0, 0, 1, 2],
[1, 2, 0, 0, 0, 7, 4, 0, 0],
[0, 4, 9, 2, 0, 6, 0, 0, 7]
]

Sample Output

[
[7, 8, 5, 4, 3, 9, 1, 2, 6],
[6, 1, 2, 8, 7, 5, 3, 4, 9],
[4, 9, 3, 6, 2, 1, 5, 7, 8],
[8, 5, 7, 9, 4, 3, 2, 6, 1],
[2, 6, 1, 7, 5, 8, 9, 3, 4],
[9, 3, 4, 1, 6, 2, 7, 8, 5],
[5, 7, 8, 3, 9, 4, 6, 1, 2],
[1, 2, 6, 5, 8, 7, 4, 9, 3],
[3, 4, 9, 2, 1, 6, 8, 5, 7]
]







Hints
Hint 1

The brute-force approach to this problem is to generate every possible Sudoku board and to check each one until you find one that's valid. The issue with this approach is that there are 9^81 possible 9x9 Sudoku boards. This is an extremely large number, which makes it practically impossible to take this approach. How can you avoid generating every possible Sudoku board?

Hint 2

Keep in mind that a Sudoku board doesn't need to be entirely filled to figure out if it's invalid and won't lead to a solution. Try generating partially filled Sudoku boards until they become invalid, thereby abandoning solutions that will never lead to a properly solved board.

Hint 3

The method described in Hint #2 is more formally known as backtracking. This involves attempting to place digits into empty positions in the Sudoku board and checking at each insertion if that newly inserted digit makes the Sudoku board invalid. If it does, then you try to insert another digit until you find one that doesn't invalidate the board. If it doesn't invalidate the board, you temporarily place that digit and continue to try to solve the rest of the board. If you ever reach a position where there are no valid digits to be inserted (every digit placed in that position leads to an invalid board), that means that one of the previously inserted digits is incorrect. Thus, you must backtrack and change previously placed digits. For more details on this approach, refer to the Conceptual Overview section of this question's video explanation.




This code first defines a function solveSudoku that takes a Sudoku board as input and tries to solve it using backtracking. The function recursively tries to place each number from 1 to 9 in each empty cell of the board. If a valid placement is found, it recursively calls itself to solve the remaining cells. If no valid placement is found, it backtracks and tries a different number or a different empty cell.

The findNextEmpty function finds the next empty cell in the Sudoku board. The isValidPlacement function checks if it is valid to place a given number in a given cell. The main function creates a Sudoku board with the given input data and calls the solveSudoku function to solve it. It then prints the solved board or a message indicating that there is no solution.







*/

package main

import (
	"fmt"
)

func solveSudoku(board [][]int) bool {
	emptyRow, emptyCol := findNextEmpty(board)

	if emptyRow == -1 && emptyCol == -1 {
		return true
	}

	for number := 1; number <= 9; number++ {
		if isValidPlacement(board, emptyRow, emptyCol, number) {
			board[emptyRow][emptyCol] = number

			if solveSudoku(board) {
				return true
			}

			board[emptyRow][emptyCol] = 0 // Backtrack
		}
	}

	return false
}

func findNextEmpty(board [][]int) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return i, j
			}
		}
	}

	return -1, -1 // No empty cells found
}

func isValidPlacement(board [][]int, row, col, number int) bool {
	// Check row
	for j := 0; j < 9; j++ {
		if board[row][j] == number {
			return false
		}
	}

	// Check column
	for i := 0; i < 9; i++ {
		if board[i][col] == number {
			return false
		}
	}

	// Check 3x3 subgrid
	subgridRowStart := (row / 3) * 3
	subgridColStart := (col / 3) * 3

	for i := subgridRowStart; i < subgridRowStart+3; i++ {
		for j := subgridColStart; j < subgridColStart+3; j++ {
			if board[i][j] == number {
				return false
			}
		}
	}

	return true
}

func main() {
	board := [][]int{
		{7, 8, 0, 4, 0, 0, 1, 2, 0},
		{6, 0, 0, 0, 7, 5, 0, 0, 9},
		{0, 0, 0, 6, 0, 1, 0, 7, 8},
		{0, 0, 7, 0, 4, 0, 2, 6, 0},
		{0, 0, 1, 0, 5, 0, 9, 3, 0},
		{9, 0, 4, 0, 6, 0, 0, 0, 5},
		{0, 7, 0, 3, 0, 0, 0, 1, 2},
		{1, 2, 0, 0, 0, 7, 4, 0, 0},
		{0, 4, 9, 2, 0, 6, 0, 0, 7},
	}

	if solveSudoku(board) {
		fmt.Println("Solved Sudoku board:")

		for _, row := range board {
			for _, cell := range row {
				fmt.Print(cell, " ")
			}
			fmt.Println()
		}
	} else {
		fmt.Println("The given Sudoku puzzle has no solution.")
	}
}
