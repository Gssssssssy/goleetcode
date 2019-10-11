package l64

/*
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。
*/

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func minPathSum(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i != 0 && j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else if i == 0 && j != 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else {
				grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			}
		}
	}
	return grid[rows-1][cols-1]
}
