package main

var (
	empty uint8 = '.'
	queue uint8 = 'Q'
)

func solveNQueens(n int) [][]string {
	broad := make([]string, n)
	for i := 0; i < n; i++ {
		row := make([]uint8, n)
		for j := 0; j < n; j++ {
			row[j] = empty
		}
		broad[i] = string(row)
	}

	isOk := func(row, col int) bool {
		for i := 0; i < row; i++ {
			if broad[i][col] == queue {
				return false
			}
		}
		tmpRow, tmpCol := row-1, col-1
		for tmpRow >= 0 && tmpCol >= 0 {
			if broad[tmpRow][tmpCol] == queue {
				return false
			}
			tmpRow--
			tmpCol--
		}
		tmpRow, tmpCol = row-1, col+1
		for tmpRow >= 0 && tmpCol < n {
			if broad[tmpRow][tmpCol] == queue {
				return false
			}
			tmpRow--
			tmpCol++
		}
		return true
	}

	setBroad := func(row, col int, value uint8) {
		t := []uint8(broad[row])
		t[col] = value
		broad[row] = string(t)
	}

	var ans [][]string
	var solveDfs func(row int)
	solveDfs = func(row int) {
		if row == n {
			ans = append(ans, append([]string{}, broad...))
			return
		}
		for col := 0; col < n; col++ {
			if !isOk(row, col) {
				continue
			}
			setBroad(row, col, queue)
			solveDfs(row + 1)
			// backtrack
			setBroad(row, col, empty)
		}
	}
	solveDfs(0)
	return ans
}
