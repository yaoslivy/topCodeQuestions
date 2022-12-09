package main

func islandPerimeter(grid [][]int) int {
	//从左往右，从上往下，先统计为1的格子的周长
	//当当前格子上或左也是1时，则需要减去双倍的重叠边的个数
	if len(grid) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				res += 4
				//判断上边和左边是否有岛屿
				if i-1 >= 0 && grid[i-1][j] == 1 {
					res -= 2
				}
				if j-1 >= 0 && grid[i][j-1] == 1 {
					res -= 2
				}
			}
		}
	}
	return res
}
