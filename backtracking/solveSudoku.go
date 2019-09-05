// https://leetcode.com/problems/sudoku-solver/
package main

func solveSudoku(board [][]byte) {
	nRow := len(board)
	nCol := len(board[0])
	isOk := func(n byte, row, col int) bool {
		for i := 0; i < nRow; i++ {
			if board[i][col] == n {
				return false
			}
		}
		for i := 0; i < nCol; i++ {
			if board[row][i] == n {
				return false
			}
		}
		row, col = row-row%3, col-col%3
		for i := row; i < row+3; i++ {
			for j := col; j < col+3; j++ {
				if board[i][j] == n {
					return false
				}
			}
		}
		return true
	}

	var bt func(r, c int) bool
	bt = func(r, c int) bool {
		if r == nRow {
			return true
		}
		if board[r][c] != '.' {
			if c == nCol-1 {
				return bt(r+1, 0)
			}
			return bt(r, c+1)
		}
		var n byte
		for n = '1'; n <= '9'; n++ {
			if isOk(n, r, c) {
				board[r][c] = n
				if c == nCol-1 {
					if bt(r+1, 0) {
						return true
					}
				} else {
					if bt(r, c+1) {
						return true
					}
				}
				board[r][c] = '.'
			}
		}
		return false
	}

	bt(0, 0)
}
