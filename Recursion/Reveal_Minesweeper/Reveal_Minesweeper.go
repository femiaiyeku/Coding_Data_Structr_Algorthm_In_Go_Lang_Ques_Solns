/*



Minesweeper is a popular video game. From Wikipedia, "The game features a grid of clickable squares, with hidden "mines" scattered
throughout the board. The objective is to clear the board without detonating any mines, with help from clues about the number of neighboring
mines in each field." Specifically, when a player clicks on a square (also called a cell) that doesn't contain a mine, the square reveals a number
representing the number of immediately adjacent mines (including diagonally adjacent mines).
You're given a two-dimensional array of strings that represents a Minesweeper board for a game in progress. You're also given a row and a
column representing the indices of the next square that the player clicks on the board. Write a function that returns an updated board after the
click (your function can mutate the input board).
The board will always contain only strings, and each string will be one of the following:
● "N":Amine that has not been clicked on.
"X": A mine that has been clicked on, indicating a lost game.
4
• "H": A cell with no mine, but whose content is still hidden to the player.
• "8-8":A cell with no mine, with an integer from 0 to 8 representing the number of adjacent mines. Note that this is a single-digit integer
represented as a string. For example "2" would mean there are 2 adjacent cells with mines. Numbered cells are not clickable as they
have already been revealed.
If the player clicks on a mine, replace the "M" with "X", Indicating the game was lost.
If the player clicks on a cell adjacent to a mine, replace the "H" with a string representing the number of adjacent mines.
If the player clicks on a cell with no adjacent mines, replace the "H" with "e". Additionally, reveal all of the adjacent hidden cells as if the player had clicked on those cells as well
If the player clicks on a cell that isn't represented by the input board, do nothing.

Sample Input:
board = [
["E", "E", "E", "E", "E"],
["E", "E", "M", "E", "E"],
["E", "E", "E", "E", "E"],
["E", "E", "E", "E", "E"]]
click = [3, 0]
Sample Output:
[
["B", "1", "E", "1", "B"],
["B", "1", "M", "1", "B"],
["B", "1", "1", "1", "B"],
["B", "B", "B", "B", "B"]]
Explanation:
5

Sample Input:
board = [
["B", "1", "E", "1", "B"],
["B", "1", "M", "1", "B"],
["B", "1", "1", "1", "B"],
["B", "B", "B", "B", "B"]]
click = [1, 2]
Sample Output:
[
["B", "1", "E", "1", "B"],
["B", "1", "X", "1", "B"],
["B", "1", "1", "1", "B"],
["B", "B", "B", "B", "B"]]
Explanation:
6



This code first defines a function updateBoard that takes a Minesweeper board and a player's click as input and updates the board accordingly. The function first checks if the clicked cell is within the bounds of the board. If it is, it checks the value of the cell. If the cell is a mine, it updates the cell to "X" to indicate a lost game. If the cell is a hidden cell, it counts the number of adjacent mines and updates the cell accordingly. If the cell has no adjacent mines, it recursively reveals all adjacent hidden cells. Finally, it prints the updated board.

The main function creates a Minesweeper board with the given input data and simulates a player's click. It then prints the updated board.











*/

package main

import (
	"fmt"
)

func updateBoard(board [][]string, click []int) {
	row := click[0]
	col := click[1]

	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) {
		return
	}

	if board[row][col] == "M" {
		board[row][col] = "X"
		return
	}

	if board[row][col] == "H" {
		countMines := countAdjacentMines(board, row, col)

		if countMines == 0 {
			revealAdjacentCells(board, row, col)
		} else {
			board[row][col] = fmt.Sprintf("%d", countMines)
		}
	}
}

func countAdjacentMines(board [][]string, row int, col int) int {
	countMines := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			newRow := row + i
			newCol := col + j

			if isValidCell(board, newRow, newCol) && board[newRow][newCol] == "M" {
				countMines++
			}
		}
	}

	return countMines
}

func revealAdjacentCells(board [][]string, row int, col int) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			newRow := row + i
			newCol := col + j

			if isValidCell(board, newRow, newCol) && board[newRow][newCol] == "H" {
				revealAdjacentCells(board, newRow, newCol)
			} else if isValidCell(board, newRow, newCol) && board[newRow][newCol] == "" {
				board[newRow][newCol] = "B"
			}
		}
	}
}

func isValidCell(board [][]string, row int, col int) bool {
	return row >= 0 && row < len(board) && col >= 0 && col < len(board[0])
}

func main() {
	board := [][]string{
		{"E", "E", "E", "E", "E"},
		{"E", "E", "M", "E", "E"},
		{"E", "E", "E", "E", "E"},
		{"E", "E", "E", "E", "E"},
	}

	click := []int{3, 0}

	updateBoard(board, click)

	for _, row := range board {
		for _, cell := range row {
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}
}
