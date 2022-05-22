package main

func numIslands2(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				grid[i][j] = '0'
				Q := make([]SaveNode, 0) //将遍历到的每一个岛屿用队列来存，再广搜
				Q = append(Q, SaveNode{i, j})
				for len(Q) != 0 {
					row := Q[0].iVal
					col := Q[0].jVal
					Q = Q[1:]
					if row-1 >= 0 && grid[row-1][col] == '1' {
						Q = append(Q, SaveNode{row - 1, col})
						grid[row-1][col] = '0'
					}
					if row+1 < len(grid) && grid[row+1][col] == '1' {
						Q = append(Q, SaveNode{row + 1, col})
						grid[row+1][col] = '0'
					}
					if col-1 >= 0 && grid[row][col-1] == '1' {
						Q = append(Q, SaveNode{row, col - 1})
						grid[row][col-1] = '0'
					}
					if col+1 < len(grid[0]) && grid[row][col+1] == '1' {
						Q = append(Q, SaveNode{row, col + 1})
						grid[row][col+1] = '0'
					}
				}
			}
		}
	}
	return res
}

type SaveNode struct {
	iVal int
	jVal int
}

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

//递归从i，j位置为起点上右下左递归，每次将位置置为'0'
func dfs(arr [][]byte, i, j int) {
	width := len(arr)
	height := len(arr[0])
	if i < 0 || i >= width || j < 0 || j >= height || arr[i][j] == '0' {
		return
	}
	arr[i][j] = '0'
	dfs(arr, i-1, j) //上
	dfs(arr, i, j+1) //右
	dfs(arr, i+1, j) //下
	dfs(arr, i, j-1) //左
}
