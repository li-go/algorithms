package main

import (
	"fmt"
)

func main() {
	fmt.Println("permuteUnique:", permuteUnique([]int{1, 2, 3, 3}))
	fmt.Println("solveNQueens:", solveNQueens(4))
	fmt.Println("generateParenthesis:", generateParenthesis(3))
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
	fmt.Println("solveSudoku:")
	for _, r := range board {
		fmt.Println(string(r))
	}
}
